package level

import (
	"fmt"
	"github.com/echocat/slf4g/level"
	"strconv"
	"strings"
)

var DefaultNames level.Names = &Names{}

type Names struct{}

func (instance *Names) ToName(lvl level.Level) (string, error) {
	switch lvl {
	case level.Trace:
		return "TRACE", nil
	case level.Debug:
		return "DEBUG", nil
	case level.Info:
		return "INFO", nil
	case level.Warn:
		return "WARN", nil
	case level.Error:
		return "ERROR", nil
	case level.Fatal:
		return "FATAL", nil
	case Panic:
		return "PANIC", nil
	default:
		return fmt.Sprintf("%d", lvl), nil
	}
}

func (instance *Names) ToLevel(name string) (level.Level, error) {
	switch strings.ToUpper(name) {
	case "TRACE":
		return level.Trace, nil
	case "DEBUG", "VERBOSE":
		return level.Debug, nil
	case "INFO", "INFORMATION":
		return level.Info, nil
	case "WARN", "WARNING":
		return level.Warn, nil
	case "ERROR", "ERR":
		return level.Error, nil
	case "FATAL":
		return level.Fatal, nil
	case "PANIC":
		return Panic, nil
	default:
		if result, err := strconv.ParseUint(name, 10, 16); err != nil {
			return 0, fmt.Errorf("illegal level: %s", name)
		} else {
			return level.Level(result), nil
		}
	}
}
