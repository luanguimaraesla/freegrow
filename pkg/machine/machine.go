package machine

import (
	"context"
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/pkg/async"
	"github.com/luanguimaraesla/freegrow/pkg/node"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Node interface {
	Put(context.Context, *node.Node) error
}

type Machine struct {
	Metadata *resource.Metadata `yaml:"metadata"`
	Spec     *MachineSpec       `yaml:"spec"`
	logger   *zap.Logger
}

type MachineSpec struct {
	Bind string `yaml:"bind"`
	Etcd *Etcd  `yaml:"etcd"`
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

	if err := m.Spec.Etcd.Init(); err != nil {
		return err
	}

	return nil
}

func (m *Machine) Run() error {
	m.Logger().Info("starting server")
	if err := m.Listen(); err != nil {
		return err
	}

	return nil
}

func (m *Machine) Node(name string) Node {
	return async.NewAsyncNode(name, m.Spec.Etcd)
}

func (m *Machine) Logger() *zap.Logger {
	if m.logger == nil {
		m.logger = global.Logger.With(
			zap.String("entity", "machine"),
		)
	}

	return m.logger
}
