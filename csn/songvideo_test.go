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
	item, err := GetSongVideo(1216124)
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
	if video.Plays < 32778 {
		panic("Plays has to be greater than 32778")
	}
	if video.Downloads < 9858 {
		panic("Plays has to be greater than 9858")
	}
	video.Plays = 32778
	video.Downloads = 9858
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if video.Formats.Length() == 0 || video.Formats.Count("http") != 5 {
		panic("Video has to have formats")
	}
	// video.Formats changing from day to day "1183/3/1182901-658f6751" => `3` means Wed
	video.Formats = "{\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/128/file-name.mp4,mp4,10810,360)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/320/file-name.mp4,mp4,14890,480)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/m4a/file-name.mp4,mp4,24820,720)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/flac/file-name.mp4,mp4,48740,1080)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/32/file-name.mp4,mp4,5630,180)\"}"
	LogStruct(video)
	// Output:
	// Id : 1216124
	// Title : "Mình Yêu Nhau Đi"
	// Artists : dna.StringArray{"Bích Phương"}
	// Authors : dna.StringArray{"Tiên Cookie"}
	// Topics : dna.StringArray{"Video Clip", "Việt Nam"}
	// Thumbnail : "http://data.chiasenhac.com/data/thumb/1217/1216124_prv.jpg"
	// Producer : "RED Team (2014)"
	// Downloads : 9858
	// Plays : 32778
	// Formats : "{\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/128/file-name.mp4,mp4,10810,360)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/320/file-name.mp4,mp4,14890,480)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/m4a/file-name.mp4,mp4,24820,720)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/flac/file-name.mp4,mp4,48740,1080)\",\"(http://data6.chiasenhac.com/downloads/1217/3/1216124-65a0d4d4/32/file-name.mp4,mp4,5630,180)\"}"
	// Href : "http://chiasenhac.com/hd/video/v-video/minh-yeu-nhau-di~bich-phuong~1216124.html"
	// IsLyric : 1
	// Lyric : "Hình như anh có điều muốn nói\nCứ ngập ngừng rồi thôi\nVà có lẽ anh không biết rằng em cũng đang chờ đợi.\n\nỞ cạnh bên anh bình yên lắm\nAnh hiền lành ấm áp\nCứ tiếp tục ngại ngùng thì ai sẽ là người đầu tiên nói ra?\n\n[ĐK 1:]\nHay là mình cứ bất chấp hết yêu nhau đi\nHay để chắc chắn anh cứ lắng nghe tim muốn gì\nRồi nói cho em nghe\nMột câu thôi.\n\n1, 2, 3, 5... anh có đánh rơi nhịp nào không?\nNếu câu trả lời là có anh hãy đến ôm em ngay đi\nEm đã chờ đợi từ anh giây phút ấy cũng lâu lắm rồi\nVà dẫu cho mai sau có ra sao\nThì em vẫn sẽ không hối tiếc vì ngày hôm nay đã nói yêu.\n\n[ĐK 2:]\nCho dù ta đã mất rất rất lâu để yêu nhau\nNhưng chẳng còn gì ý nghĩa nếu như chúng ta không hiểu nhau\nMình Yêu Nhau Đi lyrics on ChiaSeNhac.com\nVà muốn quan tâm nhau, phải không anh?\nVà em xin hứa sẽ mãi mãi yêu một mình anh.\n\nCho dù ngày sau dẫu có nắng hay mưa trên đầu\nEm chẳng ngại điều gì đâu chỉ cần chúng ta che chở nhau\nCó anh bên em là em yên lòng\nKể từ hôm nay em sẽ chính thức được gọi anh: Anh yêu."
	// DateReleased : ""
	// DateCreated : "2014-02-09 21:10:00"
	// Type : false
	// Checktime : "2013-11-21 00:00:00"
}
