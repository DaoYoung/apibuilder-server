package rest

import "strconv"

type Value struct {
	Val string
}

func (v *Value) Defined() bool {
	return v.Val != ""
}

func (v *Value) Default(value string) *Value {
	if v.Val == "" {
		v.Val = value
	}
	return v
}

func (v *Value) String() string {
	return v.Val
}

func (v *Value) Int() int {
	num, _ := strconv.Atoi(v.Val)

	return num
}

func (v *Value) IntP() *int {
	if v.Val == "" {
		return nil
	}

	num, _ := strconv.Atoi(v.Val)
	return &num
}

func (v *Value) Int64() int64 {
	num, _ := strconv.ParseInt(v.Val, 10, 64)

	return num
}

func (v *Value) Int64P() *int64 {
	if v.Val == "" {
		return nil
	}

	num, _ := strconv.ParseInt(v.Val, 10, 64)
	return &num
}

func (v *Value) Bool() bool {
	b, _ := strconv.ParseBool(v.Val)

	return b
}

func (v *Value) BoolP() *bool {
	if v.Val == "" {
		return nil
	}

	b, _ := strconv.ParseBool(v.Val)
	return &b
}

func (v *Value) Float() float64 {
	num, _ := strconv.ParseFloat(v.Val, 64)

	return num
}

func (v *Value) FloatP() *float64 {
	if v.Val == "" {
		return nil
	}

	num, _ := strconv.ParseFloat(v.Val, 64)
	return &num
}