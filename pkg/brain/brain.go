package brain

import (
	"github.com/luanguimaraesla/freegrow/internal/log"
)

type Brain struct {
	*log.Logger
}

func New() *Brain {
	return &Brain{log.NewLogger()}
}
