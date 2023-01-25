package slf4g2logrus

import (
	"context"
	"github.com/echocat/slf4g"
	blevel "github.com/echocat/slf4g-logrus/level"
	"github.com/echocat/slf4g/fields"
	"github.com/echocat/slf4g/level"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

const (
	RootLoggerName = "ROOT"
)

var (
	callerKey = struct {
		slf4gLogrusLoggerCallerKey struct{}
	}{}
)

type CoreLogger struct {
	Provider *Provider
	Name     string
}

func (instance *CoreLogger) GetName() string {
	return instance.Name
}

func (instance *CoreLogger) Log(e log.Event, skipFrames uint16) {
	le := logrus.NewEntry(instance.Provider.Target)
	le.Level = blevel.Slf4gToLogrus(e.GetLevel())
	if err := e.ForEach(func(key string, value interface{}) error {
		le.Data[key] = value
		return nil
	}); err != nil {
		panic(err)
	}

	if v := log.GetTimestampOf(e, instance.Provider); v != nil {
		le.Time = *v
		delete(le.Data, instance.Provider.GetFieldKeysSpec().GetTimestamp())
	}

	var msg string
	if v := log.GetMessageOf(e, instance.Provider); v != nil {
		msg = *v
		delete(le.Data, instance.Provider.GetFieldKeysSpec().GetMessage())
	}

	if v := log.GetLoggerOf(e, instance.Provider); v == nil && instance.GetName() != RootLoggerName {
		le.Data[instance.Provider.GetFieldKeysSpec().GetLogger()] = instance.GetName()
	} else if v != nil && *v == RootLoggerName {
		delete(le.Data, instance.Provider.GetFieldKeysSpec().GetLogger())
	}

	if le.Logger.ReportCaller {
		le.Caller = instance.getCaller(skipFrames + 1)
		le.Context = context.WithValue(context.Background(), callerKey, le.Caller)
	}

	le.Log(le.Level, msg)
}

func (instance *CoreLogger) getCaller(skipFrames uint16) *runtime.Frame {
	pcs := make([]uintptr, 2)
	depth := runtime.Callers(int(skipFrames)+2, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	frame, ok := frames.Next()
	if !ok {
		return nil
	}
	return &frame
}

func (instance *CoreLogger) getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func (instance *CoreLogger) NewEvent(l level.Level, values map[string]interface{}) log.Event {
	return instance.NewEventWithFields(l, fields.WithAll(values))
}

// NewEventWithFields provides a shortcut if an event should directly be created
// from fields.
func (instance *CoreLogger) NewEventWithFields(l level.Level, f fields.ForEachEnabled) log.Event {
	asFields, err := fields.AsFields(f)
	if err != nil {
		panic(err)
	}
	return &event{
		provider: instance.Provider,
		fields:   asFields,
		level:    l,
	}
}

func (instance *CoreLogger) Accepts(e log.Event) bool {
	return e != nil
}

func (instance *CoreLogger) IsLevelEnabled(l level.Level) bool {
	return instance.Provider.getLevel().CompareTo(l) <= 0
}

func (instance *CoreLogger) GetProvider() log.Provider {
	return instance.Provider
}
