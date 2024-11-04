package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/mdmoshiur/example-go/internal/cache"
)

const ContextCache = "ctx_cache_"

// WithCache middleware add cacheSvc to the http request.
func WithCache(cs cache.Cacher) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ContextCache, cs)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

// ContextCacheSvc returns context cache service
func ContextCacheSvc(ctx context.Context) (cache.Cacher, error) {
	ctxVal := ctx.Value(ContextCache)
	if ctxVal == nil {
		return nil, errors.New("contextcache: failed to extract cache service from request-context")
	}
	return ctxVal.(cache.Cacher), nil
}
