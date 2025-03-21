package event

import "github.com/sirupsen/logrus"

var TextFormatter = &logrus.TextFormatter{
	DisableColors: false,
	FullTimestamp: true,
}
