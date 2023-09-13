package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/tuingking/sdkgo/appcontext"
)

func RequestID(appName string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			requestID := r.Header.Get(string(appcontext.RequestID))
			if requestID == "" {
				myid := uuid.New().String()
				requestID = fmt.Sprintf("%s-%s", appName, myid)
			}
			ctx = context.WithValue(ctx, appcontext.RequestID, requestID)
			ctx = context.WithValue(ctx, appcontext.AppName, appName)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
