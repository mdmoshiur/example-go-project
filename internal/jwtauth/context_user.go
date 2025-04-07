package jwtauth

import (
	"context"
	"errors"
)

type ContextAuthUser struct{}

type AuthUser struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ContextUser returns context user information
func ContextUser(ctx context.Context) (*AuthUser, error) {
	ctxVal := ctx.Value(ContextAuthUser{})
	if ctxVal == nil {
		return nil, errors.New("contextuser: failed to extract user info from request-context")
	}
	return ctxVal.(*AuthUser), nil
}
