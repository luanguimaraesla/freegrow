package scheduler

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron   *cron.Cron
	events []*Event
}

func New() *Scheduler {
	return &Scheduler{
		cron: cron.New(),
	}
}

func (s *Scheduler) Add(e *Event) error {
	id, err := s.cron.AddFunc(e.Schedule, e.Func)
	if err != nil {
		return err
	}

	e.entryID = id
	s.events = append(s.events, e)

	return nil
}

func (s *Scheduler) Start() error {
	if s.cron == nil {
		return fmt.Errorf("cron controller not initialized")
	}

	s.cron.Start()

	return nil
}

func (s *Scheduler) Stop() error {
	if s.cron == nil {
		return fmt.Errorf("cron controller not initialized")
	}

	ctx := s.cron.Stop()
	ctx.Done()

	return nil
}
