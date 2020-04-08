// Package relay is the RPi interface to control relays
package relay

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/internal/board/board"
	"github.com/luanguimaraesla/freegrow/internal/board/device"
	"github.com/luanguimaraesla/freegrow/internal/controller"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

type Relay struct {
	id     board.DeviceID
	logger *zap.Logger
}

func RelayName(port string) string {
	return fmt.Sprintf("relay_%s", port)
}

// NewRelay creates a new Relay device
func NewRelay(port string) (*Relay, error) {
	powerOnState := device.NewDigitalDeviceState("on")
	if err := powerOnState.Ports().Append(
		device.NewDigitalPort(port, device.DigitalPortStateHigh),
	); err != nil {
		return nil, err
	}

	powerOffState := device.NewDigitalDeviceState("off")
	if err := powerOffState.Ports().Append(
		device.NewDigitalPort(port, device.DigitalPortStateLow),
	); err != nil {
		return nil, err
	}

	d := device.NewDigitalDevice(RelayName(port))

	if err := d.States().Append(
		powerOnState,
		powerOffState,
	); err != nil {
		return nil, err
	}

	id := controller.Controller.RegisterDigitalDevice(d)

	return &Relay{
		id: id,
	}, nil
}

// Activate function turns the Relay device on
func (r *Relay) Activate() error {
	r.Logger().Debug("activating")
	dev, err := controller.Controller.DigitalDevice(r.id)
	if err != nil {
		return err
	}

	state, err := dev.States().Get("on")
	if err != nil {
		return err
	}

	state.Activate()

	return nil
}

// Deactivate function turns the Relay device off
func (r *Relay) Deactivate() error {
	r.Logger().Debug("deactivating")
	dev, err := controller.Controller.DigitalDevice(r.id)
	if err != nil {
		return err
	}

	state, err := dev.States().Get("off")
	if err != nil {
		return err
	}

	state.Activate()

	return nil
}

func (r *Relay) Logger() *zap.Logger {
	if r.logger == nil {
		log := global.Logger.With(
			zap.String("entity", "relay"),
			zap.String("id", r.id.String()),
		)

		r.logger = log
	}

	return r.logger
}
