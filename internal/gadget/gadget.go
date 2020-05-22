package gadget

import (
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

type GadgetID string

type Gadget struct {
	logger  *zap.Logger
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled"`
}

func (g *Gadget) Logger() *zap.Logger {
	if g.logger == nil {
		log := global.Logger.With(
			zap.String("name", g.Name),
		)

		g.logger = log
	}

	return g.logger
}
