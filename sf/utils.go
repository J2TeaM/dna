package sf

import (
	"dna"
	"fmt"
	"reflect"
	"time"
)

func getColumn(f reflect.StructField, structValue interface{}) (dna.String, dna.String) {
	var columnName, columnValue dna.String
	switch f.Type.String() {
	case "dna.Int":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("%v", structValue))

	case "dna.Float":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("%v", structValue))

	case "dna.Bool":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("%v", structValue))

	case "dna.String":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", structValue))

	case "dna.StringArray":
		var tempStr dna.String = dna.String(fmt.Sprintf("%#v", structValue)).Replace("dna.StringArray", "")
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", tempStr))

	case "dna.IntArray":
		var tempStr dna.String = dna.String(fmt.Sprintf("%#v", structValue)).Replace("dna.IntArray", "")
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", tempStr))
	case "time.Time":
		columnName = dna.String(f.Name).ToSnakeCase()
		datetime := structValue.(time.Time)
		if !datetime.IsZero() {
			columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", dna.String(datetime.String()).ReplaceWithRegexp(`\+.+$`, ``).Trim()))
		} else {
			columnValue = dna.String(fmt.Sprintf("%v", "NULL"))
		}

	default:
		// panic("A Field of struct is not dna basic type")
	}
	return columnName, columnValue
}

func getInsertStatement(tbName dna.String, structValue interface{}, condStr dna.String, isPrintable dna.Bool) dna.String {
	var realKind string
	var columnNames, columnValues dna.StringArray
	tempintslice := []int{0}
	var ielements int
	var kind string = reflect.TypeOf(structValue).Kind().String()
	if kind == "ptr" {
		realKind = reflect.TypeOf(structValue).Elem().Kind().String()

	} else {
		realKind = reflect.TypeOf(structValue).Kind().String()

	}

	if realKind != "struct" {
		panic("Param has to be struct")
	}

	if kind == "ptr" {
		ielements = reflect.TypeOf(structValue).Elem().NumField()
	} else {
		ielements = reflect.TypeOf(structValue).NumField()
	}

	for i := 0; i < ielements; i++ {
		tempintslice[0] = i
		if kind == "ptr" {
			f := reflect.TypeOf(structValue).Elem().FieldByIndex(tempintslice)
			v := reflect.ValueOf(structValue).Elem().FieldByIndex(tempintslice)
			clName, clValue := getColumn(f, v.Interface())
			columnNames.Push(clName)
			columnValues.Push(clValue)
		} else {
			f := reflect.TypeOf(structValue).FieldByIndex(tempintslice)
			v := reflect.ValueOf(structValue).FieldByIndex(tempintslice)
			clName, clValue := getColumn(f, v.Interface())
			columnNames.Push(clName)
			columnValues.Push(clValue)
		}

	}
	if isPrintable == true {
		return "INSERT INTO " + tbName + "\n(" + columnNames.Join(",") + ")\n" + " SELECT " + columnValues.Join(",\n") + " \n" + condStr
	} else {
		return "INSERT INTO " + tbName + "(" + columnNames.Join(",") + ")" + " SELECT " + columnValues.Join(",") + " " + condStr
	}
}

// GetTableName returns table name from a struct.
// Ex: An instance of ns.Song will return nssongs
// An instance of ns.Album will return nsalbums
func getTableName(structValue interface{}) dna.String {
	val := reflect.TypeOf(structValue)
	if val.Kind() != reflect.Ptr {
		panic("StructValue has to be pointer")
		if val.Elem().Kind() != reflect.Struct {
			panic("StructValue has to be struct type")
		}
	}
	return dna.String(val.Elem().String()).Replace(".", "").ToLowerCase() + "s"
}

func getInsertStmt(structValue interface{}, condStr dna.String) dna.String {
	return getInsertStatement(getTableName(structValue), structValue, condStr, false) + ";"
}
