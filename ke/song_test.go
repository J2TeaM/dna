package ke

import (
	"dna"
	"time"
)

func ExampleGetSong() {
	_, err := GetSong(1966782)
	if err == nil {
		panic("Song has to have an error")
	} else {
		if err.Error() != "Keeng - Song ID: 1966782 not found" {
			panic("Wrong error message!")
		}
	}
	song, err := GetSong(1967829)
	dna.PanicError(err)
	song.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if song.Plays < 126222 {
		panic("Wrong play")
	}
	song.Plays = 126222
	dna.LogStruct(song)
	// Output:
	// Id : 1967829
	// Key : "LKD633FY"
	// Title : "Em Phải Làm Sao"
	// Artists : dna.StringArray{"Mỹ Tâm"}
	// Plays : 126222
	// ListenType : 0
	// HasLyric : true
	// Lyric : "<p>\r\n\t<span class=\"lyric\">Làm sao để cho tình ta giờ đây vui như lúc đầu<br />\r\n\tLàm sao để cho nụ hôn nồng say thôi mang nỗi sầu<br />\r\n\tLàm sao để cho từng đêm mình em thôi không đớn đau vì nhau<br />\r\n\tLàm sao để anh sẽ tin rằng em chỉ yêu mỗi anh.<br />\r\n\t<br />\r\n\tLàm sao để cho tình anh giờ đây thôi không hững hờ<br />\r\n\tLàm sao để cho tình không là mơ em thôi thẫn thờ<br />\r\n\tVì em ngày xưa từng mang lầm lỡ nên đâu dám mơ mộng gì<br />\r\n\tLàm sao để cho thời gian đừng bôi xóa đi ngày thơ.<br />\r\n\t<br />\r\n\t[ĐK:]<br />\r\n\tNgười hỡi, người có biết em vẫn luôn cần anh<br />\r\n\tĐể sưởi ấm cho trái tim mỏng manh<br />\r\n\tMùa đông có anh sẽ qua thật nhanh<br />\r\n\tLời yêu thương đó, cố giữ trên môi<br />\r\n\tMà sao tiếng yêu trong anh như mây chiều trôi<br />\r\n\tTình mình sao cứ mãi xa xôi...<br />\r\n\t<br />\r\n\tNgười hỡi, người có biết em vẫn luôn còn đây<br />\r\n\tDù bao giấc mơ cuốn theo làn mây<br />\r\n\tDù anh có quên hết bao nồng say<br />\r\n\tTình em vẫn thế, dẫu chẳng ai hay<br />\r\n\tVà bao ước mơ bên anh nay xa tầm tay<br />\r\n\tNguyện rằng em sẽ mãi yêu không đổi thay.</span></p>"
	// Link : "http://streaming1.keeng.vn/1967829_1_wap.mp3"
	// MediaUrlMono : "http://media2.keeng.vn/medias/audio/2013/12/09/54ec35cd058b422c1b4b9f812cdb5726dee66f9c_24.mp3"
	// MediaUrlPre : "http://streaming1.keeng.vn/1967829_1_wap.mp3"
	// DownloadUrl : "http://streaming1.keeng.vn/1967829_1.mp3"
	// IsDownload : 1
	// RingbacktoneCode : "7333357"
	// RingbacktonePrice : 5000
	// Price : 1000
	// Copyright : 1
	// CrbtId : 0
	// Coverart : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias2/singer/2013/01/21/673a437247a1f039bbde97119385c16524131f8f_103_103.jpg"
	// Coverart310 : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias2/singer/2013/01/21/673a437247a1f039bbde97119385c16524131f8f_310_310.jpg"
	// DateCreated : "0001-01-01 00:00:00"
	// Checktime : "2013-11-21 00:00:00"
}
