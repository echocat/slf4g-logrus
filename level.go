package sbl

import (
	"fmt"
	log "github.com/echocat/slf4g"
	"github.com/sirupsen/logrus"
)

const (
	LevelPanic = log.LevelFatal + 1000
)

func LevelLogrusToSlf4g(in logrus.Level) log.Level {
	switch in {
	case logrus.TraceLevel:
		return log.LevelTrace
	case logrus.DebugLevel:
		return log.LevelDebug
	case logrus.InfoLevel:
		return log.LevelInfo
	case logrus.WarnLevel:
		return log.LevelWarn
	case logrus.ErrorLevel:
		return log.LevelError
	case logrus.FatalLevel:
		return log.LevelFatal
	case logrus.PanicLevel:
		return LevelPanic
	default:
		panic(fmt.Sprintf("unknown logrus level %v", in))
	}
}

func LevelSlf4gToLogrus(in log.Level) logrus.Level {
	switch in {
	case log.LevelTrace:
		return logrus.TraceLevel
	case log.LevelDebug:
		return logrus.DebugLevel
	case log.LevelInfo:
		return logrus.InfoLevel
	case log.LevelWarn:
		return logrus.WarnLevel
	case log.LevelError:
		return logrus.ErrorLevel
	case log.LevelFatal:
		return logrus.FatalLevel
	case LevelPanic:
		return logrus.PanicLevel
	default:
		panic(fmt.Sprintf("unknown log level %v", in))
	}
}

func LevelProvider() []log.Level {
	return []log.Level{
		log.LevelTrace,
		log.LevelDebug,
		log.LevelInfo,
		log.LevelWarn,
		log.LevelError,
		log.LevelFatal,
		LevelPanic,
	}
}
