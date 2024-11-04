package service

import (
	"context"
	"net/http"

	"github.com/mdmoshiur/example-go/domain"
)

// CDNSvc ...
type CDNSvc interface {
	Upload(ctx context.Context, criteria *domain.FileUploadCriteria) (*http.Response, error)
	Delete(r *http.Request) (*http.Response, error)
}
