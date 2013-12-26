package nct

import (
	. "dna"
	"time"
)

func ExampleGetSong() {
	song, err := GetSong("WviZ6aais76C", 0)
	PanicError(err)
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if song.Plays < 13 {
		panic("Plays has to be greater than 13")
	}
	song.Plays = 13

	if len(*RelevantSongs) == 0 {
		panic(Sprintf("RelevantSongs has to be greater than 0, but got: %v", len(*RelevantSongs)).String())
	}

	// reset to avoid conflicted testings afterwards
	ResetRelevantPortions()
	LogStruct(song)
	// Output:
	// Id : 2870710
	// Key : "WviZ6aais76C"
	// Title : "Hương Sen Đồng Tháp (Tân Cổ)"
	// Artists : dna.StringArray{"Trọng Hữu", "Kiều Hoa"}
	// Topics : dna.StringArray{"Thể Loại Khác"}
	// Plays : 13
	// Type : "song"
	// Bitrate : 128
	// Official : 1
	// LinkKey : "f27678b52583250ee3e67b13f9e795f5"
	// Lyric : ""
	// Checktime : "2013-11-21 00:00:00"
}
