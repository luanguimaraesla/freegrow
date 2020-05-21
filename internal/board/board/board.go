package board

import (
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

type DeviceID string

type DigitalBoardInterface interface {
	RegisterDigitalDevice(DigitalDevice) DeviceID
	DigitalDevice(DeviceID) (DigitalDevice, error)
}

type PortID interface {
	Uint8() uint8
}

type Port interface {
	Output()
	High()
	Low()
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
		log := global.Logger.With(
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
