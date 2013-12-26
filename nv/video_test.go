package nv

import (
	. "dna"
	"testing"
	"time"
)

func TestGetVideo(t *testing.T) {
	_, err := GetVideo(15)
	if err == nil {
		t.Error("Video 15 has to have an error")
	}
	if err.Error() != "Nhacvui - Video 15 : Link not found" {
		t.Errorf("Error message has to be: %v", err.Error())
	}
}
func ExampleGetVideo() {
	video, err := GetVideo(476465)
	PanicError(err)
	if video.Plays < 0 {
		panic("Plays has to be greater than 0")
	}
	video.Plays = 0
	video.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if !video.Link.Match("mp4") {
		panic("NO link found")
	}
	video.Link = "http://stream306.nhac.vui.vn/uploadmusic2/ce52c1a3f8af86140763dd7885690f0f/52a5f841....mp4"
	LogStruct(video)

	// Output:
	// Id : 476465
	// Title : "Phải Chi Lúc Trước Anh Sai"
	// Artists : dna.StringArray{"Mr.Siro"}
	// Authors : dna.StringArray{}
	// Topics : dna.StringArray{"Việt Nam"}
	// Plays : 0
	// Lyric : "Lời bài hát: Phải Chi Lúc Trước Anh Sai - Mr.Siro\nNgười đóng góp: AdministratorNgày mai em nói sẽ xa căn phòng này\n Nhìn theo con sóng biết bao nhiêu ngày rồi ngồi đây với em\n Người từng đã nói \"em chỉ cần sống gần anh\"\n Lời nói ấy còn đó nhưng trong em dường như đã khác rồi\n Nhớ không? Ai hứa yêu ai dài lâu rồi đổi thay ...\n Thế gian này nào ai muốn trái tim tổn thương thêm một lần nữa\n Một khi người đã thấy được hạnh phúc thì đừng nên dối lòng mình\n Lòng em đã quên thì đừng cố nhớ\n Nụ hôn dưới mái hiên xưa\n Kỷ niệm xé nát đêm mưa\n Gom từng hơi thở để nuôi nỗi đau quá dài\n Tại sao phải đến bên ai, phải chi lúc trước anh sai\n Tình yêu đâu ai chấp nhận... Dốc hết lòng yêu nhau vẫn mất nhau... phải không?\n \nBabe. babe, ngay bây giờ em đang ở đâu babe?\n Có biết anh rất nhớ em, mình làm lại từ đầu nhé ...\n Tại sao phải đến bên ai ... phải chi lúc trước anh sai (thì em mới nên như thế !)\n Look back before you leave, you leave my life\n \nNhớ không? Ai hứa yêu ai dài lâu rồi đổi thay ...\n Thế gian này nào ai muốn trái tim tổn thương thêm một lần nữa\n Một khi người đã thấy được hạnh phúc thì đừng nên dối lòng mình\n Lòng em đã quên thì đừng cố nhớ\n Nụ hôn dưới mái hiên xưa\n Kỷ niệm xé nát đêm mưa\n Gom từng hơi thở để nuôi nỗi đau quá dài\n Tại sao phải đến bên ai, phải chi lúc trước anh sai\n Để em ngày mai bước cùng ai đó thật hạnh phúc trên nỗi đau của anh\n \nPhải chi lúc trước anh sai\n Phải chi lúc trước anh sai thì em mới nên như thế\n \nGom từng hơi thở để nuôi nỗi đau quá dài"
	// Link : "http://stream306.nhac.vui.vn/uploadmusic2/ce52c1a3f8af86140763dd7885690f0f/52a5f841....mp4"
	// Link320 : ""
	// Thumbnail : "http://nhac.vui.vn/imageupload/upload2013/2013-4/video_clip/2013-12-12/1386839989_Phai-Chi-Luc-Truoc-Anh-Sai-Video-Clip-Mrsiro.jpg"
	// Type : "video"
	// Checktime : "2013-11-21 00:00:00"
}
