package store

import (
	"context"
	"gorm.io/gorm"

	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/internal/logger"
	userrepo "github.com/mdmoshiur/example-go/user/repository"
)

type Store interface {
	Atomic(ctx context.Context, fn func(ds Store) error) error
	UserRepository() domain.UserRepository
}

type DataStore struct {
	db       *gorm.DB
	userRepo domain.UserRepository
}

func New(db *gorm.DB) Store {
	return &DataStore{
		db:       db,
		userRepo: userrepo.New(db),
	}
}

func (s *DataStore) Atomic(ctx context.Context, fn func(ds Store) error) (err error) {
	tx := s.db.WithContext(ctx).Begin()

	defer func() {
		if p := recover(); p != nil {
			if rbErr := tx.Rollback().Error; rbErr != nil {
				logger.Error(rbErr)
			}
			panic(p)
		}
		if err != nil {
			if rbErr := tx.Rollback().Error; rbErr != nil {
				logger.Error(rbErr)
			}
		} else {
			if comErr := tx.Commit().Error; comErr != nil {
				logger.Error(comErr)
			}
		}
	}()

	newStore := New(tx)
	err = fn(newStore)
	return
}

func (s *DataStore) UserRepository() domain.UserRepository {
	return s.userRepo
}
