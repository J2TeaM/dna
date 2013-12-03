package zi

import (
	. "dna"
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
