package sqlpg

import (
	. "dna"
	"dna/ns"
)

func ExampleRows_StructScan() {
	db, err := Connect(NewSQLConfig("./config.ini"))
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("Select * from nsalbums WHERE id=817916")
	if err != nil {
		Log(err)
	} else {
		for rows.Next() {
			album := ns.NewAlbum()
			err = rows.StructScan(album)
			if err != nil {
				Log(err)
			}
			LogStruct(album)
		}
	}
	// Output:
	// Id : 817916
	// Title : "Full Frequency"
	// Artists : dna.StringArray{"Sean Paul"}
	// Artistid : 3427
	// Topics : dna.StringArray{"Nhạc Âu Mỹ"}
	// Genres : dna.StringArray{"Pop", "Music"}
	// Category : dna.StringArray{"Nhạc Âu Mỹ", "Dance", "Electronic"}
	// Coverart : "http://st.nhacso.net/images/album/2013/11/07/1154016226/138382977413_3076_120x120.jpg"
	// Nsongs : 14
	// Plays : 115
	// Songids : dna.IntArray{1312944, 1301320, 1312945, 1312946, 1312947, 1312948, 1312949, 1312950, 1312951, 1312952, 1312953, 1312954, 1312955, 1309287}
	// Description : "Genres: Pop, Music\r\nExpected Release: 21 February 2014\r\n℗ 2013 Atlantic Recording Corporation for the United"
	// Label : "Atlantic Recording Corporation for the United"
	// DateReleased : "2013"
	// Checktime : "2013-11-10 01:05:01.012776"
}
