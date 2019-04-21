package controller

import (
        "fmt"

        "github.com/sirupsen/logrus"

        "github.com/luanguimaraesla/freegrow/controller/device"
        "github.com/luanguimaraesla/freegrow/controller/raspberry"
)

type Controller interface {
        RegisterDigitalDevice(*device.DigitalDevice) (int, error)
        ChangeState(int, string) error
        GetDigitalDeviceState(int) (device.DigitalDeviceState, error)
}

var (
        boardController Controller
        log *logrus.Entry
)

func SetLogger(logger *logrus.Entry){
        log = logger
}

func StartController (board string) error {
        var err error
        log.Info(fmt.Sprintf("configuring (%s) controller", board))

        switch board {
        case "raspberry":
                raspberry.SetLogger(log)
                boardController, err = raspberry.NewRaspberry()
                return err
        default:
                return fmt.Errorf("board not supported: %s", board)
        }
}

func RegisterDigitalDevice(d *device.DigitalDevice) (int, error) {
        return boardController.RegisterDigitalDevice(d)
}

func ChangeState (id int, stateName string) error {
        return boardController.ChangeState(id, stateName)
}

func GetDigitalDeviceState(id int) (device.DigitalDeviceState, error) {
        return boardController.GetDigitalDeviceState(id)
}
