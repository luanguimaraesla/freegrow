package machine

import (
	"context"
	"io/ioutil"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/luanguimaraesla/freegrow/internal/etcd"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/pkg/async"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Storage interface {
	Init() error
	Put(context.Context, string, string) error
	Get(context.Context, string) ([]*mvccpb.KeyValue, error)
	Delete(context.Context, string) error
}

type Machine struct {
	Kind     string             `yaml:"kind"`
	Metadata *resource.Metadata `yaml:"metadata"`
	Spec     *MachineSpec       `yaml:"spec"`
	storage  Storage
	logger   *zap.Logger
}

type MachineSpec struct {
	Bind string    `yaml:"bind"`
	Etcd *EtcdSpec `yaml:"etcd"`
}

type EtcdSpec struct {
	Endpoints []string `yaml:"endpoints"`
}

func New() *Machine {
	return &Machine{}
}

func (m *Machine) Load(path string) error {
	m.Logger().With(
		zap.String("path", path),
	).Debug("loading machine manifest")

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(content, m); err != nil {
		return err
	}

	return nil
}

func (m *Machine) Init() error {
	m.Logger().With(
		zap.Strings("endpoints", m.Spec.Etcd.Endpoints),
	).Info("connecting to etcd")

	m.storage = etcd.New(m.Spec.Etcd.Endpoints)

	if err := m.storage.Init(); err != nil {
		return err
	}

	return nil
}

func (m *Machine) Storage() Storage {
	if m.storage == nil {
		m.Logger().Fatal("storage should be inilialized")
	}

	return m.storage
}

func (m *Machine) Run() error {
	m.Logger().Info("starting server")
	if err := m.Listen(); err != nil {
		return err
	}

	return nil
}

func (m *Machine) Node(name string) *async.Node {
	return async.NewNode(name, m.Storage())
}

func (m *Machine) NodeList() *async.NodeList {
	return async.NewNodeList(m.Storage())
}

func (m *Machine) Resource(kind, name string) *async.Resource {
	return async.NewResource(kind, name, m.Storage())
}

func (m *Machine) Logger() *zap.Logger {
	if m.logger == nil {
		m.logger = global.Logger.With(
			zap.String("entity", "machine"),
		)
	}

	return m.logger
}
