package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// Scan scan value into Json, implements sql.Scanner interface
func (p *Permissions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	result := Permissions{}
	err := json.Unmarshal(bytes, &result)
	*p = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (p Permissions) Value() (driver.Value, error) {
	if p == (Permissions{}) {
		return nil, nil
	}

	jb, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Failed to marshal JSON value:", p))
	}

	return jb, nil
}
