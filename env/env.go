package env

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	errInvalidVar  = "env: %s is not set or empty"
	errInvalidVars = "env: found invalid environment variables"
	boolIsTrueRgx  = "true|True|TRUE"
)

func Check(c map[string]string) (err error) {
	errd := false
	msg := errInvalidVars

	for k, t := range c {
		switch t {
		case "string":
			_, err = S(k)
		case "bool":
			_, err = B(k)
		case "int":
			_, err = I(k)
		case "int32":
			_, err = I32(k)
		case "int64":
			_, err = I64(k)
		case "float32":
			_, err = F32(k)
		case "float64":
			_, err = F64(k)
		}

		if err != nil {
			errd = true
			msg = msg + "\n  * " + err.Error()
		}
	}

	if errd {
		return errors.New(msg)
	}

	return nil
}

func S(k string) (v string, err error) {
	v = os.Getenv(k)

	if v == "" {
		err = errors.New(fmt.Sprintf(errInvalidVar, k))
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
