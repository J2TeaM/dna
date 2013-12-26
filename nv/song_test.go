package nv

import (
	. "dna"
	"testing"
	"time"
)

func TestGetSong(t *testing.T) {
	_, err := GetSong(136)
	if err == nil {
		t.Error("Song 136 has to have an error")
	}
	if err.Error() != "Nhacvui - Song 136: Mp3 link not found" {
		t.Errorf("Error message has to be: %v", err.Error())
	}
}
func ExampleGetSong() {
	song, err := GetSong(472092)
	PanicError(err)
	if song.Plays < 267090 {
		panic("Plays has to be greater than 267090")
	}
	song.Plays = 267090
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if !song.Link.Match("mp3") {
		panic("NO link found")
	}
	song.Link = "http://stream306.nhac.vui.vn/uploadmusic2/ce52c1a3f8af86140763dd7885690f0f/52a5f841....mp3"
	LogStruct(song)

	// Output:
	// Id : 472092
	// Title : "Giá Có Thể Ôm Ai Và Khóc"
	// Artists : dna.StringArray{"Phạm Hồng Phước"}
	// Authors : dna.StringArray{"Phạm Hồng Phước"}
	// Topics : dna.StringArray{"Nhạc Hot Nhất"}
	// Plays : 267090
	// Lyric : "Lời bài hát: Giá Có Thể Ôm Ai Và Khóc - Phạm Hồng Phước\nNgười đóng góp: hoangthieplbGiá như có ai ngồi ở đằng sau xe\nÔm thật lâu, tựa đầu vào lưng của tôi\nGiá như có ai cùng dạo quanh phố xá\nTôi sẽ vui, sẽ hạnh phúc biết bao\n\nHẹn hò với những cô đơn riêng tôi\nMột chiều lang thang\nHẹn hò với những miên man trong tôi\nTình Yêu vỡ nát\nCần tin nhắn mỗi ngày\nCần những hỏi han\nCho tôi yên bình thêm chút\n\nĐk: \nGiá như … Có người đợi tôi đâu đó giữa cuộc đời\nGiá như … Có người ôm tôi mỗi tối\nGiá như ... Có người ngồi nghe tôi kể bao vui buồn\nGiá như ... Có thể ôm lấy ai và khóc lên."
	// Link : "http://stream306.nhac.vui.vn/uploadmusic2/ce52c1a3f8af86140763dd7885690f0f/52a5f841....mp3"
	// Link320 : ""
	// Type : "song"
	// Checktime : "2013-11-21 00:00:00"
}
