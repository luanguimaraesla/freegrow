package fakeboard

import (
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

type FakePin struct {
	pin    uint8
	logger *zap.Logger
}

func (fp *FakePin) High() {
	fp.Logger().With(
		zap.String("state", "HIGH"),
	).Info("changing state")
}

func (fp *FakePin) Low() {
	fp.Logger().With(
		zap.String("state", "LOW"),
	).Info("changing state")
}

func (fp *FakePin) Output() {
	fp.Logger().Info("setting pin output mode")
}

func (fp *FakePin) Logger() *zap.Logger {
	if fp.logger == nil {
		log := global.Logger.With(
			zap.String("entity", "FakePin"),
			zap.Uint8("id", fp.pin),
		)

		fp.logger = log
	}

	return fp.logger
}
