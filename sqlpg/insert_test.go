package sqlpg

import (
	. "dna"
	"dna/ns"
)

func ExampleGetInsertStatement() {
	album := ns.NewAlbum()
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
	// 	INSERT INTO test
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
	// $binhdna$0001-01-01 00:00:00$binhdna$
	// )
}
