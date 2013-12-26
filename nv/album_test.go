package nv

import (
	. "dna"
	"testing"
	"time"
)

func TestGetAlbum(t *testing.T) {
	_, err := GetAlbum(15)
	if err == nil {
		t.Error("Album 15 has to have an error")
	}
	if err.Error() != "Nhacvui - Album 15: No song found" {
		t.Errorf("Error message has to be: %v", err.Error())
	}
}
func ExampleGetAlbum() {
	album, err := GetAlbum(47895)
	PanicError(err)
	if album.Plays < 63867 {
		panic("Plays has to be greater than 63867")
	}
	album.Plays = 63867
	album.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	LogStruct(album)

	// Output:
	// Id : 47895
	// Title : "Không Cần Thêm Một Ai Nữa (Single)"
	// Artists : dna.StringArray{"Mr.Siro", "BigDaddy"}
	// Topics : dna.StringArray{"Nhạc Trẻ"}
	// Nsongs : 3
	// Plays : 63867
	// Coverart : "http://nv-ad-hcm.24hstatic.com/imageupload/upload2013/2013-4/album/2013-12-04/1386141910_Khong-Can-Them-Mot-Ai-Nua-Album-Mrsiro.jpg"
	// Description : "Đây là ca khúc được Mr. Siro sáng tác từ rất lâu nhưng đến năm 2013 nay anh mới tìm được rapper hợp ý để cùng hợp tác với nhau, đó là BigDaddy, rapper đang lên tại Hà Nội."
	// DateCreated : "2013-12-04 14:25:10"
	// Songids : dna.IntArray{473348, 458845, 275555}
	// Checktime : "2013-11-21 00:00:00"
}
