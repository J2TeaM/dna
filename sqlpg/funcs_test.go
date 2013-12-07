package sqlpg

import (
	. "dna"
	"time"
)

func ExampleGetInsertStatement() {
	album := NewAlbum()
	album.Id = 359294
	album.Title = "Voices Of Romance"
	album.Artists = StringArray{"Various Artists"}
	album.Topics = StringArray{"Nhạc Các Nước Khác"}
	album.Coverart = "http://st.nhacso.net/images/album/2012/07/08/1202066467/13417589320_6160_120x120.jpg"
	album.Songids = IntArray{1217599, 1217600, 1217601, 1217602, 1217603, 1217604, 1217605, 1217606, 1217607, 1217608, 1217609, 1217610, 1217611, 1217612, 1217613, 1217614, 1217615}
	album.Description = "Âm nhạc luôn là nơi chấp cánh tình yêu. Đến với VOICES OF ROMANCE bạn sẽ cảm nhận được sự lãng mạn của tình yêu, sự thăng hoa của cảm xúc, sự cô đơn của chia ly... Các cung bậc của tình yêu đều thể hiện rõ qua từng bài hát."
	album.DateReleased = "2007"
	Log(GetInsertStatement("test", album, true))

	//Output:
	// INSERT INTO test
	// (id,title,artists,artistid,topics,genres,category,coverart,nsongs,plays,songids,description,label,date_released,checktime)
	// VALUES (
	// 359294,
	// $binhdna$Voices Of Romance$binhdna$,
	// $binhdna${"Various Artists"}$binhdna$,
	// 0,
	// $binhdna${"Nhạc Các Nước Khác"}$binhdna$,
	// $binhdna${}$binhdna$,
	// $binhdna${}$binhdna$,
	// $binhdna$http://st.nhacso.net/images/album/2012/07/08/1202066467/13417589320_6160_120x120.jpg$binhdna$,
	// 0,
	// 0,
	// $binhdna${1217599, 1217600, 1217601, 1217602, 1217603, 1217604, 1217605, 1217606, 1217607, 1217608, 1217609, 1217610, 1217611, 1217612, 1217613, 1217614, 1217615}$binhdna$,
	// $binhdna$Âm nhạc luôn là nơi chấp cánh tình yêu. Đến với VOICES OF ROMANCE bạn sẽ cảm nhận được sự lãng mạn của tình yêu, sự thăng hoa của cảm xúc, sự cô đơn của chia ly... Các cung bậc của tình yêu đều thể hiện rõ qua từng bài hát.$binhdna$,
	// $binhdna$$binhdna$,
	// $binhdna$2007$binhdna$,
	// NULL
	// )
}

func ExampleGetTableName() {
	table := GetTableName(NewAlbum())
	table1 := GetTableName(NewSong())
	Log(table)
	Log(table1)
	// Output:
	// sqlpgalbums
	// sqlpgsongs
}

type Album struct {
	Id           Int
	Title        String
	Artists      StringArray
	Artistid     Int
	Topics       StringArray
	Genres       StringArray
	Category     StringArray
	Coverart     String
	Nsongs       Int
	Plays        Int
	Songids      IntArray
	Description  String
	Label        String
	DateReleased String
	Checktime    time.Time
}

// NewAlbum return default new album
func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Title = ""
	album.Artists = StringArray{}
	album.Artistid = 0
	album.Topics = StringArray{}
	album.Genres = StringArray{}
	album.Category = StringArray{}
	album.Coverart = ""
	album.Nsongs = 0
	album.Plays = 0
	album.Songids = IntArray{}
	album.Description = ""
	album.Label = ""
	album.DateReleased = ""
	album.Checktime = time.Time{}
	return album
}
func ExampleGetUpdateStatement() {
	type Test struct {
		Album
		BoolValue  Bool  // Type dna.Bool
		FloatValue Float // Type dna.Float
	}
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	// Update different types from testStruct demo
	testStruct := &Test{*NewAlbum(), true, 10.1 / 3}
	testStruct.Id = 345399                                      // Type dna.Int
	testStruct.Title = "NEW TITLE"                              // Type dna.String
	testStruct.Artists = StringArray{"FIRST", "SECOND"}         // Type dna.StringArray
	testStruct.Artistid = 999                                   // Type dna.Int
	testStruct.Songids = IntArray{1, 2, 3, 4}                   // Type dna.IntArray
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)") // Type time.Time
	testStruct.Checktime = t
	ret, err := GetUpdateStatement("test", testStruct, "id", "title", "artists", "artistid", "bool_value", "float_value", "songids", "checktime")
	if err != nil {
		panic(err.Error())
	}
	Log(ret)
	//Output:
	// UPDATE test SET
	// title=$binhdna$NEW TITLE$binhdna$,
	// artists=$binhdna${"FIRST", "SECOND"}$binhdna$,
	// artistid=999,
	// bool_value=true,
	// float_value=3.3666666666666667,
	// songids=$binhdna${1, 2, 3, 4}$binhdna$,
	// checktime=$binhdna$2013-02-03 19:54:00$binhdna$
	// WHERE id=345399
}
