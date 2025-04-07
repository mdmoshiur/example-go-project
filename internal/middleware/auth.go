package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/internal/jwtauth"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/internal/responder"
	"gorm.io/gorm"
)

// Auth AuthMiddleware authenticate http request.
func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) < 7 {
			responder.AccessDeniedErr(w)
			return
		}
		token := auth[7:]

		claims, err := jwtauth.ValidateToken(token)
		if err != nil {
			responder.AccessDeniedErr(w)
			return
		}

		// check token is not revoked
		isRevoked, err := isTokenRevoked(r.Context(), claims.StandardClaims.Id, claims.User.ID)
		if err != nil {
			if errors.Is(err, domain.ErrUserTokenNotFound) {
				responder.AccessDeniedErr(w)
				return
			}

			logger.Error(err)
			responder.InternalServerErr(w, err)
			return
		}

		if isRevoked { // token is revoked
			responder.AccessDeniedErr(w)
			return
		}

		ctx := context.WithValue(r.Context(), jwtauth.ContextAuthUser{}, claims.User)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// isTokenRevoked checks token is revoked or not
func isTokenRevoked(ctx context.Context, tokenID string, userID uint32) (bool, error) {
	cacheSvc, err := ContextCacheSvc(ctx)
	if err != nil {
		return true, err
	}

	key := fmt.Sprintf("auth_token_%d_%s", userID, tokenID)
	cachedValue, err := cacheSvc.GetCache(key)
	if err != nil {
		return true, err
	}

	if cachedValue != "" { // cache hit
		isRevoked, _ := strconv.ParseBool(cachedValue)
		return isRevoked, nil
	}

	// fetch from database
	isRevoked, err := isAuthTokenRevoked(ctx, tokenID)
	if err != nil {
		return true, err
	}

	// set cache for future uses
	if err := cacheSvc.SetCache(key, strconv.FormatBool(isRevoked), 6*time.Hour); err != nil {
		return true, domain.ErrCacheSet
	}

	return isRevoked, nil
}

// isAuthTokenRevoked ...
func isAuthTokenRevoked(ctx context.Context, tokenID string) (bool, error) {
	db, err := ContextDB(ctx)
	if err != nil {
		return true, err
	}
	revoked := true
	q := db.WithContext(ctx).Table("auth_tokens").
		Select("revoked").
		Where("token_id = ?", tokenID)

	if err = q.Take(&revoked).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, domain.ErrUserTokenNotFound
		}
		return true, fmt.Errorf("middleware:auth:fetch auth token: %w", err)
	}

	return revoked, nil
}
