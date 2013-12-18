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
	tv, err := GetTV(GetId("IWZAI70B"))
	PanicError(err)
	if tv.Plays < 30242 {
		panic("Plays has to be greater than 30242")
	}
	if tv.Likes < 50 {
		panic("Likes has to be greater than or equal to 50")
	}
	if tv.Comments < 9 {
		panic("Comments has to be greater than or equal to 9")
	}
	if tv.Rating < 0 {
		panic("Rating has to be greater than or equal to 0")
	}
	if tv.FileUrl == "" {
		panic("File URL has to be valid")
	}
	tv.Plays = 30242
	tv.Likes = 50
	tv.Comments = 9
	tv.Rating = 8.927536231884059
	tv.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	tv.FileUrl = "stream6.tv.zdn.vn/streaming/ed283ead88766c5a8ed4a82ee4abf2f4/52a2af3c/2013/1125/91/2eddbfaa80233df649d9c6f2dcf2c214.mp4?format=f360&device=ios"
	LogStruct(tv)
	// Output:
	// Id : 307894027
	// Key : "IWZAI70B"
	// Title : "Full Show"
	// Fullname : "MTV Europe Music Awards (EMA) - Full Show"
	// Episode : 0
	// DateReleased : "2013-11-25 00:00:00"
	// Duration : 6296
	// Thumbnail : "2013/1125/91/0e8396a11de4bca6bf2c7329400ee2db_1385438291.jpg"
	// FileUrl : "stream6.tv.zdn.vn/streaming/ed283ead88766c5a8ed4a82ee4abf2f4/52a2af3c/2013/1125/91/2eddbfaa80233df649d9c6f2dcf2c214.mp4?format=f360&device=ios"
	// ResolutionFlags : 15
	// ProgramId : 2048
	// ProgramName : "MTV Europe Music Awards (EMA)"
	// ProgramThumbnail : "channel/e/0/e06d061a77a7bde916b8a91163029d41_1385368981.jpg"
	// ProgramGenreIds : dna.IntArray{78}
	// ProgramGenres : dna.StringArray{"TV Show"}
	// Plays : 30242
	// Comments : 9
	// Likes : 50
	// Rating : 8.927536231884059
	// Subtitle : ""
	// Tracking : ""
	// Signature : "ec657cea3927205faab0c933f8ebdef2"
	// Checktime : "2013-11-21 00:00:00"
}
