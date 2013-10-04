package zi

import (
	. "dna"
)

func ExampleNewSong() {
	var x *Song = NewSong("ZW67FWWF")
	Logv(*x)
	// Output: zi.Song{Id:1382543919, Key:"ZW67FWWF", Title:"", Artists:dna.StringArray{}, Authors:dna.StringArray{}, Plays:0, Topics:dna.StringArray(nil), Link:"", Path:"", Lyric:"", DateCreated:""}
}

func ExampleNewSongWithId() {
	var x *Song = NewSongWithId(1382543919)
	Logv(*x)
	// Output: zi.Song{Id:1382543919, Key:"ZW67FWWF", Title:"", Artists:dna.StringArray{}, Authors:dna.StringArray{}, Plays:0, Topics:dna.StringArray(nil), Link:"", Path:"", Lyric:"", DateCreated:""}
}

func ExampleSong_GetEncodedKey() {
	var x String = "ZW67FWWF"
	// Get the same result with different bitrates
	Logv(DecodeEncodedKey(NewSong(x).GetEncodedKey(Bitrate320)))
	Logv(DecodeEncodedKey(NewSong(x).GetEncodedKey(Bitrate128)))
	Logv(DecodeEncodedKey(NewSong(x).GetEncodedKey(Bitrate256)))
	Logv(DecodeEncodedKey(NewSong(x).GetEncodedKey(Lossless)))
	// Output:
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
	// "ZW67FWWF"
}

func ExampleSong_GetDirectLink() {
	var x String = "ZW67FWWF"
	// Get the same result with different bitrates
	x = NewSong(x).GetDirectLink(Bitrate128)
	// x has form of "http://mp3.zing.vn/download/song/joke-link/ZmJGTLnsSanGsLEtkvJyDGZm"
}
