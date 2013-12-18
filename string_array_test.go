package dna

import (
	"fmt"
	"testing"
)

func TestStringArray_Map(t *testing.T) {
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	y := IntArray(x.Map(func(v String, i Int) Int {
		return v.ToInt()
	}).([]Int))
	z := IntArray{1, 2, 3, 4, 5}
	for i, v := range y {
		if v != z[i] {
			t.Errorf("%v (IntArray) cannot be converted to StringArray", x)
		}
	}
}

func TestStringArray_ParseStringArray(t *testing.T) {

	var str String
	var expectedArray StringArray = StringArray{}
	var resultedArray StringArray = StringArray{}

	// Set 1
	str = `{"Lê Việt Anh",HOLETHU,"tran,thai,thinh",YYYYYY,"tap nhu\" lai"}`
	expectedArray = ParseStringArray(str)
	resultedArray = StringArray{"Lê Việt Anh", "HOLETHU", "tran,thai,thinh", "YYYYYY", "tap nhu\" lai"}
	for idx, val := range resultedArray {
		if val != expectedArray[idx] {
			t.Error("sfsd")
		}
	}

	// Set 2
	str = `{HOLETHU,"Lê Việt Anh","tran,thai,thinh","tap nhu\" lai"}`
	expectedArray = ParseStringArray(str)
	resultedArray = StringArray{"HOLETHU", "Lê Việt Anh", "tran,thai,thinh", "tap nhu\" lai"}
	for idx, val := range resultedArray {
		if val != expectedArray[idx] {
			t.Error("sfsd")
		}
	}

	// Set 3
	str = `{"HO LE THU","Lê Việt Anh","tran,thai,thinh","tap nhu\" lai",popopopdas}`
	expectedArray = ParseStringArray(str)
	resultedArray = StringArray{"HO LE THU", "Lê Việt Anh", "tran,thai,thinh", "tap nhu\" lai", "popopopdas"}
	for idx, val := range resultedArray {
		if val != expectedArray[idx] {
			t.Error("sfsd")
		}
	}

	// Set 4
	str = `{xxxxx,"HO LE's THU","Lê, Việt Anh","tran,thai,thinh","tap nhu\" lai",popopopdas}`
	expectedArray = ParseStringArray(str)
	resultedArray = StringArray{"xxxxx", "HO LE's THU", "Lê, Việt Anh", "tran,thai,thinh", "tap nhu\" lai", "popopopdas"}
	for idx, val := range resultedArray {
		if val != expectedArray[idx] {
			t.Error("sfsd")
		}
	}

}

// Example cases
func ExampleStringArray() {
	x := StringArray{"1", "2", "3", "4", "5"}                          // literal
	var y StringArray = StringArray([]String{"1", "2", "3", "4", "5"}) // type conversion
	var z StringArray = []String{"1", "2", "3", "4", "5"}
	Logv(x)
	Logv(y)
	Logv(z)
	// Output: dna.StringArray{"1", "2", "3", "4", "5"}
	// dna.StringArray{"1", "2", "3", "4", "5"}
	// dna.StringArray{"1", "2", "3", "4", "5"}
}

func ExampleStringArray_Map() {
	// Convert StringArray to IntArray
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	y := IntArray(x.Map(func(v String, i Int) Int {
		return v.ToInt()
	}).([]Int))
	Logv(y)
	// Output: dna.IntArray{1, 2, 3, 4, 5}
}

func ExampleStringArray_Filter() {
	// Filter all elements whose value converted to integer greater than 2
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	y := x.Filter(func(v String, i Int) Bool {
		if v.ToInt() > 2 {
			return true
		} else {
			return false
		}
	})
	Logv(y)
	// Output: dna.StringArray{"3", "4", "5"}
}

func ExampleStringArray_ForEach() {
	// Loop thorugh every element
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	x.ForEach(func(v String, i Int) {
		fmt.Printf("{%#v-%#v}\n", i, v)
	})
	// Output:
	// {0-"1"}
	// {1-"2"}
	// {2-"3"}
	// {3-"4"}
	// {4-"5"}
}

func ExampleStringArray_IndexOf() {
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	Logv(x.IndexOf("3"))
	Logv(x.IndexOf("6"))
	// Output: 2
	// -1
}

func ExampleStringArray_Join() {
	// Join array with "," delimiter
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	Logv(x.Join(","))
	// Output: "1,2,3,4,5"
}

func ExampleStringArray_Length() {
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	Logv(x.Length())
	// Output: 5
}

func ExampleStringArray_ToIntArray() {
	// Notice: Every element of the array has to be convertible to integer or error will occurs at runtime
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	Logv(x.ToIntArray())
	// Output: dna.IntArray{1, 2, 3, 4, 5}
}

func ExampleStringArray_Reverse() {
	// Notice: Every element of the array has to be convertible to integer or error will occurs at runtime
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	Logv(x.Reverse())
	// Output: dna.StringArray{"5", "4", "3", "2", "1"}
}

func ExampleStringArray_Pop() {
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	x.Pop()
	Logv(x)
	// Output: dna.StringArray{"1", "2", "3", "4"}
}

func ExampleStringArray_Push() {
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	x.Push("6")
	Logv(x)
	// Output: dna.StringArray{"1", "2", "3", "4", "5", "6"}
}

func ExampleStringArray_Shift() {
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	x.Shift()
	Logv(x)
	// Output: dna.StringArray{"2", "3", "4", "5"}
}

func ExampleStringArray_Unique() {
	var x StringArray = StringArray{"1", "1", "2", "3", "5", "6", "3"}
	Logv(x.Unique())
	// Output: dna.StringArray{"1", "2", "3", "5", "6"}
}

func ExampleStringArray_Concat() {
	var x StringArray = StringArray{"1", "2", "3", "4", "5"}
	Logv(x.Concat(StringArray{"6", "7", "8", "9", "10"}))
	// Output: dna.StringArray{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
}

func ExampleStringArray_SplitWithRegexp() {
	var x StringArray = StringArray{"Hôm nay", "This is fun"}
	var y StringArray = StringArray{"Joey Yung / Dung Tổ Nhi"}
	Logv(x.SplitWithRegexp("\\s+"))
	Logv(x.SplitWithRegexp("\\d+"))
	Logv(y.SplitWithRegexp(" / "))
	// Output: dna.StringArray{"Hôm", "nay", "This", "is", "fun"}
	// dna.StringArray{"Hôm nay", "This is fun"}
	// dna.StringArray{"Joey Yung", "Dung Tổ Nhi"}
}

func ExampleStringArray_IndexWithRegexp() {
	var x StringArray = StringArray{"Hôm nay là thứ 2", "This is fun"}
	Logv(x.IndexWithRegexp("\\d+"))
	Logv(x.IndexWithRegexp("i.+un"))
	Logv(x.IndexWithRegexp("[0-9]{3}"))
	// Output: 0
	// 1
	// -1
}

func ExampleParseStringArray() {
	var str String
	var expectedArray StringArray = StringArray{}
	str = `{"Lê Việt Anh",HOLETHU,"tran,thai,thinh",YYYYYY,"tap nhu\" lai"}`
	expectedArray = ParseStringArray(str)
	Logv(expectedArray)
	// Output:
	// dna.StringArray{"Lê Việt Anh", "HOLETHU", "tran,thai,thinh", "YYYYYY", "tap nhu\" lai"}
}
