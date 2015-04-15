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
		t.Error("should return empty string")
	}

	os.Setenv(key, "val")
	val, err = S(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != "val" {
		t.Error("should return correct string")
	}
}

func randstr() string {
	n := rand.Int()
	return fmt.Sprintf("K%d", n)
}
