package site

import (
	"dna"
	"dna/sqlpg"
	"io/ioutil"
)

// RecoverSQLLogError re-executes failed sql queries in sql error log file from specified path.
// It returns the number of failed -reexec queries, and new failed
// queries will be written to the file
//
// The format of error file is:
// 	Error description - $$$error$$$SQL_QUERY$$$error$$$
// Therefore only get statements enclosed by special `$$$error$$$`
func RecoverSQLLogError(sqlErrFilePath dna.String, db *sqlpg.DB) dna.Int {
	var errCount = 0
	var errStrings = dna.StringArray{}
	b, err := ioutil.ReadFile(sqlErrFilePath.String())
	if err != nil {
		panic(err)
	}
	data := dna.String(string(b))
	// dna.Log("\n", data.Length())
	sqlArr := data.FindAllString(`(?mis)\$\$\$error\$\$\$.+?\$\$\$error\$\$\$`, -1)
	// dna.Log("\nTOTAL SQL STATEMENTS FOUND:", sqlArr.Length())
	for _, val := range sqlArr {
		sqlStmtArr := val.FindAllStringSubmatch(`(?mis)\$\$\$error\$\$\$(.+?)\$\$\$error\$\$\$`, -1)
		if len(sqlStmtArr) > 0 {
			_, err := db.Exec(sqlStmtArr[0][1].String())
			if err != nil {
				if dna.String(err.Error()).Contains(`duplicate key value violates unique constraint`) == false {
					errCount += 1
					errStrings.Push("$$$error$$$" + sqlStmtArr[0][1] + "$$$error$$$")
				}
			}
		}
	}
	if errCount == 0 {
		err = ioutil.WriteFile(sqlErrFilePath.String(), []byte{}, 0644)
	} else {
		err = ioutil.WriteFile(sqlErrFilePath.String(), []byte(errStrings.Join("\n").String()), 0644)
	}
	if err != nil {
		panic(err)
	}
	return dna.Int(errCount)
}
