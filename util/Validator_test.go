package util

import (
	"fmt"
	"testing"
)

const name = "a-name"
const age = 10
const ageStr = "10"

var validProvider = &provider{map[string]string{
	"name": name,
	"age":  ageStr,
}}

var emptyProvider = &provider{}

func TestValidate_String_Pass(t *testing.T) {
	v := NewValidator(validProvider)
	s := v.String("name")

	if s != name {
		t.Error("Did not return correct value")
	}
}

func TestValidate_StringO_Fail(t *testing.T) {
	v := NewValidator(emptyProvider)
	v.String("name")

	if _, err := v.Passed(); err != ErrMissingKey {
		t.Error("Should have set an error for missing key")
	}
}

func TestValidate_StringO_Pass(t *testing.T) {
	v := NewValidator(validProvider)
	d := "default-value"
	s := v.StringO("missing value", d)

	if s != d {
		t.Error("Should set default value")
	}

	if ok, _ := v.Passed(); !ok {
		t.Error("Should still pass validation for optionals")
	}
}

func TestValidate_String_Fail(t *testing.T) {
	v := NewValidator(emptyProvider)
	v.String("name")

	if _, err := v.Passed(); err != ErrMissingKey {
		t.Error("Should have set an error for missing key")
	}
}

func TestValidate_Int_Pass(t *testing.T) {
	v := NewValidator(validProvider)
	a := v.Int("age")

	if a != age {
		t.Error("Did not return correct value")
	}
}

func TestValidate_Int_Fail_missing(t *testing.T) {
	v := NewValidator(emptyProvider)
	v.Int("age")

	if _, err := v.Passed(); err != ErrMissingKey {
		t.Error("Should have set ErrMissingKey")
	}
}

func TestValidate_Int_Fail_cannot_convert(t *testing.T) {
	v := NewValidator(&provider{map[string]string{
		"age": "not-a-number",
	}})
	v.Int("age")

	if _, err := v.Passed(); err != ErrCannotConvertStringToInt {
		t.Error("Should have set ErrCannotConvertStringToInt")
	}
}

func TestValidate_IntO_missing(t *testing.T) {
	v := NewValidator(emptyProvider)
	d := -1
	age := v.IntO("age", d)

	if age != d {
		t.Error("Should return default value")
	}

	if ok, err := v.Passed(); !ok {
		fmt.Println("Err = ", err)
		t.Error("should still pass validation")
	}
}

func TestValidate_IntO_cannot_convert(t *testing.T) {
	v := NewValidator(&provider{map[string]string{
		"age": "not-a-number",
	}})
	d := -1
	age := v.IntO("age", d)

	if age != d {
		t.Error("Should set default value")
	}

	if ok, _ := v.Passed(); !ok {
		t.Error("Should still pass validation for optionals")
	}
}

type provider struct {
	vals map[string]string
}

func (tp *provider) FormValue(key string) string {
	return tp.vals[key]
}
