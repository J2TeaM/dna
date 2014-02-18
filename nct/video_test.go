package nct

import (
	. "dna"
	"time"
)

func ExampleGetVideo() {
	video, err := GetVideo(2876055)
	PanicError(err)
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if video.Plays < 104632 || video.Likes < 1487 {
		panic("Plays has to be greater than 104632 or likes .... ")
	}
	video.Plays = 104632
	video.Likes = 1487
	if video.StreamUrl.EndsWith("mp4") == false {
		panic("File has to be mp4 format")
	} else {
		video.StreamUrl = "http://nplus.nixcdn.com/93cfd7ee3743f301029bd59e789a10fd/530187c7/PreNCT7/GuiChoAnhPhan2-KhoiMy-2876055.mp4"
	}
	LogStruct(video)
	// Output:
	// Id : 2876055
	// Key : "N5QeESGm7ICBt"
	// Title : "Gửi Cho Anh (Phần 2)"
	// Artists : dna.StringArray{"Khởi My"}
	// Artistid : 12987
	// Topics : dna.StringArray{"Việt Nam", "Nhạc Trẻ"}
	// Plays : 104632
	// Likes : 1487
	// Duration : 2768
	// Thumbnail : "http://avatar.nct.nixcdn.com/mv/2013/12/10/e/0/5/3/1386640122904_536.jpg"
	// Image : "http://avatar.nct.nixcdn.com/mv/2013/12/10/e/0/5/3/1386640122904.jpg"
	// Type : "VIDEO"
	// LinkKey : "f9652760275d5777e5516f812b840097"
	// LinkShare : "http://www.nhaccuatui.com/video/gui-cho-anh-phan-2-khoi-my.N5QeESGm7ICBt.html"
	// Lyric : ""
	// StreamUrl : "http://nplus.nixcdn.com/93cfd7ee3743f301029bd59e789a10fd/530187c7/PreNCT7/GuiChoAnhPhan2-KhoiMy-2876055.mp4"
	// Relatives : dna.StringArray{}
	// DateCreated : "2013-12-10 08:48:42"
	// Checktime : "2013-11-21 00:00:00"
}
