package irrigator

import (
	"time"

	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/internal/system/relay"
	"go.uber.org/zap"
)

type Relay interface {
	Activate() error
	Deactivate() error
}

type Irrigator struct {
	Kind     string             `yaml:"kind" json:"kind"`
	Metadata *resource.Metadata `yaml:"metadata" json:"metadata"`
	Spec     *IrrigatorSpec     `yaml:"spec" json:"spec"`
	relay    Relay
	logger   *zap.Logger
}

type IrrigatorSpec struct {
	Enabled bool                  `yaml:"enabled" json:"enabled"`
	Port    uint8                 `yaml:"port" json:"port"`
	States  []*IrrigatorStateSpec `yaml:"states" json:"states"`
}

type IrrigatorStateSpec struct {
	Name     string `yaml:"name" json:"name"`
	Schedule string `yaml:"schedule" json:"schedule"`
}

func New() *Irrigator {
	return &Irrigator{}
}

func (i *Irrigator) Init() error {
	i.Logger().Debug("initializing relay")
	r, err := relay.New(i.Spec.Port)
	if err != nil {
		return err
	}

	i.relay = r

	return nil
}

func (i *Irrigator) Run() error {
	i.Logger().Info("starting")

	for {
		err := i.relay.Activate()
		if err != nil {
			i.Logger().Error("failed activating relay", zap.Error(err))
			return err
		}

		time.Sleep(time.Second * 2)

		err = i.relay.Deactivate()
		if err != nil {
			i.Logger().Error("failed deactivating relay", zap.Error(err))
			return err
		}
		i.Logger().Info("finishing irrigator")

		time.Sleep(time.Second * 2)
	}
}

func (i *Irrigator) Name() string {
	return i.Metadata.Name
}

func (i *Irrigator) Logger() *zap.Logger {
	if i.logger == nil {
		log := global.Logger.With(
			zap.String("name", i.Metadata.Name),
			zap.String("entity", i.Kind),
		)

		i.logger = log
	}

	return i.logger
}
