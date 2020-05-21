package device

import "go.uber.org/zap"

type PortID uint8

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

func (pid PortID) Uint8() uint8 {
	return uint8(pid)
}
