package async

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/pkg/node"
)

const nodesPrefix string = "nodes"

type Storage interface {
	Put(context.Context, string, string) error
	Get(context.Context, string) ([]*mvccpb.KeyValue, error)
}

type Node struct {
	name    string
	storage Storage
}

type NodeList struct {
	storage Storage
}

func NewNode(name string, storage Storage) *Node {
	return &Node{name, storage}
}

func NewNodeList(storage Storage) *NodeList {
	return &NodeList{
		storage: storage,
	}
}

func (nl *NodeList) List(ctx context.Context) (*resource.List, error) {
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

	resources := resource.NewList(nodes)

	return resources, nil
}

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

func (a *Node) Get(ctx context.Context) (*node.Node, error) {
	kvs, err := a.storage.Get(ctx, a.Key())
	if err != nil {
		return nil, err
	}

	data := kvs[0].Value
	n := node.New()

	if err := json.Unmarshal(data, n); err != nil {
		return nil, err
	}

	return n, nil
}

func (a *Node) Key() string {
	return filepath.Join("nodes", a.name)
}
