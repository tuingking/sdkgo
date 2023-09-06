package logger

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tuingking/sdkgo/appcontext"
)

func (l *logger) parseContextFields(ctx context.Context) *logrus.Entry {
	entry := l.entry
	if ctx != nil {
		for k, v := range l.opt.ContextFields {
			if val := ctx.Value(v); val != nil {
				entry = entry.WithField(k, val)
			}
		}

		requestID := appcontext.GetAppRequestID(ctx)
		if requestID != "" {
			entry = entry.WithField("request_id", requestID)
		}

		method := appcontext.GetRequestMethod(ctx)
		if requestID != "" {
			entry = entry.WithField("method", method)
		}

		userID := appcontext.GetUserID(ctx)
		if userID != 0 {
			entry = entry.WithField("user_id", userID)
		}

		appName := appcontext.GetAppName(ctx)
		if appName != "" {
			entry = entry.WithField("app_name", appName)
		}
	}
	return entry
}

func (l *logger) Trace(v ...interface{}) {
	l.TraceWithContext(context.TODO(), v...)
}

func (l *logger) Debug(v ...interface{}) {
	l.DebugWithContext(context.TODO(), v...)
}

func (l *logger) Info(v ...interface{}) {
	l.InfoWithContext(context.TODO(), v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.WarnWithContext(context.TODO(), v...)
}

func (l *logger) Error(v ...interface{}) {
	l.ErrorWithContext(context.TODO(), v...)
}

func (l *logger) Fatal(v ...interface{}) {
	l.FatalWithContext(context.TODO(), v...)
}

func (l *logger) TraceWithContext(ctx context.Context, v ...interface{}) {
	l.parseContextFields(ctx).Trace(v...)
}

func (l *logger) DebugWithContext(ctx context.Context, v ...interface{}) {
	l.parseContextFields(ctx).Debug(v...)
}

func (l *logger) InfoWithContext(ctx context.Context, v ...interface{}) {
	l.parseContextFields(ctx).Info(v...)
}

func (l *logger) WarnWithContext(ctx context.Context, v ...interface{}) {
	l.parseContextFields(ctx).Warn(v...)
}

func (l *logger) ErrorWithContext(ctx context.Context, v ...interface{}) {
	l.parseContextFields(ctx).Error(v...)
}

func (l *logger) FatalWithContext(ctx context.Context, v ...interface{}) {
	l.parseContextFields(ctx).Fatal(v...)
}
