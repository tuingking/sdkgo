package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

func (l *logger) PipeWriter() io.Writer {
	l.writer = l.entry.WriterLevel(logrus.WarnLevel)
	return l.writer
}
