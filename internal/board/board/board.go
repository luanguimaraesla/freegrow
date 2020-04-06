package board

import (
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

type DeviceID string

type DigitalBoardInterface interface {
	RegisterDigitalDevice(DigitalDevice) DeviceID
	DigitalDevice(DeviceID) (DigitalDevice, error)
}

type Board struct {
	DigitalBoardInterface
	name   string
	model  string
	logger *zap.Logger
}

func New(board, model string) *Board {
	return &Board{
		DigitalBoardInterface: NewDigitalBoard(),
		name:                  board,
		model:                 model,
	}
}

func (id DeviceID) String() string {
	return string(id)
}

func (b *Board) Logger() *zap.Logger {
	if b.logger == nil {
		log := logger.With(
			zap.String("board", b.name),
			zap.String("model", b.model),
		)

		b.logger = log
	}

	return b.logger
}

func NewUniqueDeviceID() DeviceID {
	return DeviceID("12345")
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
