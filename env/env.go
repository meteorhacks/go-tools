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

func S(k string) (v string, err error) {
	v = os.Getenv(k)

	if res := validStr(v); res != NO_ERR {
		msg := fmt.Sprintf(FORMAT, k, v, res)
		return "", errors.New(msg)
	}

	return v, nil
}

func validStr(v string) (res string) {
	if v == "" {
		return "value not set or empty"
	}

	return NO_ERR
}
