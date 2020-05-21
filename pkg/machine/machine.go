package machine

import (
	"fmt"
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

func (g *Gadgets) UnmarshalYAML(value *yaml.Node) error {
	var temp interface{}

	if err := value.Decode(temp); err != nil {
		return err
	}

	fmt.Printf("%v\n", value)

	*g = Gadgets{
		Items: []interface{}{},
	}

	return nil
}
