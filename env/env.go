package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	errEmptyVar    = "env: %s is not set or empty"
	errInvalidVars = "env: found invalid environment variables"
	boolIsTrueRgx  = "true|True|TRUE"
)

var (
	Delimiter = ","
)

func Get(spec interface{}) (err error) {
	v := reflect.ValueOf(spec).Elem()
	m := ""

	v.FieldByNameFunc(func(name string) bool {
		f := v.FieldByName(name)

		switch f.Interface().(type) {
		case string:
			if val, err := S(name); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.SetString(val)
			}

		case bool:
			if val, err := B(name); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.SetBool(val)
			}

		case int, int32, int64:
			if val, err := I64(name); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.SetInt(val)
			}

		case float32, float64:
			if val, err := F64(name); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.SetFloat(val)
			}

		case []string:
			if val, err := SS(name, Delimiter); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.Set(reflect.ValueOf(val))
			}

		case []bool:
			if val, err := SB(name, Delimiter); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.Set(reflect.ValueOf(val))
			}

		case []int:
			if val, err := SI(name, Delimiter); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.Set(reflect.ValueOf(val))
			}

		case []int32:
			if val, err := SI32(name, Delimiter); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.Set(reflect.ValueOf(val))
			}

		case []int64:
			if val, err := SI64(name, Delimiter); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.Set(reflect.ValueOf(val))
			}

		case []float32:
			if val, err := SF32(name, Delimiter); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.Set(reflect.ValueOf(val))
			}

		case []float64:
			if val, err := SF64(name, Delimiter); err != nil {
				m = m + "\n" + err.Error()
			} else {
				f.Set(reflect.ValueOf(val))
			}
		}

		return false
	})

	if m == "" {
		return nil
	} else {
		return errors.New(errInvalidVars + m)
	}
}

func S(k string) (v string, err error) {
	v = os.Getenv(k)

	if v == "" {
		err = errors.New(fmt.Sprintf(errEmptyVar, k))
		return "", err
	}

	return v, nil
}

func B(k string) (v bool, err error) {
	s, err := S(k)
	if err != nil {
		return false, err
	}

	return regexp.MatchString(boolIsTrueRgx, s)
}

func I(k string) (v int, err error) {
	s, err := S(k)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

func I32(k string) (v int32, err error) {
	s, err := S(k)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(i), nil
}

func I64(k string) (v int64, err error) {
	s, err := S(k)

	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(s, 10, 64)
}

func F32(k string) (v float32, err error) {
	s, err := S(k)
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}

	return float32(f), nil
}

func F64(k string) (v float64, err error) {
	s, err := S(k)

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(s, 64)
}

func SS(k string, d string) (v []string, err error) {
	s, err := S(k)
	if err != nil {
		return nil, err
	}

	v = strings.Split(s, d)
	return v, nil
}

func SB(k string, d string) (v []bool, err error) {
	s, err := SS(k, d)
	if err != nil {
		return nil, err
	}

	l := len(s)
	v = make([]bool, l)

	for i := 0; i < l; i++ {
		b, err := regexp.MatchString(boolIsTrueRgx, s[i])
		if err != nil {
			return nil, err
		}

		v[i] = b
	}

	return v, nil
}

func SI(k string, d string) (v []int, err error) {
	s, err := SS(k, d)
	if err != nil {
		return nil, err
	}

	l := len(s)
	v = make([]int, l)

	for i := 0; i < l; i++ {
		n, err := strconv.ParseInt(s[i], 10, 0)
		if err != nil {
			return nil, err
		}

		v[i] = int(n)
	}

	return v, nil
}

func SI32(k string, d string) (v []int32, err error) {
	s, err := SS(k, d)
	if err != nil {
		return nil, err
	}

	l := len(s)
	v = make([]int32, l)

	for i := 0; i < l; i++ {
		n, err := strconv.ParseInt(s[i], 10, 0)
		if err != nil {
			return nil, err
		}

		v[i] = int32(n)
	}

	return v, nil
}

func SI64(k string, d string) (v []int64, err error) {
	s, err := SS(k, d)
	if err != nil {
		return nil, err
	}

	l := len(s)
	v = make([]int64, l)

	for i := 0; i < l; i++ {
		n, err := strconv.ParseInt(s[i], 10, 0)
		if err != nil {
			return nil, err
		}

		v[i] = n
	}

	return v, nil
}

func SF32(k string, d string) (v []float32, err error) {
	s, err := SS(k, d)
	if err != nil {
		return nil, err
	}

	l := len(s)
	v = make([]float32, l)

	for i := 0; i < l; i++ {
		n, err := strconv.ParseFloat(s[i], 64)
		if err != nil {
			return nil, err
		}

		v[i] = float32(n)
	}

	return v, nil
}

func SF64(k string, d string) (v []float64, err error) {
	s, err := SS(k, d)
	if err != nil {
		return nil, err
	}

	l := len(s)
	v = make([]float64, l)

	for i := 0; i < l; i++ {
		n, err := strconv.ParseFloat(s[i], 64)
		if err != nil {
			return nil, err
		}

		v[i] = n
	}

	return v, nil
}
