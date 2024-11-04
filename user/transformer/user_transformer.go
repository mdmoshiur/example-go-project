package transformer

import (
	"github.com/mdmoshiur/example-go/domain"
)

// UserTransformer represents the user transformer.
type UserTransformer struct{}

// New is the factory function for the user transformer.
func New() domain.UserTransformer {
	return &UserTransformer{}
}
