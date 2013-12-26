package cc

import (
	. "dna"
	"testing"
	"time"
)

func TestGetVideo(t *testing.T) {
	_, err := GetVideo(70641)
	if err == nil {
		t.Error("Video 70641 has to have an error")
	}
	if err.Error() != "Chacha - Video 70641: Link not found" {
		t.Errorf("Error message has to be: %v", err.Error())
	}
}

func ExampleGetVideo() {
	video, err := GetVideo(70370)
	PanicError(err)
	if video.Plays < 274 {
		panic("Plays has to be greater than 274")
	}
	video.Plays = 274
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	LogStruct(video)
	// Output:
	// Id : 70370
	// Title : "Yêu Trong Ảo Mộng"
	// Artists : dna.StringArray{"Lâm Chấn Huy"}
	// Topics : dna.StringArray{"Nhạc trẻ"}
	// Plays : 274
	// ResolutionFlags : 7
	// Thumbnail : "http://s2.chacha.vn/videos/img/s1/8/70370/70370.jpg"
	// Lyric : "Bài Hát: Yêu Trong Ảo Mộng\nTác Giả: Phạm Khánh Hưng\nCố khép đôi mắt cay lại, để lệ sẽ không tuôn trào\nHôm nay người ta về đón rước nàng dâu, vỡ nát tim đau thật đau\nChấm dứt cho cuộc tình đầu, khi ta đã không còn nhau\nHôm nay người đi là mãi mãi biệt li mãi mãi sẽ chẳng còn chi.\n\nBao mộng mơ lúc xưa ta mơ chỉ là những ảo mộng\nMà mơ làm chi đến khi biệt li tình đau.\n\n[ĐK:]\nGạt đi nước mắt con tim tôi ơi thôi đừng khóc\nVì khóc cũng chẳng thể mang yêu thương xưa quay về\nGiờ lòng không muốn ở lại thì xin em cứ ra đi\nAnh không cầu mong ước chi.\n\nTình yêu xưa đó dẫu cho số kiếp chỉ là mộng mơ\nMà sao con tim ta vẫn yêu mãi mãi\nSẽ chẳng bao giờ tỉnh giấc sau cơn mơ\nYêu em ta mãi yêu em trong ảo mộng"
	// Links : dna.StringArray{"http://video.chacha.vn/a9495345/videos/output/8/70370/1.mp4", "http://video.chacha.vn/e47ba9cd/videos/output/8/70370/7.mp4", "http://video.chacha.vn/67715410/videos/output/8/70370/8.mp4"}
	// YearReleased : 2013
	// Checktime : "2013-11-21 00:00:00"
}
