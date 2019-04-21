package device

import (
        "sync"
)

var (
        deviceMutex sync.Mutex
        currentDeviceId = 0

        On = State{"on"}
        Off = State{"off"}
        StandBy = State{"standby"}
        Damaged = State{"damaged"}
)

type State struct {
        state string
}

type Device struct {
        Id int
        State *State
        Kind string
        Ports map[string]int
}

func NewDevice(kind string, state *State, ports map[string]int) *Device {
        return &Device{
                Id: getNewDeviceId(),
                Kind: kind,
                State: state,
                Ports: ports,
        }
}

func getNewDeviceId() int {
        deviceMutex.Lock()
        defer func () {
                currentDeviceId += 1
                deviceMutex.Unlock()
        }()
        return currentDeviceId
}
