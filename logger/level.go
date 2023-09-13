package logger

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	LevelTrace string = "trace"
	LevelDebug string = "debug"
	LevelInfo  string = "info"
	LevelWarn  string = "warn"
	LevelError string = "error"
	LevelFatal string = "fatal"
	LevelPanic string = "panic"

	LevelTraceMsg   string = "[TRACE]"
	LevelDebugMsg   string = "[DEBUG]"
	LevelInfoMsg    string = "[INFO]"
	LevelWarnMsg    string = "[WARNING]"
	LevelErrorMsg   string = "[ERROR]"
	LevelFatalMsg   string = "[FATAL]"
	LevelPanicMsg   string = "[PANIC]"
	LevelUnknownMsg string = "[UNKNOWN LOG LEVEL]"
)

var (
	ErrUnknownLevel error = fmt.Errorf(`Unknown Level`)
)

func (l *logger) convertAndSetLevel() {
	l.setLevelLogrus()
}

func (l *logger) setLevelLogrus() {
	var lrLevel logrus.Level
	switch l.opt.Level {
	case LevelTrace:
		lrLevel = logrus.TraceLevel
		l.entry.Info(OK, infoLogger+"level:", LevelTraceMsg)
	case LevelDebug:
		lrLevel = logrus.DebugLevel
		l.entry.Info(OK, infoLogger+"level:", LevelDebugMsg)
	case LevelInfo:
		lrLevel = logrus.InfoLevel
		l.entry.Info(OK, infoLogger+"level:", LevelInfoMsg)
	case LevelWarn:
		lrLevel = logrus.WarnLevel
		l.entry.Info(OK, infoLogger+"level:", LevelWarnMsg)
	case LevelError:
		lrLevel = logrus.ErrorLevel
		l.entry.Info(OK, infoLogger+"level:", LevelErrorMsg)
	case LevelFatal:
		lrLevel = logrus.FatalLevel
		l.entry.Info(OK, infoLogger+"level:", LevelFatalMsg)
	case LevelPanic:
		lrLevel = logrus.PanicLevel
		l.entry.Info(OK, infoLogger+"level:", LevelPanicMsg)
	default:
		err := errors.Wrapf(ErrUnknownLevel, errLogger, FAILED)
		l.entry.Panic(err)
	}
	//set logrus log level
	l.logger.SetLevel(lrLevel)
}
