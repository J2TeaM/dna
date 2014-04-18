package csn

import (
	"dna"
)

// SongUrl defines a struct for format field of csnsongs
type SongUrl struct {
	Link    dna.String `json:"link"`
	Type    dna.String `json:"type"`
	Size    dna.Int    `json:"file_size"`
	Bitrate dna.String `json:"bitrate"`
}

// VideoUrl defines a struct for format field of csnsongs
type VideoUrl struct {
	Link       dna.String `json:"link"`
	Type       dna.String `json:"type"`
	Size       dna.Int    `json:"file_size"`
	Resolution dna.String `json:"resolution"`
}

func NewSongUrl() *SongUrl {
	su := new(SongUrl)
	su.Link = ""
	su.Type = ""
	su.Size = 0
	su.Bitrate = ""
	return su
}

func NewVideoUrl() *VideoUrl {
	su := new(VideoUrl)
	su.Link = ""
	su.Type = ""
	su.Size = 0
	su.Resolution = ""
	return su
}

func getType(str dna.String) dna.String {
	switch {
	case str.Match(`(?mis)MP3`) == true:
		return "mp3"
	case str.Match(`(?mis)M4A`) == true:
		return "m4a"
	case str.Match(`(?mis)MP4`) == true:
		return "mp4"
	case str.Match(`(?mis)FLAC`) == true:
		return "flac"
	case str.Match(`(?mis)FLV`) == true:
		return "flv"
	default:
		dna.Log("No type found at: " + str.String())
		return ""
	}
}

func getSongUrl(str, bitrate dna.String) *SongUrl {
	su := NewSongUrl()
	su.Bitrate = bitrate
	su.Type = getType(str)
	su.Link = str.GetTagAttributes("href").ReplaceWithRegexp(`(http.+/).+(\..+$)`, "${1}file-name${2}")
	size := str.FindAllString(`[0-9\.]+ MB`, -1)
	if size.Length() > 0 {
		su.Size = size[0].ParseBytesFormat() / 1000
	}
	return su

}

func getVideoUrl(str, resolution dna.String) *VideoUrl {
	su := NewVideoUrl()
	su.Resolution = resolution
	su.Type = getType(str)
	su.Link = str.GetTagAttributes("href").ReplaceWithRegexp(`(http.+/).+(\..+$)`, "${1}file-name${2}")
	size := str.FindAllString(`[0-9\.]+ MB`, -1)
	if size.Length() > 0 {
		su.Size = size[0].ParseBytesFormat() / 1000
	}
	return su

}

func getStringifiedSongUrls(urls dna.StringArray) dna.String {
	var baseLink = dna.String("")
	songUrls := []SongUrl{}
	urls.ForEach(func(val dna.String, idx dna.Int) {
		// dna.Log(val)
		// Finding bitrate
		switch {
		case val.Match(`128kbps`) == true:
			songUrl := getSongUrl(val, "128kbps")
			baseLink = songUrl.Link.ReplaceWithRegexp(`[0-9]+/file-name.+`, "")
			songUrls = append(songUrls, *songUrl)
		case val.Match(`320kbps`) == true:
			songUrl := getSongUrl(val, "320kbps")
			baseLink = songUrl.Link.ReplaceWithRegexp(`[0-9]+/file-name.+`, "")
			songUrls = append(songUrls, *songUrl)
		case val.Match(`32kbps`) == true:
			songUrl := getSongUrl(val, "32kbps")
			baseLink = songUrl.Link.ReplaceWithRegexp(`[0-9]+/file-name.+`, "")
			songUrls = append(songUrls, *songUrl)
		case val.Match(`500kbps`) == true:
			songUrl := getSongUrl(val, "500kbps")
			songUrl.Link = baseLink + "m4a/file-name.m4a"
			songUrls = append(songUrls, *songUrl)
		case val.Match(`Lossless`) == true:
			songUrl := getSongUrl(val, "Lossless")
			songUrl.Link = baseLink + "flac/file-name.flac"
			songUrls = append(songUrls, *songUrl)
		}
	})
	// http://data.chiasenhac.com/downloads/1184/2/1183017-cfc5f7df/flac/file-name.flac
	// replace the link 500kps and lossless with available link,  apply for registered user only
	// and reduce the link length
	var ret = dna.StringArray{}
	for _, songUrl := range songUrls {
		var br dna.String
		if songUrl.Bitrate == "Lossless" {
			br = "1411"
		} else {
			br = songUrl.Bitrate.Replace("kbps", "")
		}
		t := `"(` + songUrl.Link + "," + songUrl.Type + "," + songUrl.Size.ToString() + "," + br + `)"`
		ret.Push(t)
	}
	// dna.Log(`{` + ret.Join(",") + `}`)
	return `{` + ret.Join(",") + `}`

}

func getStringifiedVideoUrls(urls dna.StringArray) dna.String {
	var baseLink = dna.String("")
	videoUrls := []VideoUrl{}
	urls.ForEach(func(val dna.String, idx dna.Int) {
		// dna.Log(val)
		// Finding bitrate
		switch {
		case val.Match(`MV 360p`) == true:
			songUrl := getVideoUrl(val, "360p")
			baseLink = songUrl.Link.ReplaceWithRegexp(`[0-9]+/file-name.+`, "")
			videoUrls = append(videoUrls, *songUrl)
		case val.Match(`MV 480p`) == true:
			songUrl := getVideoUrl(val, "480p")
			baseLink = songUrl.Link.ReplaceWithRegexp(`[0-9]+/file-name.+`, "")
			videoUrls = append(videoUrls, *songUrl)
		case val.Match(`MV 180p`) == true:
			songUrl := getVideoUrl(val, "180p")
			baseLink = songUrl.Link.ReplaceWithRegexp(`[0-9]+/file-name.+`, "")
			videoUrls = append(videoUrls, *songUrl)
		case val.Match(`HD 720p`) == true:
			songUrl := getVideoUrl(val, "720p")
			songUrl.Link = baseLink + "m4a/file-name.mp4"
			videoUrls = append(videoUrls, *songUrl)
		case val.Match(`HD 1080p`) == true:
			songUrl := getVideoUrl(val, "1080p")
			songUrl.Link = baseLink + "flac/file-name.mp4"
			videoUrls = append(videoUrls, *songUrl)
		}
	})

	var ret = dna.StringArray{}
	for _, videoUrl := range videoUrls {
		t := `"(` + videoUrl.Link + "," + videoUrl.Type + "," + videoUrl.Size.ToString() + "," + videoUrl.Resolution.Replace("p", "") + `)"`
		ret.Push(t)
	}
	// dna.Log(`{` + ret.Join(",") + `}`)
	return `{` + ret.Join(",") + `}`
}

// GetFormats returns proper formatStr for a song or a video.
// If it is a song, IsSong will be set to true. Otherwise, it will set to false false.
func GetFormats(urls dna.StringArray) (formatStr dna.String, IsSong dna.Bool) {
	switch getType(urls.Join("")) {
	case "mp3", "m4a", "flac":
		formatStr, IsSong = getStringifiedSongUrls(urls), IS_SONG
		return formatStr, IsSong
	case "mp4", "flv":
		formatStr, IsSong = getStringifiedVideoUrls(urls), IS_VIDEO
		return formatStr, IsSong
	default:
		panic("Wrong type. Cannot indentify song or video")
	}
	return "", false
}
