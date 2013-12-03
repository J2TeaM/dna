package dna

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Redefine new IntArray which is from []Int
type IntArray []Int

/*
Map method for IntArray.
The returned array has the same or different type comparing to the old one.

Notice: The real return value is interface{}. So it has to be converted into IntArray Type
	var x := dna.IntArray{1,3,5,6}
	x = dna.IntArray(x.Map(func(value ,index dna.Int) dna.Int{
		return value*index
	}).([]dna.Int))
*/
func (a IntArray) Map(fn interface{}) interface{} {
	return Map(a, fn)
}

// ForEach loops through every element in array. It does not return any value
func (a IntArray) ForEach(fn func(v, i Int)) {
	for index, value := range a {
		fn(value, Int(index))
	}
}

// Filter loops through every element in array and returns only elements which make
// input func true.
func (a IntArray) Filter(fn func(v, i Int) Bool) IntArray {
	var result = IntArray{}
	for index, value := range a {
		if fn(value, Int(index)) {
			result = append(result, value)
		}
	}
	return result
}

// Reduce returns aggregated result. **** Need to rewrite
func (a IntArray) Reduce(fn func(previousValue, currentValue, index Int) Int) Int {
	var t Int
	t = 0
	for i, v := range a {
		t = fn(t, v, Int(i))
	}
	return t
}

// Length returns a length of the array
func (a IntArray) Length() Int {
	return Int(len(a))
}

// Push inserts new value to the existing array
func (a *IntArray) Push(value Int) {
	slice := append(*a, value)
	*a = slice
}

// Pop removes the last element of the existing array
func (a *IntArray) Pop() {
	slice := *a
	slice = slice[0 : len(slice)-1]
	*a = slice
}

// Shift removes the first element of the existing array
func (a *IntArray) Shift() {
	slice := *a
	slice = slice[1:len(slice)]
	*a = slice
}

// Sort gets the existing array in order
func (a *IntArray) Sort() {
	tmp := make(sort.IntSlice, len(*a))
	for index, v := range *a {
		tmp[index] = int(v)
	}
	sort.Sort(tmp)
	result := make(IntArray, len(*a))
	for index := range *a {
		result[index] = Int(tmp[index])
	}
	*a = result
}

// Reverse returns a new reversed array
func (a IntArray) Reverse() IntArray {
	length := len(a)
	tmp := make(IntArray, length)
	for i, v := range a {
		tmp[length-i-1] = v
	}
	return tmp
}

// IndexOf returns an index of an element in an array
func (a IntArray) IndexOf(n Int) Int {
	for i, v := range a {
		if v == n {
			return Int(i)
		}
	}
	return -1
}

// Join returns new string from an array
func (a IntArray) Join(sep String) String {
	sarr := make([]string, len(a))
	for i, v := range a {
		sarr[i] = fmt.Sprint(v)
	}
	return String(strings.Join(sarr, fmt.Sprint(sep)))
}

// ToString returns  a new string by simply joining all elements in array with no delimiter
func (a IntArray) ToString() String {
	return a.Join("")
}

// Concat returns a new concatenated array
func (a IntArray) Concat(arr IntArray) IntArray {
	return append(a, arr...)
}

// Unique returns unique IntArray
func (a IntArray) Unique() IntArray {
	var tmp IntArray = IntArray{}
	for _, v := range a {
		if tmp.IndexOf(v) == -1 {
			tmp.Push(v)
		}
	}
	return tmp
}

// Value implements the Valuer interface in database/sql/driver package.
func (a IntArray) Value() (driver.Value, error) {
	return driver.Value(string(String(fmt.Sprintf("%#v", a)).Replace("dna.IntArray", ""))), nil
}

// parseStringArray returns dna.IntArray-typed array from postgresql-array-typed int
// Ex: {123,456} => dna.IntArray{123, 456}
func ParseIntArray(str String) IntArray {
	if str.Match(`^{[0-9,]+}$`) == true {
		return str.Replace("{", "").Replace("}", "").Split(",").ToIntArray()
	} else {
		panic("Int array from sql is not in correct format!")
	}
}

// Scan implements the Scanner interface in database/sql package.
func (a *IntArray) Scan(src interface{}) error {
	var source IntArray
	switch src.(type) {
	case string:
		source = ParseIntArray(String(string(src.(string))))
		Logv("adasdas")
	case []byte:
		source = ParseIntArray(String(string(src.([]byte))))
	case nil:
		source = IntArray{}
	default:
		return errors.New("Incompatible type for dna.IntArray type")
	}
	*a = source
	return nil
}
