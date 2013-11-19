package sqlpg

import (
	"dna"
	"fmt"
	"reflect"
	"time"
)

func getColumn(f reflect.StructField, v interface{}) (dna.String, dna.String) {
	var columnName, columnValue dna.String
	switch f.Type.String() {
	case "dna.Int":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("%v", v))

	case "dna.Float":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("%v", v))

	case "dna.Bool":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("%v", v))

	case "dna.String":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", v))

	case "dna.StringArray":
		var tempStr dna.String = dna.String(fmt.Sprintf("%#v", v)).Replace("dna.StringArray", "")
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", tempStr))

	case "dna.IntArray":
		var tempStr dna.String = dna.String(fmt.Sprintf("%#v", v)).Replace("dna.IntArray", "")
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", tempStr))
	case "time.Time":
		columnName = dna.String(f.Name).ToSnakeCase()
		columnValue = dna.String(fmt.Sprintf("$binhdna$%v$binhdna$", dna.String(v.(time.Time).String()).ReplaceWithRegexp(`\+.+$`, ``).Trim()))
	default:
		// panic("A Field of struct is not dna basic type")
	}
	return columnName, columnValue
}

// GetInsertQuery returns insert statement from a struct. If input value is not struct, it will panic.
//	* tableName : A name of table in database you want to insert
//	* v : A struct type you want to insert. Fields of the struct has to be dna basic types (dna.Int, dna.String, dna.StringArray...)
//	* isPrintable: A param determines where to print the pretty result statement
// Notice:  Insert statement uses Dollar-quoted String Constants with special tag "binhdna". So string or array is contained between $binhdna$ symbols
func GetInsertStatement(tableName dna.String, v interface{}, isPrintable dna.Bool) dna.String {
	var realKind string
	var columnNames, columnValues dna.StringArray
	tempintslice := []int{0}
	var ielements int
	var kind string = reflect.TypeOf(v).Kind().String()
	if kind == "ptr" {
		realKind = reflect.TypeOf(v).Elem().Kind().String()

	} else {
		realKind = reflect.TypeOf(v).Kind().String()

	}

	if realKind != "struct" {
		panic("Param has to be struct")
	}

	if kind == "ptr" {
		ielements = reflect.TypeOf(v).Elem().NumField()
	} else {
		ielements = reflect.TypeOf(v).NumField()
	}

	for i := 0; i < ielements; i++ {
		tempintslice[0] = i
		if kind == "ptr" {
			f := reflect.TypeOf(v).Elem().FieldByIndex(tempintslice)
			v := reflect.ValueOf(v).Elem().FieldByIndex(tempintslice)
			clName, clValue := getColumn(f, v.Interface())
			columnNames.Push(clName)
			columnValues.Push(clValue)
		} else {
			f := reflect.TypeOf(v).FieldByIndex(tempintslice)
			v := reflect.ValueOf(v).FieldByIndex(tempintslice)
			clName, clValue := getColumn(f, v.Interface())
			columnNames.Push(clName)
			columnValues.Push(clValue)
		}

	}
	if isPrintable == true {
		return "INSERT INTO " + tableName + "\n(" + columnNames.Join(",") + ")\n" + "VALUES (\n" + columnValues.Join(",\n") + "\n)"
	} else {
		return "INSERT INTO " + tableName + "(" + columnNames.Join(",") + ")" + " VALUES (" + columnValues.Join(",") + ")"
	}

}
