package controller

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/internal/board/board"
	"github.com/luanguimaraesla/freegrow/internal/board/raspberry"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

type Board interface {
	RegisterDigitalDevice(board.DigitalDevice) board.DeviceID
	DigitalDevice(board.DeviceID) (board.DigitalDevice, error)
}

var (
	Controller Board
)

func DefineController(board string) error {
	global.Logger.With(
		zap.String("board", board),
	).Info("configuring global board controller")

	switch board {
	case "raspberry":
		Controller = raspberry.New()
	default:
		return fmt.Errorf("board not supported: %s", board)
	}

	return nil
}
