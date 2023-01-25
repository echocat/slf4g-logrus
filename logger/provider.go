package logrus

import (
	"github.com/echocat/slf4g"
	"github.com/echocat/slf4g-logrus"
	"github.com/echocat/slf4g/fields"
	"github.com/echocat/slf4g/level"
	"github.com/sirupsen/logrus"
)

var DefaultProvider = NewProvider("logrus", logrus.StandardLogger())

func NewProvider(name string, target *logrus.Logger) *Provider {
	if target.Hooks == nil {
		target.Hooks = make(logrus.LevelHooks)
	}
	target.Hooks.Add(&hook{})
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

func (instance *Provider) GetRootLogger() log.Logger {
	return instance.GetLogger(RootLoggerName)
}

func (instance *Provider) GetLogger(name string) log.Logger {
	return log.NewLogger(&CoreLogger{
		Provider: instance,
		Name:     name,
	})
}

func (instance *Provider) GetAllLevels() level.Levels {
	return sbl.LevelProvider()
}

func (instance *Provider) GetFieldKeysSpec() fields.KeysSpec {
	return fieldKeysSpecV
}

func (instance *Provider) getLevel() level.Level {
	return sbl.LevelLogrusToSlf4g(instance.Target.Level)
}
