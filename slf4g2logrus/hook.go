package logrus

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

type hook struct {
}

func (instance hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (instance hook) Fire(entry *logrus.Entry) error {
	if entry == nil || entry.Context == nil {
		return nil
	}
	if caller, ok := entry.Context.Value(callerKey).(*runtime.Frame); ok {
		entry.Caller = caller
	}
	return nil
}
