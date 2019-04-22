package device

import (
        "fmt"
)

type DigitalState string

const (
        DigitalStateHigh DigitalState = "HIGH"
        DigitalStateLow  DigitalState = "LOW"
)

type DigitalDeviceState map[Port]DigitalState

type DigitalStateMap map[string]DigitalDeviceState

type DigitalDevice struct {
        device
        CurrentStateName string
        states DigitalStateMap
}

func NewDigitalDevice (name string, states DigitalStateMap) *DigitalDevice {
        return &DigitalDevice{
                device: device{name, getNewDeviceId(), "digital"},
                states: states,
                CurrentStateName: "",
        }
}

func (d *DigitalDevice) ChangeState (stateName string) error {
        l := d.getLogger()

        if d.CurrentStateName == stateName {
                l.Warn(fmt.Sprintf("state (%s) is already the current state"))
                return nil
        }

        l.Debug(fmt.Sprintf("changing device state from (%s) to (%s)", d.CurrentStateName, stateName))
        if _, ok := d.states[stateName]; ok {
                d.CurrentStateName = stateName
                return nil
        }
        return fmt.Errorf("state not found")
}

func (d *DigitalDevice) GetCurrentState () (DigitalDeviceState, error) {
        if d.CurrentStateName == "" {
                return nil, fmt.Errorf("current state not found")
        }
        return d.states[d.CurrentStateName], nil
}
