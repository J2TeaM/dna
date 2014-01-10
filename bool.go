package dna

import (
	"database/sql/driver"
	"errors"
)

type Bool bool

// Value implements the Valuer interface in database/sql/driver package.
func (b Bool) Value() (driver.Value, error) {
	return driver.Value(bool(b)), nil
}

func (b Bool) String() string {
	return Sprintf("%v", bool(b)).String()
}

// Scan implements the Scanner interface in database/sql package.
// Default value for nil is false
func (b *Bool) Scan(src interface{}) error {
	var source Bool
	switch src.(type) {
	case bool:
		source = Bool(src.(bool))
	case nil:
		source = false
	default:
		return errors.New("Incompatible type for dna.Bool type")
	}
	*b = source
	return nil
}
