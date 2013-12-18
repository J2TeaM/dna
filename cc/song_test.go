package cc

import (
	. "dna"
	"time"
)

func ExampleGetSong() {
	song, err := GetSong(872358)
	PanicError(err)
	if song.Plays < 30 {
		panic("Plays has to be greater than 30")
	}
	song.Plays = 30
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	LogStruct(song)

	// Output:
	// Id : 872358
	// Title : "Con Nhà Nghèo"
	// Artists : dna.StringArray{"Lương Bích Hữu"}
	// Artistid : 139063
	// Topics : dna.StringArray{"Nhạc trẻ"}
	// Plays : 30
	// Duration : 0
	// Bitrate : 128
	// Coverart : "http://s2.chacha.vn/artists//s1/16/139063/139063.jpg"
	// Lyrics : ""
	// Link : "http://audio.chacha.vn/songs/output/106/872358/2/s/con-nha-ngheo - Luong Bich Huu.mp3?s=1"
	// Checktime : "2013-11-21 00:00:00"
}
