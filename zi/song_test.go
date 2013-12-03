package zi

import (
	. "dna"
)

func ExampleSong_GetEncodedKey() {
	var key String = "ZW67FWWF"
	song := NewSong()
	song.Key = key
	song.Id = GetId(key)
	// Get the same result with different bitrates
	Logv(DecodeEncodedKey(song.GetEncodedKey(Bitrate320)))
	Logv(DecodeEncodedKey(song.GetEncodedKey(Bitrate128)))
	Logv(DecodeEncodedKey(song.GetEncodedKey(Bitrate256)))
	Logv(DecodeEncodedKey(song.GetEncodedKey(Lossless)))
	// Output:
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
}

func ExampleSong_GetDirectLink() {
	var key String = "ZW67FWWF"
	song := NewSong()
	song.Key = key
	song.Id = GetId(key)
	// Get the same result with different bitrates
	var ret = song.GetDirectLink(Bitrate128)
	Log(ret)
	// ret has form of "http://mp3.zing.vn/download/song/joke-link/ZmJGTLnsSanGsLEtkvJyDGZm"
}
