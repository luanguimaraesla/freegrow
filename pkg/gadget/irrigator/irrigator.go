package irrigator

import (
	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/internal/system/relay"
	"github.com/luanguimaraesla/freegrow/pkg/scheduler"
	"go.uber.org/zap"
)

type IrrigatorList struct {
	Resources []*Irrigator `yaml:"resources" json:"resources"`
}

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

func NewIrrigatorList() *IrrigatorList {
	return &IrrigatorList{}
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

func (i *Irrigator) On() {
	err := i.relay.Activate()
	if err != nil {
		i.Logger().Error("failed activating relay", zap.Error(err))
	}
}

func (i *Irrigator) Off() {
	err := i.relay.Deactivate()
	if err != nil {
		i.Logger().Error("failed deactivating relay", zap.Error(err))
	}
}

func (i *Irrigator) Events() []*scheduler.Event {
	events := []*scheduler.Event{}

	for _, state := range i.Spec.States {
		switch state.Name {
		case "on":
			events = append(events, scheduler.NewEvent(state.Schedule, i.On))
		case "off":
			events = append(events, scheduler.NewEvent(state.Schedule, i.Off))
		default:
			i.Logger().Error("state not found", zap.String("state", state.Name))
		}
	}

	return events
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
