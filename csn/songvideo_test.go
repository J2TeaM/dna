package csn

import (
	. "dna"
	"testing"
	"time"
)

func TestGetSongVideo(t *testing.T) {
	_, err := GetSongVideo(2172636)
	if err == nil {
		t.Error("Video 2172636 has to have an error")
	}
	if err.Error() != "Chiasenhac - Song 2172636: Mp3 link not found" {
		t.Errorf("Error message has to be: %v", err.Error())
	}
}
func ExampleGetSongVideo() {
	var video *Video
	item, err := GetSongVideo(1182901)
	if err != nil {
		panic(err.Error())
	} else {
		switch item.(type) {
		case Song:
			panic("It has to be video, not song")
		case *Song:
			panic("It has to be video, not song")
		case Video:
			panic("It has to be pointer")
		case *Video:
			video = item.(*Video)
		default:
			panic("no type found")
		}
	}
	PanicError(err)
	if video.Plays < 168297 {
		panic("Plays has to be greater than 168297")
	}
	if video.Downloads < 5541 {
		panic("Plays has to be greater than 5541")
	}
	video.Plays = 168297
	video.Downloads = 5541
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	// video.Formats changing from day to day "1183/3/1182901-658f6751" => `3` means Wed
	video.Formats = "[{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/128/file-name.mp4\",\"type\":\"mp4\",\"file_size\":21720,\"resolution\":\"360p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/320/file-name.mp4\",\"type\":\"mp4\",\"file_size\":31150,\"resolution\":\"480p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/m4a/file-name.mp4\",\"type\":\"mp4\",\"file_size\":52740,\"resolution\":\"720p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/32/file-name.mp4\",\"type\":\"mp4\",\"file_size\":11580,\"resolution\":\"180p\"}]"
	LogStruct(video)
	// Output:
	// Id : 1182901
	// Title : "Tam Giác Tình"
	// Artists : dna.StringArray{"Lâm Chấn Khang", "Saka Trương Tuyền"}
	// Authors : dna.StringArray{"Lại Hoàng Sang"}
	// Topics : dna.StringArray{"Video Clip", "Việt Nam"}
	// Thumbnail : "http://data.chiasenhac.com/data/thumb/1183/1182901_prv.jpg"
	// Producer : "Eye (2013)"
	// Downloads : 5541
	// Plays : 168297
	// Formats : "[{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/128/file-name.mp4\",\"type\":\"mp4\",\"file_size\":21720,\"resolution\":\"360p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/320/file-name.mp4\",\"type\":\"mp4\",\"file_size\":31150,\"resolution\":\"480p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/m4a/file-name.mp4\",\"type\":\"mp4\",\"file_size\":52740,\"resolution\":\"720p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1183/3/1182901-658f6751/32/file-name.mp4\",\"type\":\"mp4\",\"file_size\":11580,\"resolution\":\"180p\"}]"
	// Href : "http://chiasenhac.com/hd/video/v-video/tam-giac-tinh~lam-chan-khang-saka-truong-tuyen~1182901.html"
	// IsLyric : 1
	// Lyric : "Đớn đau lòng em khi giờ đây em nhận ra \nNgười em yêu giờ đã có một hạnh phúc\nXin em đừng nói những lời buồn đau\nĐừng để cho hai ta phải mất nhau.\n\nEm lo sợ một ngày người tình ơi\nCuộc tình ngang trái gieo yêu đau cho người đến sau\nThật tội nghiệp em giữa chốn hoang tình \nThế nhưng em vẫn đứng sau tình yêu.\n\nCuộc tình của anh có phải em đây chia lìa\nThật lòng giờ đây xin lỗi anh vì quá yêu anh.\n\n[ĐK:]\nĐừng nghi oan nữa đừng khóc em ơi\nTam Giác Tình lyrics on ChiaSeNhac.com\nĐừng trách than chỉ thêm đau lòng anh\nVì yêu nên anh không muốn cho em \nMang buồn đau tổn thương.\n\nSao duyên ta không tròn \nKhi bên anh em đứng sau một người\nĐường hoài nghi nữa anh chỉ yêu em\nYêu em xin anh đừng bên ai."
	// DateReleased : ""
	// DateCreated : "2013-12-11 16:50:00"
	// Type : false
	// Checktime : "2013-11-21 00:00:00"
}
