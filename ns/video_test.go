package ns

import (
	. "dna"
	"testing"
)

//Testing video with fail link
func TestGetVideo(t *testing.T) {
	_, err := GetVideo(19306)
	if err == nil {
		t.Error("Video has to have an error")
	} else {
		if err.Error() != "Video 19306 : Link not found" {
			t.Error("Wrong error message")
		}
	}
	video, err := GetVideo(16280)
	if err != nil {
		t.Error("Video has to have no error")
	} else {
		if video.Title != "Giữ Lấy Nhau" {
			t.Error("Video 16280 has wrong title")
		}
		if video.Artists.Length() != 2 || video.Artists[0] != "Khánh Phương" || video.Artists[1] != "Đan Thùy" {
			t.Error("Video 16280 has wrong artists")
		}
		if video.Topics.Length() != 1 || video.Topics[0] != "Nhạc Trẻ" {
			t.Error("Video 16280 has wrong topics ")
		}
		if video.Plays < 970 {
			t.Error("Video 16280 has wrong plays ")
		}
		if video.Duration != 444 {
			t.Error("Video 16280 has wrong duration ")
		}
		if video.Official != 1 || video.Proceducerid != 71 {
			t.Error("Video 16280 has wrong official or proceducerid ")
		}
		if video.Link != "http://st02.freesocialmusic.com/mp4/2013/05/30/1430055571/13698818900_3033.mp4" {
			t.Error("Video 16280 has wrong link ")
		}
		if Int(video.DateCreated.Unix()).ToTimeFormat() != "2013-5-30 9:44:50" {
			t.Error("Video 16280 has wrong date ")
		}

	}
}

func ExampleGetVideo() {
	video, err := GetVideo(18272)
	if err != nil {
		Log(err.Error())
	} else {
		Log("Id:", video.Id)
		Log("Title:", video.Title)
		Log("Artists:", video.Artists)
		Log("Topics:", video.Topics)
		Log("Duration:", video.Duration)
		Log("Official:", video.Official)
		Log("Proceducerid:", video.Proceducerid)
		Log("Link:", video.Link)
		Log("Thumbnail:", video.Thumbnail)
		Log("DateCreated:", Int(video.DateCreated.Unix()).ToTimeFormat())
		if video.Plays > 0 {
			Log("Plays > 0")
		}
		Log("Sublink:", video.Sublink)

	}
	// Output:
	// Id: 18272
	// Title: Còn Có Anh
	// Artists: dna.StringArray{"Mạnh Quân", "Khang Việt"}
	// Topics: dna.StringArray{"Nhạc Trẻ"}
	// Duration: 339
	// Official: 1
	// Proceducerid: 43
	// Link: http://st02.freesocialmusic.com/mp4/2013/11/06/1178050012/138370336610_2499.mp4
	// Thumbnail: http://st.nhacso.net/images/video/2013/11/06/1178050012/138370345410_6711_190x110.jpg
	// DateCreated: 2013-11-6 9:2:46
	// Plays > 0
	// Sublink:
}
