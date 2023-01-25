package level

import "github.com/echocat/slf4g/level"

var DefaultProvider level.Provider = &Provider{}

type Provider struct{}

func (instance *Provider) GetName() string {
	return "logrus"
}

func (instance *Provider) GetLevels() level.Levels {
	return []level.Level{
		level.Trace,
		level.Debug,
		level.Info,
		level.Warn,
		level.Error,
		level.Fatal,
		Panic,
	}
}
