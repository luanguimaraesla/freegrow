package fakeboard

import (
	"github.com/luanguimaraesla/freegrow/internal/board/board"
	"go.uber.org/zap"
)

type Board interface {
	Logger() *zap.Logger
	RegisterDigitalDevice(board.DigitalDevice) board.DeviceID
	DigitalDevice(board.DeviceID) (board.DigitalDevice, error)
}

type FakeBoard struct {
	Board
	logger *zap.Logger
}

func New() (*FakeBoard, error) {
	b := board.New("fake", "default")

	return &FakeBoard{
		Board:  b,
		logger: b.Logger(),
	}, nil
}

// Close unmaps gpio memory
func (f *FakeBoard) Close() {
	f.Logger().Info("closing")
}

// Pin returs a pin interface for a specific number
func (f *FakeBoard) Pin(port board.PortID) board.Port {
	return &FakePin{
		pin: port.Uint8(),
	}
}
