package sqlpg

import (
	"dna"
	"fmt"
	"reflect"
	"time"
)

// getPairValue returns something like `id=123` from a struct
func getPairValue(structValue interface{}, column dna.String) dna.String {
	fieldName := ("_" + column).Camelize()
	val := reflect.ValueOf(structValue).Elem()
	x, ok := val.Type().FieldByName(fieldName.ToPrimitiveValue())
	if ok == true {
		field := val.FieldByName(fieldName.ToPrimitiveValue())
		switch x.Type.String() {
		case "dna.Int":
			return dna.String(fmt.Sprintf("%v=%v", column, field.Interface().(dna.Int)))
		case "dna.Bool":
			return dna.String(fmt.Sprintf("%v=%v", column, field.Interface().(dna.Bool)))
		case "dna.Float":
			return dna.String(fmt.Sprintf("%v=%v", column, field.Interface().(dna.Float)))
		case "dna.String":
			return dna.String(fmt.Sprintf("%v=$binhdna$%v$binhdna$", column, field.Interface().(dna.String)))
		case "dna.StringArray":
			var tempStr dna.String = dna.String(fmt.Sprintf("%#v", field.Interface().(dna.StringArray))).Replace("dna.StringArray", "")
			return dna.String(fmt.Sprintf("%v=$binhdna$%v$binhdna$", column, tempStr))
		case "dna.IntArray":
			var tempStr dna.String = dna.String(fmt.Sprintf("%#v", field.Interface().(dna.StringArray))).Replace("dna.StringArray", "")
			return dna.String(fmt.Sprintf("%v=$binhdna$%v$binhdna$", column, tempStr))
		case "time.Time":
			datetime := field.Interface().(time.Time)
			if !datetime.IsZero() {
				return dna.String(fmt.Sprintf("%v=$binhdna$%v$binhdna$", column, datetime))
			} else {
				return dna.String(fmt.Sprintf("%v=%v", column, "NULL"))
			}

		}
	}
	return ""
}

// GetUpdateStatement returns an update statement from specified snake-case columns.
// If columns's names are not found, it will return an error.
// It updates some fields from a struct.
//
// 	* tbName : A name of update table.
// 	* structValue : A struct-typed value being scanned. Its fields have to be dna basic type or time.Time.
// 	* conditionColumn : A snake-case column name in the condition, usually it's an id
// 	* columns : A list of args of column names in the table being updated.
// 	* Returns an update statement.
func GetUpdateStatement(tbName dna.String, structValue interface{}, conditionColumn dna.String, columns ...dna.String) (dna.String, error) {
	if reflect.TypeOf(structValue).Kind() != reflect.Ptr {
		panic("StructValue has to be pointer")
		if reflect.TypeOf(structValue).Elem().Kind() != reflect.Struct {
			panic("StructValue has to be struct type")
		}
	}
	query := "UPDATE " + tbName + " SET "
	result := dna.StringArray{}
	for _, column := range columns {
		result.Push(getPairValue(structValue, column))
	}
	conditionRet := " WHERE " + getPairValue(structValue, conditionColumn)
	return query + result.Join(`,`) + conditionRet, nil
}
