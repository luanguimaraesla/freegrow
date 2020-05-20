package raspberry

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/internal/board/board"
	rpio "github.com/stianeikeland/go-rpio/v4"
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

func New() (*Raspberry, error) {
	b := board.New("raspberry", "default")

	rpi := &Raspberry{
		Board:  b,
		logger: b.Logger(),
	}

	rpi.Logger().Info("openning and mapping memory to access gpio")
	if err := rpio.Open(); err != nil {
		return nil, fmt.Errorf("failed initializing gpio")
	}

	return rpi, nil
}

// Close unmaps gpio memory
func (rpi *Raspberry) Close() {
	rpio.Close()
}
