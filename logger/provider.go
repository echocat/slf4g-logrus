package logrus

import (
	"github.com/echocat/slf4g"
	"github.com/echocat/slf4g-logrus"
	"github.com/echocat/slf4g/fields"
	"github.com/sirupsen/logrus"
)

var DefaultProvider = NewProvider("logrus", logrus.StandardLogger())

func NewProvider(name string, target *logrus.Logger) *Provider {
	return &Provider{
		Name:   name,
		Target: target,
	}
}

type Provider struct {
	Name   string
	Target *logrus.Logger
}

func (instance *Provider) GetName() string {
	return instance.Name
}

func (instance *Provider) GetLogger(name string) log.Logger {
	return log.NewLogger(&CoreLogger{
		Provider: instance,
		Name:     name,
	})
}

func (instance *Provider) GetAllLevels() []log.Level {
	return []log.Level{
		log.LevelTrace,
		log.LevelDebug,
		log.LevelInfo,
		log.LevelWarn,
		log.LevelError,
		log.LevelFatal,
		sbl.LevelPanic,
	}
}

func (instance *Provider) GetFieldKeySpec() fields.KeysSpec {
	return fieldKeysSpecV
}

func (instance *Provider) getLevel() log.Level {
	return sbl.LevelLogrusToSlf4g(instance.Target.Level)
}
