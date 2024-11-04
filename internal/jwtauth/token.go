package jwtauth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/internal/config"
	"github.com/mdmoshiur/example-go/internal/logger"
)

// ValidateToken validates json web token
func ValidateToken(t string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(t, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWT().SecretKey, nil
	})
	if err != nil {
		logger.Error(fmt.Errorf("token parse failed: err: %w", err))
		return nil, err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	return claims, nil
}

// GenAccessToken generates json web token
func GenAccessToken(user *AuthUser) (string, string, error) {
	id := uuid.New()
	now := time.Now().Unix()
	// Create the Claims
	claims := AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        id.String(),
			ExpiresAt: time.Now().Add(config.JWT().ExpirationDurationInHours).Unix(),
			IssuedAt:  now,
			NotBefore: now,
		},
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	st, err := token.SignedString(config.JWT().SecretKey)
	if err != nil {
		return "", "", fmt.Errorf("jwt signing failed: err: %w", err)
	}

	return st, id.String(), nil
}
