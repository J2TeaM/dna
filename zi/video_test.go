package zi

import (
	. "dna"
	"time"
)

func ExampleVideo_GetEncodedKey() {
	var x String = "ZW67FWWF"
	video := NewVideo()
	video.Id = GetId(x)
	// Get the same result with different resolution
	// "ZW67FWWF" =>"ZmcnTZnslNHnsZETdgmtbGkn" => "ZW67FWWF"
	Logv(DecodeEncodedKey(video.GetEncodedKey(Resolution240p)))
	Logv(DecodeEncodedKey(video.GetEncodedKey(Resolution360p)))
	Logv(DecodeEncodedKey(video.GetEncodedKey(Resolution480p)))
	Logv(DecodeEncodedKey(video.GetEncodedKey(Resolution720p)))
	Logv(DecodeEncodedKey(video.GetEncodedKey(Resolution1080p)))
	// Output: "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
}

func ExampleVideo_GetDirectLink() {
	var x String = "ZW67FWWF"
	video := NewVideo()
	video.Id = GetId(x)
	// Get the same result with different bitrates
	x = video.GetDirectLink(Resolution720p)
	// x has form of "http://mp3.zing.vn/html5/video/ZmJGyLHaAaHmaLuTsDnyvGkm"
}

func ExampleGetVideo() {
	video, err := GetVideo(1382633268)
	PanicError(err)
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if video.Plays < 263 {
		panic("Plays has to be greater than 263")
	}
	video.Plays = 350
	if !video.Link.Match(`/[a-zA-Z]{24}`) {
		panic("Video link is not valid")
	}
	// video.Link has the format http://api.mp3.zing.vn/api/mobile/source/video/LnxnyLHsANinGgJyNvmyvmLH
	// EncodedKey `LnxnyLHsANinGgJyNvmyvmLH` changing every its called
	video.Link = "http://api.mp3.zing.vn/api/mobile/source/video/LnxnyLHsANinGgJyNvmyvmLH"
	LogStruct(video)
	// Output:
	// Id : 1382633268
	// Title : "I Love You Because"
	// Artists : dna.StringArray{"Elvis Presley", "Lisa Marie Presley"}
	// Topics : dna.StringArray{"Âu Mỹ", "Pop"}
	// Plays : 350
	// Thumbnail : "http://image.mp3.zdn.vn/thumb_video/2013/11/20/3/3/33f7a203a057517fac802c8cd696595b_3.jpg"
	// Link : "http://api.mp3.zing.vn/api/mobile/source/video/LnxnyLHsANinGgJyNvmyvmLH"
	// Lyric : ""
	// DateCreated : "2013-11-20 00:00:00"
	// Checktime : "2013-11-21 00:00:00"
	// ArtistIds : dna.IntArray{5281, 2556}
	// Duration : 175
	// StatusId : 1
	// ResolutionFlags : 15
	// Likes : 0
	// Comments : 0
}
