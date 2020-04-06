package device

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/internal/board/board"
	"go.uber.org/zap"
)

type digitalPortState string

const (
	DigitalPortStateHigh digitalPortState = "HIGH"
	DigitalPortStateLow  digitalPortState = "LOW"
)

type DigitalPort struct {
	id     PortID
	State  digitalPortState
	logger *zap.Logger
}

type DigitalPortSet struct {
	ports []*DigitalPort
}

type DigitalDeviceState struct {
	alias   string
	portSet *DigitalPortSet
}

type DigitalDeviceStateSet struct {
	states []board.DigitalDeviceState
}

type DigitalDevice struct {
	*Device
	stateSet *DigitalDeviceStateSet
}

func (dps digitalPortState) String() string {
	return string(dps)
}

func NewDigitalPort(id string, state digitalPortState) *DigitalPort {
	return &DigitalPort{
		id:    PortID(id),
		State: state,
	}
}

func NewDigitalDeviceState(alias string) *DigitalDeviceState {
	return &DigitalDeviceState{
		alias:   alias,
		portSet: &DigitalPortSet{},
	}
}

func (dps *DigitalPortSet) Append(ports ...*DigitalPort) error {
	for _, port := range ports {
		for _, registeredPort := range dps.ports {
			if port.id == registeredPort.id {
				return fmt.Errorf("port %s is duplicated", port.id.String())
			}
		}

		dps.ports = append(dps.ports, port)
	}

	return nil
}

func (ddss *DigitalDeviceStateSet) Append(states ...board.DigitalDeviceState) error {
	for _, state := range states {
		for _, registeredState := range ddss.states {
			if state.Alias() == registeredState.Alias() {
				return fmt.Errorf("state alias %s is duplicated", state.Alias())
			}
		}

		ddss.states = append(ddss.states, state)
	}

	return nil
}

func (ddss *DigitalDeviceStateSet) Get(alias string) (board.DigitalDeviceState, error) {
	var dds board.DigitalDeviceState

	for _, registeredState := range ddss.states {
		if alias == registeredState.Alias() {
			dds = registeredState
		}
	}

	if dds == nil {
		return nil, fmt.Errorf("state alias %s not found", alias)
	}

	return dds, nil
}

func NewDigitalDevice(alias string) *DigitalDevice {
	return &DigitalDevice{
		Device:   NewDevice(alias),
		stateSet: &DigitalDeviceStateSet{},
	}
}

func (d *DigitalDevice) States() board.DigitalDeviceStateSet {
	return d.stateSet
}

func (dds *DigitalDeviceState) Activate() {
	dds.portSet.Activate()
}

func (dds *DigitalDeviceState) Alias() string {
	return dds.alias
}

func (dds *DigitalDeviceState) Ports() *DigitalPortSet {
	return dds.portSet
}

func (dps *DigitalPortSet) Activate() {
	for _, port := range dps.ports {
		port.Activate()
	}
}

func (dp *DigitalPort) Activate() {
	dp.Logger().Debug("activating")
}

func (dp *DigitalPort) Logger() *zap.Logger {
	if dp.logger == nil {
		log := logger.With(
			zap.String("entity", "DigitalPort"),
			zap.String("id", dp.id.String()),
			zap.String("state", dp.State.String()),
		)

		dp.logger = log
	}

	return dp.logger
}

func (ddss *DigitalDeviceStateSet) Logger() *zap.Logger {
	return logger.With(
		zap.String("entity", "DigitalDeviceStateSet"),
	)
}
