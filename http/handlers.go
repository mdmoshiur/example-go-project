package http

import (
	"gorm.io/gorm"

	"github.com/mdmoshiur/example-go/internal/cache"
	"github.com/mdmoshiur/example-go/store"

	userdeliveryhttp "github.com/mdmoshiur/example-go/user/delivery/http"
	usertransformer "github.com/mdmoshiur/example-go/user/transformer"
	userusecase "github.com/mdmoshiur/example-go/user/usecase"
)

type Handler struct {
	UserHandler *userdeliveryhttp.UserHandler
}

func RegisterHandlers(db *gorm.DB, cs cache.Cacher) *Handler {
	// initialize datastore
	s := store.New(db)

	// initialize necessary transformers
	ut := usertransformer.New()

	// initialize necessary usecases
	uc := userusecase.New(s, cs)

	// initialize necessary http handlers
	return &Handler{
		UserHandler: userdeliveryhttp.New(uc, ut),
	}
}
