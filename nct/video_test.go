package nct

import (
	. "dna"
	"time"
)

func ExampleGetVideo() {
	video, err := GetVideo("N5QeESGm7ICBt")
	PanicError(err)
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if video.Plays < 104632 {
		panic("Plays has to be greater than 104632")
	}
	video.Plays = 104632

	if len(*RelevantSongs) != 10 {
		panic("RelevantSongs has to be 10")
	}

	// reset to avoid conflicted testings afterwards
	ResetRelevantPortions()
	LogStruct(video)
	// Output:
	// Id : 2876055
	// Key : "N5QeESGm7ICBt"
	// Title : "Gửi Cho Anh (Phần 2)"
	// Artists : dna.StringArray{"Khởi My"}
	// Topics : dna.StringArray{"Âm Nhạc", "Việt Nam", "Nhạc Trẻ"}
	// Plays : 104632
	// Duration : 0
	// Thumbnail : "http://m.img.nct.nixcdn.com/mv/2013/12/10/e/0/5/3/1386640122904_640.jpg"
	// Type : "mv"
	// LinkKey : "f9652760275d5777e5516f812b840097"
	// Lyric : ""
	// DateCreated : "2013-12-10 08:48:42"
	// Checktime : "2013-11-21 00:00:00"
}
