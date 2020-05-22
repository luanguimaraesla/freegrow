package machine

import (
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
}

type Machine struct {
	resource.Meta `yaml:"metadata,inline"`
	Spec          *MachineSpec `yaml:"spec"`
	logger        *zap.Logger
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

	return nil
}

func (m *Machine) Run() error {
	m.Logger().Info("starting server")
	if err := m.Listen(); err != nil {
		return err
	}

	return nil
}

func (m *Machine) Logger() *zap.Logger {
	if m.logger == nil {
		m.logger = global.Logger.With(
			zap.String("entity", "machine"),
		)
	}

	return m.logger
}
