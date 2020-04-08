package irrigator

import (
	"time"

	"github.com/luanguimaraesla/freegrow/internal/gadget"
	"github.com/luanguimaraesla/freegrow/internal/system/relay"
	"go.uber.org/zap"
)

type Relay interface {
	Activate() error
	Deactivate() error
}

type Gadget interface {
	Logger() *zap.Logger
}

type Irrigator struct {
	Gadget
	relay         Relay
	operationTime time.Duration
}

func New(name, port string, operationTime time.Duration) (*Irrigator, error) {
	r, err := relay.NewRelay(port)
	if err != nil {
		return nil, err
	}

	return &Irrigator{
		Gadget:        gadget.New("irrigator", name),
		relay:         r,
		operationTime: operationTime,
	}, nil
}

func (i *Irrigator) Start() error {
	i.Logger().Info("starting")
	err := i.relay.Activate()
	if err != nil {
		i.Logger().Error("failed activating relay", zap.Error(err))
		return err
	}

	i.Logger().Debug("starting operation")
	time.Sleep(i.operationTime)

	err = i.relay.Deactivate()
	if err != nil {
		i.Logger().Error("failed deactivating relay", zap.Error(err))
		return err
	}
	i.Logger().Info("finishing irrigator")

	return nil
}
