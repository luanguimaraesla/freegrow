package scheduler

import "github.com/robfig/cron/v3"

type Event struct {
	Schedule string
	Func     func()
	entryID  cron.EntryID
}

func NewEvent(schedule string, f func()) *Event {
	return &Event{
		Schedule: schedule,
		Func:     f,
	}
}
