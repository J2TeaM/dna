package nct

import (
	"dna"
)

// GetRelevantPortions returns relevant songs, albums or videos
func GetRelevantPortions(data *dna.String) {
	if EnableRelevantPortionsMode {
		var body *dna.String
		// This is is supposed to get relevant or similar songs'
		// var body *dna.String
		songPart := data.FindAllString(`(?mis)^.+BXH Bài Hát`, 1)
		if songPart.Length() > 0 {
			body = &songPart[0]
		} else {
			body = data
		}
		similarSongs := body.FindAllString(`(?mis)<div class="info_data">.+?0</span>`, -1)
		similarSongs.ForEach(func(val dna.String, idx dna.Int) {
			if val.Match(`/bai-hat/`) == true {
				keyArr := val.FindAllStringSubmatch(`\.([0-9a-zA-Z]+)\.html" class="name_song"`, 1)
				if len(keyArr) > 0 {
					RelevantSongs.Push(keyArr[0][1].Trim())
				}
			}
		})

		similarSongs = data.FindAllString(`(?mis)<div class="info_song">.+?0</span>`, -1)
		similarSongs.ForEach(func(val dna.String, idx dna.Int) {
			keyArr := val.FindAllStringSubmatch(`\.([0-9a-zA-Z]+)\.html" class="name_song"`, 1)
			if len(keyArr) > 0 {
				RelevantSongs.Push(keyArr[0][1].Trim())
			}
		})

		// }

		// This part to find similar albums

		albumPart := data.FindAllString(`(?mis)^.+BXH Playlist`, 1)
		if albumPart.Length() > 0 {
			body = &albumPart[0]
		} else {
			body = data
		}
		albumKeyArr := body.FindAllString(`<a href="http://www.nhaccuatui.com/playlist/.+html.+`, -1)
		albumKeys := dna.StringArray(albumKeyArr.Map(func(val dna.String, idx dna.Int) dna.String {
			keyArr := val.GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
			if len(keyArr) > 0 {
				return keyArr[0][1]
			} else {
				return ""
			}
		}).([]dna.String)).Unique().Filter(func(val dna.String, idx dna.Int) dna.Bool {
			return !val.Contains("-")
		})
		RelevantAlbums = RelevantAlbums.Concat(albumKeys)

		// This part to find similar videos
		videoPart := data.FindAllString(`(?mis)^.+BXH Video`, 1)
		if videoPart.Length() > 0 {
			body = &videoPart[0]
		} else {
			body = data
		}
		videoKeyArr := body.FindAllString(`<a href="http://www.nhaccuatui.com/video/.+html.+`, -1)
		videoKeys := dna.StringArray(videoKeyArr.Map(func(val dna.String, idx dna.Int) dna.String {
			keyArr := val.GetTagAttributes("href").FindAllStringSubmatch(`/.+\.(.+)\.html`, -1)
			if len(keyArr) > 0 {
				return keyArr[0][1]
			} else {
				return ""
			}
		}).([]dna.String)).Unique().Filter(func(val dna.String, idx dna.Int) dna.Bool {
			if val != "com/video/top-20" {
				return true
			} else {
				return false
			}
		})
		RelevantVideos = RelevantVideos.Concat(videoKeys)

		RelevantSongs = RelevantSongs.Unique()
		RelevantAlbums = RelevantAlbums.Unique()
		RelevantVideos = RelevantVideos.Unique()

	}

}
