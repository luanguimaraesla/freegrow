package gadgets

import (
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry
)

func SetLogger(logger *logrus.Entry) {
	log = logger
}

func GetLogger() *logrus.Entry {
	return log
}
