package zi

import (
	. "dna"
	"time"
)

func ExampleGetArtist() {
	artist, err := GetArtist(1134)
	PanicError(err)
	artist.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	LogStruct(artist)
	// Output:
	// Id : 1134
	// Name : "Britt Nicole"
	// Alias : ""
	// Birthname : "Brittany Nicole Waddell"
	// Birthday : "02/08/1984"
	// Sex : 2
	// Link : "/nghe-si/Britt-Nicole"
	// Topics : dna.StringArray{"Âu Mỹ", "Pop", "Rock"}
	// Avatar : "avatars/b/d/bdc6f49840bd51e69516a86f204a8be3_1285928914.jpg"
	// Coverart : "cover_artist/1/5/156005c5baf40ff51a327f1c34f2975b_1379182334.jpg"
	// Coverart2 : "cover2_artist/7/9/799bad5a3b514f096e69bbc4a7896cd9_1379182339.jpg"
	// ZmeAcc : ""
	// Role : "1"
	// Website : "http://www.brittnicole.com/"
	// Biography : "Brittany Nicole Waddell (sinh 02 tháng 8 năm 1985) là một nghệ sĩ nhạc Christian những người thực hiện theo tên Britt Nicole. Cô ký hợp đồng với EMI CMG / Sparrow. Cô hát nhạc Contemporary Christian."
	// Publisher : "EMI"
	// Country : "United States"
	// IsOfficial : 1
	// YearActive : "2004"
	// StatusId : 1
	// DateCreated : "0001-01-01 00:00:00"
	// Checktime : "2013-11-21 00:00:00"
}
