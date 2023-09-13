package logger

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tuingking/sdkgo/appcontext"
)

func (l *logger) parseContextFields(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{}
	if ctx != nil {
		for k, v := range l.opt.ContextFields {
			if val := ctx.Value(v); val != nil {
				fields[k] = v
			}
		}

		requestID := appcontext.GetAppRequestID(ctx)
		if requestID != "" {
			fields["request_id"] = requestID
		}

		method := appcontext.GetRequestMethod(ctx)
		if method != "" {
			fields["method"] = method
		}

		userID := appcontext.GetUserID(ctx)
		if userID != 0 {
			fields["user_id"] = userID
		}

		appName := appcontext.GetAppName(ctx)
		if appName != "" {
			fields["app_name"] = appName
		}

		packageName := appcontext.GetPackageName(ctx)
		if packageName != "" {
			fields["pkg"] = packageName
		}
	}
	return fields
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
	l.logger.WithFields(l.parseContextFields(ctx)).Trace(v...)
}

func (l *logger) DebugWithContext(ctx context.Context, v ...interface{}) {
	l.logger.WithFields(l.parseContextFields(ctx)).Debug(v...)
}

func (l *logger) InfoWithContext(ctx context.Context, v ...interface{}) {
	l.logger.WithFields(l.parseContextFields(ctx)).Info(v...)
}

func (l *logger) WarnWithContext(ctx context.Context, v ...interface{}) {
	l.logger.WithFields(l.parseContextFields(ctx)).Warn(v...)
}

func (l *logger) ErrorWithContext(ctx context.Context, v ...interface{}) {
	l.logger.WithFields(l.parseContextFields(ctx)).Error(v...)
}

func (l *logger) FatalWithContext(ctx context.Context, v ...interface{}) {
	l.logger.WithFields(l.parseContextFields(ctx)).Fatal(v...)
}
