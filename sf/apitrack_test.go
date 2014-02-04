package sf

import (
	"dna"
)

func ExampleAPISongFreaksTrack() {
	apitrack, err := GetSongFreaksTrack(4715354)
	if err != nil {
		dna.Log(err.Error())
	} else {
		apitrack.Response.RenderTime = 0
		dna.Log("Id:", apitrack.Id)
		dna.Log("XMLName:", apitrack.XMLName)
		dna.Log("Response:", apitrack.Response)
		dna.Log("Track:", apitrack.Track)
		if len(apitrack.Videos) < 5 {
			panic("No videos length")
		}
		// dna.Log("Videos:", apitrack.Videos)
		dna.Log("Comments:", apitrack.Comments)
	}
	// Output:
	// Id: 4715354
	// XMLName: xml.Name{Space:"", Local:"songfreaks"}
	// Response: sf.APIResponse{Code:101, RenderTime:0, Value:"SUCCESS: LICENSE, LYRICS"}
	// Track: sf.APITrack{Id:4715354, TrackGroupId:42143, UrlSlug:"dont-dream-its-over-lyrics-crowded-house-1", AMG:235384, IsInstrumental:false, Viewable:true, Duration:"", LyricId:3908215, HasLrc:false, TrackNumber:3, DiscNumber:1, Title:"Don't Dream It's Over", Rating:sf.APIRating{AverageRating:0, UserRating:0, TotalRatings:0}, Album:sf.APITrackAlbum{Id:6314, AMG:210763, UrlSlug:"pineapple-head-album-crowded-house", Year:1993, Coverart:"http://www.lyricfind.com/images/not_available_cov75.jpg", CoverartLarge:"http://www.lyricfind.com/images/not_available_cov200.jpg", Title:"Pineapple Head", Artist:sf.APIArtist{Id:3301, AMG:3998, UrlSlug:"crowded-house", Image:"http://www.lyricfind.com/images/amg/pic200/drp200/p212/p21233k4k6b.jpg", Genres:"Rock", Name:"Crowded House", Link:"http://www.songfreaks.com/crowded-house", Rating:sf.APIRating{AverageRating:0, UserRating:0, TotalRatings:0}, Bio:sf.APIArtistBio{Author:"", Value:""}}, Link:"http://www.songfreaks.com/pineapple-head-album-crowded-house"}, Artists:[]sf.APIArtist{sf.APIArtist{Id:3301, AMG:3998, UrlSlug:"crowded-house", Image:"http://www.lyricfind.com/images/amg/pic200/drp200/p212/p21233k4k6b.jpg", Genres:"Rock", Name:"Crowded House", Link:"http://www.songfreaks.com/crowded-house", Rating:sf.APIRating{AverageRating:0, UserRating:0, TotalRatings:0}, Bio:sf.APIArtistBio{Author:"", Value:""}}}, Lrc:sf.APILrc{Lines:[]sf.APILrcLine(nil)}, Link:"http://www.songfreaks.com/dont-dream-its-over-lyrics-crowded-house-1", Lyrics:"There is freedom within, there is freedom without\nTry to catch the deluge in a paper cup\nThere's a battle ahead, many battles are lost\nBut you'll never see the end of the road\nWhile you're traveling with me\n\n[Chorus]\nHey now, hey now\nDon't dream it's over\nHey now, hey now\nWhen the world comes in\nThey come, they come\nTo build a wall between us\nWe know they won't win\n\nNow I'm towing my car, there's a hole in the roof\nMy possessions are causing me suspicion but there's no proof\nIn the paper today tales of war and of waste\nBut you turn right over to the T.V. page\n\n[Chorus]\n\nNow I'm walking again to the beat of a drum\nAnd I'm counting the steps to the door of your heart\nOnly shadows ahead barely clearing the roof\nGet to know the feeling of liberation and release\n\n[Chorus]\n\nDon't let them win\n(Hey now, hey now, hey now, hey now)\nHey now, hey now\nDon't let them win\n(They come, they come)\nDon't let them win\n(Hey now, hey now, hey now, hey now)", Copyright:"Lyrics © Universal Music Publishing Group", Writer:"FINN, NEIL MULLANE", SubmittedLyric:false, FoxMobile:sf.APIFoxMobile{Available:false}}
	// Comments: sf.APIComments{NResult:0, NPages:0, Values:[]sf.APIComment(nil)}
}
