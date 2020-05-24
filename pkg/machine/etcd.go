package machine

import (
	"context"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
	client    *clientv3.Client
}

func (e *Etcd) Init() error {
	cli, err := e.Client()
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	response, err := cli.Get(
		ctx,
		"",
		clientv3.WithCountOnly(),
		clientv3.WithPrefix(),
	)
	if err != nil {
		return err
	}

	global.Logger.Debug("connected", zap.Int64("objects", response.Count))

	return nil
}

func (e *Etcd) Client() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   e.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (e *Etcd) Put(ctx context.Context, key, value string) error {
	cli, err := e.Client()
	if err != nil {
		return err
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)
	if _, err := kv.Put(ctx, key, value); err != nil {
		return err
	}

	return nil
}

func (e *Etcd) Get(ctx context.Context, key string) ([]*mvccpb.KeyValue, error) {
	cli, err := e.Client()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)
	gr, err := kv.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	return gr.Kvs, nil
}

func (e *Etcd) Delete(ctx context.Context, key string) error {
	cli, err := e.Client()
	if err != nil {
		return err
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)
	if _, err := kv.Delete(ctx, key, clientv3.WithPrefix()); err != nil {
		return err
	}

	return nil

}
