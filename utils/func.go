package utils

import (
	"dna"
	"dna/sqlpg"
	"errors"
	"time"
)

// GetMaxId returns max id of a specified table.
func GetMaxId(tableName dna.String, db *sqlpg.DB) (dna.Int, error) {
	var maxid dna.Int
	err := db.QueryRow("SELECT max(id) FROM " + tableName).Scan(&maxid)
	switch {
	case err == sqlpg.ErrNoRows:
		return 0, err
	case err != nil:
		return 0, err
	default:
		return maxid, nil
	}
}

// SelectUnavailableIds accepts a table name as an input and a list of ids as a source.
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
func SelectUnavailableIds(tblName dna.String, srcIds *dna.IntArray, db *sqlpg.DB) (*dna.IntArray, error) {

	if srcIds.Length() > 0 {

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
	} else {
		return nil, errors.New("Empty input array")
	}

}

// SelectUnavailableKeys accepts a table name as an input and a list of keys as a source.
// It returns a new list of keys that do not exist in the destination table
//
// 	* tblName : a table name
// 	* srcKeys : a source keys
// 	* db : a pointer to connected databased
// 	* Returns a new list of keys which are not from the specified table
//
// Notice: Only applied to a table having a column named "key".
// The column has to be indexed to ensure good performance
//
// The format of sql statement is:
//	with dna (key) as (values ('43f3HhhU6DGV'),('uFfgQhKbwAfN'),('RvFDlckJB5QU'),('uIF7rwd5wo4p'),('Kveukbhre1ry'),('oJ1lzAlKwJX6'),('43f3HhhU6DGV'),('uFfgQhKbwAfN'),('hfhtyMdywMau'),('PpZuccjYqy1b'))
//	select key from dna where key not in
//	(select key from nctalbums where key in ('43f3HhhU6DGV','uFfgQhKbwAfN','RvFDlckJB5QU','uIF7rwd5wo4p','Kveukbhre1ry','oJ1lzAlKwJX6','43f3HhhU6DGV','uFfgQhKbwAfN','hfhtyMdywMau','PpZuccjYqy1b'))
func SelectUnavailableKeys(tblName dna.String, srcKeys *dna.StringArray, db *sqlpg.DB) (*dna.StringArray, error) {
	if srcKeys.Length() > 0 {
		val := dna.StringArray(srcKeys.Map(func(val dna.String, idx dna.Int) dna.String {
			return `('` + val + `')`
		}).([]dna.String))
		val1 := dna.StringArray(srcKeys.Map(func(val dna.String, idx dna.Int) dna.String {
			return `'` + val + `'`
		}).([]dna.String))
		selectStmt := "with dna (key) as (values " + val.Join(",") + ") \n"
		selectStmt += "select key from dna where key not in \n(select key from " + tblName + " where key in (" + val1.Join(",") + "))"
		keys := &[]dna.String{}
		err := db.Select(keys, selectStmt)
		switch {
		case err != nil:
			return nil, err
		case err == nil && keys != nil:
			slice := dna.StringArray(*keys)
			return &slice, nil
		case err == nil && keys == nil:
			return &dna.StringArray{}, nil
		default:
			panic("Default case triggered. Case is not expected. Cannot select non existed keys")
		}
	} else {
		return nil, errors.New("Empty input array")
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
