package ke

import (
	"dna"
)

func ExampleGetAPILyric() {
	lyric, err := GetAPILyric(1944090)
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
		dna.Log("Lenght :", length)
	} else {
		panic("Error has to be nil")
	}
	// Output:
	// Id : 86682
	// Title : "Chờ Hoài Giấc Mơ"
	// Artists : "Akio Lee ft Akira Phan"
	// Coverart : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias_6/album/image/2014/01/10/27a4cdcbb7f60aa123529a49d76888708ec8872d_103_103.jpg"
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
		dna.Log("SIMILAR SONGS LENGTH:", length)
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
		if apiArtistEntry.Nsongs < 46 {
			panic("Nsongs has to be greater than 46")
		}
		if apiArtistEntry.Nalbums < 7 {
			panic("Nalbums has to be greater than 7")
		}
		if apiArtistEntry.Nvideos < 24 {
			panic("Nvideos has to be greater than 24")
		}
		apiArtistEntry.Nsongs = 46
		apiArtistEntry.Nalbums = 7
		apiArtistEntry.Nvideos = 24
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

func ExampleGetAPIArtistSongs() {
	apiArtistSongs, err := GetAPIArtistSongs(1394, 1, 1)
	if err == nil {
		if len(apiArtistSongs.Data) != 1 {
			panic("Lenght of artist songs has to be one")
		}
		song := &apiArtistSongs.Data[0]
		if song.Plays == 0 {
			panic("Plays has to be different from 0, GOT:" + song.Plays.ToString().String())
		}
		song.Plays = 0
		switch {
		case song.Link == "", song.MediaUrlMono == "", song.MediaUrlPre == "", song.DownloadUrl == "":
			panic("Empty URLS")
		}
	} else {
		panic("Error has to be nil")
	}
}

func ExampleGetAPIArtistAlbums() {
	apiArtistAlbums, err := GetAPIArtistAlbums(1394, 1, 1)
	if err == nil {
		if len(apiArtistAlbums.Data) != 1 {
			panic("Lenght of artist songs has to be one")
		}
		album := &apiArtistAlbums.Data[0]
		if album.Plays == 0 {
			panic("Plays has to be different from 0")
		}
		album.Plays = 0
		switch {
		case album.Coverart == "", album.Url == "":
			panic("Empty urls")
		}
		// dna.LogStruct(album)
	} else {
		panic("Error has to be nil")
	}

}

func ExampleGetAPIArtistVideos() {
	apiArtistVideos, err := GetAPIArtistVideos(1394, 1, 1)
	if err == nil {
		if len(apiArtistVideos.Data) != 1 {
			panic("Lenght of artist songs has to be one")
		}
		video := &apiArtistVideos.Data[0]
		if video.Plays == 0 {
			dna.Log("Plays has to be greater different from 0")
		}
		video.Plays = 1010
		switch {
		case video.Link == "", video.DownloadUrl == "", video.Url == "":
			panic("Wrong urls")
		}
		dna.LogStruct(video)
	} else {
		panic("Error has to be nil")
	}
}
