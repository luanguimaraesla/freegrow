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

	var err error = nil

	switch board {
	case "raspberry":
		Controller, err = raspberry.New()
		return err
	default:
		return fmt.Errorf("board not supported: %s", board)
	}
}
