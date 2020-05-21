package machine

import (
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Gadgets struct {
	Items []interface{}
}

type Machine struct {
	logger  *zap.Logger
	Board   string   `mapstructure:"board"`
	Gadgets *Gadgets `mapstructure:"gadgets"`
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

func (m *Machine) Logger() *zap.Logger {
	if m.logger == nil {
		m.logger = global.Logger.With(
			zap.String("entity", "machine"),
		)
	}

	return m.logger
}

type A struct{}

func (g *Gadgets) UnmarshalYAML(value *yaml.Node) error {
	*g = Gadgets{
		Items: []interface{}{},
	}

	return nil
}
