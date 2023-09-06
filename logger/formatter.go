package logger

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	FormatJSON string = "json"
	FormatText string = "text"

	formatJSON          string = "[JSON]"
	formatText          string = "[TEXT]"
	formatUnknownstring        = "[UNKNOWN LOG FORMAT]"
	formatTimeMs               = "2006-01-02T15:04:05.000Z07:00"
)

var (
	errUnknownFormat = fmt.Errorf(`Unknown log format`)
	ErrUnknownFormat = errors.Wrapf(errUnknownFormat, errLogger, FAILED)
)

func (l *logger) convertAndSetFormatter() {
	switch l.opt.Formatter {
	case FormatText:
		l.logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: formatTimeMs})
		l.entry.Info(OK, infoLogger, formatText)
	case FormatJSON:
		l.logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: formatTimeMs})
		l.entry.Info(OK, infoLogger, formatJSON)
	default:
		l.entry.Panic(ErrUnknownFormat)
	}
}
