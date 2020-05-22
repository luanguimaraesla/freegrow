package machine

import (
	"fmt"
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/pkg/gadget/irrigator"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Gadget interface {
	Run() error
}

type Runner struct {
	logger *zap.Logger
	Class  string     `mapstructure:"class"`
	Spec   *yaml.Node `mapstructure:"spec"`
	Gadget Gadget
}

type Machine struct {
	logger  *zap.Logger
	Board   string    `mapstructure:"board"`
	Runners []*Runner `mapstructure:"gadgets"`
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

func (r *Runner) Run() error {
	switch class := r.Class; class {
	case "irrigator":
		gadget := irrigator.New()

		if err := r.Spec.Decode(&gadget); err != nil {
			return err
		}

		r.Gadget = gadget
	default:
		return fmt.Errorf("no runner found for class %s", class)
	}

	r.Logger().Info("running gadget")
	r.Gadget.Run()

	return nil
}

func (r *Runner) Logger() *zap.Logger {
	if r.logger == nil {
		r.logger = global.Logger.With(
			zap.String("entity", "runner"),
		)
	}

	return r.logger
}
