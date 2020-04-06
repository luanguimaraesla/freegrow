package gadget

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

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
		log := logger.With(
			zap.String("name", g.name),
			zap.String("id", g.id.String()),
			zap.String("class", g.class),
		)

		g.logger = log
	}

	return g.logger
}

func initLogger() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger = log
}

func init() {
	initLogger()
}
