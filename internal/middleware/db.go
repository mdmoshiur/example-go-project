package middleware

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type ContextDBKey struct{}

// WithDB middleware add db connection to the http request.
func WithDB(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ContextDBKey{}, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

// ContextDB returns context database connection
func ContextDB(ctx context.Context) (*gorm.DB, error) {
	ctxVal := ctx.Value(ContextDBKey{})
	if ctxVal == nil {
		return nil, errors.New("contextdb: failed to extract context database connection from request-context")
	}
	return ctxVal.(*gorm.DB), nil
}
