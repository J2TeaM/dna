package zi

import (
	. "dna"
)

func ExampleVideo_GetEncodedKey() {
	var x String = "ZW67FWWF"
	// Get the same result with different resolution
	// "ZW67FWWF" =>"ZmcnTZnslNHnsZETdgmtbGkn" => "ZW67FWWF"
	Logv(DecodeEncodedKey(NewVideo(x).GetEncodedKey(Resolution240p)))
	Logv(DecodeEncodedKey(NewVideo(x).GetEncodedKey(Resolution360p)))
	Logv(DecodeEncodedKey(NewVideo(x).GetEncodedKey(Resolution480p)))
	Logv(DecodeEncodedKey(NewVideo(x).GetEncodedKey(Resolution720p)))
	Logv(DecodeEncodedKey(NewVideo(x).GetEncodedKey(Resolution1080p)))
	// Output: "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
}

func ExampleVideo_GetDirectLink() {
	var x String = "ZW67FWWF"
	// Get the same result with different bitrates
	x = NewVideo(x).GetDirectLink(Resolution720p)
	// x has form of "http://mp3.zing.vn/html5/video/ZmJGyLHaAaHmaLuTsDnyvGkm"
}
