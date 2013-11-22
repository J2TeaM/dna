package sqlpg

import (
	. "dna"
	"dna/ns"
	"testing"
)

func TestDB(t *testing.T) {
	song := ns.NewSong()
	song.Id = 2
	song.Title = "Example title"
	song.Artists = StringArray{"First artists", "Second Artitsts"}
	db, err := Connect(DefaultConfig)
	if err != nil {
		t.Error("DB has to have no connection error")
	} else {
		insertErr := db.Insert(song)
		if insertErr != nil {
			t.Error("Insert has to be complete")
		}
	}

	insertIgnoreErr := db.InsertIgnore(song)
	if insertIgnoreErr != nil {
		t.Error("Insert has to be ignored")
	}

	song.Artists = StringArray{"Third artists", "Fourth Artitsts"}
	song.Authors = StringArray{"My authors"}
	updateErr := db.Update(song, "id", "artists", "authors")
	if updateErr != nil {
		t.Error("update has to be updated")
	}

	rows, selectError := db.Query("Select * from nssongs where id=2")
	if selectError != nil {
		t.Error("select has to be done")
	}

	for rows.Next() {
		song1 := ns.NewSong()
		err = rows.StructScan(song1)
		if err != nil {
			t.Error("Row has to be scan")
		}

		if song1.Artists[0] != "Third artists" || song1.Artists[1] != "Fourth Artitsts" {
			t.Error("artists wrong")
		}

		if song1.Id != 2 {
			t.Error("wrong song id")
		}

		if song1.Authors.Length() != 1 && song1.Authors[0] != "My authors" {
			t.Error("wrong authors")
		}

		if song1.Title != "Example title" {
			t.Error("wrong title")
		}

		if !song1.DateUpdated.IsZero() {
			t.Error("date udpated hs to be zero")
		}
	}
	_, queryErr := db.Query("Delete from nssongs where id=2")
	if queryErr != nil {
		t.Error("query has to be done")
	}
}

func ExampleDB() {
	// CODE is exactly the same as the one in TestDB()
	// Initialize some fake values for song
	song := ns.NewSong()
	song.Id = 2
	song.Title = "Example title"
	song.Artists = StringArray{"First artists", "Second Artitsts"}
	db, err := Connect(DefaultConfig)
	if err != nil {
		panic("DB has to have no connection error")
	}

	// Insert a new song
	insertErr := db.Insert(song)
	if insertErr != nil {
		panic("Insert has to be complete")
	}

	// Insert and ignore the available song
	insertIgnoreErr := db.InsertIgnore(song)
	if insertIgnoreErr != nil {
		panic("Insert has to be ignored")
	}

	// Update fields: aritsts & authors with primary key: id
	song.Artists = StringArray{"Third artists", "Fourth Artitsts"}
	song.Authors = StringArray{"My authors"}
	updateErr := db.Update(song, "id", "artists", "authors")
	if updateErr != nil {
		panic("update has to be updated")
	}

	// Select a row based on the id of the song above
	rows, selectError := db.Query("Select * from nssongs where id=2")
	if selectError != nil {
		panic("select has to be done")
	}
	for rows.Next() {
		scannedSong := ns.NewSong()
		err = rows.StructScan(scannedSong)
		if err != nil {
			panic("Row has to be scan")
		}
		if scannedSong.Artists[0] != "Third artists" || scannedSong.Artists[1] != "Fourth Artitsts" {
			panic("artists wrong")
		}
		if scannedSong.Id != 2 {
			panic("wrong song id")
		}
		if scannedSong.Authors.Length() != 1 && scannedSong.Authors[0] != "My authors" {
			panic("wrong authors")
		}
		if scannedSong.Title != "Example title" {
			panic("wrong title")
		}
		if !scannedSong.DateUpdated.IsZero() {
			panic("date udpated hs to be zero")
		}
	}

	// Delete a row
	_, queryErr := db.Query("Delete from nssongs where id=2")
	if queryErr != nil {
		panic("query has to be done")
	}
}
