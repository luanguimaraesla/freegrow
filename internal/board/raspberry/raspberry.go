package raspberry

import (
	"github.com/luanguimaraesla/freegrow/internal/board/board"
	"go.uber.org/zap"
)

type Board interface {
	Logger() *zap.Logger
	RegisterDigitalDevice(board.DigitalDevice) board.DeviceID
	DigitalDevice(board.DeviceID) (board.DigitalDevice, error)
}

type Raspberry struct {
	Board
	logger *zap.Logger
}

func New() *Raspberry {
	b := board.New("raspberry", "default")

	return &Raspberry{
		Board:  b,
		logger: b.Logger(),
	}
}
