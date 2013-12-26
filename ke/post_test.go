package ke

import (
	"dna"
)

func ExampleGetLyric() {
	lyric, err := GetLyric(1944090)
	if err == nil {
		dna.LogStruct(lyric)
	} else {
		// dna.Log(lyric)
		panic("Error has to be nil")
	}
	// Output:
	// Data : "<p>\r\n\t<strong>Em Sẽ Hạnh Phúc<br />\r\n\t</strong></p>\r\n<p>\r\n\t---</p>\r\n<p>\r\n\tNgười yêu ơi anh biết mình sai nên để cho em rời xa yêu dấu hôm qua hãy mang em đi thật xa</p>\r\n<p>\r\n\tDù anh biết năm tháng dần qua sẽ xóa đi bóng hình em trong trái tim anh</p>\r\n<p>\r\n\tRồi em sẽ hạnh phúc thôi!</p>\r\n<p>\r\n\tNhững khó khăn em mong anh biết dù rằng mình đã cách xa nhưng anh vẫn nhiều lo lắng</p>\r\n<p>\r\n\tNắm tay nhau qua bao tháng ngày để đến hôm nay lạc bước yêu dấu phai tàn</p>\r\n<p>\r\n\tNgày không em anh như mất lối, không em anh như chơi vơi</p>\r\n<p>\r\n\tKhông em bên anh không có ai kề môi</p>\r\n<p>\r\n\tSẽ bên nhau khi duyên đã lỡ trao nhau yêu thương đã vỡ</p>\r\n<p>\r\n\tChỉ vì yêu em anh sẽ chỉ yêu mình em</p>\r\n<p>\r\n\tNgười yêu ơi anh biết mình sai nên để cho em rời xa yêu dấu hôm qua hãy mang em đi thật xa</p>\r\n<p>\r\n\tDù anh biết năm tháng dần qua sẽ xóa đi bóng hình em trong trái tim anh</p>\r\n<p>\r\n\tRồi em sẽ hạnh phúc thôi!</p>\r\n<p>\r\n\t<em>Và sẽ có người thay thế anh trong giấc mơ</em></p>"
	// Status : 1
}

func ExampleGetAPIAlbum() {
	album, err := GetAPIAlbum(86682)

	if err == nil {
		if album.Plays < 44450 {
			panic("Plays has to be greater than 44450")
		}
		album.Plays = 44450
		for _, song := range album.SongList {
			song.Plays = 10129
		}
		var length = len(album.SongList)
		album.SongList = nil
		dna.LogStruct(album)
		dna.Log("Lenght :",length)
	} else {
		panic("Error has to be nil")
	}
	// Output:
	// Id : 86682
	// Title : "Chờ Hoài Giấc Mơ"
	// Artists : "Akio Lee ft Akira Phan"
	// Coverart : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias_6/album/image/2013/12/17/ea441cd919ead9a14ad155b184c6a4370e50c0b6_103_103.jpg"
	// Url : "http://keeng.vn/album/Cho-Hoai-Giac-Mo-Akio-Lee/2K2O4QG8.html"
	// Plays : 44450
	// SongList : []ke.APISong(nil)
// Lenght : 5

}

func ExampleGetAPISongEntry() {
	apiSongEntry, err := GetAPISongEntry(1968535)
	if err == nil {
		if apiSongEntry.MainSong.Plays < 14727 {
			panic("Plays has to be greater than 14727")
		}
		apiSongEntry.MainSong.Plays = 14727
		for _, song := range apiSongEntry.RelevantSongs {
			song.Plays = 14727
		}
		var length = len(apiSongEntry.RelevantSongs)
		dna.Log("MAIN SONG:")
		dna.LogStruct(&apiSongEntry.MainSong)
		dna.Log("SIMILAR SONGS LENGTH:",length)
	} else {
		panic("Error has to be nil")
	}
	// Output:
// MAIN SONG:
// Id : 1968535
// Title : "Giáng Sinh Không Nhà"
// Artists : "Hồ Quang Hiếu"
// Plays : 14727
// ListenType : 0
// Lyric : "<p>\r\n\t<strong>Giáng Sinh Không Nhà </strong></p>\r\n<p>\r\n\t1. Chân bước đi dưới muôn ánh đèn đêm<br />\r\n\tNhưng cớ sao vẫn luôn thấy quạnh hiu<br />\r\n\tThêm Giáng Sinh nữa con đã không ở nhà.</p>\r\n<p>\r\n\tTrên phố đông tấp nập người lại qua<br />\r\n\tNhưng trái tim vẫn luôn nhớ nơi xa<br />\r\n\tCon muốn quay bước chân muốn trở về nhà.</p>\r\n<p>\r\n\t[ĐK:]<br />\r\n\tVề nghe gió đông đất trời giá lạnh<br />\r\n\tĐể ngồi nép bên nhau lòng con ấm hơn<br />\r\n\tNhìn theo ánh sao đêm gửi lời chúc lành<br />\r\n\tGiờ tâm trí con mong...về nhà.</p>\r\n<p>\r\n\t2. Khi tiếng chuông ngân lên lúc nửa đêm<br />\r\n\tThấy xuyến xao giống như những ngày xưa<br />\r\n\tTheo lũ bạn tung tăng đi xem nhà thờ.</p>\r\n<p>\r\n\tNhư cánh chim đến lúc cũng bay xa<br />\r\n\tCon đã mang theo mình những ước vọng<br />\r\n\tNhưng lúc này bâng khuâng con nhớ mọi người.</p>\r\n<p>\r\n\t[ĐK]<br />\r\n\tNhững ký ức ấm áp mãi như đang còn<br />\r\n\tVà sẽ giúp con đứng vững trước những phong ba<br />\r\n\tTrong tim con luôn yêu và nhớ thiết tha<br />\r\n\tMarry christmas, giáng sinh bình an.</p>"
// Link : "http://media2.keeng.vn/medias/audio/2013/12/19/a5cc9183b876c288f099e97aecc189f004b2137a_128.mp3"
// MediaUrlMono : "http://media2.keeng.vn/medias/audio/2013/12/19/a5cc9183b876c288f099e97aecc189f004b2137a_24.mp3"
// MediaUrlPre : "http://media2.keeng.vn/medias/audio/2013/12/19/a5cc9183b876c288f099e97aecc189f004b2137a_128.mp3"
// DownloadUrl : "http://media2.keeng.vn/medias/audio/2013/12/19/a5cc9183b876c288f099e97aecc189f004b2137a.mp3"
// IsDownload : 1
// RingbacktoneCode : ""
// RingbacktonePrice : 0
// Url : "http://keeng.vn/audio/Giang-Sinh-Khong-Nha-Ho-Quang-Hieu-320Kbps/N9TUO6DI.html"
// Price : 1000
// Copyright : 1
// CrbtId : 0
// Coverart : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias2/singer/2013/06/13/a05019cceb7742d108159d661d894f19bc886eb1_103_103.jpg"
// Coverart310 : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias2/singer/2013/06/13/a05019cceb7742d108159d661d894f19bc886eb1_310_310.jpg"
// SIMILAR SONGS LENGTH: 10

}

func ExampleGetAPIArtistEntry() {
	apiArtistEntry, err := GetAPIArtistEntry(1394)
	if err == nil {
		dna.LogStruct(apiArtistEntry)
	} else {
		panic("Error has to be nil")
	}
	// Output:
	// Artist : ke.APIArtistProfile{Id:1394, Title:"Minh Hằng", Coverart:"http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias_6/singer/2013/11/19/1226beceb5d049998976a08f822c0cc9037c0a32_103_103.jpg"}
	// Nsongs : 46
	// Nalbums : 7
	// Nvideos : 24
}

func ExampleGetAPIArtistSongs(){
	apiArtistSongs, err := GetAPIArtistSongs(1394,1,1)
	if err == nil{
		if len(apiArtistSongs.Data) != 1 {
			panic("Lenght of artist songs has to be one")
		}
		song := &apiArtistSongs.Data[0]
		if song.Plays < 2152 {
			panic("Plays has to be greater than 2152")
		}
		song.Plays = 2152
		dna.LogStruct(song)
	} else {
		panic("Error has to be nil")
	}
	// Output:
// Id : 553150
// Title : "Biết Yêu"
// Artists : "Thủy Tiên ft Đông Nhi ft Yến Trang ft Minh Hằng"
// Plays : 2152
// ListenType : 0
// Lyric : ""
// Link : "http://media.keeng.vn/medias_8/audio/2011/08/19/14/biet-yeu-553150.mp3"
// MediaUrlMono : "http://media.keeng.vn/medias_8/audio/2011/08/19/14/biet-yeu-553150_24.mp3"
// MediaUrlPre : "http://media.keeng.vn/medias_8/audio/2011/08/19/14/biet-yeu-553150.mp3"
// DownloadUrl : "http://media.keeng.vn/medias_8/audio/2011/08/19/14/biet-yeu-553150.mp3"
// IsDownload : 1
// RingbacktoneCode : ""
// RingbacktonePrice : 0
// Url : "http://keeng.vn/audio/Biet-Yeu-Thuy-Tien-192Kbps/H1REY5PC.html"
// Price : 1000
// Copyright : 1
// CrbtId : 0
// Coverart : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias2/singer/2013/02/22/f8898ddb7d8838a83c19f1c121b6a6082d14a03d_103_103.jpg"
// Coverart310 : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias2/singer/2013/02/22/f8898ddb7d8838a83c19f1c121b6a6082d14a03d_310_310.jpg"
}

func ExampleGetAPIArtistAlbums(){
	apiArtistAlbums, err := GetAPIArtistAlbums(1394,1,1)
	if err == nil{
		if len(apiArtistAlbums.Data) != 1 {
			panic("Lenght of artist songs has to be one")
		}
		album := &apiArtistAlbums.Data[0]
		if album.Plays < 48763 {
			panic("Plays has to be greater than 48763")
		}
		album.Plays = 48763
		dna.LogStruct(album)
	}else {
		panic("Error has to be nil")
	}
 	// Output:
// Id : 86646
// Title : "Hit Tháng 11/2013"
// Artists : "Hồ Quang Hiếu ft MTV ft Hương Tràm ft Chi Dân ft Minh Hằng ft Khởi My ft Thùy Chi ft Trà My Idol ft Vy Oanh ft Nam Cường"
// Coverart : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias_6/album/image/2013/11/29/14d50794fc202324de800fa336c3b5cccb35888b_103_103.jpg"
// Url : "http://keeng.vn/album/Hit-Thang-11-2013-Ho-Quang-Hieu/3JQN0O88.html"
// Plays : 48763
// SongList : []ke.APISong(nil)
}

func ExampleGetAPIArtistVideos(){
	apiArtistVideos, err := GetAPIArtistVideos(1394,1,1)
	if err == nil{
		if len(apiArtistVideos.Data) != 1 {
			panic("Lenght of artist songs has to be one")
		}
		video := &apiArtistVideos.Data[0]
		if video.Plays < 1010 {
			dna.Log("Plays has to be greater than 1010")
		}
		video.Plays = 1010
		dna.LogStruct(video)
	} else {
		panic("Error has to be nil")
	}
 	// Output:
// Id : 200452
// Title : "Ngày Tết Quê Em"
// Artists : "Hồ Ngọc Hà ft Minh Hằng ft V.music"
// Plays : 1010
// ListenType : 0
// Link : "http://media.keeng.vn/medias_2/video/2013/05/20/5599e98d2e1e0eb69e15fc37cb2a7d1b4beba0b5_mp4_640_360.mp4"
// IsDownload : 0
// DownloadUrl : "http://media.keeng.vn/medias_2/video/2013/05/20/5599e98d2e1e0eb69e15fc37cb2a7d1b4beba0b5.mp4"
// RingbacktoneCode : ""
// RingbacktonePrice : 0
// Url : "http://keeng.vn/video/Ngay-Tet-Que-Em-Ho-Ngoc-Ha/4F5DB8F9.html"
// Price : 0
// Copyright : 0
// CrbtId : 0
// Thumbnail : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias/images/video/video20120501/ngay-tet-que-em-473_147_83.jpg"
}
