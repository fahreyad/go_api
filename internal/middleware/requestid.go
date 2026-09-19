package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey int

const (
	requestIDHeader            = "X-Request-ID"
	requestIDKey    contextKey = iota
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a unique request ID (you can use a UUID or any other method)
		requestID := r.Header.Get(requestIDHeader)
		if requestID == "" {
			requestID = uuid.New().String() // Replace with actual request ID generation logic
		}

		// Set the request ID in the response header
		w.Header().Set("X-Request-ID", requestID)
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		// Call the next handler in the chain
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	requestID := ctx.Value(requestIDKey).(string)
	return requestID

}
