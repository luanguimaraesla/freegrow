package raspberry

import (
        "fmt"

        "github.com/sirupsen/logrus"

        "github.com/luanguimaraesla/freegrow/controller/device"
)

type digitalDevice interface {
       ChangeState(string) error
       GetCurrentState() (device.DigitalDeviceState, error)
       GetKind() string
       GetName() string
       GetId()  int
}

type Raspberry struct {
        model string
        devices map[int]digitalDevice
}

var (
        log *logrus.Entry
        supportedModels = []string{"default"}
)

func SetLogger(logger *logrus.Entry) {
        log = logger.WithFields(logrus.Fields{
                "board": "raspberry",
        })
}


func NewRaspberry () (*Raspberry, error) {
        model := "default" // This should be a configuration

        log = log.WithFields(logrus.Fields{
                "model": model,
        })
        device.SetLogger(log)

        if !validateRaspberry(model) {
                log.Error("unsupported model")
                return nil, fmt.Errorf("invalid board")
        }

        return &Raspberry{
                model: model,
                devices: make(map[int]digitalDevice),
        }, nil
}

func validateRaspberry(model string) bool {
        for _, m := range supportedModels {
                if m == model {
                        return true
                }
        }
        return false
}

func (r *Raspberry) RegisterDigitalDevice(d *device.DigitalDevice) (int, error) {
        r.devices[d.GetId()] = d
        // [TODO] check ports
        return d.GetId(), nil
}

func (r *Raspberry) ChangeState (deviceId int, stateName string) error {
        d := r.devices[deviceId]

        switch kind := d.GetKind(); kind {
        case "digital":
                //[TODO] Do some stuff to activate the device
                d.ChangeState(stateName)

                return nil
        default:
                return fmt.Errorf("unknown device kind (%s)", kind)
        }
}

func (r *Raspberry) GetDigitalDeviceState (deviceId int) (device.DigitalDeviceState, error) {
        return r.devices[deviceId].GetCurrentState()
}
