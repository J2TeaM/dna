package nct

import (
	"dna"
)

func ExampleGetAPIVideo() {
	_, err0 := GetAPIVideo(28760555)
	if err0 == nil {
		panic("Video id 28760555 has to have an error")
	} else {
		if err0.Error() != "NCT - Video ID:28760555 not found" {
			panic("Wrong error message")
		}
	}
	video, err := GetAPIVideo(2876055)
	if err == nil {
		if video.Likes < 1474 || video.Plays < 556700 {
			panic("Cannot get video likes or plays")
		} else {
			video.Likes = 1474
			video.Plays = 556700
		}
		if video.StreamUrl.Match(`\.mp4`) == false {
			panic("Video link has to be mp4 format")
		} else {
			video.StreamUrl = "http://nplus.nixcdn.com/bd4aa514acf9173ca298424dce992639/52f83228/PreNCT7/GuiChoAnhPhan2-KhoiMy-2876055.mp4"
		}
		dna.LogStruct(video)
	} else {
		dna.Log(err.Error())
	}
	// Output:
	// Id : 2876055
	// Key : "N5QeESGm7ICBt"
	// Title : "Gửi Cho Anh (Phần 2)"
	// Thumbnail : "http://avatar.nct.nixcdn.com/mv/2013/12/10/e/0/5/3/1386640122904.jpg"
	// Image : "http://avatar.nct.nixcdn.com/mv/2013/12/10/e/0/5/3/1386640122904_536.jpg"
	// Artist : "Khởi My"
	// Time : "46:08"
	// Artistid : 12987
	// Likes : 1474
	// Plays : 556700
	// Linkshare : "http://www.nhaccuatui.com/video/gui-cho-anh-phan-2-khoi-my.N5QeESGm7ICBt.html"
	// StreamUrl : "http://nplus.nixcdn.com/bd4aa514acf9173ca298424dce992639/52f83228/PreNCT7/GuiChoAnhPhan2-KhoiMy-2876055.mp4"
	// ObjType : "VIDEO"
}

func ExampleGetAPISong() {
	_, err0 := GetAPISong(28760555)
	if err0 == nil {
		panic("Song id 28760555 has to have an error")
	} else {
		if err0.Error() != "NCT - Song ID:28760555 not found" {
			panic("Wrong error message")
		}
	}
	song, err := GetAPISong(2854574)
	if err == nil {
		if song.Likes < 277 || song.Plays < 416163 {
			panic("Cannot get song likes or plays")
		} else {
			song.Likes = 277
			song.Plays = 416163
		}
		if song.StreamUrl.Match(`\.mp3`) == false || song.Linkdown.Match(`\.mp3`) == false || song.LinkdownHQ.Match(`\.mp3`) == false {
			panic("Song link has to be mp3 format")
		} else {
			song.StreamUrl = "http://a.nixcdn.com/96e84c4eb0e143c6259c6ab2cb533c7a/52f83930/NhacCuaTui844/AnhBiet-HoQuangHieu-2854574.mp3"
			song.Linkdown = "http://download.a.nixcdn.com/96e84c4eb0e143c6259c6ab2cb533c7a/52f83930/NhacCuaTui844/AnhBiet-HoQuangHieu-2854574.mp3"
			song.LinkdownHQ = "http://download.a.nixcdn.com/5ee0ab8b469dc2c9cf9e93a98f46fa0b/52f83930/NhacCuaTui844/AnhBiet-HoQuangHieu-2854574_hq.mp3"
		}
		if song.Image != "" {
			song.Image = "http://avatar.nct.nixcdn.com/singer/avatar/2013/11/28/b/f/8/d/1385631181033.jpg"
		}
		dna.LogStruct(song)
	} else {
		dna.Log(err.Error())
	}
	// Output:
	// Id : 2854574
	// Key : "uUPvpEvU4CmH"
	// Title : "Anh Biết"
	// Artist : "Hồ Quang Hiếu"
	// Likes : 277
	// Plays : 416163
	// LinkShare : "http://www.nhaccuatui.com/bai-hat/anh-biet-ho-quang-hieu.uUPvpEvU4CmH.html"
	// StreamUrl : "http://a.nixcdn.com/96e84c4eb0e143c6259c6ab2cb533c7a/52f83930/NhacCuaTui844/AnhBiet-HoQuangHieu-2854574.mp3"
	// Image : "http://avatar.nct.nixcdn.com/singer/avatar/2013/11/28/b/f/8/d/1385631181033.jpg"
	// Coverart : ""
	// ObjType : "SONG"
	// Duration : 280
	// Linkdown : "http://download.a.nixcdn.com/96e84c4eb0e143c6259c6ab2cb533c7a/52f83930/NhacCuaTui844/AnhBiet-HoQuangHieu-2854574.mp3"
	// LinkdownHQ : "http://download.a.nixcdn.com/5ee0ab8b469dc2c9cf9e93a98f46fa0b/52f83930/NhacCuaTui844/AnhBiet-HoQuangHieu-2854574_hq.mp3"
}

func ExampleGetAPIAlbum() {
	_, err0 := GetAPIAlbum(28760555)
	if err0 == nil {
		panic("Album id 28760555 has to have an error")
	} else {
		if err0.Error() != "NCT - Album ID:28760555 not found" {
			panic("Wrong error message")
		}
	}
	album, err := GetAPIAlbum(12255234)
	if err == nil {
		if album.Likes < 128 || album.Plays < 320384 {
			panic("Cannot get album likes or plays")
		} else {
			album.Likes = 128
			album.Plays = 320384
		}
		if len(album.Listsong) != 3 {
			panic("Songs of the album has to have length equal to 3")
		} else {
			album.Listsong = []APISong{}
		}
		dna.LogStruct(album)
	} else {
		dna.Log(err.Error())
	}
	// Output:
	// Id : 12255234
	// Key : "nsnkteavOHbX"
	// Title : "Biết Trước Sẽ Không Mất Nhau (Single)"
	// Thumbnail : "http://avatar.nct.nixcdn.com/playlist/2013/11/28/4/e/9/b/1385644142690_300.jpg"
	// Coverart : "http://avatar.nct.nixcdn.com/playlist/2013/11/28/4/e/9/b/1385644142690_500.jpg"
	// Image : "http://avatar.nct.nixcdn.com/playlist/2013/11/28/4/e/9/b/1385644142690_300.jpg"
	// Artist : "Vĩnh Thuyên Kim, Hồ Quang Hiếu"
	// Likes : 128
	// Plays : 320384
	// Linkshare : "http://www.nhaccuatui.com/playlist/biet-truoc-se-khong-mat-nhau-single-vinh-thuyen-kim-ho-quang-hieu.nsnkteavOHbX.html"
	// Listsong : []nct.APISong{}
	// Description : "Sau sự kết hợp thành công cùng các nữ ca sĩ xinh đẹp trong thời gian gần đây như: Bảo Thy, Nhật Kim Anh, Lương Khánh Vy... Chàng hotboy Hồ Quang Hiếu có sự kết hợp mới cùng Vĩnh Thuyên Kim trong một sáng tác của Lê Chí Trung \"Biết Trước Sẽ Không Mất Nhau\"."
	// Genre : "Nhạc Trẻ"
	// ObjType : "PLAYLIST"
}

func ExampleGetAPILyric() {
	_, err0 := GetAPILyric(29097271)
	if err0 == nil {
		panic("SongLyric id 29097271 has to have an error")
	} else {
		if err0.Error() != "NCT - Song ID:29097271 Lyric not found" {
			panic("Wrong error message")
		}
	}
	lyric, err := GetAPILyric(2909727)
	if err == nil {
		dna.LogStruct(lyric)
	} else {
		dna.Log(err.Error())
	}
	// Output:
	// Lyricid : 950870
	// Lyric : "Bài hát: Mảnh Ghép Đã Vỡ - Minh Vương M4U \nSáng tác: Hoàng Bảo Nam \n\nNgày tháng anh lệ rơi \nVì khóc bao đêm nhớ mong một người \nNiềm đau chôn sâu cay đắng khi em vô tình lãng quên \nGiờ anh biết phải làm sao \nĐể xoá đi bao kỷ niệm \n\nThật quá khó khi anh vẫn còn yêu em \nỞ nơi phương trời xa kia \nĐã khiến trong em đổi thay thật rồi\nTình yêu bao năm đậm sâu sẽ mãi chỉ là giấc mơ \nGiờ em đã có người yêu thay thế trong anh rất nhiều \nCòn đâu nữa thời gian như trước em dành cho anh \n\n[ĐK:] \nLời nói chia tay hôm qua sao quá nghiệt ngã \nBật khóc trong đêm anh nghe tiếng em lần cuối \nChẳng lẽ anh chỉ là người lấp chổ khoảng trống trong em \nMỗi khi em buồn \n\nHạnh phúc nay đã vỡ tan như hoa thuỷ tinh \nMảnh ghép yêu thương trong anh sẽ không bao giờ lành \nHọc cách quên đi một người quá khó \nVì nỗi đau ngày qua cứ mãi để lại \nSẽ không bao giờ phôi phai"
	// Songid : 2909727
	// Status : 3
	// TimedLyric : "\ufeff[00:15.92]Bài hát: Mảnh Ghép Đã Vỡ\n[00:16.93]Ca sĩ: Minh Vương M4U \n[00:18.36]Sáng tác: Hoàng Bảo Nam \n[00:19.53]\n[00:27.71]Ngày tháng anh lệ rơi \n[00:29.91]Vì khóc bao đêm\n[00:31.23] nhớ mong một người \n[00:33.98]Niềm đau chôn sâu\n[00:35.06] cay đắng khi em\n[00:36.22] vô tình lãng quên \n[00:38.38]Giờ anh biết\n[00:39.28] phải làm sao \n[00:40.97]Để xoá đi bao kỷ niệm \n[00:43.53]Thật quá khó\n[00:44.58] khi anh vẫn\n[00:45.26] còn yêu em \n[00:46.26]\n[02:25.65][00:47.54]Ở nơi phương trời xa kia \n[02:27.98][00:49.94]Đã khiến trong em\n[02:29.20][00:51.22] đổi thay thật rồi \n[02:31.63][00:53.73]Tình yêu bao năm\n[02:32.83][00:55.14] đậm sâu sẽ mãi \n[02:34.11][00:56.24]chỉ là giấc mơ \n[02:36.40][00:58.51]Giờ em đã có \n[02:37.65][00:59.79]người yêu thay thế \n[02:39.11][01:01.27]hơn anh rất nhiều \n[02:41.55][01:03.58]Còn đâu nữa \n[02:42.43][01:04.46]thời gian như trước\n[02:43.56][01:05.74] em dành cho anh \n[02:45.84][01:07.92]\n[03:34.55][02:48.26][01:10.18]Lời nói chia tay hôm qua\n[03:36.59][02:50.17][01:12.30] sao quá nghiệt ngã \n[03:39.51][02:52.97][01:15.07]Bật khóc trong đêm \n[03:40.76][02:54.21][01:16.33]anh nghe tiếng\n[03:41.98][02:55.48][01:17.63] em lần cuối \n[03:44.58][02:58.13][01:20.07]Chẳng lẽ anh\n[03:45.38][02:58.93][01:20.94] chỉ là người lấp chổ\n[03:48.28][03:01.84][01:23.92] khoảng trống trong em \n[03:50.55][03:04.18][01:26.28]Mỗi khi em buồn \n[03:53.61][03:06.37][01:29.14]\n[03:54.75][03:08.21][01:30.23]Hạnh phúc nay \n[03:55.53][03:08.99][01:30.98]đã vỡ tan \n[03:56.67][03:10.17][01:32.17]như hoa thuỷ tinh \n[03:59.57][03:13.32][01:35.20]Mảnh ghép yêu thương \n[04:00.73][03:14.48][01:36.36]trong anh sẽ không \n[04:02.69][03:16.20][01:38.33]bao giờ lành \n[04:04.62][03:17.99][01:40.23]Học cách quên đi \n[04:05.97][03:19.32][01:41.41]một người quá khó \n[04:08.47][03:21.88][01:44.12]Vì nỗi đau\n[04:09.53][03:23.03][01:45.27] ngày qua cứ mãi để lại \n[04:13.81][03:27.38][01:49.50]Sẽ không\n[04:14.82][03:28.28][01:50.62] bao giờ phôi phai\n[03:29.51][01:52.70]\n"
	// TimedLyricFile : "2014/01/13/d/0/5/b/1389584944210.lrc"
	// UsernameCreated : "nct_official"
}

func ExampleGetAPIArtist() {
	_, err0 := GetAPIArtist(496741)
	if err0 == nil {
		panic("SongArtist id 496741 has to have an error")
	} else {
		if err0.Error() != "NCT - Artist ID:496741 not found" {
			panic("Wrong error message")
		}
	}
	artist, err := GetAPIArtist(49674)
	if err == nil {
		dna.LogStruct(artist)
	} else {
		dna.Log(err.Error())
	}
	if artist.Avatar != "" {
		artist.Avatar = "http://avatar.nct.nixcdn.com/singer/avatar/2013/12/16/7/f/f/7/1387176808800.jpg"
	}
	// Output:
	// Id : 49674
	// Name : "Đàm Vĩnh Hưng"
	// Avatar : "http://avatar.nct.nixcdn.com/singer/avatar/2014/13/7E4149A9_2.jpg"
	// NSongs : 0
	// NAlbums : 0
	// NVideos : 0
	// ObjType : "ARTIST"
}
