package logger


import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logrus *logrus.Entry
}

func New() Logger {
	l := initLogger()
	return Logger{Logrus: l}
}

func initLogger() *logrus.Entry {
	logger := logrus.WithFields(logrus.Fields{
		"logger":   "LOGRUS",
	})
	logger.Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:             true,
		FullTimestamp:             true,
	})
	logger.Logger.SetLevel(logrus.TraceLevel)
	return logger
}

