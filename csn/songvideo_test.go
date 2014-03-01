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
	item, err := GetSongVideo(1209100)
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
	if video.Plays < 342712 {
		panic("Plays has to be greater than 342712")
	}
	if video.Downloads < 9858 {
		panic("Plays has to be greater than 9858")
	}
	video.Plays = 342712
	video.Downloads = 9858
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	// video.Formats changing from day to day "1183/3/1182901-658f6751" => `3` means Wed
	video.Formats = "[{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/128/file-name.mp4\",\"type\":\"mp4\",\"file_size\":45440,\"resolution\":\"360p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/320/file-name.mp4\",\"type\":\"mp4\",\"file_size\":67490,\"resolution\":\"480p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/m4a/file-name.mp4\",\"type\":\"mp4\",\"file_size\":122300,\"resolution\":\"720p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/flac/file-name.mp4\",\"type\":\"mp4\",\"file_size\":228310,\"resolution\":\"1080p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/32/file-name.mp4\",\"type\":\"mp4\",\"file_size\":21270,\"resolution\":\"180p\"}]"
	LogStruct(video)
	// Output:
	// Id : 1209100
	// Title : "Con Gái Thời Nay"
	// Artists : dna.StringArray{"Lý Hải", "Bảo Chung"}
	// Authors : dna.StringArray{"Văn Hoà"}
	// Topics : dna.StringArray{"Video Clip", "Việt Nam"}
	// Thumbnail : "http://data.chiasenhac.com/data/thumb/1210/1209100_prv.jpg"
	// Producer : ""
	// Downloads : 9858
	// Plays : 342712
	// Formats : "[{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/128/file-name.mp4\",\"type\":\"mp4\",\"file_size\":45440,\"resolution\":\"360p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/320/file-name.mp4\",\"type\":\"mp4\",\"file_size\":67490,\"resolution\":\"480p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/m4a/file-name.mp4\",\"type\":\"mp4\",\"file_size\":122300,\"resolution\":\"720p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/flac/file-name.mp4\",\"type\":\"mp4\",\"file_size\":228310,\"resolution\":\"1080p\"},{\"link\":\"http://data6.chiasenhac.com/downloads/1210/3/1209100-56d5b4b1/32/file-name.mp4\",\"type\":\"mp4\",\"file_size\":21270,\"resolution\":\"180p\"}]"
	// Href : "http://chiasenhac.com/hd/video/v-video/con-gai-thoi-nay~ly-hai-bao-chung~1209100.html"
	// IsLyric : 1
	// Lyric : "Con gái thời nay trông thật là xinh \nNhư hoa vừa nở ngày xuân đầu mùa\nLàn da trắng tựa như mây \nLại thêm nụ cười duyên dáng \nLàm cho lòng anh đây nhớ.\n\nCon gái thời nay trông thật xì teen \nTung tăng ngoài phố làm bao người nhìn\nNhiều anh cứ muốn làm quen \nLàm sao được lòng em yêu \nÔi con gái nhà ai cưng ghê.\n\nDáng em tựa như nàng tiên đang về trên phố\nThì cho anh đi theo cùng \nAnh đây vẫn còn cô đơn nguyện xin dâng tròn cuộc đời anh \nCon Gái Thời Nay lyrics on ChiaSeNhac.com\nNếu em chưa chồng xin em trả lời đi để anh thương.\n\nCon gái thời nay trông thật là yêu \nYêu em từ phút gặp em lần đâu\nVậy là anh đã gặp may \nĐôi ta là do duyên số \nYêu ai cũng vậy yêu dùm anh đi."
	// DateReleased : "2014"
	// DateCreated : "2014-01-25 01:11:00"
	// Type : false
	// Checktime : "2013-11-21 00:00:00"
}
