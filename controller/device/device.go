package device

import (
        "sync"
)

var (
        deviceMutex sync.Mutex
        currentDeviceId = 0
)

type Port int

type device struct {
        name string
        id int
        kind string
}

func getNewDeviceId() int {
        deviceMutex.Lock()
        defer func () {
                currentDeviceId += 1
                deviceMutex.Unlock()
        }()
        return currentDeviceId
}

func (d *device) GetName() string {
        return d.name
}

func (d *device) GetId() int {
        return d.id
}

func (d *device) GetKind() string {
        return d.kind
}
