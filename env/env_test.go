package env

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type Spec struct {
		ENV_S    string
		ENV_B    bool
		ENV_I    int
		ENV_I32  int32
		ENV_I64  int64
		ENV_F32  float32
		ENV_F64  float64
		ENV_SS   []string
		ENV_SB   []bool
		ENV_SI   []int
		ENV_SI32 []int32
		ENV_SI64 []int64
		ENV_SF32 []float32
		ENV_SF64 []float64
	}

	spec := Spec{}

	err := Get(&spec)
	if err == nil {
		t.Error("should return error")
	}

	expc := Spec{
		ENV_S:    "str",
		ENV_B:    true,
		ENV_I:    1,
		ENV_I32:  1,
		ENV_I64:  1,
		ENV_F32:  0.1,
		ENV_F64:  0.1,
		ENV_SS:   []string{"str"},
		ENV_SB:   []bool{true},
		ENV_SI:   []int{1},
		ENV_SI32: []int32{1},
		ENV_SI64: []int64{1},
		ENV_SF32: []float32{0.1},
		ENV_SF64: []float64{0.1},
	}

	os.Setenv("ENV_S", "str")
	os.Setenv("ENV_B", "true")
	os.Setenv("ENV_I", "1")
	os.Setenv("ENV_I32", "1")
	os.Setenv("ENV_I64", "1")
	os.Setenv("ENV_F32", "0.1")
	os.Setenv("ENV_F64", "0.1")
	os.Setenv("ENV_SS", "str")
	os.Setenv("ENV_SB", "true")
	os.Setenv("ENV_SI", "1")
	os.Setenv("ENV_SI32", "1")
	os.Setenv("ENV_SI64", "1")
	os.Setenv("ENV_SF32", "0.1")
	os.Setenv("ENV_SF64", "0.1")

	if err = Get(&spec); err != nil {
		t.Error("should not return error")
	} else if !reflect.DeepEqual(spec, expc) {
		t.Error("should return correct values")
	}
}

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

func TestSB(t *testing.T) {
	key := randstr()
	sep := ","

	val, err := SB(key, sep)
	if err == nil {
		t.Error("should return error")
	} else if val != nil {
		t.Error("should return empty value")
	}

	req := "true,True,TRUE,something-else"
	res := []bool{true, true, true, false}

	os.Setenv(key, req)
	val, err = SB(key, sep)
	if err != nil {
		t.Error("should not return error")
	} else if !reflect.DeepEqual(val, res) {
		fmt.Println(val, res)
		t.Error("should return correct value")
	}
}

func TestSI(t *testing.T) {
	key := randstr()
	sep := ","

	val, err := SI(key, sep)
	if err == nil {
		t.Error("should return error")
	} else if val != nil {
		t.Error("should return empty value")
	}

	req := "1,2,3"
	res := []int{1, 2, 3}

	os.Setenv(key, req)
	val, err = SI(key, sep)
	if err != nil {
		t.Error("should not return error")
	} else if !reflect.DeepEqual(val, res) {
		t.Error("should return correct value")
	}
}

func TestSI32(t *testing.T) {
	key := randstr()
	sep := ","

	val, err := SI32(key, sep)
	if err == nil {
		t.Error("should return error")
	} else if val != nil {
		t.Error("should return empty value")
	}

	req := "1,2,3"
	res := []int32{1, 2, 3}

	os.Setenv(key, req)
	val, err = SI32(key, sep)
	if err != nil {
		t.Error("should not return error")
	} else if !reflect.DeepEqual(val, res) {
		t.Error("should return correct value")
	}
}

func TestSI64(t *testing.T) {
	key := randstr()
	sep := ","

	val, err := SI64(key, sep)
	if err == nil {
		t.Error("should return error")
	} else if val != nil {
		t.Error("should return empty value")
	}

	req := "1,2,3"
	res := []int64{1, 2, 3}

	os.Setenv(key, req)
	val, err = SI64(key, sep)
	if err != nil {
		t.Error("should not return error")
	} else if !reflect.DeepEqual(val, res) {
		t.Error("should return correct value")
	}
}

func TestSF32(t *testing.T) {
	key := randstr()
	sep := ","

	val, err := SF32(key, sep)
	if err == nil {
		t.Error("should return error")
	} else if val != nil {
		t.Error("should return empty value")
	}

	req := "1.2,2.3,3.45"
	res := []float32{1.2, 2.3, 3.45}

	os.Setenv(key, req)
	val, err = SF32(key, sep)
	if err != nil {
		t.Error("should not return error")
	} else if !reflect.DeepEqual(val, res) {
		t.Error("should return correct value")
	}
}

func TestSF64(t *testing.T) {
	key := randstr()
	sep := ","

	val, err := SF64(key, sep)
	if err == nil {
		t.Error("should return error")
	} else if val != nil {
		t.Error("should return empty value")
	}

	req := "1.2,2.3,3.45"
	res := []float64{1.2, 2.3, 3.45}

	os.Setenv(key, req)
	val, err = SF64(key, sep)
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
