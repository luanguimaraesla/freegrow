package controller

import (
        "fmt"

        "github.com/luanguimaraesla/freegrow/controller/device"
        "github.com/luanguimaraesla/freegrow/controller/raspberry"
)

type Controller interface {
        RegisterDigitalDevice(int, *device.State) (int, error)
        Activate(int) error
        Deactivate(int) error
}

var (
        boardController Controller
)

func StartController (board string) error {
        switch board {
        case "raspberry":
                boardController = raspberry.NewRaspberry()
                return nil
        default:
                return fmt.Errorf("error configuring new board: %s", board)
        }
}

func RegisterDigitalDevice(port int, state *device.State) (int, error) {
        return boardController.RegisterDigitalDevice(port, state)
}

func Activate(id int) error {
        return boardController.Activate(id)
}

func Deactivate(id int) error {
        return boardController.Deactivate(id)
}
