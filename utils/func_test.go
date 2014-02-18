package utils

import (
	"dna"
	"dna/sqlpg"
)

func ExampleSelectMissingIds() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig("./app.ini"))
	dna.PanicError(err)
	ids, err := SelectMissingIds("ziartists", &dna.IntArray{5, 6, 7, 8, 9}, db)
	dna.PanicError(err)
	dna.Log(ids)
	db.Close()
	// Output:
	// &dna.IntArray{8}
}

// The statement above means: Select all artistids which are not available
// from ziartists table from 5->9.
// Result is 8.
func ExampleSelectMissingIdsWithRange() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig("./app.ini"))
	dna.PanicError(err)
	ids, err := SelectMissingIdsWithRange("ziartists", 5, 9, db)
	dna.PanicError(err)
	dna.Log(ids)
	db.Close()
	// Output:
	// &dna.IntArray{8}
}
