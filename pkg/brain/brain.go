package brain

import (
	"github.com/luanguimaraesla/freegrow/internal/global"
)

type Brain struct {
	*global.Logger
}

func New() *Brain {
	return &Brain{global.NewLogger()}
}

func (b *Brain) Init() error {
	b.L.Info("checking connection with database")

	return nil
}
