package middleware

import (
	"context"
	"net/http"

	"github.com/tuingking/sdkgo/appcontext"
)

func AppName(name string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if name == "" {
				name = "UNKNOWN"
			}
			ctx = context.WithValue(ctx, appcontext.AppName, name)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
