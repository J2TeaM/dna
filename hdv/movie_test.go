package hdv

import (
	"dna"
	"time"
)

func ExampleGetMovie() {
	movie, err := GetMovie(5561)
	dna.PanicError(err)
	movie.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	dna.LogStruct(movie)
	// Output:
	// Id : 5561
	// Title : "The Scandal - Sự Thật Nghiệt Ngã - Tập 36/36"
	// AnotherTitle : "The Scandal"
	// ForeignTitle : "The Scandal"
	// VnTitle : "Sự Thật Nghiệt Ngã"
	// Topics : dna.StringArray{"HDViệt Đề Cử", "Hàn quốc", "Gia Đình", "Lãng Mạn"}
	// Actors : dna.StringArray{"Kim Jae Won", "Jung Yoon Suk", "Kim Hwi Soo"}
	// Directors : dna.StringArray{"Kim Jin Man", "Park Jae Bum"}
	// Countries : dna.StringArray{"Hàn Quốc"}
	// Description : "Bộ phim kể về câu chuyện của một cảnh sát hình sự (Kim Jae Won - Ha Eun Joong) tình cờ biết được sự thật, người cha của anh hiện tại thật ra là kẻ bắt cóc và chính ông đã bắt cóc anh lúc anh còn nhỏ. Cuộc sống của anh đã có nhiều thay đổi khi gặp gỡ người phụ nữ Woo Ah Mi (Jo Yoon Hee), một người mẹ đơn thân 26 tuổi. Kể từ giây phút ấy, anh bắt đầu hành trình hàn gắn nỗi đau trong trái tim mình cùng Woo Ah Mi."
	// YearReleased : 2013
	// IMDBRating : dna.IntArray{0, 0}
	// Similars : dna.IntArray{}
	// Thumbnail : "http://t.hdviet.com/thumbs/214x321/ca8507356775119a9f183e2abe95fe02.jpg"
	// MaxResolution : 720
	// IsSeries : true
	// SeasonId : 0
	// Seasons : dna.IntArray{}
	// Epid : 0
	// CurrentEps : 36
	// MaxEp : 36
	// Checktime : "2013-11-21 00:00:00"
}
