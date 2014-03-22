package zi

import (
	. "dna"
)

func ExampleGetAPISong() {
	apisong, err := GetAPISong(1381645456)
	PanicError(err)
	if apisong.Plays < 0 {
		panic("Plays has to be greater than 0")
	}
	if apisong.Likes < 0 {
		panic("Likes has to be greater than or equal to 0")
	}
	if apisong.Comments < 0 {
		panic("Comments has to be greater than or equal to 0")
	}
	apisong.Plays = 2000
	apisong.Likes = 0
	apisong.Comments = 0
	switch {
	case apisong.LinkDownload["128"] == "" || apisong.LinkDownload["320"] == "" || apisong.LinkDownload["lossless"] == "":
		panic("Link Download has some values")
	case apisong.Source["128"] == "" || apisong.Source["320"] == "" || apisong.Source["lossless"] == "":
		panic("Source has some values")
	}
	apisong.LinkDownload = map[String]String{"128": "http://api.mp3.zing.vn/api/mobile/download/song/kmxHyLnsdxHvbQgtLDJTvHkH", "lossless": "http://api.mp3.zing.vn/api/mobile/download/song/LmcnyZGsdxnbvQXyIqffroffTvnLn", "320": "http://api.mp3.zing.vn/api/mobile/download/song/LGcHyZHsBJnbDQCyVFntDGLn"}
	apisong.Source = map[String]String{"128": "http://api.mp3.zing.vn/api/mobile/source/song/knJHTLmsdcmbbQCyLFJyvHLG", "lossless": "http://api.mp3.zing.vn/api/mobile/source/song/ZmcGtkGsdJHFbWCyrPefrMfftvHkm", "320": "http://api.mp3.zing.vn/api/mobile/source/song/kmJHtkGNdxnbvQhtdbmTbmLH"}
	LogStruct(apisong)
	// Output:
	// Id : 1073802256
	// Key : "ZWZAOC90"
	// Title : "美少女战士/ Sailor Moon"
	// ArtistIds : "136,4781"
	// Artists : "Châu Huệ Mẫn,Vương Hinh Bình"
	// AlbumId : 1073747508
	// Album : "’94 美的化身演唱会/ Incarnation of Beauty Live 1994 (CD1)"
	// AuthorId : 0
	// Authors : ""
	// GenreId : "4,33"
	// Zaloid : 0
	// Username : "mp3"
	// IsHit : 0
	// IsOfficial : 1
	// DownloadStatus : 0
	// Copyright : ""
	// Thumbnail : "avatars/f/3/f3ccdd27d2000e3f9255a7e3e2c48800_1291614343.jpg"
	// Plays : 2000
	// Link : "/bai-hat/Sailor-Moon-Chau-Hue-Man-Vuong-Hinh-Binh/ZWZAOC90.html"
	// Source : map[dna.String]dna.String{"128":"http://api.mp3.zing.vn/api/mobile/source/song/knJHTLmsdcmbbQCyLFJyvHLG", "lossless":"http://api.mp3.zing.vn/api/mobile/source/song/ZmcGtkGsdJHFbWCyrPefrMfftvHkm", "320":"http://api.mp3.zing.vn/api/mobile/source/song/kmJHtkGNdxnbvQhtdbmTbmLH"}
	// LinkDownload : map[dna.String]dna.String{"128":"http://api.mp3.zing.vn/api/mobile/download/song/kmxHyLnsdxHvbQgtLDJTvHkH", "lossless":"http://api.mp3.zing.vn/api/mobile/download/song/LmcnyZGsdxnbvQXyIqffroffTvnLn", "320":"http://api.mp3.zing.vn/api/mobile/download/song/LGcHyZHsBJnbDQCyVFntDGLn"}
	// AlbumCover : "covers/2/3/233d32ad129990d4c583c6db55ea5e17_1290438828.jpg"
	// Likes : 0
	// LikeThis : false
	// Favourites : 0
	// FavouritesThis : false
	// Comments : 0
	// Topics : "Hoa Ngữ"
	// Video : zi.APIVideo{Id:0, Title:"", ArtistIds:"", Artists:"", GenreId:"", Thumbnail:"", Duration:0, StatusId:0, Link:"", Source:map[dna.String]dna.String(nil), Plays:0, Likes:0, LikeThis:false, Favourites:0, FavouritesThis:false, Comments:0, Topics:"", Response:zi.APIResponse{MsgCode:0}}
	// Response : zi.APIResponse{MsgCode:1}

}

func ExampleGetAPISongLyric() {
	apiSongLyric, err := GetAPISongLyric(1381645456)
	PanicError(err)
	LogStruct(apiSongLyric)
	// Output:
	// Id : ""
	// Content : "飞身到天边为这世界一战\r\n红日在夜空天际出现\r\n抛出救生圈雾里舞我的剑\r\n邪道外魔星际飞闪\r\n周:你你你快跪下\r\n看我引弓千里箭\r\n汤:你你你快跪下\r\n勿要我放出了魔毯\r\n王:你你你快跪下\r\n勿要我手握天血剑\r\n你你你快跪下\r\n狂风扫落雷电\r\n美少女转身变\r\n已变成战士\r\n以爱凝聚力量救世人跳出生天\r\n身体套光圈合上两眼都见\r\n明亮像佛光天际初现\r\n虽诡计多端但美少女一变\r\n邪道外魔都企一边"
	// Mark : 0
	// Author : "buzzed"
	// Response : zi.APIResponse{MsgCode:1}

}

func ExampleGetAPIAlbum() {
	apiAlbum, err := GetAPIAlbum(1381684168)
	PanicError(err)
	if apiAlbum.Plays < 360 {
		panic("Plays has to be greater than 360")
	}
	if apiAlbum.Likes < 0 {
		panic("Likes has to be greater than or equal to 0")
	}
	if apiAlbum.Comments < 0 {
		panic("Comments has to be greater than or equal to 0")
	}
	apiAlbum.Plays = 400
	apiAlbum.Likes = 0
	apiAlbum.Comments = 0
	LogStruct(apiAlbum)
	// Output:
	// Id : 1381684168
	// Title : "Good Bye..."
	// ArtistIds : "37831"
	// Artists : "C.S.C→luv"
	// GenreId : "38,5"
	// Zaloid : 0
	// Username : ""
	// Coverart : "covers/8/4/84366884afe11cc37fd3e37166dcde0a_1374395207.jpg"
	// Description : "Good Bye... là album của nhóm nhạc doujin C.S.C→luv phát hành vào ngày 14/03/2010 tại lễ hội Reitaisai 7. Album bao gồm các ca khúc hòa âm từ nhạc của trò chơi Shooting nổi tiếng Touhou."
	// IsHit : 0
	// IsOfficial : 1
	// IsAlbum : 1
	// YearReleased : "2010"
	// StatusId : 1
	// Link : "/album/Good-Bye-C-S-C-luv/ZWZADOC8.html"
	// Plays : 400
	// Topics : "Pop / Dance, Nhật Bản"
	// Likes : 0
	// LikeThis : false
	// Comments : 0
	// Favourites : 0
	// FavouritesThis : 0
	// Response : zi.APIResponse{MsgCode:1}

}

func ExampleGetAPIVideo() {
	apiVideo, err := GetAPIVideo(1381585674)
	PanicError(err)
	if apiVideo.Source["480"] == "" {
		panic("Source has not to be empty")
	}
	apiVideo.Source = map[String]String{"480": "http://api.mp3.zing.vn/api/mobile/source/video/LncmtZnsBalvzNlTzxnTvHkH"}
	if apiVideo.Plays < 1032532 {
		panic("Plays has to be greater than 360")
	}
	if apiVideo.Likes < 82 {
		panic("Likes has to be greater than or equal to 82")
	}
	if apiVideo.Comments < 3 {
		panic("Comments has to be greater than or equal to 3")
	}
	apiVideo.Plays = 1032532
	apiVideo.Likes = 82
	apiVideo.Comments = 3
	LogStruct(apiVideo)
	// Output:
	// Id : 1073742474
	// Title : "Người Là Niềm Đau"
	// ArtistIds : "212"
	// Artists : "Lâm Hùng"
	// GenreId : "1,8"
	// Thumbnail : "thumb_video/d/e/deb452e41ec76fa05cc12710981a6380_1340686117.jpg"
	// Duration : 311
	// StatusId : 1
	// Link : "/video-clip/Nguoi-La-Niem-Dau-Lam-Hung/ZWZ9ZO0A.html"
	// Source : map[dna.String]dna.String{"480":"http://api.mp3.zing.vn/api/mobile/source/video/LncmtZnsBalvzNlTzxnTvHkH"}
	// Plays : 1032532
	// Likes : 82
	// LikeThis : false
	// Favourites : 0
	// FavouritesThis : false
	// Comments : 3
	// Topics : "Việt Nam, Nhạc Trẻ"
	// Response : zi.APIResponse{MsgCode:1}

}

func ExampleGetAPIVideoLyric() {
	apiVideoLyric, err := GetAPIVideoLyric(1381585674)
	PanicError(err)
	LogStruct(apiVideoLyric)
	// Output:
	// Id : ""
	// Content : "Em yêu ơi sao quên đi từng đêm mưa ướt cuộc tình\r\nMặn nồng ân ái đôi ta cùng say quên đêm giá băng\r\nNay em đã ra đi vùi chôn dĩ vãng ngày nào\r\nHỡi em yêu, người là niềm đau\r\nEm dối gian anh, em dối gian anh\r\nLời yêu nồng say ngày nào\r\nĐã quên lãng sao hỡi người\r\nEm đã xa rồi, em mãi xa rồi\r\nĐể từng đêm ngồi chờ mong người\r\nMà người chẳng quay về\r\nKhi anh đây xa em cõi lòng giá băng\r\nTrong đêm thâu anh mãi gọi tên em, hỡi người\r\nBao nhiêu đêm lang thang, mình đã vùi chôn nỗi buồn\r\nTìm vào quên lãng bằng muôn xót xa\r\nEm yêu ơi sao quên đi từng đêm mưa ướt cuộc tình\r\nMặn nồng ân ái đôi ta cùng say quên đêm giá băng\r\nNay em đã ra đi vùi chôn dĩ vãng ngày nào\r\nHỡi em yêu, người là niềm đau"
	// Mark : 0
	// StatusId : 0
	// Author : "freshyidol"
	// DateCreated : 0
	// Response : zi.APIResponse{MsgCode:1}

}

func ExampleGetAPIArtist() {
	apiArtist, err := GetAPIArtist(828)
	PanicError(err)
	LogStruct(apiArtist)
	// Output:
	// Id : 828
	// Name : "Quang Lê"
	// Alias : ""
	// Birthname : "Leon Quang Lê"
	// Birthday : "24/01/1981"
	// Sex : 1
	// GenreId : "1,11,13"
	// Avatar : "avatars/9/6/96c7f8568cdc943997aace39708bf7b6_1376539870.jpg"
	// Coverart : "cover_artist/9/9/9920ce8b6c7eb43328383041acb58e76_1376539928.jpg"
	// Coverart2 : ""
	// ZmeAcc : ""
	// Role : "1"
	// Website : ""
	// Biography : "Quang Lê sinh ra tại Huế, trong gia đình gồm 6 anh em và một người chị nuôi, Quang Lê là con thứ 3 trong gia đình.\r\nĐầu những năm 1990, Quang Lê theo gia đình sang định cư tại bang Missouri, Mỹ.\r\nHiện nay Quang Lê sống cùng gia đình ở Los Angeles, nhưng vẫn thường xuyên về Việt Nam biểu diễn.\r\n\r\nSự nghiệp:\r\n\r\nSay mê ca hát từ nhỏ và niềm say mê đó đã cho Quang Lê những cơ hội để đi đến con đường ca hát ngày hôm nay. Có sẵn chất giọng Huế ngọt ngào, Quang Lê lại được cha mẹ cho theo học nhạc từ năm lớp 9 đến năm thứ 2 của đại học khi gia đình chuyển sang sống ở California . Anh từng đoạt huy chương bạc trong một cuộc thi tài năng trẻ tổ chức tại California. Thời gian đầu, Quang Lê chỉ xuất hiện trong những sinh hoạt của cộng đồng địa phương, mãi đến năm 2000 mới chính thức theo nghiệp ca hát. Nhưng cũng phải gần 2 năm sau, Quang Lê mới tạo được chỗ đứng trên sân khấu ca nhạc của cộng đồng người Việt ở Mỹ. Và từ đó, Quang Lê liên tục nhận được những lời mời biểu diễn ở Mỹ, cũng như ở Canada, Úc...\r\nLà một ca sĩ trẻ, cùng gia đình định cư ở Mỹ từ năm 10 tuổi, Quang Lê đã chọn và biểu diễn thành công dòng nhạc quê hương. Nhạc sĩ lão thành Châu Kỳ cũng từng khen Quang Lê là ca sĩ trẻ diễn đạt thành công nhất những tác phẩm của ông…\r\nQuang Lê rất hạnh phúc và anh xem lời khen tặng đó là sự khích lệ rất lớn để anh cố gắng nhiều hơn nữa trong việc diễn đạt những bài hát của nhạc sĩ Châu Kỳ cũng như những bài hát về tình yêu quê hương đất nước. 25 tuổi, được xếp vào số những ca sĩ trẻ thành công, nhưng Quang Lê luôn khiêm tốn cho rằng thành công thường đi chung với sự may mắn, và điều may mắn của anh là được lớn lên trong tiếng đàn của cha, giọng hát của mẹ.\r\nTiếng hát, tiếng đàn của cha mẹ anh quyện lấy nhau, như một sợi dây vô hình kết nối mọi người trong gia đình lại với nhau. Những âm thanh ngọt ngào đó chính là dòng nhạc quê hương mà Quang Lê trình diễn ngày hôm nay. Quang Lê cho biết: \"Mặc dù sống ở Mỹ đã lâu nhưng hình ảnh quê hương không bao giờ phai mờ trong tâm trí Quang Lê, nên mỗi khi hát những nhạc phẩm quê hương, những hình ảnh đó lại như hiện ra trước mắt\". Có lẽ vì thế mà giọng hát của Quang Lê như phảng phất cái không khí êm đềm của thành phố Huế.\r\nQuang Lê là con thứ 3 trong gia đình gồm 6 anh em và một người chị nuôi. Từ nhỏ, Quang Lê thường được người chung quanh khen là có triển vọng. Cậu bé chẳng hiểu \"có triển vọng\" là gì, chỉ biết là mình rất thích hát, và thích được cất tiếng hát trước người thân, để được khen ngợi và cổ vũ.\r\nĐầu những năm 1990, Quang Lê theo gia đình sang định cư tại bang Missouri, Mỹ. Một hôm, nhân có buổi lễ được tổ chức ở ngôi chùa gần nhà, một người quen của gia đình đã đưa Quang Lê đến để giúp vui cho chương trình sinh hoạt của chùa, và anh đã nhận được sự đón nhận nhiệt tình của khán giả. Quang Lê nhớ lại, \"người nghe không chỉ vỗ tay hoan hô mà còn thưởng tiền nữa\". Đối với một đứa trẻ 10 tuổi, thì đó quả là một niềm hạnh phúc lớn lao, khi nghĩ rằng niềm đam mê của mình lại còn có thể kiếm tiền giúp đỡ gia đình.\r\nQuan điểm của Quang Lê là khi dự định làm một việc gì thì hãy cố gắng hết mình để đạt được những điều mà mình mơ ước. Quang Lê cho biết anh toàn tâm toàn ý với dòng nhạc quê hương trữ tình mà anh đã chọn lựa và được đón nhận, nhưng anh tiết lộ là những lúc đi hát vũ trường, vì muốn thay đổi và để hòa đồng với các bạn trẻ, anh cũng trình bày những ca khúc \"Techno\" và cũng nhảy nhuyễn không kém gì vũ đoàn minh họa.\r\n\r\nAlbum:\r\n\r\nSương trắng miền quê ngoại (2003)\r\nXin gọi nhau là cố nhân (2004)\r\nHuế đêm trăng (2004)\r\nKẻ ở miền xa (2004)\r\n7000 đêm góp Lại (2005)\r\nĐập vỡ cây đàn (2007)\r\nHai quê (2008)\r\nTương tư nàng ca sĩ (2009)\r\nĐôi mắt người xưa (2010)\r\nPhải lòng con gái bến tre (2011)\r\nKhông phải tại chúng mình (2012)"
	// Publisher : "Ca sĩ Tự Do"
	// Country : "Việt Nam"
	// IsOfficial : 1
	// YearActive : "2000"
	// StatusId : 1
	// DateCreated : 0
	// Link : "/nghe-si/Quang-Le"
	// Topics : "Việt Nam, Nhạc Trữ Tình"
	// Response : zi.APIResponse{MsgCode:1}

}

func ExampleGetAPITV() {
	apiTV, err := GetAPITV(51173 + ID_DIFFERENCE)
	PanicError(err)
	if apiTV.Plays < 9933 {
		// Log(apiTV)
		panic("Plays has to be greater than 9933, GOT:" + apiTV.Plays.ToString().String())
	}
	if apiTV.Likes < 71 {
		panic("Likes has to be greater than or equal to 71")
	}
	if apiTV.Comments < 6 {
		// 	// LogStruct(apiTV)
		panic("Comments has to be greater than or equal to 6, GOT:" + apiTV.Comments.ToString().String())
	}
	if apiTV.Rating < 0 {
		// panic("Rating has to be greater than or equal to 0")
	}
	if apiTV.FileUrl == "" {
		// panic("File URL has to be valid")
	}
	if apiTV.OtherUrl["Video3GP"] == "" || apiTV.OtherUrl["Video480"] == "" || apiTV.OtherUrl["Video720"] == "" {
		panic("OtherUrl has to be valid")
	}
	apiTV.Plays = 9933
	apiTV.Likes = 71
	apiTV.Comments = 6
	apiTV.Rating = 9.775862068965518
	apiTV.FileUrl = "stream6.tv.zdn.vn/streaming/ed283ead88766c5a8ed4a82ee4abf2f4/52a2af3c/2013/1125/91/2eddbfaa80233df649d9c6f2dcf2c214.mp4?format=f360&device=ios"
	apiTV.OtherUrl = map[String]String{"Video3GP": "stream.m.tv.zdn.vn/tv/8da2d224e6ac5321f4daece3810ff137/52a2af3c/Video3GP/2013/1125/91/a05a3df4e113ce0e19ca99ecd6ae59c4.3gp?format=f3gp&device=ios", "Video720": "stream6.tv.zdn.vn/streaming/d2078fd68f65c0b56f195c181d108029/52a2af3c/Video720/2013/1125/91/f692b3dfac156aa86b51e916154a57a9.mp4?format=f720&device=ios", "Video480": "stream6.tv.zdn.vn/streaming/9133791e61a7c0459899018a22833cc8/52a2af3c/Video480/2013/1125/91/0cdcb216227b9d1898d48ec4deb85b51.mp4?format=f480&device=ios"}
	LogStruct(apiTV)
	// Output:
	// Id : 51173
	// Title : "The End of Twerk"
	// Fullname : "Glee - Season 5 - Tập 5 - The End of Twerk"
	// Episode : 5
	// DateReleased : "10/02/2014"
	// Duration : 2552
	// Thumbnail : "2013/1201/d1/1cb54a6001ccc301682f73e48e36f92c_1386256771.jpg"
	// FileUrl : "stream6.tv.zdn.vn/streaming/ed283ead88766c5a8ed4a82ee4abf2f4/52a2af3c/2013/1125/91/2eddbfaa80233df649d9c6f2dcf2c214.mp4?format=f360&device=ios"
	// OtherUrl : map[dna.String]dna.String{"Video3GP":"stream.m.tv.zdn.vn/tv/8da2d224e6ac5321f4daece3810ff137/52a2af3c/Video3GP/2013/1125/91/a05a3df4e113ce0e19ca99ecd6ae59c4.3gp?format=f3gp&device=ios", "Video720":"stream6.tv.zdn.vn/streaming/d2078fd68f65c0b56f195c181d108029/52a2af3c/Video720/2013/1125/91/f692b3dfac156aa86b51e916154a57a9.mp4?format=f720&device=ios", "Video480":"stream6.tv.zdn.vn/streaming/9133791e61a7c0459899018a22833cc8/52a2af3c/Video480/2013/1125/91/0cdcb216227b9d1898d48ec4deb85b51.mp4?format=f480&device=ios"}
	// LinkUrl : "http://tv.zing.vn/video/glee---season-5-tap-5-the-end-of-twerk/IWZAI86Z.html"
	// ProgramId : 1910
	// ProgramName : "Glee - Season 5"
	// ProgramThumbnail : "channel/9/9/9954115f583d4d40ec6428061a97fb12_1382526168.jpg"
	// ProgramGenres : []zi.APIProgramGenre{zi.APIProgramGenre{Id:82, Name:"Phim Truyền Hình"}}
	// Plays : 9933
	// Comments : 6
	// Likes : 71
	// Rating : 9.775862068965518
	// SubTitle : ""
	// Tracking : ""
	// Signature : "582e5fa51e234e052b0570eac9997539"

}
