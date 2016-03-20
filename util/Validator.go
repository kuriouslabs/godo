package util

import (
	"errors"
	"strconv"
)

var (
	ErrMissingKey               = errors.New("Missing Key")
	ErrCannotConvertStringToInt = errors.New("Cannot convert string to int")
)

type ValidatorProviding interface {
	FormValue(string) string
}

type Validator struct {
	err error
	p   ValidatorProviding
}

func NewValidator(p ValidatorProviding) *Validator {
	return &Validator{
		p: p,
	}
}

func (v *Validator) String(key string) string {
	if v.err != nil {
		return ""
	}
	s := v.p.FormValue(key)
	if s == "" {
		v.err = ErrMissingKey
		return ""
	}
	return s
}

func (v *Validator) StringO(key string, defaultStr string) string {
	if v.err != nil {
		return ""
	}

	s := v.p.FormValue(key)
	if s == "" {
		return defaultStr
	}
	return s
}

// Int attempts to extract an int for the given key
// defaults to 0 if not present
func (v *Validator) Int(key string) int {
	d := 0
	return v.convert(key, d, func(s string) interface{} {
		i, err := strconv.Atoi(s)
		if err != nil {
			v.err = ErrCannotConvertStringToInt
			return d
		}
		return i
	}).(int)
}

func (v *Validator) IntO(key string, defaultInt int) int {
	d := defaultInt
	return v.convertO(key, d, func(s string) interface{} {
		i, err := strconv.Atoi(s)
		if err != nil {
			return defaultInt
		}
		return i
	}).(int)
}

type convertAction func(val string) interface{}

func (v *Validator) convert(key string, defaultValue interface{}, f convertAction) interface{} {
	c := converter{
		validator:    v,
		key:          key,
		defaultValue: defaultValue,
		isOptional:   false,
	}
	return c.convert(f)
}

func (v *Validator) convertO(key string, defaultValue interface{}, f convertAction) interface{} {
	c := converter{
		validator:    v,
		key:          key,
		defaultValue: defaultValue,
		isOptional:   true,
	}
	return c.convert(f)
}

type converter struct {
	validator    *Validator
	key          string
	defaultValue interface{}
	isOptional   bool
}

func (c converter) convert(f convertAction) interface{} {
	var s string
	if c.isOptional {
		s = c.validator.StringO(c.key, "")
	} else {
		s = c.validator.String(c.key)
	}

	if c.validator.err != nil {
		return c.defaultValue
	}

	return f(s)
}

func (v *Validator) Passed() (bool, error) {
	if v.err == nil {
		return true, nil
	}
	return false, v.err
}
