package dna

import (
	"fmt"
	"testing"
)

func TestIntArray_Map(t *testing.T) {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	y := StringArray(x.Map(func(v, i Int) String {
		return v.ToString()
	}).([]String))
	z := StringArray{"1", "2", "3", "4", "5"}
	for i, v := range y {
		if v != z[i] {
			t.Errorf("%v (IntArray) cannot be converted to StringArray", x)
		}
	}
}

func TestIntArray_PopPushShift(t *testing.T) {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	y := IntArray{2, 3, 4, 6}
	x.Pop()
	x.Push(6)
	x.Shift()
	for i, v := range y {
		if v != x[i] {
			t.Errorf("%v (IntArray) cannot be converted to StringArray", x)
		}
	}
}

// Example cases

func ExampleIntArray() {
	x := IntArray{1, 2, 3}                    // literal
	var y IntArray = IntArray([]Int{1, 2, 3}) // type conversion
	var z IntArray = []Int{1, 2, 3}
	Logv(x)
	Logv(y)
	Logv(z)
	Logv(z.Join(""))
	// Output: dna.IntArray{1, 2, 3}
	// dna.IntArray{1, 2, 3}
	// dna.IntArray{1, 2, 3}
	// "123"
}

func ExampleIntArray_Map() {
	// Convert IntArray to StringArray
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	y := StringArray(x.Map(func(v, i Int) String {
		return v.ToString()
	}).([]String))
	Logv(y)
	// Output: dna.StringArray{"1", "2", "3", "4", "5"}
}

func ExampleIntArray_Filter() {
	// Filter all elements greater than 2
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	y := x.Filter(func(v, i Int) Bool {
		if v > 2 {
			return true
		} else {
			return false
		}
	})
	Logv(y)
	// Output: dna.IntArray{3, 4, 5}
}

func ExampleIntArray_ForEach() {
	// Loop thorugh every element
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	x.ForEach(func(v, i Int) {
		fmt.Printf("{%#v-%#v}\n", i, v)
	})
	// Output:
	// {0-1}
	// {1-2}
	// {2-3}
	// {3-4}
	// {4-5}
}

func ExampleIntArray_Pop() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	x.Pop()
	Logv(x)
	// Output: dna.IntArray{1, 2, 3, 4}
}

func ExampleIntArray_Push() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	x.Push(6)
	Logv(x)
	// Output: dna.IntArray{1, 2, 3, 4, 5, 6}
}

func ExampleIntArray_Shift() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	x.Shift()
	Logv(x)
	// Output: dna.IntArray{2, 3, 4, 5}
}

func ExampleIntArray_Reverse() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	Logv(x.Reverse())
	// Output: dna.IntArray{5, 4, 3, 2, 1}
}

func ExampleIntArray_Join() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	y := x.Join("")
	Logv(y)
	// Output: "12345"
}

func ExampleIntArray_IndexOf() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	Logv(x.IndexOf(3))
	Logv(x.IndexOf(6))
	// Output: 2
	// -1
}

func ExampleIntArray_Length() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	Logv(x.Length())
	// Output: 5
}

func ExampleIntArray_Concat() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	y := x.Concat(IntArray{6, 7, 8, 9, 10})
	Logv(y)
	// Output: dna.IntArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func ExampleIntArray_ToString() {
	var x IntArray = IntArray{1, 2, 3, 4, 5}
	Logv(x.ToString())
	// Output: "12345"
}

func ExampleIntArray_Sort() {
	var x IntArray = IntArray{3, 1, 5, 2, 4}
	x.Sort()
	Logv(x)
	// Output: dna.IntArray{1, 2, 3, 4, 5}
}

func ExampleIntArray_Unique() {
	var x IntArray = IntArray{1, 1, 2, 3, 5, 6, 3}
	Logv(x.Unique())
	// Output: dna.IntArray{1, 2, 3, 5, 6}
}

func ExampleParseIntArray() {
	var str String
	var expectedArray IntArray = IntArray{}
	str = `{123,123,123123,123}`
	expectedArray = ParseIntArray(str)
	Logv(expectedArray)
	// Output:
	// dna.IntArray{123, 123, 123123, 123}
}
