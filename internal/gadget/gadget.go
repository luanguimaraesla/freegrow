package gadget

import (
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

type GadgetID string

type Gadget struct {
	logger *zap.Logger
	name   string
	class  string
	id     GadgetID
}

func NewGadgetID() GadgetID {
	return GadgetID("gadget-id-123456")
}

func (gid GadgetID) String() string {
	return string(gid)
}

func New(class, name string) *Gadget {
	return &Gadget{
		class: class,
		name:  name,
		id:    NewGadgetID(),
	}
}

func (g *Gadget) Logger() *zap.Logger {
	if g.logger == nil {
		log := global.Logger.With(
			zap.String("name", g.name),
			zap.String("id", g.id.String()),
			zap.String("class", g.class),
		)

		g.logger = log
	}

	return g.logger
}
