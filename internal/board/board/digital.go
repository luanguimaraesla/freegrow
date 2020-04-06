package board

import (
	"fmt"

	"go.uber.org/zap"
)

type DigitalDevice interface {
	Alias() string
	States() DigitalDeviceStateSet
}

type DigitalDeviceStateSet interface {
	Get(string) (DigitalDeviceState, error)
	Append(...DigitalDeviceState) error
}

type DigitalDeviceState interface {
	Alias() string
	Activate()
}

type DigitalDeviceMap map[DeviceID]DigitalDevice

type DigitalBoard struct {
	DigitalDevices DigitalDeviceMap
	logger         *zap.Logger
}

func NewDigitalBoard() *DigitalBoard {
	return &DigitalBoard{
		DigitalDevices: DigitalDeviceMap{},
	}
}

func (db *DigitalBoard) RegisterDigitalDevice(dev DigitalDevice) DeviceID {
	l := db.Logger().With(
		zap.String("deviceAlias", dev.Alias()),
	)

	l.Debug("registering device")

	id := NewUniqueDeviceID()

	db.DigitalDevices[id] = dev

	l.With(
		zap.String("deviceID", id.String()),
	).Debug("device registered")

	return id
}

func (db *DigitalBoard) DigitalDevice(id DeviceID) (DigitalDevice, error) {
	dev, ok := db.DigitalDevices[id]
	if !ok {
		return nil, fmt.Errorf("device not found")
	}

	return dev, nil
}

func (db *DigitalBoard) Logger() *zap.Logger {
	if db.logger == nil {
		log := logger.With(
			zap.String("mode", "digital"),
		)

		db.logger = log
	}

	return db.logger
}
