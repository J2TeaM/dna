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
	song, err := GetSong(1182753)
	PanicError(err)
	if song.Plays < 274133 {
		panic("Plays has to be greater than 274133")
	}
	if song.Downloads < 6856 {
		panic("Plays has to be greater than 6856")
	}
	song.Plays = 274133
	song.Downloads = 6856
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	// song.Formats changing from day to day "1183/3/1182753-ef690820" => `3` means Wed
	song.Formats = "[{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/128/file-name.mp3\",\"type\":\"mp3\",\"file_size\":3770,\"bitrate\":\"128kbps\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/320/file-name.mp3\",\"type\":\"mp3\",\"file_size\":9400,\"bitrate\":\"320kbps\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/m4a/file-name.m4a\",\"type\":\"m4a\",\"file_size\":11630,\"bitrate\":\"500kbps\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/flac/file-name.flac\",\"type\":\"flac\",\"file_size\":25500,\"bitrate\":\"Lossless\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/32/file-name.m4a\",\"type\":\"m4a\",\"file_size\":1110,\"bitrate\":\"32kbps\"}]"
	LogStruct(song)

	// Output:
	// Id : 1182753
	// Title : "Chỉ Cần Em Vui Để Anh Được Vui"
	// Artists : dna.StringArray{"Ưng Đại Vệ"}
	// Authors : dna.StringArray{"Ưng Đại Vệ"}
	// Topics : dna.StringArray{"Việt Nam", "Pop", "Rock", "Nhạc Trẻ"}
	// AlbumTitle : "Chỉ Cần Em Vui Để Anh Được Vui (Single)"
	// AlbumHref : "http://playlist.chiasenhac.com/nghe-album/chi-can-em-vui-de-anh-duoc-vui~ung-dai-ve~1182753.html"
	// AlbumCoverart : "http://data.chiasenhac.com/data/cover/14/13388.jpg"
	// Producer : "HIT (2013)"
	// Downloads : 6856
	// Plays : 274133
	// Formats : "[{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/128/file-name.mp3\",\"type\":\"mp3\",\"file_size\":3770,\"bitrate\":\"128kbps\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/320/file-name.mp3\",\"type\":\"mp3\",\"file_size\":9400,\"bitrate\":\"320kbps\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/m4a/file-name.m4a\",\"type\":\"m4a\",\"file_size\":11630,\"bitrate\":\"500kbps\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/flac/file-name.flac\",\"type\":\"flac\",\"file_size\":25500,\"bitrate\":\"Lossless\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182753-ef690820/32/file-name.m4a\",\"type\":\"m4a\",\"file_size\":1110,\"bitrate\":\"32kbps\"}]"
	// Href : "http://chiasenhac.com/mp3/vietnam/v-pop/chi-can-em-vui-de-anh-duoc-vui~ung-dai-ve~1182753.html"
	// IsLyric : 1
	// Lyric : "1. Đã không tin một điều\nEm đã nói không còn tình yêu\nMột câu nói khiến anh đau lòng\nMột vết thương không lành sâu trong tim anh.\n\nVì sao vẫn mãi không tin một điều\nLà em kết thúc trong một tình yêu\nLà anh cố chấp đơn phương\nCố gắng thật nhiều vì em...tự dối lòng!\n\n[ĐK:]\nBởi vì một người và chỉ một người anh tin suốt đời\nDù là thiệt thòi nguyện làm tất cả để em được vui\nNước mắt rơi sau tiếng cười\nChỉ cần em vui để anh được vui.\n\nBởi vì một lời và chỉ một lời em nói suốt đời\nLặng thầm đợi chờ một ngày em sẽ thấu hiểu lòng anh\nChỉ Cần Em Vui Để Anh Được Vui lyrics on ChiaSeNhac.com\nVẫn biết giấc mơ đó khó thành\nNhận ra em đã hết yêu anh rồi.\n\n2. Vì sao vẫn mãi không tin một điều\nEm đã nói không còn tình yêu\nMột câu nói khiến anh đau lòng\nMột vết thương không lành sâu trong tim anh."
	// DateReleased : ""
	// DateCreated : "2013-12-11 12:32:00"
	// Type : true
	// Checktime : "2013-11-21 00:00:00"

}
