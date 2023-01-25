package slf4g2logrus

import "github.com/sirupsen/logrus"

type KeysSpec interface {
	GetTimestamp() string
	GetMessage() string
	GetError() string
	GetLogger() string
}

var fieldKeysSpecV = &fieldKeysSpec{}

type fieldKeysSpec struct{}

func (instance *fieldKeysSpec) GetTimestamp() string {
	return logrus.FieldKeyTime
}

func (instance *fieldKeysSpec) GetMessage() string {
	return logrus.FieldKeyMsg
}

func (instance *fieldKeysSpec) GetError() string {
	return logrus.ErrorKey
}

func (instance *fieldKeysSpec) GetLogger() string {
	return "logger"
}
