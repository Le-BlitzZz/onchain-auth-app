package event

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	hooks := make(logrus.LevelHooks)

	Log = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    TextFormatter,
		Hooks:        hooks,
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

}
