package logrus

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/echocat/slf4g"
	"github.com/echocat/slf4g-logrus"
	"github.com/sirupsen/logrus"
	"reflect"
)

func Configure() {
	ConfigureWith(log.GetRootLogger())
}

func ConfigureWith(target log.CoreLogger) {
	faw := &FormatterAndWriter{
		Target: target,
		Magic:  NewMagic(16),
	}
	logrus.SetFormatter(faw)
	logrus.SetOutput(faw)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(false)
}

func CreateFor(target log.CoreLogger) *logrus.Logger {
	result := &logrus.Logger{
		Hooks: make(logrus.LevelHooks),
	}
	ConfigureExistingWith(result, target)
	return result
}

func ConfigureExistingWith(existing logrus.FieldLogger, target log.CoreLogger) {
	faw := &FormatterAndWriter{
		Target: target,
		Magic:  NewMagic(16),
	}
	var elr *logrus.Logger
	switch v := existing.(type) {
	case *logrus.Logger:
		elr = v
	case *logrus.Entry:
		elr = v.Logger
	default:
		panic(fmt.Errorf("cannot handle logger of type %v; has to be either %v or %v",
			reflect.TypeOf(existing), reflect.TypeOf(&logrus.Logger{}), reflect.TypeOf(&logrus.Entry{})))
	}
	elr.Out = faw
	elr.Formatter = faw
	elr.ReportCaller = false
	elr.Level = logrus.TraceLevel
	elr.ExitFunc = func(int) {}
}

type FormatterAndWriter struct {
	Target log.CoreLogger
	Magic  Magic
}

func (instance *FormatterAndWriter) Format(le *logrus.Entry) ([]byte, error) {
	le.Data[instance.Target.GetProvider().GetFieldKeysSpec().GetMessage()] = le.Message
	if v := le.Time; !v.IsZero() {
		le.Data[instance.Target.GetProvider().GetFieldKeysSpec().GetTimestamp()] = v
	}

	e := instance.Target.NewEvent(sbl.LevelLogrusToSlf4g(le.Level), nil).
		WithAll(le.Data)

	instance.Target.Log(e, 5)
	return instance.magic(), nil
}

func (instance *FormatterAndWriter) Write(p []byte) (n int, err error) {
	magic := instance.magic()
	if bytes.Equal(p, magic) {
		return len(p), nil
	}
	return 0, fmt.Errorf("expected magic %q; but got: %q", magic.String(), hex.EncodeToString(p))
}

func (instance *FormatterAndWriter) magic() Magic {
	if v := instance.Magic; v != nil {
		return v
	}
	return defaultMagic
}
