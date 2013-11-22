package sqlpg

import (
	"dna"
	"reflect"
)

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
