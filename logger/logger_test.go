package logger_test

import (
	"context"
	"os"
	"testing"

	"github.com/tuingking/sdkgo/appcontext"
	"github.com/tuingking/sdkgo/logger"
)

func Test_Logger(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		l := logger.New(logger.Option{})
		ctx := context.WithValue(context.Background(), appcontext.RequestID, "foo")
		ctx = context.WithValue(ctx, appcontext.PackageName, "logger")
		l.InfoWithContext(ctx, "hello logger")
	})

	t.Run("Init", func(t *testing.T) {
		logger.Init(logger.Option{
			Level:         "debug",
			Formatter:     "json",
			Output:        "file",
			LogOutputPath: "./example.log",
		})
		defer func() {
			logger.Log.Stop()
			os.Remove("./example.log")
		}()

		ctx := context.WithValue(context.Background(), appcontext.RequestID, "foo")
		ctx = context.WithValue(ctx, appcontext.PackageName, "logger")
		logger.Log.InfoWithContext(ctx, "hello logger")
	})
}
