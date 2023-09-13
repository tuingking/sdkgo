package logger

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

const (
	OutputStdout  string = `stdout`
	OutputFile    string = `file`
	OutputDiscard string = `discard`

	outputStdout  string = `[STDOUT]`
	outputFile    string = `[FILE]`
	outputDiscard string = `[DISCARD]`
	outputUnknown string = `[UNKNOWN LOG OUTPUT]`
)

var (
	errUnknownOutput = errors.New(`Unknown log Output`)
	ErrUnknownOutput = errors.Wrapf(errUnknownOutput, errLogger, FAILED)
)

func (l *logger) convertAndSetOutput() {
	switch l.opt.Output {
	case OutputDiscard:
		l.logger.SetOutput(io.Discard)
		l.entry.Info(OK, infoLogger+"output:", outputDiscard)
	case OutputStdout:
		l.logger.SetOutput(os.Stdout)
		l.entry.Info(OK, infoLogger+"output:", outputStdout)
	case OutputFile:
		f, err := os.OpenFile(l.opt.LogOutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			err = errors.Wrapf(err, errLogger, FAILED)
			l.entry.Panic(err)
		}
		l.file = f
		l.logger.SetOutput(f)
		l.entry.Info(OK, infoLogger+"output:", outputFile, l.opt.LogOutputPath)
	default:
		l.entry.Panic(ErrUnknownOutput)
	}
}
