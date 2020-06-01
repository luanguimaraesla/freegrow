package scheduler

import (
	"fmt"
	"time"

	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Scheduler struct {
	cron   *cron.Cron
	events []*Event
	logger *zap.Logger
}

func New() *Scheduler {
	return &Scheduler{
		cron: cron.New(),
	}
}

func (s *Scheduler) Add(e *Event) error {
	if event, ok := s.eventExists(e); ok {
		s.Logger().Debug("updating event which is already registered")
		event.Update()
		return nil
	}

	s.Logger().Debug("registering a new crontab function")
	id, err := s.cron.AddFunc(e.Schedule, e.Func)
	if err != nil {
		return err
	}

	e.entryID = id
	s.events = append(s.events, e)

	return nil
}

func (s *Scheduler) eventExists(e *Event) (*Event, bool) {
	for i, event := range s.events {
		if e.IsEqual(event) {
			return s.events[i], true
		}
	}

	return nil, false
}

func (s *Scheduler) Refresh() {
	s.Logger().Debug("refreshing events", zap.Int("total", len(s.events)))
	updated := []*Event{}

	for i, event := range s.events {
		if time.Now().Sub(event.updatedAt) < 30*time.Second {
			updated = append(updated, s.events[i])
		} else {
			s.Logger().Debug("removing crontab", zap.Int("entryID", int(event.entryID)))
			s.cron.Remove(event.entryID)
		}
	}

	s.Logger().Debug(
		"events updated",
		zap.Int("remaining", len(updated)),
		zap.Int("deleted", len(s.events)-len(updated)),
	)

	s.events = updated
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

func (s *Scheduler) Logger() *zap.Logger {
	if s.logger == nil {
		s.logger = global.Logger.With(
			zap.String("entity", "scheduler"),
		)
	}

	return s.logger
}
