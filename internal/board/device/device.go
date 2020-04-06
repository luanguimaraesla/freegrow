package device

import "go.uber.org/zap"

var logger *zap.Logger

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

func initLogger() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger = log
}

func init() {
	initLogger()
}
