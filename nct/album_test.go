package nct

import (
	. "dna"
	"time"
)

func ExampleGetAlbum() {
	album, err := GetAlbum(12255234)
	PanicError(err)
	album.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if album.Plays < 220787 {
		panic("Plays has to be greater than 220787, GET:" + album.Plays.ToString().String())
	}
	if album.Likes < 130 {
		panic("Likes has to be greater than 130, GET:" + album.Likes.ToString().String())
	}
	album.Plays = 220787
	album.Likes = 130

	LogStruct(album)
	// Output:
	// Id : 12255234
	// Key : "nsnkteavOHbX"
	// Title : "Biết Trước Sẽ Không Mất Nhau (Single)"
	// Artists : dna.StringArray{"Vĩnh Thuyên Kim", "Hồ Quang Hiếu"}
	// Topics : dna.StringArray{"Nhạc Trẻ"}
	// Likes : 130
	// Plays : 220787
	// Songids : dna.IntArray{2853789, 2808083, 2854574}
	// Nsongs : 3
	// Description : "Sau sự kết hợp thành công cùng các nữ ca sĩ xinh đẹp trong thời gian gần đây như: Bảo Thy, Nhật Kim Anh, Lương Khánh Vy... Chàng hotboy Hồ Quang Hiếu có sự kết hợp mới cùng Vĩnh Thuyên Kim trong một sáng tác của Lê Chí Trung \"Biết Trước Sẽ Không Mất Nhau\"."
	// Coverart : "http://avatar.nct.nixcdn.com/playlist/2013/11/28/4/e/9/b/1385644142690_500.jpg"
	// LinkKey : "6d0815b24a40117506fe5f12f3234846"
	// LinkShare : "http://www.nhaccuatui.com/playlist/biet-truoc-se-khong-mat-nhau-single-vinh-thuyen-kim-ho-quang-hieu.nsnkteavOHbX.html"
	// Type : "PLAYLIST_SONG"
	// Official : false
	// HasFeature : true
	// DateCreated : "2013-11-28 20:09:02"
	// Checktime : "2013-11-21 00:00:00"
}
