// Package relay is the RPi interface to control relays
package relay

import (
        "github.com/luanguimaraesla/freegrow/controller"
        "github.com/luanguimaraesla/freegrow/controller/device"
)

type Relay struct {
        Id int
        State *device.State
}

// NewRelay creates a new Relay device
func NewRelay (port int) (*Relay, error) {
        state := device.Off

        id, err := controller.RegisterDigitalDevice(port, &state)
        if err != nil {
                return nil, err
        }

        return &Relay{
                Id: id,
                State: &state,
        }, nil
}

// Activate function turns the Relay device on
func (r *Relay) Activate () error {
        return controller.Activate(r.Id)
}

// Deactivate function turns the Relay device off
func (r *Relay) Deactivate () error {
        return controller.Deactivate(r.Id)
}
