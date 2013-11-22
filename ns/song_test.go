package ns

import (
	. "dna"
	"testing"
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
	if err != nil {
		Log(err.Error())
	} else {
		Log("Id:", song.Id)
		Log("Title:", song.Title)
		Log("Artists:", song.Artists)
		Log("Artistid:", song.Artistid)
		Log("Authors:", song.Authors)
		Log("Authorid:", song.Authorid)
		// Log("Plays:", song.Plays)
		Log("Duration:", song.Duration)
		Log("Link:", song.Link)
		Log("Topics:", song.Topics)
		Log("Category:", song.Category)
		Log("Bitrate:", song.Bitrate)
		Log("Official:", song.Official)
		Log("Islyric:", song.Islyric)
		Log("Lyric:", song.Lyric)
		Log("DateCreated:", Int(song.DateCreated.Unix()).ToTimeFormat())
		// Log("DateUpdated:", song.DateUpdated)
		Log("SameArtist:", song.SameArtist)

		if song.Plays > 0 {
			Log("Plays > 0")
		}

		if !song.DateUpdated.IsZero() {
			Log("Date updated is correct")
		}
	}
	// Output:
	// Id: 1312937
	// Title: Nếu Như Ta Cách Xa
	// Artists: dna.StringArray{"Bảo Thy", "Hồ Quang Hiếu"}
	// Artistid: 7125
	// Authors: dna.StringArray{"Nhạc Hoa"}
	// Authorid: 13299
	// Duration: 245
	// Link: http://st02.freesocialmusic.com/mp3/2013/11/07/1430055571/138380871215_657.mp3
	// Topics: dna.StringArray{"Nhạc Trẻ"}
	// Category: dna.StringArray{}
	// Bitrate: 320
	// Official: 1
	// Islyric: 1
	// Lyric: HQH :
	// Anh cứ ngỡ mặt xa cách lòng, niềm cô đơn mỗi đêm sẽ lướt qua trong nỗi nhớ
	// Bao lâu ta cách xa, mà cứ như vừa hôm qua
	// Từ thâm tâm anh trách mình, đã để em ra đi
	// Không thể nào phai phôi, đã quá yêu rồi, người ơi đừng vội buông lơi
	// Hạnh phúc trong cuộc đời, chỉ một lần mà thôi
	// Em nói anh nghe đi, cớ sao bây giờ mỗi người một nơi như thế?
	// Hãy cho nhau thời gian để quay lại.. Được yêu thêm lần nữa, người yêu hỡi...
	// BT :
	// Em vẫn luôn tự hỏi chính mình, rằng chia tay với anh là một quyết định sai hay đúng?
	// Không gian như vỡ tan, hạnh phúc đang dần phai nhoà
	// Nhẹ quay lưng nhìn tháng ngày, mình đã tay trong tay
	// Quá khứ như cơn mơ, nhói đay vô bờ, giờ em một mình bơ vơ
	// Khoảnh khắc em quay đi, sao anh lạnh lùng đến thế?
	// Anh nói em nghe đi, dẫu chỉ thầm thì..rằng anh cần em mọi khi
	// Vì tình yêu thì không đúng sai gì
	// Buồn làm chi để đêm đêm kí ức hoen bờ mi
	// Anh/em cứ ngỡ mặt xa cách lòng..
	// Niềm cô đơn mỗi đêm sẽ lướt qua trong nỗi nhớ...
	// Niềm cô đơn mỗi đêm sẽ tan biến...trong kỉ niệm...
	// DateCreated: 2013-11-7 14:18:32
	// SameArtist: 0
	// Plays > 0
	// Date updated is correct

}
