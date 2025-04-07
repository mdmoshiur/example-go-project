package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/internal/cache"
	"github.com/mdmoshiur/example-go/internal/jwtauth"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/store"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase represents the user usecase
type UserUseCase struct {
	Store    store.Store
	CacheSvc cache.Cacher
}

// New return user usecase implementation
func New(s store.Store, cs cache.Cacher) domain.UserUseCase {
	return &UserUseCase{
		Store:    s,
		CacheSvc: cs,
	}
}

// StoreOrUpdate stores/updates user in database
func (u *UserUseCase) StoreOrUpdate(ctx context.Context, user *domain.User) error {
	if len(user.Password) > 0 {
		hp, err := GenPasswordHash(user.Password)
		if err != nil {
			logger.Error(err)
			return err
		}

		user.Password = hp // set hashed password
	}

	if user.ID > 0 { // update user
		err := u.Store.UserRepository().Update(ctx, user)
		if err != nil {
			return err
		}

		// if deactivate or suspend user revoke all user tokens
		if user.Status != nil && *user.Status != domain.UserStatusActive {
			err = u.Store.UserRepository().RevokeAllAuthToken(ctx, user.ID)
			if err != nil {
				return err
			}

			// clear auth cache
			key := fmt.Sprintf("auth_token_%d_*", user.ID)
			err = u.CacheSvc.ClearCache(key)
			if err != nil {
				logger.Error(err)
				return err
			}
		}

		return nil
	}

	// create user
	atomicCallback := func(ds store.Store) error {
		if err := ds.UserRepository().Store(ctx, user); err != nil {
			return err
		}

		return nil
	}

	err := u.Store.Atomic(ctx, atomicCallback)
	return err
}

func (u *UserUseCase) Login(ctx context.Context, ctr *domain.LoginCriteria) (string, error) {
	user, err := u.Store.UserRepository().FetchUserByEmail(ctx, ctr.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return "", domain.ErrWrongEmail
		}
		logger.Error(err)
		return "", err
	}

	if !IsMatchedPasswordHash(ctr.Password, user.Password) {
		return "", domain.ErrWrongPassword
	}

	if user.Status != nil && *user.Status != domain.UserStatusActive {
		return "", domain.ErrDeactivatedUser
	}

	ctxUser := &jwtauth.AuthUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	token, tokenID, err := jwtauth.GenAccessToken(ctxUser)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	// save token into database
	authToken := &domain.AuthToken{
		TokenID: tokenID,
		UserID:  user.ID,
		Revoked: false,
	}

	err = u.Store.UserRepository().StoreAuthToken(ctx, authToken)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) Logout(ctx context.Context, token string) error {
	claims, err := jwtauth.ValidateToken(token)
	if err != nil {
		logger.Error(err)
		return err
	}

	tokenID := claims.StandardClaims.Id
	err = u.Store.UserRepository().RevokeAuthToken(ctx, tokenID)
	if err != nil {
		logger.Error(err)
		return err
	}

	// clear cache
	key := fmt.Sprintf("auth_token_%d_%s", claims.User.ID, tokenID)
	err = u.CacheSvc.ClearCache(key)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// GenPasswordHash generates password hash from plain password
func GenPasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// IsMatchedPasswordHash checks password with its hashed password
func IsMatchedPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
