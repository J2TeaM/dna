package nct

import (
	. "dna"
	"time"
)

func ExampleGetSong() {
	song, err := GetSong(2870710)
	PanicError(err)
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if song.Plays < 261 {
		panic("Plays has to be greater than 261")
	}
	song.Plays = 261
	song.Likes = 0
	if song.StreamUrl.Match(`\.mp3`) == false || song.Linkdown.Match(`\.mp3`) == false || song.LinkdownHQ.Match(`\.mp3`) == false {
		panic("Song link has to be mp3 format")
	} else {
		song.StreamUrl = "http://a.nixcdn.com/cb64c1030d184db608e7629ebfa2cb55/530187c7/NhacCuaTui845/HuongSenDongThapTanCo-TrongHuuKieuHoa-2870710.mp3"
		song.Linkdown = "http://a.nixcdn.com/cb64c1030d184db608e7629ebfa2cb55/530187c7/NhacCuaTui845/HuongSenDongThapTanCo-TrongHuuKieuHoa-2870710.mp3"
		song.LinkdownHQ = "http://download.a.nixcdn.com/cb64c1030d184db608e7629ebfa2cb55/530187c7/NhacCuaTui845/HuongSenDongThapTanCo-TrongHuuKieuHoa-2870710.mp3"
	}
	LogStruct(song)
	// Output:
	// Id : 2870710
	// Key : "WviZ6aais76C"
	// Title : "Hương Sen Đồng Tháp (Tân Cổ)"
	// Artists : dna.StringArray{"Trọng Hữu", "Kiều Hoa"}
	// Topics : dna.StringArray{"Thể Loại Khác"}
	// LinkKey : "f27678b52583250ee3e67b13f9e795f5"
	// Type : "SONG"
	// Bitrate : 128
	// Official : false
	// Likes : 0
	// Plays : 261
	// LinkShare : "http://www.nhaccuatui.com/bai-hat/huong-sen-dong-thap-tan-co-trong-huu-kieu-hoa.WviZ6aais76C.html"
	// StreamUrl : "http://a.nixcdn.com/cb64c1030d184db608e7629ebfa2cb55/530187c7/NhacCuaTui845/HuongSenDongThapTanCo-TrongHuuKieuHoa-2870710.mp3"
	// Image : "http://avatar.nct.nixcdn.com/singer/avatar/2013/8393_Trong_Huu_A.jpg"
	// Coverart : ""
	// Duration : 406
	// Linkdown : "http://a.nixcdn.com/cb64c1030d184db608e7629ebfa2cb55/530187c7/NhacCuaTui845/HuongSenDongThapTanCo-TrongHuuKieuHoa-2870710.mp3"
	// LinkdownHQ : "http://download.a.nixcdn.com/cb64c1030d184db608e7629ebfa2cb55/530187c7/NhacCuaTui845/HuongSenDongThapTanCo-TrongHuuKieuHoa-2870710.mp3"
	// Lyricid : 0
	// HasLyric : false
	// Lyric : ""
	// LyricStatus : 0
	// HasLrc : false
	// Lrc : ""
	// LrcUrl : ""
	// UsernameCreated : ""
	// Checktime : "2013-11-21 00:00:00"
}
