package env

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestS(t *testing.T) {
	key := randstr()

	val, err := S(key)
	if err == nil {
		t.Error("should return error")
	} else if val != "" {
		t.Error("should return empty value")
	}

	req := "str"
	res := "str"

	os.Setenv(key, req)
	val, err = S(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}
}

func TestB(t *testing.T) {
	key := randstr()

	val, err := B(key)
	if err == nil {
		t.Error("should return error")
	} else if val != false {
		t.Error("should return empty value")
	}

	req := "true"
	res := true

	os.Setenv(key, req)
	val, err = B(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}

	req = "True"
	res = true

	os.Setenv(key, req)
	val, err = B(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}

	req = "TRUE"
	res = true

	os.Setenv(key, req)
	val, err = B(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}
}

func randstr() string {
	n := rand.Int()
	return fmt.Sprintf("K%d", n)
}
