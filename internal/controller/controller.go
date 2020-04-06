package controller

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/internal/board/board"
	"github.com/luanguimaraesla/freegrow/internal/board/raspberry"
	"go.uber.org/zap"
)

type Board interface {
	RegisterDigitalDevice(board.DigitalDevice) board.DeviceID
	DigitalDevice(board.DeviceID) (board.DigitalDevice, error)
}

var (
	Controller Board
	logger     *zap.Logger
)

func DefineController(board string) error {
	logger.With(
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
