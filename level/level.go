package level

import (
	"fmt"
	"github.com/echocat/slf4g/level"
	"github.com/sirupsen/logrus"
)

const (
	Panic = level.Fatal + 1000
)

func LogrusToSlf4g(in logrus.Level) level.Level {
	switch in {
	case logrus.TraceLevel:
		return level.Trace
	case logrus.DebugLevel:
		return level.Debug
	case logrus.InfoLevel:
		return level.Info
	case logrus.WarnLevel:
		return level.Warn
	case logrus.ErrorLevel:
		return level.Error
	case logrus.FatalLevel:
		return level.Fatal
	case logrus.PanicLevel:
		return Panic
	default:
		panic(fmt.Sprintf("unknown logrus level %v", in))
	}
}

func Slf4gToLogrus(in level.Level) logrus.Level {
	switch in {
	case level.Trace:
		return logrus.TraceLevel
	case level.Debug:
		return logrus.DebugLevel
	case level.Info:
		return logrus.InfoLevel
	case level.Warn:
		return logrus.WarnLevel
	case level.Error:
		return logrus.ErrorLevel
	case level.Fatal:
		return logrus.FatalLevel
	case Panic:
		return logrus.PanicLevel
	default:
		panic(fmt.Sprintf("unknown log level %v", in))
	}
}
