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
	errfmt  = "env: invalid %s (%s) - %s"
	boolrgx = "true|True|TRUE"
)

func Check(c map[string]string) (err error) {
	// TODO
	return nil
}

func S(k string) (v string, err error) {
	v = os.Getenv(k)

	if v == "" {
		m := "value not set or empty"
		err = errors.New(fmt.Sprintf(errfmt, k, v, m))
		return "", err
	}

	return v, nil
}

func B(k string) (v bool, err error) {
	s, err := S(k)
	if err != nil {
		return false, err
	}

	return regexp.MatchString(boolrgx, s)
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
		b, err := regexp.MatchString(boolrgx, s[i])
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
