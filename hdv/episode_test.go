package hdv

import (
	"dna"
	"time"
)

func ExampleGetEpisode() {
	// Only to renew ACCESS_TOKEN_KEY
	_, err := GetMovie(5571)
	dna.PanicError(err)
	episode, err := GetEpisode(5571, 1)
	dna.PanicError(err)
	episode.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if episode.PlaylistM3u8.Match(`(?mis)^#EXTM3U\r\n#EXT-X-VERSION:3.+480_hdv_app.m3u8\r\n$`) == false {
		panic("PlaylistM3u8 is wrong")
	} else {
		episode.PlaylistM3u8 = ""
	}

	if episode.ViSrt.Match(`^77u/MQ0KMDA6MDA.+OgIHLhu5NpLg0KDQo=$`) == false {
		panic("Vietnamese subtitle is wrong")
	} else {
		episode.ViSrt = ""
	}

	if episode.EnSrt.Match(`^77u/MQ0KMDA6MDA6MDAsNTg.+MA0KKioqDQoNCg==$`) == false {
		panic("English subtitle is wrong")
	} else {
		episode.EnSrt = ""
	}

	if episode.EpisodeM3u8.Match(`(?mis)^#EXTM3U\r\n#EXT-X-VERSION:3.+#EXT-X-ENDLIST\r\n$`) == false {
		panic("M3u8 file is wrong")
	} else {
		episode.EpisodeM3u8 = ""
	}

	if episode.LinkPlayBackup.Match(`playlist_480_hdv_app.m3u8$`) == false {
		panic("Wrong link backup")
	} else {
		episode.LinkPlayBackup = "http://f-11.vn-hd.com/c52447879ebfd67168d82f16f1cb374a/onair_2014/Grimm_S03_HDTV_AC3/E001/playlist_480_hdv_app.m3u8"
	}

	if episode.Link.Match(`playlist_480_hdv_app.m3u8$`) == false {
		panic("Wrong link backup")
	} else {
		episode.Link = "http://f-11.vn-hd.com/c52447879ebfd67168d82f16f1cb374a/onair_2014/Grimm_S03_HDTV_AC3/E001/playlist_480_hdv_app.m3u8"
	}

	dna.LogStruct(episode)
	// Output:
	// MovieId : 5571
	// EpId : 1
	// Title : "Grimm (Season 3) - Tập 1"
	// LinkPlayBackup : "http://f-11.vn-hd.com/c52447879ebfd67168d82f16f1cb374a/onair_2014/Grimm_S03_HDTV_AC3/E001/playlist_480_hdv_app.m3u8"
	// Link : "http://f-11.vn-hd.com/c52447879ebfd67168d82f16f1cb374a/onair_2014/Grimm_S03_HDTV_AC3/E001/playlist_480_hdv_app.m3u8"
	// LinkPlayOther : "http://125.212.216.74/vod/_definst_/smil:mp4_02/store_01_2014/onair_2014/Grimm_S03_HDTV_AC3/E001/Grimm_S03_HDTV_AC3_E001.smil/Manifest"
	// SubtitleExt : dna.StringArray{"http://s.vn-hd.com/mp4_02/store_01_2014/onair_2014/Grimm_S03_HDTV_AC3/E001/Grimm_S03_HDTV_AC3_E001_VIE.srt", "http://s.vn-hd.com/mp4_02/store_01_2014/onair_2014/Grimm_S03_HDTV_AC3/E001/Grimm_S03_HDTV_AC3_E001_ENG.srt"}
	// SubtitleExtSe : dna.StringArray{"http://s.vn-hd.com/mp4_02/store_01_2014/onair_2014/Grimm_S03_HDTV_AC3/E001/Grimm_S03_HDTV_AC3_E001_VIE.srt", "http://s.vn-hd.com/mp4_02/store_01_2014/onair_2014/Grimm_S03_HDTV_AC3/E001/Grimm_S03_HDTV_AC3_E001_ENG.srt"}
	// Subtitle : dna.StringArray{"", "http://s.vn-hd.com/mp4_02/store_01_2014/onair_2014/Grimm_S03_HDTV_AC3/E001/Grimm_S03_HDTV_AC3_E001_ENG.srt"}
	// EpisodeId : 17
	// Audiodub : 0
	// Audio : 0
	// Season : "[{\"MovieID\":\"473\",\"Name\":\"Grimm (Season 1)\"},{\"MovieID\":\"2364\",\"Name\":\"Grimm (Season 2)\"}]"
	// PlaylistM3u8 : ""
	// ViSrt : ""
	// EnSrt : ""
	// EpisodeM3u8 : ""
	// Checktime : "2013-11-21 00:00:00"
}
