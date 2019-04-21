// Package relay is the RPi interface to control relays
package relay

import (
        "github.com/luanguimaraesla/freegrow/controller"
        "github.com/luanguimaraesla/freegrow/controller/device"
)

type Relay struct {
        Id int
}

// NewRelay creates a new Relay device
func NewRelay (name string, port int) (*Relay, error) {
        d := device.NewDigitalDevice(
                name,
                device.DigitalStateMap{
                        "on": device.DigitalDeviceState{
                                device.Port(port): device.DigitalStateHigh,
                        },
                        "off": device.DigitalDeviceState{
                                device.Port(port): device.DigitalStateLow,
                        },
                },
        )

        id, err := controller.RegisterDigitalDevice(d)
        if err != nil {
                return nil, err
        }

        return &Relay{
                Id: id,
        }, nil
}

// Activate function turns the Relay device on
func (r *Relay) Activate () error {
        return controller.ChangeState(r.Id, "on")
}

// Deactivate function turns the Relay device off
func (r *Relay) Deactivate () error {
        return controller.ChangeState(r.Id, "off")
}

// GetState informs the state of the relay
func (r *Relay) GetState () (device.DigitalDeviceState, error) {
        return controller.GetDigitalDeviceState(r.Id)
}
