package dna

import (
	"database/sql/driver"
	"errors"
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

// ToPrimitiveValue returns primitive type float64
func (f Float) ToPrimitiveValue() float64 {
	return float64(f)
}

// Value implements the Valuer interface in database/sql/driver package.
func (f Float) Value() (driver.Value, error) {
	return driver.Value(float64(f)), nil
}

// Scan implements the Scanner interface in database/sql package.
// Default value for nil is 0
func (f *Float) Scan(src interface{}) error {
	var source Float
	switch src.(type) {
	case float64:
		source = Float(src.(float64))
	case nil:
		source = 0
	default:
		return errors.New("Incompatible type for dna.Float type")
	}
	*f = source
	return nil
}
