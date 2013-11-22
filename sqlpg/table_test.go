package sqlpg

import (
	. "dna"
	"dna/ns"
)

func ExampleGetTableName() {
	table := GetTableName(ns.NewAlbum())
	table1 := GetTableName(ns.NewSong())
	Log(table)
	Log(table1)
	// Output:
	// nsalbums
	// nssongs
}
