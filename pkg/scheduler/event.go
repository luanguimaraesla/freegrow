package scheduler

import (
	"reflect"
	"time"

	"github.com/robfig/cron/v3"
)

type Owner interface {
	Name() string
	Class() string
}

type Event struct {
	Owner     Owner
	Schedule  string
	Func      func()
	updatedAt time.Time
	entryID   cron.EntryID
}

func NewEvent(owner Owner, schedule string, f func()) *Event {
	return &Event{
		Owner:     owner,
		Schedule:  schedule,
		Func:      f,
		updatedAt: time.Now(),
	}
}

func (e *Event) IsEqual(other *Event) bool {
	return other.Schedule == e.Schedule &&
		other.Owner.Name() == e.Owner.Name() &&
		other.Owner.Class() == e.Owner.Class() &&
		reflect.ValueOf(other.Func).Pointer() == reflect.ValueOf(e.Func).Pointer()
}

func (e *Event) Update() {
	e.updatedAt = time.Now()
}
