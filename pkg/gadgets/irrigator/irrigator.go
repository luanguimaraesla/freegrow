package irrigator

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/luanguimaraesla/freegrow/internal/system/relay"
	"github.com/luanguimaraesla/freegrow/pkg/gadgets"
)

var (
	log *logrus.Entry
)

type gate interface {
	Activate() error
	Deactivate() error
}

type Irrigator struct {
	Name          string
	gate          gate
	operationTime time.Duration
}

func New(name string, port int, operationTime time.Duration) (*Irrigator, error) {
	r, err := relay.NewRelay(fmt.Sprintf("%s_relay_%d", name, port), port)
	if err != nil {
		return nil, err
	}

	return &Irrigator{
		Name:          name,
		gate:          r,
		operationTime: operationTime,
	}, nil
}

func (i *Irrigator) Start() error {
	l := i.getLogger()
	l.Info("starting irrigator")
	err := i.gate.Activate()
	if err != nil {
		l.WithError(err).Error("failed activating gate relay")
		return err
	}

	l.Debug(fmt.Sprintf("will operate for %v seconds", i.operationTime))
	time.Sleep(i.operationTime)

	err = i.gate.Deactivate()
	if err != nil {
		l.WithError(err).Error("failed deactivating gate relay")
		return err
	}
	l.Info("finishing irrigator")

	return nil
}

func (i *Irrigator) getLogger() (logger *logrus.Entry) {
	return gadgets.GetLogger().WithFields(logrus.Fields{
		"gadget":        "irrigator",
		"irrigatorName": i.Name,
	})
}
