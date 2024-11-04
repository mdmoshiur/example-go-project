package paginator

import (
	"net/http"
	"strconv"

	"github.com/mdmoshiur/example-go/internal/config"
	"gorm.io/gorm"
)

// Page ...
type Page struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

// Paginate ...
func (p *Page) Paginate(db *gorm.DB) *gorm.DB {
	// set total row count
	db.Count(&p.Total)

	offset := (p.Page - 1) * p.PageSize
	return db.Offset(offset).Limit(p.PageSize)
}

// PaginateWithoutCount paginates for find() method where above paginate method won't work
func (p *Page) PaginateWithoutCount(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.PageSize
	return db.Offset(offset).Limit(p.PageSize)
}

// NewPager is the factory function a new page.
func NewPager(r *http.Request) *Page {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	switch {
	case pageSize < 1:
		pageSize = config.App().PaginationPageSize
	case pageSize > 200:
		pageSize = 200
	}

	return &Page{
		Page:     page,
		PageSize: pageSize,
	}
}
