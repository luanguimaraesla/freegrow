package async

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/luanguimaraesla/freegrow/pkg/node"
)

type Storage interface {
	Put(context.Context, string, string) error
	Get(context.Context, string) ([]*mvccpb.KeyValue, error)
}

type asyncNode struct {
	name    string
	storage Storage
}

func NewAsyncNode(name string, storage Storage) *asyncNode {
	return &asyncNode{name, storage}
}

func (a *asyncNode) Put(ctx context.Context, n *node.Node) error {
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}

	if err := a.storage.Put(ctx, a.Key(), string(data)); err != nil {
		return err
	}

	return nil
}

func (a *asyncNode) Get(ctx context.Context) (*node.Node, error) {
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

func (a *asyncNode) Key() string {
	return filepath.Join("nodes", a.name)
}
