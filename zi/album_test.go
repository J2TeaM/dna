package zi

import (
	. "dna"
	"time"
)

func ExampleGetAlbum() {
	album, err := GetAlbum(1381697198)
	PanicError(err)
	album.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if album.Plays < 340 {
		panic("Plays has to be greater than 340")
	}
	album.Plays = 350
	album.EncodedKey = "" // Set to empty to prevent getting error because EncodedKey is changing every time
	LogStruct(album)
	// Output:
	// Id : 1381697198
	// Key : "ZWZB06AE"
	// EncodedKey : ""
	// Title : "Sibelius - The Symphonies CD 2"
	// Artists : dna.StringArray{"Gennady Rozhdestvensky", "Moscow Radio Symphony Orchestra"}
	// Coverart : "http://image.mp3.zdn.vn/thumb/240_240/covers/f/3/f3ccdd27d2000e3f9255a7e3e2c48800_1384684754.jpg"
	// Topics : dna.StringArray{"Hòa Tấu", "Classical"}
	// Plays : 350
	// Songids : dna.IntArray{1382630665, 1382630666, 1382630667, 1382630668, 1382630669, 1382630670, 1382630671, 1382630672}
	// YearReleased : "2012"
	// Nsongs : 8
	// Description : "Nhà soạn nhạc Jean Sibelius đã viết nên những bản nhạc cổ điễn lãng mạn tuyệt vời với album gồm 2 đĩa CD Sibelius - The Symphonies. Hãy lắng nghe và cảm nhận bạn nhé!"
	// DateCreated : "2013-11-17 17:39:14"
	// Checktime : "2013-11-21 00:00:00"
	// IsAlbum : 1
	// IsHit : 0
	// IsOfficial : 1
	// Likes : 0
	// Comments : 0
	// StatusId : 1
	// ArtistIds : dna.IntArray{32664, 32665}
}
