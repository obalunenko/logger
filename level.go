package logger

import (
	"github.com/sirupsen/logrus"
)

type Level struct {
	logrus.Level
}

func (l Level) IsDebug() bool {
	return l.Level >= logrus.DebugLevel
}
