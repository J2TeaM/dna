package csn

import (
	. "dna"
	"testing"
	"time"
)

func TestGetVideo(t *testing.T) {
	_, err := GetVideo(1190840)
	if err == nil {
		t.Error("Video 1190840 has to have an error")
	}
	// if err.Error() != "It has to be video, not song" {
	// 	t.Errorf("Error message has to be: %v", err.Error())
	// }
}

func ExampleGetVideo() {
	video, err := GetVideo(1213739)
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
	if video.Formats == "" || video.Formats.Count("http") != 5 {
		panic("Wrong formats")
	}
	// video.Formats changing from day to day "1183/3/1182901-658f6751" => `3` means Wed
	video.Formats = "{\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/128/file-name.mp4,mp4,35290,360)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/320/file-name.mp4,mp4,51240,480)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/m4a/file-name.mp4,mp4,88840,720)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/flac/file-name.mp4,mp4,167770,1080)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/32/file-name.mp4,mp4,17060,180)\"}"
	LogStruct(video)
	// Output:
	// Id : 1213739
	// Title : "Thương Vợ"
	// Artists : dna.StringArray{"Lý Hải"}
	// Authors : dna.StringArray{"Phi Bằng"}
	// Topics : dna.StringArray{"Video Clip", "Việt Nam"}
	// Thumbnail : "http://data.chiasenhac.com/data/thumb/1214/1213739_prv.jpg"
	// Producer : ""
	// Downloads : 5541
	// Plays : 168297
	// Formats : "{\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/128/file-name.mp4,mp4,35290,360)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/320/file-name.mp4,mp4,51240,480)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/m4a/file-name.mp4,mp4,88840,720)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/flac/file-name.mp4,mp4,167770,1080)\",\"(http://data6.chiasenhac.com/downloads/1214/3/1213739-a77ed0e5/32/file-name.mp4,mp4,17060,180)\"}"
	// Href : "http://chiasenhac.com/hd/video/v-video/thuong-vo~ly-hai~1213739.html"
	// IsLyric : 1
	// Lyric : "Muốn chơi cho hoài tôi cứ để vợ lo hoài\nSáng trưa hay chiều cả ngày tôi cứ nhậu say\nTội nghiệp vợ tôi 1 lòng 1 dạ với tôi\nHôm sớm lo cho chồng chưa 1 lần vợ than bất công.\n\nCó khi ra hoài tôi cứ lén vợ ra ngoài \nVới bao cô nàng tóc dài xoả tới ngan vai\nTại gì tôi say nên chẳng lấy lòng được ai\nMới hiểu ra chân tình chỉ có vợ là yêu mình.\n \n[ĐK:]\nƠi vợ vợ ơi...sao mà em hổng cười\nƠi vợ vợ ơi...thương vợ nhất trên đời\nAnh thề từ nay sẽ không còn nhậu say\nThương Vợ lyrics on ChiaSeNhac.com\nAnh thề từ đây sẽ không để ý ai.\n\nƠi vợ vợ ơi...bây giờ anh hiểu rồi\nAnh thiệt không nên,để vợ khổ 1 đời\nAnh thề từ nay chăm lo dựng lại tương lai\nAnh thề suốt kiếp trọn đời không hề đổi thay."
	// DateReleased : "2014"
	// DateCreated : "2014-02-06 10:47:00"
	// Type : false
	// Checktime : "2013-11-21 00:00:00"
}
