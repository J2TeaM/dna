package csn

import (
	. "dna"
	"testing"
	"time"
)

func TestGetSong(t *testing.T) {
	_, err := GetSong(2171936)
	if err == nil {
		t.Error("Song 2171936 has to have an error")
	}
	if err.Error() != "Chiasenhac - Song 2171936: Mp3 link not found" {
		t.Errorf("Error message has to be: %v", err.Error())
	}
}
func ExampleGetSong() {
	song, err := GetSong(1222808)
	PanicError(err)
	if song.Plays < 1205 {
		panic("Plays has to be greater than 1205")
	}
	if song.Downloads < 104 {
		panic("Plays has to be greater than 104")
	}
	if song.AlbumCoverart == "" {
		panic("Has to have coverart")
	}
	if song.Formats == "" || song.Formats.Count("http") != 2 {
		panic("Has to have Formats")
	}

	if song.AlbumHref == "" {
		panic("Album href have value")
	}
	song.AlbumHref = "http://playlist.chiasenhac.com/nghe-album/hoa-no-ve-dem~giang-tu-huong-lan~1222808.html"
	song.AlbumCoverart = "http://data.chiasenhac.com/data/cover/17/16766.jpg"
	song.Plays = 1205
	song.Downloads = 104
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	// song.Formats changing from day to day "1183/3/1182753-ef690820" => `3` means Wed
	song.Formats = "{\"(http://data6.chiasenhac.com/downloads/1223/3/1222808-e6d92163/128/file-name.mp3,mp3,5520,128)\",\"(http://data6.chiasenhac.com/downloads/1223/3/1222808-e6d92163/32/file-name.m4a,m4a,1520,32)\"}"
	LogStruct(song)

	// Output:
	// Id : 1222808
	// Title : "Hoa Nở Về Đêm"
	// Artists : dna.StringArray{"Giang Tử", "Hương Lan"}
	// Authors : dna.StringArray{"Mạnh Phát"}
	// Topics : dna.StringArray{"Việt Nam", "Pop", "Rock", "Trữ Tình"}
	// AlbumTitle : "Duyên Quê"
	// AlbumHref : "http://playlist.chiasenhac.com/nghe-album/hoa-no-ve-dem~giang-tu-huong-lan~1222808.html"
	// AlbumCoverart : "http://data.chiasenhac.com/data/cover/17/16766.jpg"
	// Producer : "Thuý Nga TNCD533 (2014)"
	// Downloads : 104
	// Plays : 1205
	// Formats : "{\"(http://data6.chiasenhac.com/downloads/1223/3/1222808-e6d92163/128/file-name.mp3,mp3,5520,128)\",\"(http://data6.chiasenhac.com/downloads/1223/3/1222808-e6d92163/32/file-name.m4a,m4a,1520,32)\"}"
	// Href : "http://chiasenhac.com/mp3/vietnam/v-pop/hoa-no-ve-dem~giang-tu-huong-lan~1222808.html"
	// IsLyric : 1
	// Lyric : "Chuyện từ một đêm cuối nẻo một người tiễn một người đi\nĐẹp tựa bài thơ nở giữa đêm sương nở tận tâm hồn\nChuyện một mình tôi chép dòng tâm tình tặng người chưa biết một lần\nVì giây phút ấy tôi tình cờ hiểu rằng\nTình yêu đẹp nghìn đời là tình yêu khi đơn côi.\n\nVào chuyện từ một đêm khoác bờ vai một mảnh áo dạ đường khuya\nBồi hồi người trai hướng nẻo đêm sâu, dấu tình yêu đầu\nVì còn tìm nhau lối về ngõ hẹp còn chờ in dấu chân anh\nNiềm thương mến đó bây giờ và nghìn đời\nDù gió đùa dạt dào còn đẹp như khi quen nhau.\n\n[ĐK:]\nAi lớn lên không từng hẹn hò không từng yêu thương\nNhưng có mấy người tìm được một tình yêu ngát hương\nMến những người chưa quen\nHoa Nở Về Đêm lyrics on ChiaSeNhac.com\nMột người ở lại đèn trăng gối mộng\nYêu ai anh băng sông dài cho đẹp lòng trai.\n\nMột người tìm vui mãi tận trời nào giá lạnh hồn đông\nMột người chợt nghe gió giữa mênh mông rót vào trong lòng\nVà một mình tôi chép dòng tâm tình tặng người chưa biết một lần\nVì trong phút ấy tôi tìm mình thì thầm giờ đã gặp được một nụ hoa nở về đêm."
	// DateReleased : ""
	// DateCreated : "2014-02-22 18:42:00"
	// Type : true
	// Checktime : "2013-11-21 00:00:00"
}
