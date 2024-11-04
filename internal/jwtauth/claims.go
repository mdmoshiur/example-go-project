package jwtauth

import "github.com/golang-jwt/jwt"

type AccessTokenClaims struct {
	jwt.StandardClaims

	User *AuthUser `json:"user,omitempty"`
}
