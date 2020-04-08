package device

import "go.uber.org/zap"

type PortID string

type Device struct {
	alias  string
	logger *zap.Logger
}

func NewDevice(alias string) *Device {
	return &Device{
		alias: alias,
	}
}

func (d *Device) Alias() string {
	return d.alias
}

func (pid PortID) String() string {
	return string(pid)
}
