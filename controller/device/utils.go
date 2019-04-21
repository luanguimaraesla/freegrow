package device

import (
        "github.com/sirupsen/logrus"
)

var log *logrus.Entry

func SetLogger (logger *logrus.Entry) {
        log = logger
}

func logWithMetadata(id int, name string, kind string) *logrus.Entry {
        return log.WithFields(logrus.Fields{
                "deviceId": id,
                "deviceName": name,
                "deviceKind": kind,
        })
}
