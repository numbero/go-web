package logger

import "github.com/sirupsen/logrus"

var Log *logrus.Logger

func Init() {
	Log = logrus.New()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
