package domain

import "mime/multipart"

// FileUploadCriteria represents file upload criteria
type FileUploadCriteria struct {
	File     multipart.File `json:"file"`
	FileName string         `json:"file_name"`
	Name     string         `json:"name"`
}
