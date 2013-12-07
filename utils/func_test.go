package utils

import (
	"dna"
	"dna/sqlpg"
)

func ExampleSelectNonExistedIds() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig("./app.ini"))
	dna.PanicError(err)
	ids, err := SelectNonExistedIds("ziartists", dna.IntArray{5, 6, 7, 8, 9}, db)
	dna.PanicError(err)
	dna.Log(ids)
	db.Close()
	// Output:
	// &dna.IntArray{8}
}
