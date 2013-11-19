package sqlpg

import (
	"database/sql"
	. "dna"
	"errors"
	"reflect"
	"time"
)

func StructScan(structValue interface{}, rows *sql.Rows) error {

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
