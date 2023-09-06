package logger

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	infoLogger string = `Logger:`
	errLogger  string = `%s Logger Error`
	OK         string = "[OK]"
	FAILED     string = "[FAILED]"
)

var Log Logger

var once = &sync.Once{}

type Logger interface {
	SetOption(opt Option)
	Stop()
	PipeWriter() io.Writer
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	TraceWithContext(ctx context.Context, v ...interface{})
	DebugWithContext(ctx context.Context, v ...interface{})
	InfoWithContext(ctx context.Context, v ...interface{})
	WarnWithContext(ctx context.Context, v ...interface{})
	ErrorWithContext(ctx context.Context, v ...interface{})
	FatalWithContext(ctx context.Context, v ...interface{})
}

type logger struct {
	mu     *sync.RWMutex
	logger *logrus.Logger
	entry  *logrus.Entry
	opt    Option
	file   *os.File
	writer *io.PipeWriter
}

type Option struct {
	Level         string
	Formatter     string
	Output        string
	LogOutputPath string
	DefaultFields map[string]string
	ContextFields map[string]string
}

func Init(opt Option) {
	once.Do(func() {
		Log = New(opt)
	})
}

func New(opt Option) Logger {
	log := &logger{
		mu:     &sync.RWMutex{},
		logger: logrus.New(),
		entry:  logrus.WithFields(logrus.Fields{}),
		opt:    opt,
	}
	log.logger.SetOutput(io.Discard)
	log.setDefaultOptions()
	log.applyOptions()

	return log
}

func (l *logger) setDefaultOptions() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.opt.Output == "" {
		//never put default to discard, error will not be displayed!
		l.opt.Output = OutputStdout
	}
	if l.opt.Formatter == "" {
		l.opt.Formatter = FormatText
	}
	if l.opt.Level == "" {
		l.opt.Level = LevelTrace
	}
}

func (l *logger) applyOptions() {
	l.convertAndSetOutput()
	l.convertAndSetFormatter()
	l.convertAndSetLevel()
}

func (l *logger) SetOption(opt Option) {
	l.mu.Lock()
	l.opt = opt
	l.mu.Unlock()
	l.applyOptions()
}

func (l *logger) Stop() {
	if l.writer != nil {
		l.writer.Close()
	}
	if l.file != nil {
		l.file.Close()
	}
}
