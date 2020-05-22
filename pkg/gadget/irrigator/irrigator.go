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

type Irrigator struct {
	gadget.Gadget    `yaml:",inline"`
	gadget.Scheduler `yaml:",inline"`
	Port             uint8 `yaml:"port"`
	relay            Relay
}

func New() *Irrigator {
	return &Irrigator{}
}

func Load(name string, port uint8, operationTime time.Duration) (*Irrigator, error) {
	return nil, nil
}

func (i *Irrigator) Init() error {
	i.Logger().Debug("initializing relay")
	r, err := relay.New(i.Port)
	if err != nil {
		return err
	}

	i.relay = r

	return nil
}

func (i *Irrigator) Run() error {
	i.Logger().Info("starting")
	err := i.relay.Activate()
	if err != nil {
		i.Logger().Error("failed activating relay", zap.Error(err))
		return err
	}

	err = i.relay.Deactivate()
	if err != nil {
		i.Logger().Error("failed deactivating relay", zap.Error(err))
		return err
	}
	i.Logger().Info("finishing irrigator")

	return nil
}
