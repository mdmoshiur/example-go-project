package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/domain/dto"
	"gorm.io/gorm"
)

// UserRepo represents mysql implementation of user repository contract
type UserRepo struct {
	DB *gorm.DB
}

// New return a mysql implementation of user storage repository
func New(db *gorm.DB) domain.UserRepository {
	return &UserRepo{
		DB: db,
	}
}

// Store inserts a new user to database
func (ur *UserRepo) Store(ctx context.Context, user *domain.User) error {
	if err := ur.DB.WithContext(ctx).Create(user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			if strings.Contains(pgErr.Message, "email") {
				return domain.ErrUserDuplicateEmail
			} else if strings.Contains(pgErr.Message, "phone") {
				return domain.ErrUserDuplicatePhone
			}
		}

		return fmt.Errorf("repository:user: user create: %w", err)
	}

	return nil
}

// Update updates user to database
func (ur *UserRepo) Update(ctx context.Context, user *domain.User) error {
	q := ur.DB.WithContext(ctx).Model(&domain.User{}).
		Where("id = ?", user.ID)

	if err := q.Omit("id").Updates(user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			if strings.Contains(pgErr.Message, "email") {
				return domain.ErrUserDuplicateEmail
			} else if strings.Contains(pgErr.Message, "phone") {
				return domain.ErrUserDuplicatePhone
			}
		}

		return fmt.Errorf("repository:user: user update: %w", err)
	}

	if q.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// StoreAuthToken stores authenticated user's token
func (ur *UserRepo) StoreAuthToken(ctx context.Context, authToken *domain.AuthToken) error {
	if err := ur.DB.WithContext(ctx).Create(authToken).Error; err != nil {
		return fmt.Errorf("repository:user: store auth token: %w", err)
	}

	return nil
}

// RevokeAuthToken revokes authenticated user's token
func (ur *UserRepo) RevokeAuthToken(ctx context.Context, tokenID string) error {
	q := ur.DB.WithContext(ctx).Model(&domain.AuthToken{}).
		Where("token_id = ?", tokenID).
		Updates(domain.AuthToken{
			Revoked: true,
		})

	if err := q.Error; err != nil {
		return fmt.Errorf("repository:user: revoke auth token: %w", err)
	}

	return nil
}

// RevokeAllAuthToken revokes specific user's all auth tokens
func (ur *UserRepo) RevokeAllAuthToken(ctx context.Context, userID uint32) error {
	q := ur.DB.WithContext(ctx).Model(&domain.AuthToken{}).
		Where("user_id = ?", userID).
		Updates(domain.AuthToken{
			Revoked: true,
		})

	if err := q.Error; err != nil {
		return fmt.Errorf("repository:user: revoke all auth token: %w", err)
	}

	return nil
}

func (ur *UserRepo) FetchUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *domain.User
	q := ur.DB.WithContext(ctx).Table("users").
		Where("users.email = ?", email)

	if err := q.Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("repository:user:fetch user: %w", err)
	}

	return user, nil
}

func (ur *UserRepo) FetchDetails(ctx context.Context, userID uint32) (*dto.UserDetails, error) {
	var details *dto.UserDetails
	q := ur.DB.WithContext(ctx).Model(&domain.User{}).
		Select(`
			id,
			name,
			phone,
			email
 		`).
		Where("users.id = ?", userID)

	if err := q.Take(&details).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("repository:user:fetch user details: %w", err)
	}

	return details, nil
}
