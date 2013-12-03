package dna

import (
	"fmt"
	"reflect"
	"time"
)

// Log prints format to screen with GO-syntax representation except for string
func Log(a ...interface{}) {
	format := ""
	for i, v := range a {
		if reflect.ValueOf(v).Kind() == reflect.String {
			format += "%v"
		} else {
			if reflect.ValueOf(v).Kind() == reflect.Float64 || reflect.ValueOf(v).Kind() == reflect.Float64 {
				format += "%f"
			} else {
				format += "%#v"
			}
		}
		if i < len(a)-1 {
			format += " "
		}
	}
	format += "\n"
	fmt.Printf(format, a...)
}

// Log prints an variable with full format: "%#v". Only accept one variable
func Logv(a interface{}) {
	fmt.Printf("%#v\n", a)
}

// LogStruct prints struct type with its field name and field value. Using for debug
func LogStruct(a interface{}) {
	tempintslice := []int{0}
	ielements := reflect.TypeOf(a).Elem().NumField()
	for i := 0; i < ielements; i++ {
		tempintslice[0] = i
		f := reflect.TypeOf(a).Elem().FieldByIndex(tempintslice)
		v := reflect.ValueOf(a).Elem().FieldByIndex(tempintslice)
		if f.Type.String() != "time.Time" {
			fmt.Printf("%v : %#v\n", f.Name, v.Interface())
		} else {
			fmt.Printf("%v : %#v\n", f.Name, String(v.Interface().(time.Time).String()).ReplaceWithRegexp(`\+.+$`, ``).Trim())
		}

	}
}

// Print outputs the values on screen
func Print(a ...interface{}) {
	fmt.Print(a...)
}

func Sprintf(format String, a ...interface{}) String {
	return String(fmt.Sprintf(format.ToPrimitiveValue(), a...))
}

// PanicError panics if err is not nil
func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
