package env

import (
	"errors"
	"fmt"
	"os"
)

const (
	NO_ERR = ""
	FORMAT = "env: invalid %s (%s) - %s"
)

func Check(c map[string]string) (err error) {
	// TODO
	return nil
}

func S(k string) (v string, err error) {
	v = os.Getenv(k)

	if v == "" {
		m := "value not set or empty"
		err = errors.New(fmt.Sprintf(FORMAT, k, v, m))
		return "", err
	}

	return v, nil
}

func B(k string) (v bool, err error) {
	// TODO
	return v, nil
}

func I(k string) (v int, err error) {
	// TODO
	return v, nil
}

func I32(k string) (v int32, err error) {
	// TODO
	return v, nil
}

func I64(k string) (v int64, err error) {
	// TODO
	return v, nil
}

func F32(k string) (v float32, err error) {
	// TODO
	return v, nil
}

func F64(k string) (v float64, err error) {
	// TODO
	return v, nil
}

func SS(k string, d string) (v []string, err error) {
	// TODO
	return v, nil
}

func SB(k string, d string) (v []bool, err error) {
	// TODO
	return v, nil
}

func SI(k string, d string) (v []int, err error) {
	// TODO
	return v, nil
}

func SI32(k string, d string) (v []int32, err error) {
	// TODO
	return v, nil
}

func SI64(k string, d string) (v []int64, err error) {
	// TODO
	return v, nil
}

func SF32(k string, d string) (v []float32, err error) {
	// TODO
	return v, nil
}

func SF64(k string, d string) (v []float64, err error) {
	// TODO
	return v, nil
}
