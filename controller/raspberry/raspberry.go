package raspberry

import (
        "fmt"

        "github.com/sirupsen/logrus"

        "github.com/luanguimaraesla/freegrow/device"
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
        supportedModels = []string{"default"}

        log *logrus.Entry
)

func (r *Raspberry) getLogger() *logrus.Entry {
        return log.WithFields(logrus.Fields{
                "board": "raspberry",
                "boardModel": r.model,
        })
}

func SetLogger(logger *logrus.Entry) {
        log = logger
}

func NewRaspberry () (*Raspberry, error) {
        model := "default" // This should be a configuration

        if !validateRaspberry(model) {
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
        l := r.getLogger()
        l.Debug(fmt.Sprintf("registering new device: (%d) %s", d.GetId(), d.GetName()))

        if err := r.checkDigitalPorts(d); err != nil {
                return -1, err
        }

        r.devices[d.GetId()] = d
        l.Debug(fmt.Sprintf("device registered: (%d) %s", d.GetId(), d.GetName()))

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

func (r *Raspberry) checkDigitalPorts(d *device.DigitalDevice) error {
        l := r.getLogger()
        l.Debug(fmt.Sprintf("checking digital ports availability for device: (%d) %s", d.GetId(), d.GetName()))
        // [TODO] implement port check
        l.Debug(fmt.Sprintf("all ports are available for device: (%d) %s", d.GetId(), d.GetName()))
        return nil
}
