package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/internal/responder"
	"github.com/mdmoshiur/example-go/internal/validation"
)

// UserHandler represents user HTTP/JSON handler.
type UserHandler struct {
	UserUseCase     domain.UserUseCase
	UserTransformer domain.UserTransformer
}

type (
	// UserReq represents create user request
	UserReq struct {
		IsCreate    bool
		Name        string  `json:"name" validate:"required,min=3,max=128"`
		Email       string  `json:"email" validate:"required"`
		PhoneNumber *string `json:"phone_number"`
		Password    string  `json:"password" validate:"required"`
		Status      *uint8  `json:"status"`
	}
)

// New will initialize the user handler.
func New(uc domain.UserUseCase, ut domain.UserTransformer) *UserHandler {
	return &UserHandler{
		UserUseCase:     uc,
		UserTransformer: ut,
	}
}

// StoreOrUpdate creates/updates a user
func (u *UserHandler) StoreOrUpdate(w http.ResponseWriter, r *http.Request) {
	req := &UserReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error(err)
		responder.BadReqErr(w, err)
		return
	}

	err := validation.Validator.Struct(req)
	if err != nil {
		responder.ValidatorValidationErr(w, err)
		return
	}

	isCreate := true
	userID := 0
	if r.Method == http.MethodPut {
		isCreate = false
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			responder.BadReqErr(w, err)
			return
		}
		userID = id
	}
	req.IsCreate = isCreate

	user := &domain.User{
		ID:       uint32(userID),
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.PhoneNumber,
		Password: req.Password,
		Status:   req.Status,
	}

	err = u.UserUseCase.StoreOrUpdate(r.Context(), user)
	if err != nil {
		if errors.Is(err, domain.ErrUserDuplicateEmail) {
			responder.ValidationErr(w, validation.Errors{
				"email": err.Error(),
			})
			return
		} else if errors.Is(err, domain.ErrUserDuplicatePhone) {
			responder.ValidationErr(w, validation.Errors{
				"phone_number": err.Error(),
			})
			return
		}

		responder.InternalServerErr(w, err)
		logger.Error(err)
		return
	}

	msg := "user updated successfully!"
	if isCreate {
		msg = "user created successfully!"
	}
	(&responder.Response{
		Status:  http.StatusOK,
		Message: msg,
	}).Render(w)
}

// Login logged in a user
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctr := &domain.LoginCriteria{}
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(ctr); err != nil {
		logger.Error(err)
		responder.BadReqErr(w, err)
		return
	}

	token, err := u.UserUseCase.Login(ctx, ctr)
	if err != nil {
		if errors.Is(err, domain.ErrWrongPassword) {
			(&responder.Response{
				Status:  http.StatusUnauthorized,
				Message: "wrong password given, please recheck your password",
				Error: validation.Errors{
					"password": err.Error(),
				},
			}).Render(w)
			return
		} else if errors.Is(err, domain.ErrWrongEmail) {
			(&responder.Response{
				Status:  http.StatusNotFound,
				Message: "email not found, please recheck your email",
				Error: validation.Errors{
					"email": err.Error(),
				},
			}).Render(w)
			return
		} else if errors.Is(err, domain.ErrDeactivatedUser) {
			(&responder.Response{
				Status:  http.StatusForbidden,
				Message: "admin deactivated your account",
				Error:   err.Error(),
			}).Render(w)
			return
		}

		responder.InternalServerErr(w, err)
		return
	}

	(&responder.Response{
		Status:  http.StatusOK,
		Message: "user logged in successfully",
		Data: map[string]string{
			"token": token,
		},
	}).Render(w)
}

// Logout logged out a user
func (u *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	err := u.UserUseCase.Logout(r.Context(), r.Header.Get("Authorization")[7:])
	if err != nil {
		responder.InternalServerErr(w, err)
		return
	}

	(&responder.Response{
		Status:  http.StatusOK,
		Message: "user logged out successfully",
	}).Render(w)
}
