package sqlpg

import (
	"database/sql"
	"dna"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"reflect"
)

// DB is a wrapper of sql.DB but some custom methods are added to enhance its functionalties.
type DB struct {
	*sql.DB
}

// Connect returns database of connected server
func Connect(cf *Config) (*DB, error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cf.Username, cf.Password, cf.Host, cf.Post, cf.Database, cf.SSLMode)
	db, err := sql.Open("postgres", connectionString)
	return &DB{db}, err
}

// QueryRecords is a reimplementation of sql.Db.Query().
// But it returns *Rows not *sql.Rows and error and takes query param as dna.String type
func (db *DB) Query(query dna.String, args ...interface{}) (*Rows, error) {
	rows, err := db.DB.Query(query.ToPrimitiveValue(), args...)
	return &Rows{rows}, err
}

// Insert inserts custom struct to a table.
// The table's name depends on the type's name and its package's name.
// Ex: Any instance of type Song from package ns will be inserted into table nssongs.
// Insert returns an error if the struct fails.
func (db *DB) Insert(structValue interface{}) error {
	tbName := GetTableName(structValue)
	insertQuery := GetInsertStatement(tbName, structValue, false)
	_, err := db.Query(insertQuery)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// InsertIgnore runs exactly the same as Insert.
// However if the insert value has already existed in a table,
// it does not return any error.
// It only returns an error if and only if other errors occur.
func (db *DB) InsertIgnore(structValue interface{}) error {
	err := db.Insert(structValue)
	if err != nil {
		if dna.String(err.Error()).Contains(`duplicate key value violates unique constraint`) {
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}

// Update updates statement from GetUpdateStatment and returns error if available
//
// 	* structValue : A struct-typed value being scanned. Its fields have to be dna basic type or time.Time.
// 	* conditionColumn : A snake-case column name in the condition, usually it's an id
// 	* columns : A list of args of column names in the table being updated.
// 	* Returns an update statement.
func (db *DB) Update(structValue interface{}, conditionColumn dna.String, columns ...dna.String) error {
	tbName := GetTableName(structValue)
	updateQuery, err0 := GetUpdateStatement(tbName, structValue, conditionColumn, columns...)
	if err0 != nil {
		return err0
	} else {
		_, err := db.Query(updateQuery)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

// Select runs an arbitrary SQL query,
// binding the columns in the result to fields on the struct specified by structValue.
//
//	* structValue: A struct value being binded. It has to be a pointer to a slice
//	* query: A query statement
//	* args: The args are for any placeholder parameters in the query.
//	* Returns error if available.
//
// It supports only custom struct, basic dna types(dna.Int,dna.String..) and time.Time
//
// Basic usage : (See example below for exact implementation):
//
// 	// Custone Struct
// 	songs := &[]ns.Song{}
// 	err := db.Select(songs, "SELECT * FROM nssongs ORDER BY id ASC LIMIT 10")
// 	// Basic dna types
// 	ids := &[]Int{}
//	err := db.Select(ids, "SELECT id FROM nssongs ORDER BY id ASC LIMIT 10")
func (db *DB) Select(structValue interface{}, query dna.String, args ...interface{}) error {
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	} else {
		for rows.Next() {
			ptrStruct := reflect.ValueOf(structValue)
			if ptrStruct.Kind() == reflect.Ptr {
				realStruct := reflect.Indirect(ptrStruct)
				if realStruct.Kind() != reflect.Slice {
					return errors.New("Select() Method only accepts slice")
				} else {
					val := reflect.New(reflect.TypeOf(structValue).Elem().Elem())
					if reflect.Indirect(val).Kind() == reflect.Struct {
						rows.StructScan(val.Interface())
					} else {
						rows.Scan(val.Interface())
					}
					// Log(val.Interface())
					realStruct.Set(reflect.Append(realStruct, reflect.Indirect(val)))
				}
			} else {
				return errors.New("Select() Method only accepts pointer")
			}

		}
	}

	return nil
}
