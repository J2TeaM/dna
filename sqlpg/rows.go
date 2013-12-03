package sqlpg

import (
	"database/sql"
	. "dna"
	"errors"
	_ "github.com/lib/pq"
	"reflect"
	"time"
)

type Row struct {
	*sql.Row
}

type Rows struct {
	*sql.Rows
}

// StructScan scans struct-typed value from *sql.Rows. It returns an error if available.
// The structValue has to be pointer.
//
// 	* structValue: Types of struct fields have to be dna basic ones (dna.String, dna.Int...) or time.Time.
//	* Return an error
//
// When scanning a struct, it will be implemented to the name convention mentioned above:
// changing from snake case to upper camel case.
// If a column's name after being camelized is not found in the struct, its value will be ignored automatically.
// Only matched names will be scanned.
//
// For example: If a struct-typed song is utilized to scan an album table, some identical columns will be matched
// such as id, artists.
func (rows *Rows) StructScan(structValue interface{}) error {
	if reflect.TypeOf(structValue).Kind() != reflect.Ptr {
		panic("StructValue has to be pointer")
		if reflect.TypeOf(structValue).Elem().Kind() != reflect.Struct {
			panic("StructValue has to be struct type")
		}
	}

	columns, err1 := rows.Columns()
	if err1 != nil {
		return errors.New("Cannot find columns")
	}
	rawResult := make([]interface{}, len(columns))
	dest := make([]interface{}, len(columns))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}
	err := rows.Scan(dest...)
	if err != nil {
		return errors.New("Cannot scan value")
	}
	for idx, rawValue := range rawResult {
		fieldName := ("_" + String(columns[idx])).Camelize()
		val := reflect.ValueOf(structValue).Elem()
		x, ok := val.Type().FieldByName(fieldName.ToPrimitiveValue())
		if ok == true {
			field := val.FieldByName(fieldName.ToPrimitiveValue())
			switch x.Type.String() {
			case "dna.Int":
				if rawValue != nil {
					field.Set(reflect.ValueOf(Int(rawValue.(int64))))
				}
			case "dna.Bool":
				if rawValue != nil {
					field.Set(reflect.ValueOf(Bool(rawValue.(bool))))
				}
			case "dna.Float":
				if rawValue != nil {
					field.Set(reflect.ValueOf(Float(rawValue.(float64))))
				}
			case "dna.String":
				if rawValue != nil {
					field.Set(reflect.ValueOf(String(string(rawValue.([]byte)))))
				}
			case "dna.StringArray":
				if rawValue != nil {
					field.Set(reflect.ValueOf(ParseStringArray(String(string(rawValue.([]byte))))))
				}
			case "dna.IntArray":
				if rawValue != nil {
					field.Set(reflect.ValueOf(ParseIntArray(String(string(rawValue.([]byte))))))
				}
			case "time.Time":
				if rawValue != nil {
					field.Set(reflect.ValueOf(rawValue.(time.Time)))
				}
			}
		}
	}
	return nil
}
