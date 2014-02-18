/*
hdviet.com.

NOTICE: CHECK TOKEN_ACCESS_KEY EXPIRED

Mobile site: http://mmovie.hdviet.com/

Decompile APK to JAR

Decompile .class files: using d4j from http://secureteam.net/d4j/d4j.tar.gz into .java files

Note 3 .java files at: com.hdvietnam.android.a
 m.java // md5 algorithm
 n.java // md5 algorithm
 e.java => containing SECRET_KEY when call e.a()

The hash string = md5(str + SECRET_KEY)

All the source code saved in file : hdviet_java_decompiled.zip


In com.hdvietnam.android.a, look at file d.java. This file builds up urls for a movie, a series or a channel.

	Id : 4748
	Title : "Castle - Nhà Văn phá Án: Phần 6 - Tập 13/24"
	AnotherTitle : "Castle (Season 6)"
	ForeignTitle : "Castle"
	VnTitle : "Nhà Văn phá Án: Phần 6"
	Topics : dna.StringArray{"Hài", "Âu Mỹ", "Bí ẩn", "Tâm lý", "Tội phạm"}
	Actors : dna.StringArray{"Nathan Fillion", "Stana Katic", "Susan Sullivan", "Jon Huertas", "Molly C. Quinn", "Tamala Jones", "Seamus Dever", "Ruben Santiago-Hudson", "Penny Johnson"}
	Directors : dna.StringArray{}
	Countries : dna.StringArray{"Mỹ"}
	Description : "Richard “Rick” Castle là một nhà văn trinh thám nổi tiếng, bị cảnh sát New York điều tra về một vụ sát nhân hàng loạt được cho là copy từ những cuốn tiểu thuyết của anh. Cùng thời điểm đó, Rick đang gặp khó khăn và anh đã “giết” nhân vật chính của cuốn tiểu thuyết hình sự thành công nhất của mình-Derrick Storm. Trong thời gian bị điều tra, anh gặp nữ thám tử xinh đẹp Kate Beckett và có cảm hứng viết trở lại , anh quyết định dùng hình tượng nữ thám tử cho cuốn sách tiếp theo của mình..."
	YearReleased : 2013
	IMDBRating : dna.IntArray{810, 59956}
	Similars : dna.IntArray{}
	Thumbnail : "http://t.hdviet.com/thumbs/214x321/893a0d023c6a60168a3a860cc4b9a169.jpg"
	MaxResolution : 720
	IsSeries : true
	SeasonId : 6
	Seasons : dna.IntArray{2019, 2864, 2879, 2897, 3299}
	Epid : 0
	CurrentEps : 13
	MaxEp : 24
	---------
	Title : "Castle (Season 6) - Tập 2"
	LinkPlayBackup : "http://v-01.vn-hd.com/c52447879ebfd67168d82f16f1cb374a/10102013/Castle_S06_HDTV_AC3/E002/playlist_480_hdv_app.m3u8"
	Link : "http://v-01.vn-hd.com/c52447879ebfd67168d82f16f1cb374a/10102013/Castle_S06_HDTV_AC3/E002/playlist_480_hdv_app.m3u8"
	LinkPlayOther : "http://125.212.216.74/vod/_definst_/smil:mp4_02/store_10_2013/10102013/Castle_S06_HDTV_AC3/E002/Castle_S06_HDTV_AC3_E002.smil/Manifest"
	SubtitleExt : hdv.APISubtitleList{Vietnamese:hdv.APISubtitle{Label:"Việt", Source:"http://s.vn-hd.com/mp4_02/store_10_2013/10102013/Castle_S06_HDTV_AC3/E002/Castle_S06_HDTV_AC3_E002_VIE.srt"}, English:hdv.APISubtitle{Label:"Anh", Source:"http://s.vn-hd.com/mp4_02/store_10_2013/10102013/Castle_S06_HDTV_AC3/E002/Castle_S06_HDTV_AC3_E002_ENG.srt"}}
	SubtitleExtSe : hdv.APISubtitleList{Vietnamese:hdv.APISubtitle{Label:"Việt", Source:"http://s.vn-hd.com/mp4_02/store_10_2013/10102013/Castle_S06_HDTV_AC3/E002/Castle_S06_HDTV_AC3_E002_VIE.srt"}, English:hdv.APISubtitle{Label:"Anh", Source:"http://s.vn-hd.com/mp4_02/store_10_2013/10102013/Castle_S06_HDTV_AC3/E002/Castle_S06_HDTV_AC3_E002_ENG.srt"}}
	Subtitle : hdv.APISubtitleList{Vietnamese:hdv.APISubtitle{Label:"Việt", Source:""}, English:hdv.APISubtitle{Label:"Anh", Source:"http://s.vn-hd.com/mp4_02/store_10_2013/10102013/Castle_S06_HDTV_AC3/E002/Castle_S06_HDTV_AC3_E002_ENG.srt"}}
	Episode : "13"
	Audiodub : "0"
	Audio : "0"
	Season : []hdv.Season{hdv.Season{Id:"2019", Title:"Castle (Season 1)"}, hdv.Season{Id:"2864", Title:"Castle (Season 2)"}, hdv.Season{Id:"2879", Title:"Castle (Season 3)"}, hdv.Season{Id:"2897", Title:"Castle (Season 4)"}, hdv.Season{Id:"3299", Title:"Castle (Season 5)"}}
	Adver : dna.StringArray{}
*/
package hdv
