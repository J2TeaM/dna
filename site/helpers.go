package site

import (
	"dna"
	"dna/sqlpg"
)

// RecoverErrorQueries is a wrapper of RecoverSQLLogError.
// It prints some useful info to console.
func RecoverErrorQueries(path dna.String, db *sqlpg.DB) {
	dna.Print("Recovering all sql queries having errors")
	errCount := RecoverSQLLogError(path, db)
	if errCount == 0 {
		dna.Print(" => Completed!")
	} else {
		dna.Log("\nTotal error count:", errCount)
	}
}
