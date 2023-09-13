package appcontext

import "context"

// SetAppLang - Setter context x-app-lang
func SetAppLang(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, AppLang, value)
}

// GetAppLang - Getter context x-app-lang
func GetAppLang(ctx context.Context) string {
	if v, ok := ctx.Value(AppLang).(string); ok {
		return v
	}
	return ""
}

// SetAppRequestID - Setter context AppRequestID
func SetAppRequestID(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, RequestID, value)
}

// GetAppRequestID - Getter context RequestID
func GetAppRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(RequestID).(string); ok {
		return v
	}
	return ""
}

// SetRequestMethod - Setter context RequestMethod
func SetRequestMethod(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, RequestMethod, value)
}

// GetRequestMethod - Getter context RequestMethod
func GetRequestMethod(ctx context.Context) string {
	if v, ok := ctx.Value(RequestMethod).(string); ok {
		return v
	}
	return ""
}

// SetUserID - Setter context UserID
func SetUserID(ctx context.Context, value int64) context.Context {
	return context.WithValue(ctx, UserID, value)
}

// GetUserID - Getter context UserID
func GetUserID(ctx context.Context) int64 {
	if v, ok := ctx.Value(UserID).(int64); ok {
		return v
	}
	return 0
}

// SetAppName - Setter context AppName
func SetAppName(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, AppName, value)
}

// GetAppName - Getter context AppName
func GetAppName(ctx context.Context) string {
	if v, ok := ctx.Value(AppName).(string); ok {
		return v
	}
	return ""
}

// SetPackageName - Setter context PackageName
func SetPackageName(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, PackageName, value)
}

// GetPackageName - Getter context PackageName
func GetPackageName(ctx context.Context) string {
	if v, ok := ctx.Value(PackageName).(string); ok {
		return v
	}
	return ""
}
