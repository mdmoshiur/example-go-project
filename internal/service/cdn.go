package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mdmoshiur/example-go/domain"
	"github.com/mdmoshiur/example-go/internal/config"
)

// CDN ...
type CDN struct {
	Host      string
	Token     string
	Directory string
	Timeout   time.Duration
	Client    HTTPClient
}

// NewCDN ...
func NewCDN(client HTTPClient, cfg config.CDNCfg) *CDN {
	return &CDN{
		Host:      cfg.Host,
		Token:     cfg.Token,
		Directory: cfg.Directory,
		Timeout:   cfg.Timeout,
		Client:    client,
	}
}

// Upload ...
func (c *CDN) Upload(ctx context.Context, ctr *domain.FileUploadCriteria) (*http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/uploads", c.Host)

	// Create a buffer to hold the multipart request body
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)

	// Create a form field for the file
	fileWriter, err := writer.CreateFormFile("file", ctr.FileName) // "file" is the field name
	if err != nil {
		return nil, err
	}

	// Copy the file content into the form field
	_, err = io.Copy(fileWriter, ctr.File)
	if err != nil {
		return nil, err
	}

	// Add extra form fields
	_ = writer.WriteField("name", ctr.Name)
	_ = writer.WriteField("dir", c.Directory)

	// Close the multipart writer to finalize the request body
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}

	// Set Header
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Token", c.Token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("service:cdn: do: %v", err.Error())
	}
	// defer resp.Body.Close()

	return resp, nil
}

// Delete ...
func (c *CDN) Delete(r *http.Request) (*http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/uploads/%s", c.Host, chi.URLParam(r, "*"))

	ctx, cancel := context.WithTimeout(r.Context(), c.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set Header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", c.Token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("service:cdn: do: %v", err.Error())
	}
	// defer resp.Body.Close()

	return resp, nil
}
