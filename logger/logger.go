package logrus

import (
	"github.com/echocat/slf4g"
	sbl "github.com/echocat/slf4g-logrus"
	"github.com/echocat/slf4g/fields"
	"github.com/sirupsen/logrus"
)

type CoreLogger struct {
	Provider *Provider
	Name     string
}

func (instance *CoreLogger) GetName() string {
	return instance.Name
}

func (instance *CoreLogger) Log(e log.Event) {
	le := logrus.NewEntry(instance.Provider.Target)
	le.Data = logrus.Fields(fields.AsMap(e.GetFields()))
	le.Level = sbl.LevelSlf4gToLogrus(e.GetLevel())

	if v := log.GetTimestampOf(e, instance.Provider); v != nil {
		le.Time = *v
		delete(le.Data, instance.Provider.GetFieldKeySpec().GetTimestamp())
	}

	var msg string
	if v := log.GetMessageOf(e, instance.Provider); v != nil {
		msg = *v
		delete(le.Data, instance.Provider.GetFieldKeySpec().GetMessage())
	}

	if v := log.GetLoggerOf(e, instance.Provider); v == nil {
		le.Data[instance.Provider.GetFieldKeySpec().GetMessage()] = instance.GetName()
	}

	le.Log(le.Level, msg)
}

func (instance *CoreLogger) IsLevelEnabled(level log.Level) bool {
	return instance.Provider.getLevel().CompareTo(level) <= 0
}

func (instance *CoreLogger) GetProvider() log.Provider {
	return instance.Provider
}
