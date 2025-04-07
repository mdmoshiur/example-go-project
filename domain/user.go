package domain

import (
	"context"
	"errors"
	"time"
)

// available User errors
var (
	// ErrUserNotFound occurs if user not found on database
	ErrUserNotFound       = errors.New("user not found")
	ErrUserTokenNotFound  = errors.New("token not found")
	ErrUserDuplicateEmail = errors.New("email already exists")
	ErrUserDuplicatePhone = errors.New("phone number already exists")
	ErrWrongPassword      = errors.New("wrong password")
	ErrWrongEmail         = errors.New("wrong email")
	ErrDeactivatedUser    = errors.New("deactivated user")
)

// user's status constants
const (
	UserStatusDeactivated = iota
	UserStatusActive
	UserStatusSuspended
)

type (
	// User represents a User.
	User struct {
		ID        uint32    `json:"id" gorm:"primaryKey"`
		Name      string    `json:"name" gorm:"type:varchar(128) not null"`
		Email     string    `json:"email" gorm:"type:varchar(64) not null;uniqueIndex"`
		Phone     *string   `json:"phone" gorm:"type:varchar(16);uniqueIndex"`
		Status    *uint8    `json:"status" gorm:"type:smallint;not null; default: 1;index"`
		RoleId    *uint8    `json:"role_id" gorm:"type:smallint;index"`
		Password  string    `json:"password" gorm:"type:varchar(128) not null"`
		CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) not null;default:current_timestamp"`
		UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) not null;default:current_timestamp"`
	}

	// Role ...
	Role struct {
		ID          uint32       `json:"id" gorm:"type:smallint"`
		Name        string       `json:"name" gorm:"type:varchar(64) not null;unique"`
		Permissions *Permissions `json:"permissions" gorm:"type:json"`
		Users       []User       `json:"users"`

		TimeStamp
	}

	// Permissions ...
	Permissions struct {
		User struct {
			Create bool `json:"create"`
			Update bool `json:"update"`
		} `json:"user"`
		Question struct {
			Create   bool `json:"create"`
			Validate bool `json:"validate"`
		} `json:"question"`
	}
)

// AuthToken represents authenticated users tokens
type AuthToken struct {
	TokenID   string    `json:"token_id" gorm:"type:varchar(64) not null;primaryKey"`
	UserID    uint32    `json:"user_id" gorm:"not null;index"`
	Revoked   bool      `json:"revoked" gorm:"not null;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) not null;default:current_timestamp"`
}

type (
	// LoginCriteria used when logged-in user
	LoginCriteria struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

// UserRepository represents User's repository contract
type UserRepository interface {
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
	Store(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	StoreAuthToken(ctx context.Context, authToken *AuthToken) error
	RevokeAuthToken(ctx context.Context, tokenID string) error
	RevokeAllAuthToken(ctx context.Context, userID uint32) error
}

// UserUseCase represents User's use case contract
type UserUseCase interface {
	Login(ctx context.Context, ctr *LoginCriteria) (string, error)
	Logout(ctx context.Context, token string) error
	StoreOrUpdate(ctx context.Context, user *User) error
}

type UserTransformer interface{}
