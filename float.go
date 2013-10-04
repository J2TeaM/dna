package dna

import (
	"fmt"
	"math"
)

// Redefine new Int Type
type Float float64

// ToString returns string from float
func (f Float) ToString() String {
	return String(fmt.Sprint(f))
}

// ToFormattedString returns a new formatted string given width and precision params.
//
// Notice: When width is positive, it fills from left with whitespace, otherwise, it fills from right to left with space
func (f Float) ToFormattedString(width, precision Int) String {
	return String(fmt.Sprintf("%[2]*.[3]*[1]f", f, int(width), int(precision)))
}

// Ceil returns the least integer value greater than or equal to the float number
func (f Float) Ceil() Int {
	return Int(math.Ceil(float64(f)))
}

// Alias of Ceil
func (f Float) Round() Int {
	return f.Ceil()
}

// Floor returns the greatest integer value less than or equal to x
func (f Float) Floor() Int {
	return Int(math.Floor(float64(f)))
}

// ToInt returns Int from Float. Alias of Floor
func (f Float) ToInt() Int {
	return f.Floor()
}

// Value returns primitive type float64
func (f Float) Value() float64 {
	return float64(f)
}
