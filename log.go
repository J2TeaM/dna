package dna

import (
	"fmt"
	"reflect"
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

// Log prints an variable with full format: "%#v"
func Logv(a interface{}) {
	fmt.Printf("%#v\n", a)
}

// Print outputs the values on screen
func Print(a ...interface{}) {
	fmt.Print(a...)
}
