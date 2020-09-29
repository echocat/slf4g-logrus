package logrus

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/echocat/slf4g"
	"github.com/echocat/slf4g-logrus"
	"github.com/sirupsen/logrus"
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
	faw := &FormatterAndWriter{
		Target: target,
		Magic:  NewMagic(16),
	}
	return &logrus.Logger{
		Out:          faw,
		Hooks:        make(logrus.LevelHooks),
		Formatter:    faw,
		ReportCaller: false,
		Level:        logrus.TraceLevel,
		ExitFunc:     func(int) {},
	}
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

	e := log.NewEvent(sbl.LevelLogrusToSlf4g(le.Level), nil, 3).
		WithAll(le.Data).
		WithContext(le.Context)

	instance.Target.Log(e)
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
