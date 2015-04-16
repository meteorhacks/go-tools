package env

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
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

func TestI(t *testing.T) {
	key := randstr()

	val, err := I(key)
	if err == nil {
		t.Error("should return error")
	} else if val != 0 {
		t.Error("should return empty value")
	}

	req := "123"
	var res int = 123

	os.Setenv(key, req)
	val, err = I(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}
}

func TestI32(t *testing.T) {
	key := randstr()

	val, err := I32(key)
	if err == nil {
		t.Error("should return error")
	} else if val != 0 {
		t.Error("should return empty value")
	}

	req := "123"
	var res int32 = 123

	os.Setenv(key, req)
	val, err = I32(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}
}

func TestI64(t *testing.T) {
	key := randstr()

	val, err := I64(key)
	if err == nil {
		t.Error("should return error")
	} else if val != 0 {
		t.Error("should return empty value")
	}

	req := "123"
	var res int64 = 123

	os.Setenv(key, req)
	val, err = I64(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}
}

func TestF32(t *testing.T) {
	key := randstr()

	val, err := F32(key)
	if err == nil {
		t.Error("should return error")
	} else if val != 0 {
		t.Error("should return empty value")
	}

	req := "12.34"
	var res float32 = 12.34

	os.Setenv(key, req)
	val, err = F32(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}
}

func TestF64(t *testing.T) {
	key := randstr()

	val, err := F64(key)
	if err == nil {
		t.Error("should return error")
	} else if val != 0 {
		t.Error("should return empty value")
	}

	req := "12.34"
	var res float64 = 12.34

	os.Setenv(key, req)
	val, err = F64(key)
	if err != nil {
		t.Error("should not return error")
	} else if val != res {
		t.Error("should return correct value")
	}
}

func TestSS(t *testing.T) {
	key := randstr()
	sep := ","

	val, err := SS(key, sep)
	if err == nil {
		t.Error("should return error")
	} else if val != nil {
		t.Error("should return empty value")
	}

	req := "a,b,c"
	res := []string{"a", "b", "c"}

	os.Setenv(key, req)
	val, err = SS(key, sep)
	if err != nil {
		t.Error("should not return error")
	} else if !reflect.DeepEqual(val, res) {
		t.Error("should return correct value")
	}
}

func randstr() string {
	n := rand.Int()
	return fmt.Sprintf("K%d", n)
}
