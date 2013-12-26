package cc

import (
	. "dna"
	"testing"
	"time"
)

func TestGetAlbum(t *testing.T) {
	_, err := GetAlbum(18449)
	if err == nil {
		t.Error("Album 18449 has to have an error")
	}
	if err.Error() != "Chacha - Album 18449: No song found" {
		t.Errorf("Error message has to be: %v", err.Error())
	}
}

func ExampleGetAlbum() {
	album, err := GetAlbum(18356)
	PanicError(err)
	if album.Plays < 28 {
		panic("Plays has to be greater than 28")
	}
	album.Plays = 28
	album.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	LogStruct(album)

	// Output:
	// Id : 18356
	// Title : "BEYONCE"
	// Artists : dna.StringArray{"Beyoncé Knowles"}
	// Topics : dna.StringArray{"Nhạc Âu Mỹ"}
	// Nsongs : 14
	// Plays : 28
	// Coverart : "http://s2.chacha.vn/albums//s2/2/18356/18356.jpg"
	// Description : "Đây là album thứ 5 của Beyoncé được cô bất ngờ tung ra vào ngày 13/12/2013. Album bao gồm 14 ca khúc. Đặc biệt hơn với 14 ca khúc này, cựu thành viên Destiny's Child đều thực hiện MV để quảng bá.   Beyonce từng được đánh giá là một trong những nghệ sỹ siêng năng trong việc thu âm và thực hiện MV. Nhưng với album lần này, cô đã làm các fan hoàn toàn bất ngờ khi có quyết định trên."
	// Songids : dna.IntArray{894228, 894252, 894250, 894248, 894246, 894244, 894242, 894240, 894238, 894236, 894234, 894232, 894230, 894254}
	// YearReleased : 2013
	// Checktime : "2013-11-21 00:00:00"
}
