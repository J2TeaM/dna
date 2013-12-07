package zi

import (
	. "dna"
	"time"
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

func ExampleGetSong() {
	song, err := GetSong(1382642591)
	PanicError(err)
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if song.Plays < 15 {
		panic("Plays has to be greater than 15")
	}
	song.Plays = 20

	LogStruct(song)

	// Output:
	// Id : 1382642591
	// Key : "ZW697O9F"
	// Title : "Bridges"
	// Artists : dna.StringArray{"Rebecca Ferguson", "John Legend"}
	// Authors : dna.StringArray{}
	// Plays : 20
	// Topics : dna.StringArray{"Âu Mỹ", "Pop"}
	// Link : "http://m.mp3.zing.vn/xml/song-load/MjAxMyUyRjEyJTJGMDElMkY4JTJGOCUyRjg4OTc0Mzg5MGJlNTlmMWVkNDgwYTZjM2Q2NWViZDNiLm1wMyU3QzI="
	// Path : "2013/12/01/8/8/889743890be59f1ed480a6c3d65ebd3b.mp3"
	// Lyric : ""
	// DateCreated : "2013-12-01 00:00:00"
	// Checktime : "2013-11-21 00:00:00"
	// ArtistIds : dna.IntArray{12556, 2072}
	// VideoId : 0
	// AlbumId : 1381698540
	// IsHit : 0
	// IsOfficial : 1
	// DownloadStatus : 1
	// Copyright : ""
	// BitrateFlags : 3
	// Likes : 0
	// Comments : 0
	// Thumbnail : "avatars/9/5/95d241f273be66577c7fe267dbb31d75_1351440915.png"
}
