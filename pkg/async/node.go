package async

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/pkg/node"
)

const nodesPrefix string = "nodes"

// Storage is the interface to manage resources within the KV Store
type Storage interface {
	Put(context.Context, string, string) error
	Get(context.Context, string) ([]*mvccpb.KeyValue, error)
	Delete(context.Context, string) error
}

// Node is the representation of a node but implements asynchronous methods
// to interact with its resources
type Node struct {
	name    string
	storage Storage
}

// NodeList is a list of asynchronous nodes
type NodeList struct {
	storage Storage
}

// NewNode receives a name and the storage reference and returns a new
// asynchronous node struct
func NewNode(name string, storage Storage) *Node {
	return &Node{name, storage}
}

// NewNodeList receives the storage reference for the KV store and returns
// an asynchronous NodeList struct to manage a group of nodes
func NewNodeList(storage Storage) *NodeList {
	return &NodeList{
		storage: storage,
	}
}

// List receives a context which is passed to the storage to handle its operations,
// and returns a list of resources containing all the node.Node objects stored in
// the KV Store
func (nl *NodeList) List(ctx context.Context) (*resource.ResourceList, error) {
	kvs, err := nl.storage.Get(ctx, nodesPrefix)
	if err != nil {
		return nil, err
	}

	nodes := []*node.Node{}

	for _, kv := range kvs {
		data := kv.Value
		n := node.New()

		if err := json.Unmarshal(data, n); err != nil {
			return nil, err
		}

		nodes = append(nodes, n)
	}

	resources := resource.NewResourceList(nodes)

	return resources, nil
}

// Get receives a context which is passed to the storage to handle its operations,
// and returns a list of asynchronous nodes, for which clients can use to get the
// respective node.Node objects later
func (nl *NodeList) Get(ctx context.Context) ([]*Node, error) {
	kvs, err := nl.storage.Get(ctx, nodesPrefix)
	if err != nil {
		return []*Node{}, err
	}

	nodes := []*Node{}

	for _, kv := range kvs {
		nodes = append(nodes, NewNode(
			filepath.Base(string(kv.Key)),
			nl.storage,
		))
	}

	return nodes, nil
}

// Put receives a context which is passed to the storage to handle its operations,
// and a node.Node reference which must be updated in the storage
func (a *Node) Put(ctx context.Context, n *node.Node) error {
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}

	if err := a.storage.Put(ctx, a.Key(), string(data)); err != nil {
		return err
	}

	return nil
}

// Get receives a context which is passed to the storage to handle its operations,
// and returns the corresponding node.Node reference loaded from the KV store
func (a *Node) Get(ctx context.Context) (*node.Node, error) {
	kvs, err := a.storage.Get(ctx, a.Key())
	if err != nil {
		return nil, err
	}

	if len(kvs) == 0 {
		return nil, fmt.Errorf("node not found")
	}

	data := kvs[0].Value
	n := node.New()

	if err := json.Unmarshal(data, n); err != nil {
		return nil, err
	}

	return n, nil
}

// Delete removes the node reference from the KV store
func (a *Node) Delete(ctx context.Context) error {
	err := a.storage.Delete(ctx, a.Key())
	if err != nil {
		return err
	}

	return nil
}

// Key returns the Node key name used in the KV store
func (a *Node) Key() string {
	return filepath.Join("nodes", a.name)
}
