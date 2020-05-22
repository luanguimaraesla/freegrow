package machine

import (
	"fmt"
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/controller"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/pkg/gadget/irrigator"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Gadget interface {
	Init() error
	Run() error
}

type Runner struct {
	logger *zap.Logger
	Class  string    `yaml:"class"`
	Spec   yaml.Node `yaml:"spec"`
	Gadget Gadget
}

type Machine struct {
	logger  *zap.Logger
	Board   string    `yaml:"board"`
	Runners []*Runner `yaml:"gadgets"`
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
	m.Logger().Info("defining controller")
	if err := controller.DefineController(m.Board); err != nil {
		return err
	}

	m.Logger().Info("initializing runners")
	for _, r := range m.Runners {
		if err := r.Init(); err != nil {
			return err
		}
	}

	return nil
}

func (m *Machine) Run() error {
	for _, r := range m.Runners {
		if err := r.Run(); err != nil {
			return err
		}
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

func (r *Runner) Init() error {
	l := r.Logger().With(
		zap.String("class", r.Class),
	)

	l.Info("initializing a new runner")

	switch class := r.Class; class {
	case "irrigator":
		gadget := irrigator.New()

		l.Debug("decoding")
		if err := r.Spec.Decode(gadget); err != nil {
			return err
		}

		r.Gadget = gadget
	default:
		return fmt.Errorf("no runner found for class %s", class)
	}

	l.Info("running gadget")
	if err := r.Gadget.Init(); err != nil {
		return err
	}

	return nil
}

func (r *Runner) Run() error {
	if r.Gadget == nil {
		return fmt.Errorf("gadget was not initialized")
	}

	if err := r.Gadget.Run(); err != nil {
		return err
	}

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
