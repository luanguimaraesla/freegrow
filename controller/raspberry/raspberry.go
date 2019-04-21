package raspberry

import (
        "fmt"

        "github.com/luanguimaraesla/freegrow/controller/device"
)

type Raspberry struct {
        model string
        devices map[int]*device.Device
}

func NewRaspberry () *Raspberry {
        return &Raspberry{
                model: "default",
                devices: make(map[int]*device.Device),
        }
}

func (r *Raspberry) RegisterDigitalDevice(port int, state *device.State) (int, error) {
        ports := map[string]int{
                "common": port,
        }
        newDevice := device.NewDevice("digital", state, ports)
        r.devices[newDevice.Id] = newDevice

        return newDevice.Id, nil
}

func (r *Raspberry) Activate (deviceId int) error {
        d := r.devices[deviceId]

        switch kind := d.Kind; kind {
        case "digital":
                //[TODO] Do some stuff to activate the device
                *(d.State) = device.On
                return nil
        default:
                return fmt.Errorf("unknown device kind (%s)")
        }
}

func (r *Raspberry) Deactivate (deviceId int) error {
        d := r.devices[deviceId]

        switch kind := d.Kind; kind {
        case "digital":
                //[TODO] Do some stuff to deactivate the device
                *(d.State) = device.On
                return nil
        default:
                return fmt.Errorf("unknown device kind (%s)")
        }
}
