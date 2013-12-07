package utils

import (
	"dna"
	"dna/sqlpg"
	"time"
)

// SelectNonExistedIds accepts a table name as an input and a list of ids as a source.
// It returns a new list of ids that do not exist in the destination table
//
// 	* tblName : a table name
// 	* srcIds : a source ids
// 	* db : a pointer to connected databased
// 	* Returns a new list of ids which are not from the specified table
//
// The format of sql statement is:
// 	with dna (id) as (values (5),(6),(7),(8),(9))
// 	select id from dna where id not in
// 	(select id from ziartists where id in (5,6,7,8,9))
func SelectNonExistedIds(tblName dna.String, srcIds *dna.IntArray, db *sqlpg.DB) (*dna.IntArray, error) {
	val := dna.StringArray(srcIds.Map(func(val dna.Int, idx dna.Int) dna.String {
		return "(" + val.ToString() + ")"
	}).([]dna.String))
	selectStmt := "with dna (id) as (values " + val.Join(",") + ") \n"
	selectStmt += "select id from dna where id not in \n(select id from " + tblName + " where id in (" + srcIds.Join(",") + "))"
	ids := &[]dna.Int{}
	err := db.Select(ids, selectStmt)
	switch {
	case err != nil:
		return nil, err
	case err == nil && ids != nil:
		slice := dna.IntArray(*ids)
		return &slice, nil
	case err == nil && ids == nil:
		return &dna.IntArray{}, nil
	default:
		panic("Default case triggered. Case is not expected. Cannot select non existed ids")
	}
}

// SelectNewSidsFromAlbums returns a slice  of songids from a table since the last specified time.
// The table has to be album type and has a column called songids.
func SelectNewSidsFromAlbums(tblName dna.String, lastTime time.Time, db *sqlpg.DB) *dna.IntArray {
	idsArrays := &[]dna.IntArray{}
	year := dna.Sprintf("%v", lastTime.Year())
	month := dna.Sprintf("%d", lastTime.Month())
	day := dna.Sprintf("%v", lastTime.Day())
	checktime := dna.Sprintf("'%v-%v-%v'", year, month, day)
	query := dna.Sprintf("SELECT songids FROM %s WHERE checktime >= %s", tblName, checktime)
	// dna.Log(query)
	err := db.Select(idsArrays, query)
	dna.PanicError(err)
	ids := &dna.IntArray{}
	if idsArrays != nil {
		for _, val := range *idsArrays {
			for _, id := range val {
				ids.Push(id)
			}
		}
		return ids
	} else {
		return nil
	}

}
