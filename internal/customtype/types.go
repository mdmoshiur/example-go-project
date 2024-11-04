package customtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type (
	SliceOfUint32 []uint32

	Object        map[string]interface{}
	SliceOfObject []Object
)

// Scan scan value into Json, implements sql.Scanner interface
func (obj *Object) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	result := Object{}
	err := json.Unmarshal(bytes, &result)
	*obj = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (obj Object) Value() (driver.Value, error) {
	if len(obj) == 0 {
		return nil, nil
	}

	jb, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Failed to marshal JSON value:", obj))
	}

	return jb, nil
}

// Scan scan value into Json, implements sql.Scanner interface
func (sob *SliceOfObject) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	result := []Object{}
	err := json.Unmarshal(bytes, &result)
	*sob = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (sob SliceOfObject) Value() (driver.Value, error) {
	if len(sob) == 0 {
		return nil, nil
	}

	jb, err := json.Marshal(sob)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Failed to marshal JSON value:", sob))
	}

	return jb, nil
}

// Scan scan value into Json, implements sql.Scanner interface
func (s *SliceOfUint32) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	result := []uint32{}
	err := json.Unmarshal(bytes, &result)
	*s = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (s SliceOfUint32) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	jb, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Failed to marshal JSON value:", s))
	}

	return jb, nil
}
