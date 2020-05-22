package node

import (
	"fmt"

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
	Class  string    `yaml:"class"`
	Spec   yaml.Node `yaml:"spec"`
	Gadget Gadget
	logger *zap.Logger
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
