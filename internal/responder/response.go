package responder

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/internal/config"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/internal/paginator"
	"github.com/mdmoshiur/example-go/internal/validation"
)

// Response ...
type Response struct {
	Status     int             `json:"-"`
	Message    interface{}     `json:"message,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Error      interface{}     `json:"error,omitempty"` // this field will be omitted from user response body base on logger level
	Pagination *paginator.Page `json:"pagination,omitempty"`
}

// Render ...
func (r *Response) Render(w http.ResponseWriter) {
	if !config.App().Verbose {
		r.Error = nil // if verbose set to false then remove the error from public response
	}
	b, err := json.Marshal(r)
	if err != nil {
		logger.Error(err)
		internalErr(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Status != 0 {
		w.WriteHeader(r.Status)
	}
	_, err = w.Write(b)
	if err != nil {
		logger.Error(err)
	}
}

func internalErr(w http.ResponseWriter, e error) {
	r := &Response{
		Status:  http.StatusInternalServerError,
		Message: domain.ErrSomethingWentWrong.Error(),
		Error:   e.Error(),
	}
	if !config.App().Verbose {
		r.Error = nil // if verbose set to false then remove the error from public response
	}
	b, err := json.Marshal(r)
	if err != nil {
		logger.Error(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Status != 0 {
		w.WriteHeader(r.Status)
	}
	_, err = w.Write(b)
	if err != nil {
		logger.Error(err)
	}
}

func AccessDeniedErr(w http.ResponseWriter) {
	(&Response{
		Status:  http.StatusUnauthorized,
		Message: "access denied",
		Error:   "invalid token",
	}).Render(w)
}

func BadReqErr(w http.ResponseWriter, e error) {
	(&Response{
		Status:  http.StatusBadRequest,
		Message: domain.ErrBadRequest.Error(),
		Error:   e.Error(),
	}).Render(w)
}

func ValidationErr(w http.ResponseWriter, e validation.Errors) {
	(&Response{
		Status:  http.StatusUnprocessableEntity,
		Message: domain.ErrValidation.Error(),
		Error:   e,
	}).Render(w)
}

func ValidatorValidationErr(w http.ResponseWriter, e error) {
	errs := e.(validator.ValidationErrors)
	(&Response{
		Status:  http.StatusUnprocessableEntity,
		Message: domain.ErrValidation.Error(),
		Error:   validation.RemovePrefixStructName(errs.Translate(validation.Translator)),
	}).Render(w)
}

func NotFoundErr(w http.ResponseWriter, e error) {
	(&Response{
		Status:  http.StatusNotFound,
		Message: e.Error(),
		Error:   e.Error(),
	}).Render(w)
}

func InternalServerErr(w http.ResponseWriter, e error) {
	(&Response{
		Status:  http.StatusInternalServerError,
		Message: domain.ErrSomethingWentWrong.Error(),
		Error:   e.Error(),
	}).Render(w)
}

func CDNResponseRender(w http.ResponseWriter, resp *http.Response) {
	defer resp.Body.Close()

	// Set response header
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	// Write the response body to the new response
	_, err := io.Copy(w, resp.Body)
	if err != nil {
		logger.Error(err)
	}
}
