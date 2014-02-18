package utils

import (
	"bufio"
	"dna"
	"dna/sqlpg"
	"errors"
	"os"
	"time"
)

// ForeachLine loops through every line a file.
// An anynomous input function has line, index as params
func ForeachLine(filePath dna.String, lineFunc func(dna.String, dna.Int)) {

	var err error
	var line []byte
	f, err := os.Open(filePath.String())
	if err != nil {
		dna.Log("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)

	i := 0
	err = nil
	for err == nil {
		i += 1
		line, _, err = r.ReadLine()
		lineFunc(dna.String(string(line)), dna.Int(i))
	}
}

func IsValidTable(tblName dna.String, db *sqlpg.DB) dna.Bool {
	_, err := db.Exec("select * from " + tblName.String() + " limit 0")
	if err == nil {
		return true
	} else {
		return false
	}
}

// ToSeconds returns total seconds from the time format "01:02:03"
func ToSeconds(str dna.String) dna.Int {
	if str == "" {
		return 0
	} else {
		intervals := dna.IntArray(str.Split(":").Map(func(val dna.String, idx dna.Int) dna.Int {
			return val.ToInt()
		}).([]dna.Int))
		switch intervals.Length() {
		case 3:
			return intervals[0]*3600 + intervals[1]*60 + intervals[2]
		case 2:
			return intervals[0]*60 + intervals[1]
		case 1:
			return intervals[0]
		default:
			return 0
		}
	}
}

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

// SelectMissingIds accepts a table name as an input and a list of ids as a source.
// It returns a new list of ids that do not exist in the destination table
//
// 	* tblName : a table name
// 	* srcIds : a source ids
// 	* db : a pointer to connected databased
// 	* Returns a new list of ids which are not from the specified table
//
// The format of sql statement is:
// 	WITH dna (id) AS (VALUES (5),(6),(7),(8),(9))
// 	SELECT id FROM dna WHERE NOT EXISTS
// 	(SELECT 1 from ziartists WHERE id=dna.id)
func SelectMissingIds(tblName dna.String, srcIds *dna.IntArray, db *sqlpg.DB) (*dna.IntArray, error) {

	if srcIds.Length() > 0 {

		val := dna.StringArray(srcIds.Map(func(val dna.Int, idx dna.Int) dna.String {
			return "(" + val.ToString() + ")"
		}).([]dna.String))
		selectStmt := "with dna (id) as (values " + val.Join(",") + ") \n"
		selectStmt += "SELECT id FROM dna WHERE NOT EXISTS\n (SELECT 1 from " + tblName + "  WHERE id=dna.id)"
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

// SelectMissingIds accepts a table name as an input and a range as a source.
// It returns a new list of ids that do not exist in the destination table
//
// 	* tblName : a table name
// 	* head, tail : first and last number defines a range
// 	* db : a pointer to connected databased
// 	* Returns a new list of ids which are not from the specified table
//
// The format of sql statement is:
// 	SELECT id FROM generate_series(5,9) id
// 	WHERE NOT EXISTS (SELECT 1 from ziartists where id = id.id)
func SelectMissingIdsWithRange(tblName dna.String, head, tail dna.Int, db *sqlpg.DB) (*dna.IntArray, error) {

	if head > tail {
		panic("Cannot create range: head has to be less than tail")
	}

	selectStmt := dna.Sprintf("SELECT id FROM generate_series(%v,%v) id \n", head, tail)
	selectStmt += "WHERE NOT EXISTS (SELECT 1 from " + tblName + " where id = id.id)"
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

// SelectMissingKeys accepts a table name as an input and a list of keys as a source.
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
func SelectMissingKeys(tblName dna.String, srcKeys *dna.StringArray, db *sqlpg.DB) (*dna.StringArray, error) {
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

// SelectLastMissingIds returns a list of ids which is missing from a table.
// nLastIds is the total number of last ids in the table.
//
// Example: 3 last ids from table nssongs are 1,4,5. Therefore, the missing
// ids whose range is from 1->5 are 2,3.
func SelectLastMissingIds(tblName dna.String, nLastIds dna.Int, db *sqlpg.DB) (*dna.IntArray, error) {
	var min, max dna.Int

	query := dna.Sprintf("SELECT min(id), max(id) FROM (SELECT id FROM %v ORDER BY id DESC LIMIT %v) as AB", tblName, nLastIds)
	// dna.Log(query)
	db.QueryRow(query).Scan(&min, &max)
	// totalItem := max - min + 1
	// var ids = make(dna.IntArray, totalItem, totalItem+100)
	// var idx = 0
	// for i := min; i < max; i++ {
	// 	ids[idx] = i
	// 	idx += 1
	// }
	return SelectMissingIdsWithRange(tblName, min, max, db)
}
