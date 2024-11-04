package migration

import (
	"github.com/mdmoshiur/example-go/domain"
)

// Models describe models list for database migration.
var Models []interface{}

func init() {
	// add models for migration
	Models = append(Models, &domain.Role{})
	Models = append(Models, &domain.User{})
	Models = append(Models, &domain.AuthToken{})
}
