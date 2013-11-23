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

func ExampleDB_Select() {
	db, err := Connect(DefaultConfig)
	if err != nil {
		panic(err)
	}
	ids := &[]Int{}
	selectErr := db.Select(ids, "SELECT id FROM nssongs ORDER BY id ASC LIMIT 10")
	if selectErr != nil {
		Log(selectErr.Error())
	} else {
		Log("Length of ids:", len(*ids))
		Log(*ids)
	}

	songs := &[]ns.Song{}
	selectErr1 := db.Select(songs, "SELECT * FROM nssongs ORDER BY id ASC LIMIT 2")
	Log("------------------------------------")
	if selectErr1 != nil {
		Log(selectErr1.Error())
	} else {
		Log("Length of songs:", len(*songs))
		for _, song := range *songs {
			LogStruct(&song)
			Log("------------------------------------")
		}
	}
	// Output:
	// Length of ids: 10
	// []dna.Int{1, 6, 10, 14, 15, 23, 27, 28, 31, 32}
	// ------------------------------------
	// Length of songs: 2
	// Id : 1
	// Title : "Chờ Em Trong Mỏi Mòn"
	// Artists : dna.StringArray{"Ngô Kiến Huy"}
	// Artistid : 2228
	// Authors : dna.StringArray{"Phúc Trường"}
	// Authorid : 2355
	// Plays : 62495
	// Duration : 248
	// Link : "http://st01.freesocialmusic.com/mp3/2012/08/22/1060052334/13456211648_6232.mp3"
	// Topics : dna.StringArray{"Nhạc Trẻ"}
	// Category : dna.StringArray(nil)
	// Bitrate : 320
	// Official : 1
	// Islyric : 1
	// DateCreated : "2012-08-22 14:39:24"
	// DateUpdated : "2013-01-17 21:29:14"
	// Lyric : "<p>Chợt nghe con tim nhói đau<br/>Khi em vội vàng bỏ lại phía sau<br/>Là anh là những hy vọng<br/>Trong màn đêm, anh bước lẽ loi<br/><br/>Về đâu khi không có em<br/>Về đâu khi đêm vắng tanh<br/>Mình anh bơ vơ trong nỗi hiu quạnh<br/>Đôi môi rét xuông lạnh câm….<br/><br/>ĐK:<br/>Chờ em trong mỏi mòn<br/>Và tình yêu xưa có còn<br/>Hỡi em!!! Sao không quay lại<br/>Đừng bỏ mặc anh trong cô đơn<br/>Khi nỗi nhớ tìm về<br/>Với anh trong đêm lạnh giá<br/><br/>Đợi em anh sẽ đợi<br/>Và chờ em anh mãi chờ<br/>Vẫn biết em sẽ không quay về<br/>Chuyện tình đôi ta nay chia ly<br/>Em xa mãi cuộc đời anh<br/>Yêu thương đã quá tầm tay…….<br/><br/>{Biết em không về lòng vẫn mang câu thề<br/>Trái tim rơi lệ tìm đâu những đam mê<br/>Hơ hờ hơơơ……Hờ hơ hớ hờ<br/>Về với anh trong đêm lạnh giá……)<br/><br/>Đợi em anh sẽ đợi<br/>Và chờ em anh mãi chờ<br/>Vẫn biết em sẽ không quay về<br/>Chuyện tình đôi ta nay chia ly<br/>Em xa mãi cuộc đời anh<br/>Yêu thương đã quá tầm tay<br/></p>"
	// SameArtist : 0
	// Checktime : "2013-02-07 00:15:58"
	// ------------------------------------
	// Id : 6
	// Title : "Mưa Kí Ức"
	// Artists : dna.StringArray{"Anh Quốc"}
	// Artistid : 4564
	// Authors : dna.StringArray(nil)
	// Authorid : 0
	// Plays : 24112
	// Duration : 252
	// Link : "http://st01.freesocialmusic.com/mp3/2010/11/20/1472051264/12904010156_6697.mp3"
	// Topics : dna.StringArray{"Nhạc Trẻ"}
	// Category : dna.StringArray(nil)
	// Bitrate : 128
	// Official : 0
	// Islyric : 1
	// DateCreated : "2010-11-22 11:43:35"
	// DateUpdated : "2013-01-17 21:29:14"
	// Lyric : "<p>Có gió mưa từng cơn trong lòng, bóng ai xa tầm tay<br />Đêm quất quây mình anh cô đơn, khóc vì em bước đi<br />Đếm tiếng mưa ngoài hiên vô tình, nghe cay đôi mắt buồn<br />Gió cuốn theo dòng đời ngược xuôi, khiến phút giây lạnh câm…<br /><br />Giờ có anh ngồi nhớ dĩ vãng, giờ có anh ngồi mơ<br />Giờ có anh ngồi khóc, bao đêm cô đơn nghe tiếng mưa âm thầm<br />Vì em yêu ra đi, vì đôi ta chia ly<br />Hạnh phúc có không người ơi!<br /><br />Tìm kiếm bao ngày tháng, tìm kiếm bóng hình ai<br />Tìm kiếm trong hồi ức, khi đôi ta chung bước trên con đường<br />Giờ cơn mưa rơi rơi, mà vì sao hôm nay<br />Chỉ riêng bóng anh mong chờ…<br /><br />Người vì sao nỡ dối gian nhau, tình giờ đây đã quá hư hao<br />Mong chờ vô vọng, ưu phiền trong lòng, làm sao để cố quên em…<br /></p>"
	// SameArtist : 0
	// Checktime : "2013-02-07 00:15:58"
	// ------------------------------------
}
