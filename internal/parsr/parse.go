package parsr

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mdmoshiur/example-go/internal/logger"
)

func ParseQueryStr(r *http.Request, key string, required bool) (*string, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok
	if !empty && len(vs) == 0 {
		empty = true
	}

	if empty {
		if !required {
			return nil, nil
		}

		return nil, fmt.Errorf("%s is required", key)
	}

	v := vs[0]
	return &v, nil
}

func ParseQueryInt(r *http.Request, key string, required bool) (*int, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok

	if !empty && (len(vs) == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	v := vs[0]

	n, err := strconv.Atoi(v)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("can't parse the %s", key)
	}

	return &n, nil
}

func ParseQueryIntSlice(r *http.Request, key string, required bool) ([]int, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok
	l := len(vs)
	if !empty && (l == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	values := make([]int, 0, l)
	for _, v := range vs {
		n, err := strconv.Atoi(v)
		if err != nil {
			logger.Error(err)
			return nil, fmt.Errorf("can't parse the %s", key)
		}
		values = append(values, n)
	}

	return values, nil
}

func ParseQueryInt32(r *http.Request, key string, required bool) (*int32, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok

	if !empty && (len(vs) == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	v := vs[0]

	in64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("can't parse the %s", key)
	}
	in32 := int32(in64)
	return &in32, nil
}

func ParseQueryInt8(r *http.Request, key string, required bool) (*int8, error) {
	vs, ok := r.URL.Query()[key]
	empty := !ok

	if !empty && (len(vs) == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	v := vs[0]

	in64, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("can't parse the %s", key)
	}
	in8 := int8(in64)
	return &in8, nil
}

func ParseQueryUint32(r *http.Request, key string, required bool) (*uint32, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok

	if !empty && (len(vs) == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	v := vs[0]

	uin64, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("can't parse the %s", key)
	}
	un32 := uint32(uin64)
	return &un32, nil
}

func ParseQueryUint8(r *http.Request, key string, required bool) (*uint8, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok

	if !empty && (len(vs) == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	v := vs[0]

	uin64, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("can't parse the %s", key)
	}
	uin8 := uint8(uin64)
	return &uin8, nil
}

func ParseQueryFloat64(r *http.Request, key string, required bool) (*float64, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok
	if !empty && (len(vs) == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	v := vs[0]

	n, err := strconv.ParseFloat(v, 64)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("can't parse the %s", key)
	}

	return &n, nil
}

func ParseQueryBool(r *http.Request, key string, required bool) (*bool, error) {
	vs, ok := r.URL.Query()[key]

	empty := !ok

	if !empty && (len(vs) == 0 || vs[0] == "") {
		empty = true
	}

	if empty {
		if required {
			return nil, fmt.Errorf("%s is required", key)
		}

		return nil, nil
	}

	v := vs[0]

	b, err := strconv.ParseBool(v)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("can't parse the %s", key)
	}

	return &b, nil
}

func URLParam(r *http.Request, param string) string {
	val := chi.URLParam(r, param)
	return val
}
