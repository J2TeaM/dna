package dna

import (
	"fmt"
	"reflect"
	"strconv"
)

// FromCharCode returns char from charcode
func FromCharCode(code Int) String {
	return String(fmt.Sprintf("%c", code))
}

// ParseInt returns Int from String with a specific base
func ParseInt(s String, base Int) Int {
	result, _ := strconv.ParseInt(string(s), int(base), 0)
	return Int(result)
}

// Implement the map utility with anonymous function.
// It is rewritten from http://godoc.org/github.com/BurntSushi/ty/fun.
// The anonymous func has the form:  func(func(A) B, []A) []B.
// It maps the every element in array A to other element type B.
// The final result is an array of elements type B
//
// It does not check the values of anonymous function fn and iterated array xs for speed advantage
func Map(xs, fn interface{}) interface{} {
	vf, vxs := reflect.ValueOf(fn), reflect.ValueOf(xs)
	xsLen := vxs.Len()
	vys := reflect.MakeSlice(reflect.SliceOf(vf.Type().Out(0)), xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		if vf.Type().NumIn() == 2 {
			vy := vf.Call([]reflect.Value{vxs.Index(i), reflect.ValueOf(Int(i))})[0]
			vys.Index(i).Set(vy)
		} else {
			vy := vf.Call([]reflect.Value{vxs.Index(i)})[0]
			vys.Index(i).Set(vy)
		}
	}
	return vys.Interface()
}
