package sqlpg

import (
	"dna"
	"errors"
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

// GetInsertStatement returns insert statement from a struct. If input value is not struct, it will panic.
//	* tbName : A name of table in database you want to insert
//	* structValue : A struct-typed value. The struct's fields has to be dna basic types (dna.Int, dna.String..) or time.Time
//	* isPrintable: A param determines where to print the pretty result statement
//	* Return an insert statement
// Notice:  Insert statement uses Dollar-quoted String Constants with special tag "binhdna".
// So string or array is contained between $binhdna$ symbols.
// Therefore no need to escape values.
func GetInsertStatement(tbName dna.String, structValue interface{}, isPrintable dna.Bool) dna.String {
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
		return "INSERT INTO " + tbName + "\n(" + columnNames.Join(",") + ")\n" + "VALUES (\n" + columnValues.Join(",\n") + "\n);"
	} else {
		return "INSERT INTO " + tbName + "(" + columnNames.Join(",") + ")" + " VALUES (" + columnValues.Join(",") + ");"
	}
}

// GetInsertIgnoreStatement returns insert ignore statement from a struct. If input value is not struct, it will panic.
//	* tbName : A name of table in database you want to insert
//	* structValue : A struct-typed value. The struct's fields has to be dna basic types (dna.Int, dna.String..) or time.Time
//	* primaryColumn : A name of primary column. If a row is duplicate, it will be discarded.
//	* primaryValue : A value of row needed to be compared.
//	* isPrintable: A param determines where to print the pretty result statement
//	* Return an insert statement
// Notice:  Insert statement uses Dollar-quoted String Constants with special tag "binhdna".
// So string or array is contained between $binhdna$ symbols.
// Therefore no need to escape values.
func GetInsertIgnoreStatement(tbName dna.String, structValue interface{}, primaryColumn dna.String, primaryValue interface{}, isPrintable dna.Bool) dna.String {
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
	condStr := dna.Sprintf("WHERE NOT EXISTS (SELECT 1 FROM %v WHERE id=%v)", tbName, primaryValue)
	if isPrintable == true {
		return "INSERT INTO " + tbName + "\n(" + columnNames.Join(",") + ")\n" + " SELECT " + columnValues.Join(",\n") + " \n" + condStr + ";" + " \n"
	} else {
		return "INSERT INTO " + tbName + "(" + columnNames.Join(",") + ")" + " SELECT " + columnValues.Join(",") + " " + condStr + ";"
	}
}

// GetTableName returns table name from a struct.
// Ex: An instance of ns.Song will return nssongs
// An instance of ns.Album will return nsalbums
func GetTableName(structValue interface{}) dna.String {
	val := reflect.TypeOf(structValue)
	if val.Kind() != reflect.Ptr {
		panic("StructValue has to be pointer")
		if val.Elem().Kind() != reflect.Struct {
			panic("StructValue has to be struct type")
		}
	}
	return dna.String(val.Elem().String()).Replace(".", "").ToLowerCase() + "s"
}

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
			var tempStr dna.String = dna.String(fmt.Sprintf("%#v", field.Interface().(dna.IntArray))).Replace("dna.IntArray", "")
			return dna.String(fmt.Sprintf("%v=$binhdna$%v$binhdna$", column, tempStr))
		case "time.Time":
			datetime := field.Interface().(time.Time)
			if !datetime.IsZero() {
				return dna.String(fmt.Sprintf("%v=$binhdna$%v$binhdna$", column, dna.String(datetime.String()).ReplaceWithRegexp(`\+.+$`, ``).Trim()))
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
	query := "UPDATE " + tbName + " SET\n"
	result := dna.StringArray{}
	for _, column := range columns {
		result.Push(getPairValue(structValue, column))
	}
	conditionRet := "\nWHERE " + getPairValue(structValue, conditionColumn) + ";"
	return query + result.Join(",\n") + conditionRet, nil
}

var globalSqlTransactoNo = 0

// ExecQueriesInTransaction executes queries in a transaction.
// If one statement fails, the whole queries cannot commit.
//
// The returned error is nil if there is no error.
// If an error occurs, each statement will be enclosed in
// format $$$error$$$.
// 	$$$error$$$ Your Custom Query $$$error$$$
//
// This function is seen in songfreaks and allmusic sites.
func ExecQueriesInTransaction(db *DB, queries *dna.StringArray) error {
	var err error
	globalSqlTransactoNo += 1
	tx, err := db.Begin()
	if err != nil {
		dna.Log("Transaction No:" + dna.Sprintf("%v", globalSqlTransactoNo).String() + err.Error() + " Could not create transaction\n")
	}

	for idx, query := range *queries {
		_, err = tx.Exec(query.String())
		if err != nil {
			dna.Log(dna.Sprintf("DNAError: Transaction No: %v - %v - %v - %v\n", dna.Sprintf("%v", globalSqlTransactoNo), idx, err.Error(), "Could not execute the statement"))
		}
		// stmt, err := tx.Prepare(query.String())
		// if err != nil {
		// 	dna.Log(dna.Sprintf("DNAError Transaction No: %v - %v - %v - %v \n", dna.Sprintf("%v", globalSqlTransactoNo), idx, err.Error(), "Could not prepare"))
		// } else {
		// 	_, err = stmt.Exec()
		// 	if err != nil {
		// 		dna.Log(dna.Sprintf("DNAError: Transaction No: %v - %v - %v - %v\n", dna.Sprintf("%v", globalSqlTransactoNo), idx, err.Error(), "Could not execute the prepared statement"))
		// 	}

		// 	err = stmt.Close()
		// 	if err != nil {
		// 		dna.Log("Transaction No:" + dna.Sprintf("%v", globalSqlTransactoNo).String() + err.Error() + " Could not close\n")
		// 	}
		// }
	}
	err = tx.Commit()
	if err != nil {
		dna.Log("Transaction No:" + dna.Sprintf("%v", globalSqlTransactoNo).String() + err.Error() + " Could not commit transaction\n")
	}

	if err != nil {
		errQueries := dna.StringArray(queries.Map(func(val dna.String, idx dna.Int) dna.String {
			return "Transaction No:" + dna.Sprintf("%v", globalSqlTransactoNo) + " $$$error$$$" + val + "$$$error$$$"
		}).([]dna.String))
		return errors.New(err.Error() + errQueries.Join("\n").String())
	} else {
		return nil
	}
}
