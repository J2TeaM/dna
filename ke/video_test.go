package ke

import (
	"dna"
	"time"
)

func ExampleGetVideo() {
	_, err := GetVideo(203952)
	if err == nil {
		panic("Video has to have an error")
	} else {
		if err.Error() != "Keeng - Video ID: 203952 not found" {
			panic("Wrong error message!")
		}
	}

	video, err := GetVideo(215236)
	dna.PanicError(err)
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if video.Plays < 6688 {
		panic("Wrong play")
	}
	video.Plays = 6688
	dna.LogStruct(video)
	// Output:
	// Id : 215236
	// Key : "Y45E8EPS"
	// Title : "Ông Bà Già Noel"
	// Artists : dna.StringArray{"Kevin Sôcôla", "Thân Nhật Huy"}
	// Plays : 6688
	// ListenType : 0
	// Link : "http://media2.keeng.vn/medias/video/2013/12/24/bc2137b2029cccf185b657aa88fc1164f5d1d605_mp4_640_360.mp4"
	// IsDownload : 0
	// DownloadUrl : "http://media2.keeng.vn/medias/video/2013/12/24/bc2137b2029cccf185b657aa88fc1164f5d1d605.mp4"
	// RingbacktoneCode : ""
	// RingbacktonePrice : 0
	// Price : 0
	// Copyright : 0
	// CrbtId : 0
	// Thumbnail : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias_6/video/images/2013/12/24/e2d4ef038c1922a23d4f82b7d4c06972638e8753_147_83.jpg"
	// DateCreated : "2013-12-24 00:00:00"
	// Checktime : "2013-11-21 00:00:00"
}
