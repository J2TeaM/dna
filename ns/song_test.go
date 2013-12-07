package ns

import (
	. "dna"
	"testing"
	"time"
)

// Testing song with fail link
func TestGetSong(t *testing.T) {
	song1, err := GetSong(1412937)
	if err == nil {
		Log(song1)
		t.Error("The song has to have an error")
	}
	song, err := GetSong(1312951)
	if err != nil {
		t.Error("An error occurs: %v", err.Error())
	} else {
		if song.Authors.Length() > 0 && song.Authorid == 0 {
			t.Error("Song's authors founded")
		}
		if song.Duration != 189 {
			t.Error("Duration has to be 189s")
		}
		if song.Bitrate != 320 {
			t.Error("Bitrate has to be 320kbps")
		}

		if song.Topics[0] != "Nhạc Âu Mỹ" {
			t.Error("Topics has to be Nhạc Âu Mỹ")
		}

		if song.Islyric != 0 {
			t.Error("Islyric has to be 0")
		}
	}

}

// TestGetSong2 tests song which has result from XML file but has 404 error code from main page
func TestGetSong2(t *testing.T) {
	song, err := GetSong(1310212)
	if err != nil {
		t.Error("The song has to have no error")
	} else {
		if song.Title != "내가 미친년이야 (I'm Crazy Girl)" {
			t.Errorf("title error. Got: %v", song.Title)
		}
		if song.Link != "http://st02.freesocialmusic.com/mp3/2013/10/18/1348055438/13820884993_5016.mp3" {
			t.Error("link error")
		}
		if song.Artists[0] != "Kim Bo Hyung (Spica)" || song.Artistid != 99421 {
			t.Error("artists error")
		}
		if song.Authors.Length() > 0 || song.Authorid != 0 {
			t.Error("authors error")
		}

		if song.Topics.Length() > 0 {
			t.Error("topics error")
		}

		if song.Bitrate != 0 || song.Islyric != 0 || song.Lyric != "" || song.SameArtist > 0 || song.Duration != 273 || song.Official > 0 {
			t.Error("song fields error")
		}
	}

}

func ExampleGetSong() {
	song, err := GetSong(1312937)
	PanicError(err)
	if song.Plays < 0 {
		panic("Plays has to be greater than 0")
	}
	if song.DateUpdated.IsZero() {
		panic("DateUpdated has not to be zero")

	}
	song.Plays = 100
	song.DateUpdated = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	LogStruct(song)

	// Output:
	// Id : 1312937
	// Title : "Nếu Như Ta Cách Xa"
	// Artists : dna.StringArray{"Bảo Thy", "Hồ Quang Hiếu"}
	// Artistid : 7125
	// Authors : dna.StringArray{"Nhạc Hoa"}
	// Authorid : 13299
	// Plays : 100
	// Duration : 245
	// Link : "http://st02.freesocialmusic.com/mp3/2013/11/07/1430055571/138380871215_657.mp3"
	// Topics : dna.StringArray{"Nhạc Trẻ"}
	// Category : dna.StringArray{}
	// Bitrate : 320
	// Official : 1
	// Islyric : 1
	// DateCreated : "2013-11-07 14:18:32"
	// DateUpdated : "2013-11-21 00:00:00"
	// Lyric : "HQH :\nAnh cứ ngỡ mặt xa cách lòng, niềm cô đơn mỗi đêm sẽ lướt qua trong nỗi nhớ\nBao lâu ta cách xa, mà cứ như vừa hôm qua\nTừ thâm tâm anh trách mình, đã để em ra đi\nKhông thể nào phai phôi, đã quá yêu rồi, người ơi đừng vội buông lơi\nHạnh phúc trong cuộc đời, chỉ một lần mà thôi\nEm nói anh nghe đi, cớ sao bây giờ mỗi người một nơi như thế?\nHãy cho nhau thời gian để quay lại.. Được yêu thêm lần nữa, người yêu hỡi...\nBT :\nEm vẫn luôn tự hỏi chính mình, rằng chia tay với anh là một quyết định sai hay đúng?\nKhông gian như vỡ tan, hạnh phúc đang dần phai nhoà\nNhẹ quay lưng nhìn tháng ngày, mình đã tay trong tay\nQuá khứ như cơn mơ, nhói đay vô bờ, giờ em một mình bơ vơ\nKhoảnh khắc em quay đi, sao anh lạnh lùng đến thế?\nAnh nói em nghe đi, dẫu chỉ thầm thì..rằng anh cần em mọi khi\nVì tình yêu thì không đúng sai gì\nBuồn làm chi để đêm đêm kí ức hoen bờ mi\nAnh/em cứ ngỡ mặt xa cách lòng..\nNiềm cô đơn mỗi đêm sẽ lướt qua trong nỗi nhớ...\nNiềm cô đơn mỗi đêm sẽ tan biến...trong kỉ niệm..."
	// SameArtist : 0
	// Checktime : "2013-11-21 00:00:00"
}
