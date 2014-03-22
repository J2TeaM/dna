package zi

import (
	. "dna"
	"time"
)

func ExampleTV_GetEncodedKey() {
	var x String = "ZW67FWWF"
	tv := NewTV()
	tv.Key = x
	tv.Id = GetId(x)
	y := DecodeEncodedKey(tv.GetEncodedKey())
	Logv(y)
	// Output: "ZW67FWWF"
}

func ExampleTV_GetDirectLink() {
	var x String = "IWZA0O0O"
	tv := NewTV()
	tv.Key = x
	tv.Id = GetId(x)
	// Get the same result with different bitrates
	x = tv.GetDirectLink()
	// x has a form of "http://tv.zing.vn/html5/video/LmcntlQhEitDHLn"
}

func ExampleGetTV() {
	tv, err := GetTV(GetId("IWZAI860"))
	PanicError(err)
	if tv.Plays < 10412 {
		// Log(tv)
		panic("Plays has to be bigger than 10412")
	}
	if tv.Likes < 19 {
		panic("Likes has to be greater than or equal to 19")
	}
	// if tv.Comments < 9 {
	// 	// panic("Comments has to be greater than or equal to 9")
	// }
	if tv.Rating < 0 {
		panic("Rating has to be greater than or equal to 0")
	}
	if tv.FileUrl == "" {
		panic("File URL has to be valid")
	}
	tv.Plays = 10412
	tv.Likes = 19
	tv.Comments = 0
	tv.Rating = 9.1
	tv.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	tv.FileUrl = "stream6.tv.zdn.vn/streaming/a19a0b455fd78fe97a3150e24d9cf721/5323cc2f/2013/1201/f4/418e13b56de6eb817d6df794ddcac630.mp4?format=f360&device=ios"
	LogStruct(tv)
	// Output:
	// Id : 307894368
	// Key : "IWZAI860"
	// Title : "Ngọc Anh Vs Ngọc Thịnh"
	// Fullname : "Thử Thách Cùng Bước Nhảy 2013 - Tập 23 - Ngọc Anh Vs Ngọc Thịnh"
	// Episode : 23
	// DateReleased : "2013-12-01 00:00:00"
	// Duration : 700
	// Thumbnail : "2013/1201/f4/4a255f168c3344eb3df21f745328cb24_1385877303.jpg"
	// FileUrl : "stream6.tv.zdn.vn/streaming/a19a0b455fd78fe97a3150e24d9cf721/5323cc2f/2013/1201/f4/418e13b56de6eb817d6df794ddcac630.mp4?format=f360&device=ios"
	// ResolutionFlags : 2
	// ProgramId : 1778
	// ProgramName : "Thử Thách Cùng Bước Nhảy 2013"
	// ProgramThumbnail : "channel/7/a/7a1acf638cf34f4fcf5104cb30202fcb_1376621966.jpg"
	// ProgramGenreIds : dna.IntArray{78}
	// ProgramGenres : dna.StringArray{"TV Show"}
	// Plays : 10412
	// Comments : 0
	// Likes : 19
	// Rating : 9.1
	// Subtitle : ""
	// Tracking : ""
	// Signature : "7e1c826b4569ad7488a2e8867378f434"
	// Checktime : "2013-11-21 00:00:00"
}
