package device

import (
        "sync"

        "github.com/sirupsen/logrus"
)

var (
        deviceMutex sync.Mutex
        currentDeviceId = 0

        log *logrus.Entry
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

func (d *device) getLogger() *logrus.Entry {
        return log.WithFields(logrus.Fields{
                "deviceId": d.id,
                "deviceName": d.name,
                "deviceKind": d.kind,
        })
}

func SetLogger (logger *logrus.Entry) {
        log = logger
}
