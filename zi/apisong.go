package zi

import (
	. "dna"
	"dna/http"
	"encoding/json"
	"time"
)

//GetAPISong fetchs a song from API url. An url pattern is:
//http://api.mp3.zing.vn/api/mobile/song/getsonginfo?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":"ZW67FWWF"}
//
//The following result:
//	{
//	  "song_id": 1074700719,
//	  "song_id_encode": "ZW67FWWF",
//	  "title": "Bức Tranh Từ Nước Mắt",
//	  "artist_id": "13705",
//	  "artist": "Mr. Siro",
//	  "album_id": 1073845392,
//	  "album": "Bức Tranh Từ Nước Mắt",
//	  "composer_id": 303,
//	  "composer": "Mr Siro",
//	  "genre_id": "1,8",
//	  "zaloid": 0,
//	  "username": "mp3",
//	  "is_hit": 1,
//	  "is_official": 1,
//	  "download_status": 1,
//	  "copyright": "",
//	  "thumbnail": "avatars/7/5/7515ed78bf8f72a08e952708cd614db0_1385205599.jpg",
//	  "total_play": 312392,
//	  "link": "/bai-hat/Buc-Tranh-Tu-Nuoc-Mat-Mr-Siro/ZW67FWWF.html",
//	  "source": {
//	    "128": "http://api.mp3.zing.vn/api/mobile/source/song/LmJnykGNlNmnNkuTZvctbGZm",
//	    "320": "http://api.mp3.zing.vn/api/mobile/source/song/LHJHTLnNANHmaLitBbGyDGLH",
//	    "lossless": "http://api.mp3.zing.vn/api/mobile/source/song/LnJmyLmazamHskRtIqefrYffTbGZG"
//	  },
//	  "link_download": {
//	    "128": "http://api.mp3.zing.vn/api/mobile/download/song/LHJntLHNlanGNLRtLFxtvGZH",
//	    "320": "http://api.mp3.zing.vn/api/mobile/download/song/LmcHTZmszNmGNLuTdbGtDnLH",
//	    "lossless": "http://api.mp3.zing.vn/api/mobile/download/song/ZmxGtZnsAsGGskuTUPeKIMfKtvGLG"
//	  },
//	  "album_cover": null,
//	  "likes": 31216,
//	  "like_this": false,
//	  "favourites": 0,
//	  "favourite_this": false,
//	  "comments": 940,
//	  "genre_name": "Việt Nam, Nhạc Trẻ",
//	  "video": {
//	    "video_id": 1074700719,
//	    "title": "Bức Tranh Từ Nước Mắt",
//	    "artist": "Mr. Siro",
//	    "thumbnail": "thumb_video/2/b/2b01f358a38b70c2761eb2913daec382_1379306977.jpg",
//	    "duration": 340
//	  },
//	  "response": {
//	    "msgCode": 1
//	  }
//	}
func GetAPISong(id Int) (*APISong, error) {
	var apisong = new(APISong)
	apisong.Id = id
	apisong.Key = GetKey(id)
	link := `http://api.mp3.zing.vn/api/mobile/song/getsonginfo?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":"` + apisong.Key + `"}`
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apisong)
		return apisong, nil
	} else {
		return nil, err
	}
}

//GetAPISongLyric fetchs a song lyric from API url. An url pattern is:
// http://api.mp3.zing.vn/api/mobile/song/getlyrics?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":"ZW67FWWF"}
//
// The following result:
//	{
//	  "id": 1188945,
//	  "content": "Chuyện hai chúng ta bây giờ khác rồi\nThật lòng anh không muốn ai phải bối rối\nSợ em nhìn thấy nên anh đành phải lẳng lặng đứng xa\nChuyện tình thay đổi nên bây giờ trở thành người thứ ba\nTrách ai bây giờ, trách mình thôi.....\n\nĐK:\nNhìn em hạnh phúc bên ai càng làm anh tan nát lòng\nMới hiểu tại sao tình yêu người ta sợ khi cách xa\nĐiều anh lo lắng cứ vẫn luôn xảy ra\nNếu không đổi thay chẳng có ai sống được vì thiếu mất yêu thương.\n\nThời gian giết chết cuộc tình còn đau hơn giết chính mình\nTại sao mọi thứ xung quanh vẫn thế chỉ lòng người thay đổi\nGiờ em chỉ là tất cả quá khứ anh phải cố xoá trong nước mắt\n\n[ Trong tình yêu, thuộc về ai không quan trọng\nMột giây mơ màng là đã mất nhau....]\n\nCàng nghĩ đến em, anh càng hối hận\nVì xa em nên mất em thật ngu ngốc\nGiờ tình anh như bức tranh bằng nước mắt không màu sắc\nNhẹ nhàng và trong suốt cho dù đau đớn vẫn lặng yên\nTrách ai bây giờ, trách mình thôi....\n\nĐK:\nNhìn em hạnh phúc bên ai càng làm anh tan nát lòng\nMới hiểu tại sao tình yêu người ta sợ khi cách xa\nĐiều anh lo lắng cứ vẫn luôn xảy ra\nNếu không đổi thay chẳng có ai sống được vì thiếu mất yêu thương.\n\nThời gian giết chết cuộc tình còn đau hơn giết chính mình\nTại sao mọi thứ xung quanh vẫn thế chỉ lòng người thay đổi\nGiờ em chỉ là tất cả quá khứ anh phải cố xoá trong nước mắt.\n\nNụ cười em vẫn như xưa mà lòng em sao khác rồi\nNỗi đau này chỉ mình anh nhận lấy vì anh đã sai\nGiờ anh phải cố giữ nước mắt đừng rơi\nBức tranh tình yêu của em từ lâu đã không hề có anh......\n\nTrong tình yêu, thuộc về ai không quan trọng, rồi cũng mất nhau…",
//	  "mark": 4664,
//	  "author": "o0cobemuaxuan0o",
//	  "response": {
//	    "msgCode": 1
//	  }
//	}
func GetAPISongLyric(id Int) (*APISongLyric, error) {
	var apisongLyric = new(APISongLyric)
	link := `http://api.mp3.zing.vn/api/mobile/song/getlyrics?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":"` + GetKey(id) + `"}`
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apisongLyric)
		return apisongLyric, nil
	} else {
		return nil, err
	}
}

//GetAPIAlbum fetchs an album from API url. An url pattern is:
//http://api.mp3.zing.vn/api/mobile/playlist/getplaylistinfo?key=fafd463e2131914934b73310aa34a23f&requestdata={"id":1073816610}
//
//The following result:
//	{
//	  "playlist_id": 1073816610,
//	  "title": "Thời Gian",
//	  "artist_id": "28768",
//	  "artist": "Stillrock",
//	  "genre_id": "1,10",
//	  "zaloid": 0,
//	  "username": "",
//	  "cover": "covers/3/e/3e2c54b351be6220c2afe114f2cc0b90_1356856078.jpg",
//	  "description": "Ý tưởng của album dựa trên những diễn biến cuộc sống đời thường của ban nhạc. Có bắt tay vào làm việc mà cụ thể là đầu tư cho một album rock trong điều kiện còn nhiều khó khăn quả thực lúc đó mới biết quí THỜI GIAN. Đam mê thứ tài sản vô giá mà mỗi rocker chân chính có được, ROCK gieo vào đầu óc những ý nghĩ tích cực: sự mạnh mẽ, sự táo bạo và cả sự nghiêm túc đàng hoàng. Quả thực, rock đã tiếp sức để mỗi người trẻ tự vượt qua những cám dỗ trong cuộc sống, những lũ nhện ma quái giăng tơ huyền ảo khắp nơi. Đó chính là TƠ NHỆN.  Niềm vui , những điều hạnh phúc vừa đến, lại phải đối mặt với khó khăn và nỗi đau. Trong văn chương nói rất nhiều cái gọi là cõi tạm, sinh ra là khóc là khổ. Vậy mới biết “Có nước mắt trôi trên cuộc đời, có sóng gió mênh mang bầu trời, hãy đứng vững bằng đôi chân không mềm yếu” (trích lời bài CHIẾN SĨ) . Phải đứng vững trên đôi chân mình, chân lí sống không bao giờ thay đổi và hãy nghe TẤT CẢ  để thấy mình trong đó, cuộc sống đôi khi vô thường, nhàm chán đến tận cùng, vậy mới biết lí trí chẳng vượt qua được sự tận cùng. Chưa hết, TRĂNG MỜ muốn đem tới cho rock fan một chút cảm thụ nhân văn mà tác giả Phan Hoàng Thái muốn gửi gắm. Quả đất là một vòng tròn, Mặt trăng là một vòng tròn, mỗi con người lần lượt  phải đi qua cái vòng tròn lẫn quẩn đó… cái chết. Để rồi khi đối mặt với nó có nghĩa là chúng ta thấy sự sống quí giá chừng nào. RUNG xô đẩy người nghe từ đất liền ra biển khơi, ở đâu cũng vậy cũng có hiền nhân và quĩ dữ, cũng có nguy nan và bình s",
//	  "is_hit": 1,
//	  "is_official": 1,
//	  "is_album": 1,
//	  "year": "2008",
//	  "status_id": 1,
//	  "link": "/album/Thoi-Gian-Stillrock/ZWZA7UAW.html",
//	  "total_play": 285289,
//	  "genre_name": "Việt Nam, Rock Việt",
//	  "likes": 102,
//	  "like_this": false,
//	  "comments": 0,
//	  "favourites": 0,
//	  "favourite_this": false,
//	  "response": {
//	    "msgCode": 1
//	  }
//	}
func GetAPIAlbum(id Int) (*APIAlbum, error) {
	var apialbum = new(APIAlbum)
	apialbum.Id = id
	link := `http://api.mp3.zing.vn/api/mobile/playlist/getplaylistinfo?key=fafd463e2131914934b73310aa34a23f&requestdata={"id":` + (apialbum.Id - 307843200).ToString() + `}`
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apialbum)
		return apialbum, nil
	} else {
		return nil, err
	}
}

//GetAPIVideo fetchs a video from API url. An url pattern is:
//http://api.mp3.zing.vn/api/mobile/video/getvideoinfo?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":1074729245}
//
//The following result:
//	{
//	  "video_id": 1074729245,
//	  "title": "Xin Anh Đừng Đến",
//	  "artist_id": "465",
//	  "artist": "Bảo Thy",
//	  "genre_id": "1,8,66",
//	  "thumbnail": "thumb_video/d/c/dcacff355635deedf62fd80de34f2346_1380622208.jpg",
//	  "duration": 307,
//	  "status_id": 1,
//	  "link": "/video-clip/Xin-Anh-Dung-Den-Bao-Thy/ZW686I9D.html",
//	  "source": {
//	    "240": "http://api.mp3.zing.vn/api/mobile/source/video/LncGtLGazaDuFAQyFlHTFGLn",
//	    "360": "http://api.mp3.zing.vn/api/mobile/source/video/LHJmTkHaAaFiDlQtdCnyDnLG",
//	    "480": "http://api.mp3.zing.vn/api/mobile/source/video/LHcHTknNlNFuFzpylJnyFHZm",
//	    "720": "http://api.mp3.zing.vn/api/mobile/source/video/ZHJHyLHNSNFRDzWyNbmTDmLm",
//	    "1080": "http://api.mp3.zing.vn/api/mobile/source/video/ZnxmTkGalNvuvzQtkGcntvHLn"
//	  },
//	  "total_play": 1684238,
//	  "likes": 7663,
//	  "like_this": false,
//	  "favourites": 0,
//	  "favourite_this": false,
//	  "comments": 246,
//	  "genre_name": "Việt Nam, Nhạc Trẻ, Nhạc Dance",
//	  "response": {
//	    "msgCode": 1
//	  }
//	}
func GetAPIVideo(id Int) (*APIVideo, error) {
	var apivideo = new(APIVideo)
	apivideo.Id = id
	link := `http://api.mp3.zing.vn/api/mobile/video/getvideoinfo?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":` + (apivideo.Id - 307843200).ToString() + `}`
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apivideo)
		return apivideo, nil
	} else {
		return nil, err
	}
}

//GetAPIVideoLyric fetchs a video from API url. An url pattern is:
//http://api.mp3.zing.vn/api/mobile/video/getlyrics?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":1074729245}
//
//The following result:
//	{
//	  "id": "1207338",
//	  "content": "Xin Anh Đừng Đến\n(Bảo Thy on the mind)\nÁnh sáng trắng xóa bỗng thấy em trong cơn mơ\nMột mình em như kề bên không còn ai\nBờ vai em đang run run khi không gian đang chìm sâu\nTrong ánh mắt em giờ đây một màu u tối.\n\nNhớ ánh mắt ấy tiếng nói ấy sao giờ đây\nChỉ còn em đang ngồi bên những niềm đau\nTự dặn mình hãy cố xóa hết những giấc mơ khi bên anh\nEm còn mơ, em còn mơ đầy thương nhớ.\n\nEm sẽ xóa hết những phút giây ta yêu thương\nĐể về sau gặp lại nhau em sẽ vơi đi nỗi đau\nHãy để nỗi nhớ khi mà em đang bơ vơ\nKhông cần anh không cần thương nhớ\nNhững ngày còn vụn vỡ\n\n(em sẽ cố quên)\n(tình yêu đó)\n\n[ĐK:]\nXin anh hãy nói, đôi ta chia tay\nCho con tim em không như bao ngày\nXin anh đừng đến trong cơn mơ\nĐể từng ngày qua em thôi trông mong chờ\nLeave me alone!\n\nHey Boy ! Shake your body x3\nHey Girl ! Let me Put your hands up in the air\n\n[ĐK:]\nXin anh hãy nói, đôi ta chia tay\nCho con tim em không như bao ngày\nXin anh đừng đến trong cơn mơ\nĐể từng ngày qua em thôi trông mong chờ\nLeave me alone!",
//	  "mark": 593,
//	  "status_id": 0,
//	  "author": "pynyuno",
//	  "created_date": 1380727569,
//	  "response": {
//	    "msgCode": 1
//	  }
//	}
func GetAPIVideoLyric(id Int) (*APIVideoLyric, error) {
	var apiVideoLyric = new(APIVideoLyric)
	link := `http://api.mp3.zing.vn/api/mobile/video/getlyrics?keycode=fafd463e2131914934b73310aa34a23f&requestdata={"id":` + (id - 307843200).ToString() + `}`
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apiVideoLyric)
		return apiVideoLyric, nil
	} else {
		return nil, err
	}
}

//GetAPIArtist fetchs an artist from API url. An url pattern is:
//http://api.mp3.zing.vn/api/mobile/artist/getartistinfo?key=fafd463e2131914934b73310aa34a23f&requestdata={"id":828}
//
//The following result:
//	{
//	  "artist_id": 828,
//	  "name": "Quang Lê",
//	  "alias": "",
//	  "birthname": "Leon Quang Lê",
//	  "birthday": "24/01/1981",
//	  "sex": 1,
//	  "genre_id": "1,11,13",
//	  "avatar": "avatars/9/6/96c7f8568cdc943997aace39708bf7b6_1376539870.jpg",
//	  "cover": "cover_artist/9/9/9920ce8b6c7eb43328383041acb58e76_1376539928.jpg",
//	  "cover2": "",
//	  "zme_acc": "",
//	  "role": "1",
//	  "website": "",
//	  "biography": "Quang Lê sinh ra tại Huế, trong gia đình gồm 6 anh em và một người chị nuôi, Quang Lê là con thứ 3 trong gia đình.\r\nĐầu những năm 1990, Quang Lê theo gia đình sang định cư tại bang Missouri, Mỹ.\r\nHiện nay Quang Lê sống cùng gia đình ở Los Angeles, nhưng vẫn thường xuyên về Việt Nam biểu diễn.\r\n\r\nSự nghiệp:\r\n\r\nSay mê ca hát từ nhỏ và niềm say mê đó đã cho Quang Lê những cơ hội để đi đến con đường ca hát ngày hôm nay. Có sẵn chất giọng Huế ngọt ngào, Quang Lê lại được cha mẹ cho theo học nhạc từ năm lớp 9 đến năm thứ 2 của đại học khi gia đình chuyển sang sống ở California . Anh từng đoạt huy chương bạc trong một cuộc thi tài năng trẻ tổ chức tại California. Thời gian đầu, Quang Lê chỉ xuất hiện trong những sinh hoạt của cộng đồng địa phương, mãi đến năm 2000 mới chính thức theo nghiệp ca hát. Nhưng cũng phải gần 2 năm sau, Quang Lê mới tạo được chỗ đứng trên sân khấu ca nhạc của cộng đồng người Việt ở Mỹ. Và từ đó, Quang Lê liên tục nhận được những lời mời biểu diễn ở Mỹ, cũng như ở Canada, Úc...\r\nLà một ca sĩ trẻ, cùng gia đình định cư ở Mỹ từ năm 10 tuổi, Quang Lê đã chọn và biểu diễn thành công dòng nhạc quê hương. Nhạc sĩ lão thành Châu Kỳ cũng từng khen Quang Lê là ca sĩ trẻ diễn đạt thành công nhất những tác phẩm của ông…\r\nQuang Lê rất hạnh phúc và anh xem lời khen tặng đó là sự khích lệ rất lớn để anh cố gắng nhiều hơn nữa trong việc diễn đạt những bài hát của nhạc sĩ Châu Kỳ cũng như những bài hát về tình yêu quê hương đất nước. 25 tuổi, được xếp vào số những ca sĩ trẻ thành công, nhưng Quang Lê luôn khiêm tốn cho rằng thành công thường đi chung với sự may mắn, và điều may mắn của anh là được lớn lên trong tiếng đàn của cha, giọng hát của mẹ.\r\nTiếng hát, tiếng đàn của cha mẹ anh quyện lấy nhau, như một sợi dây vô hình kết nối mọi người trong gia đình lại với nhau. Những âm thanh ngọt ngào đó chính là dòng nhạc quê hương mà Quang Lê trình diễn ngày hôm nay. Quang Lê cho biết: \"Mặc dù sống ở Mỹ đã lâu nhưng hình ảnh quê hương không bao giờ phai mờ trong tâm trí Quang Lê, nên mỗi khi hát những nhạc phẩm quê hương, những hình ảnh đó lại như hiện ra trước mắt\". Có lẽ vì thế mà giọng hát của Quang Lê như phảng phất cái không khí êm đềm của thành phố Huế.\r\nQuang Lê là con thứ 3 trong gia đình gồm 6 anh em và một người chị nuôi. Từ nhỏ, Quang Lê thường được người chung quanh khen là có triển vọng. Cậu bé chẳng hiểu \"có triển vọng\" là gì, chỉ biết là mình rất thích hát, và thích được cất tiếng hát trước người thân, để được khen ngợi và cổ vũ.\r\nĐầu những năm 1990, Quang Lê theo gia đình sang định cư tại bang Missouri, Mỹ. Một hôm, nhân có buổi lễ được tổ chức ở ngôi chùa gần nhà, một người quen của gia đình đã đưa Quang Lê đến để giúp vui cho chương trình sinh hoạt của chùa, và anh đã nhận được sự đón nhận nhiệt tình của khán giả. Quang Lê nhớ lại, \"người nghe không chỉ vỗ tay hoan hô mà còn thưởng tiền nữa\". Đối với một đứa trẻ 10 tuổi, thì đó quả là một niềm hạnh phúc lớn lao, khi nghĩ rằng niềm đam mê của mình lại còn có thể kiếm tiền giúp đỡ gia đình.\r\nQuan điểm của Quang Lê là khi dự định làm một việc gì thì hãy cố gắng hết mình để đạt được những điều mà mình mơ ước. Quang Lê cho biết anh toàn tâm toàn ý với dòng nhạc quê hương trữ tình mà anh đã chọn lựa và được đón nhận, nhưng anh tiết lộ là những lúc đi hát vũ trường, vì muốn thay đổi và để hòa đồng với các bạn trẻ, anh cũng trình bày những ca khúc \"Techno\" và cũng nhảy nhuyễn không kém gì vũ đoàn minh họa.\r\n\r\nAlbum:\r\n\r\nSương trắng miền quê ngoại (2003)\r\nXin gọi nhau là cố nhân (2004)\r\nHuế đêm trăng (2004)\r\nKẻ ở miền xa (2004)\r\n7000 đêm góp Lại (2005)\r\nĐập vỡ cây đàn (2007)\r\nHai quê (2008)\r\nTương tư nàng ca sĩ (2009)\r\nĐôi mắt người xưa (2010)\r\nPhải lòng con gái bến tre (2011)\r\nKhông phải tại chúng mình (2012)",
//	  "agency_name": "Ca sĩ Tự Do",
//	  "national_name": "Việt Nam",
//	  "is_official": 1,
//	  "year_active": "2000",
//	  "status_id": 1,
//	  "created_date": 0,
//	  "link": "/nghe-si/Quang-Le",
//	  "genre_name": "Việt Nam, Nhạc Trữ Tình",
//	  "response": {
//	    "msgCode": 1
//	  }
// }
func GetAPIArtist(id Int) (*APIArtist, error) {
	var apiArtist = new(APIArtist)
	apiArtist.Id = id
	link := `http://api.mp3.zing.vn/api/mobile/artist/getartistinfo?key=fafd463e2131914934b73310aa34a23f&requestdata={"id":` + apiArtist.Id.ToString() + `}`
	result, err := http.Get(link)
	if err == nil {
		data := &result.Data
		json.Unmarshal([]byte(*data), apiArtist)
		return apiArtist, nil
	} else {
		return nil, err
	}
}

//GetSongFromAPI gets a song from API. It does not get content from main site.
func GetSongFromAPI(id Int) (*Song, error) {
	var song *Song = NewSong()
	song.Id = id

	asong, err := GetAPISong(id)

	if err != nil {
		return nil, err
	} else {

		if asong.Video.Id > 0 {
			song.VideoId = asong.Video.Id + 307843200
		}
		if asong.AlbumId > 0 {
			song.AlbumId = asong.AlbumId + 307843200
		}
		song.ArtistId = asong.ArtistId.ToInt()
		song.IsHit = asong.IsHit
		song.IsOfficial = asong.IsOfficial
		song.DownloadStatus = asong.DownloadStatus
		song.Copyright = asong.Copyright
		flags := 0
		for key, val := range asong.Source {
			switch {
			case key == "128" && val != "":
				flags = flags | LBr128
			case key == "320" && val != "":
				flags = flags | LBr320
			case key == "lossless" && val != "":
				flags = flags | LBrLossless
			}
		}
		song.BitrateFlags = Int(flags)
		song.Likes = asong.Likes
		song.Comments = asong.Comments
		song.Thumbnail = asong.Thumbnail
		song.Checktime = time.Now()
		return song, nil
	}

}
