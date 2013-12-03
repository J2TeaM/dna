package dna

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type StringArray []String

func (sa StringArray) Length() Int {
	return Int(len(sa))
}

/*
Map method for StringArray.
The returned array has the same or different type comparing to the old one.

Notice: The real return value is interface{}. So it has to be converted into StringArray Type
	var x := dna.StringArray{"anbinh","ducbinh","binhdna","1232"}
	x = dna.StringArray(x.Map(func(value dna.String ,index dna.Int) dna.String{
		return fmt.Sprint("{",value,":",index,"}")
	}).([]dna.String))
*/
func (sa StringArray) Map(fn interface{}) interface{} {
	return Map(sa, fn)
}

// ForEach loops through every element in string array and does not return anything
func (sa StringArray) ForEach(fn func(v String, i Int)) {
	for index, value := range sa {
		fn(value, Int(index))
	}
}

// Join returns a new string by joining all elements in array with no seperator
func (sa StringArray) Join(sep String) String {
	sarr := make([]string, len(sa))
	for i, v := range sa {
		sarr[i] = fmt.Sprint(v)
	}
	return String(strings.Join(sarr, fmt.Sprint(sep)))
}

// Filter loops through a string array and return a new string array whose elements are filtered by anonymous function
func (sa StringArray) Filter(fn func(v String, i Int) Bool) StringArray {
	var result = StringArray{}
	for index, value := range sa {
		if bool(fn(value, Int(index))) {
			result = append(result, value)
		}
	}
	return result
}

// IndexOf returns an index of a string in array
func (sa StringArray) IndexOf(value String) Int {
	for i, v := range sa {
		if v == value {
			return Int(i)
		}
	}
	return -1
}

// ToIntArray return a new int Array from a string array
// Make sure that every element whose type is String can be convertible to Int
func (sa StringArray) ToIntArray() IntArray {
	arr := make(IntArray, len(sa))
	for i, v := range sa {
		arr[i] = v.ToInt()
	}
	return arr
}

// Reverse returns new string array with reversed order
func (sa StringArray) Reverse() StringArray {
	length := len(sa)
	tmp := make(StringArray, length)
	for i, v := range sa {
		tmp[length-i-1] = v
	}
	return tmp
}

// Push inserts new value to the existing array
func (sa *StringArray) Push(value String) {
	slice := append(*sa, value)
	*sa = slice
}

// Pop removes the last element of the existing array
func (sa *StringArray) Pop() {
	slice := *sa
	slice = slice[0 : len(slice)-1]
	*sa = slice
}

// Shift removes the first element of the existing array
func (sa *StringArray) Shift() {
	slice := *sa
	slice = slice[1:len(slice)]
	*sa = slice
}

// Unique returns unique StringArray
func (sa StringArray) Unique() StringArray {
	var tmp StringArray = StringArray{}
	for _, v := range sa {
		if tmp.IndexOf(v) == -1 {
			tmp.Push(v)
		}
	}
	return tmp
}

// Concat returns a new concatenated array
func (sa StringArray) Concat(arr StringArray) StringArray {
	return append(sa, arr...)
}

// SplitWithRegexp returns a new array whose element is splitted by pattern
func (sa StringArray) SplitWithRegexp(pattern String) StringArray {
	var ret StringArray = StringArray{}
	var tmp StringArray
	for _, value := range sa {
		tmp = value.SplitWithRegexp(pattern, -1)
		ret = ret.Concat(tmp)
	}
	return ret
}

// IndexWithRegexp returns the first index of element satisfied the pattern
func (sa StringArray) IndexWithRegexp(pattern String) Int {
	for index, value := range sa {
		if value.Match(pattern) {
			return Int(index)
		}
	}
	return -1
}

// Value implements the Valuer interface in database/sql/driver package.
func (sa StringArray) Value() (driver.Value, error) {
	return driver.Value(string(String(fmt.Sprintf("%#v", sa)).Replace("dna.StringArray", ""))), nil
}

// ParseStringArray returns dna.StringArray-typed array from postgresql-array-typed string
// Ex: {"abc","def"} => dna.StringArray{"abc", "def"}
func ParseStringArray(str String) StringArray {
	var i Int = 0
	var length Int = str.Length()
	if str.CharAt(0) != "{" || str.CharAt(length-1) != "}" {
		panic("The input is not postgresql array")
	}
	// openQIdx is the open quote index
	// closeQIdx is the close quote index
	var openQIdx Int = 0
	var closeQIdx Int = 0

	// indexes of values not inside quotes
	// openNoqIdx is open no quote index
	// closeNoQIdx is close no quote index
	var openNoQIdx Int = 0
	var closeNoQIdx Int = 0

	// used to check if first value is not quoted & is the first element of an array
	// fcNoQV is the first comma of no quoted value
	var fcNoQV Bool = true

	// return value
	var result StringArray

	// to check if comma (,) is inside quotes
	// isVRecoreded is the value inside quotes being recorded
	var isVRecoreded = false
	// Loop through pos 1 to pos length-2
	i++
	for i < length {
		switch str.CharAt(i) {
		case `\`:
			i = i + 2
		case `"`:
			closeQIdx = i
			if closeQIdx > openQIdx && openQIdx != 0 && isVRecoreded == true {
				result.Push(str.Substring(openQIdx+1, closeQIdx).Replace(`\"`, `"`))
				isVRecoreded = false
				// reset no quoted value
				openNoQIdx = 0
				closeNoQIdx = 0
				fcNoQV = false
				// Log(str.Substring(openQIdx+1, closeQIdx))
			} else {
				isVRecoreded = true
			}
			openQIdx = closeQIdx
			i++
		case `,`:
			if isVRecoreded == false {
				closeNoQIdx = i
				if closeNoQIdx > openNoQIdx && (openNoQIdx != 0 || fcNoQV == true) {
					result.Push(str.Substring(openNoQIdx+1, closeNoQIdx).Replace(`\"`, `"`))
					// Log("---", str.Substring(openNoQIdx+1, closeNoQIdx))
				}
				openNoQIdx = closeNoQIdx
			}
			i++
		case `}`:
			if isVRecoreded == false && str.CharAt(i-1) != `"` {
				closeNoQIdx = i
				result.Push(str.Substring(openNoQIdx+1, closeNoQIdx).Replace(`\"`, `"`))
			}
			i++
		default:
			i++
		}
	}
	return result
}

// Scan implements the Scanner interface in database/sql package.
func (sa *StringArray) Scan(src interface{}) error {
	var source StringArray
	switch src.(type) {
	case string:
		source = ParseStringArray(String(string(src.(string))))
	case []byte:
		source = ParseStringArray(String(string(src.([]byte))))
	case nil:
		source = StringArray{}
	default:
		return errors.New("Incompatible type for dna.StringArray type")
	}
	*sa = source
	return nil
}
