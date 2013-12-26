/*
This package implements some functionalities regarding chiasenhac.com.

Given an example link such as http://chiasenhac.com/google-bot~1184919.html,
we can not be sure whether it is a song or a video link based on an id input.
Therefore, there is a general type representing Song and Video type callled SongVideo.

It fetches a link, decides what kind of link is. If it is a song link, it will be converted to Song type.
Otherwise, it will be converted to Video type. It implements item.Item interface.

There are 2 basic kinds of links used in this package to get a complete SongVideo instance. For example:

	http://chiasenhac.com/google-bot~1184919.html
	http://chiasenhac.com/google-bot~1184919_download.html

The first one is supposed to get basic info (Artists, Authors...) The second is to find all "url formats".
"The formats" of a link is just a json-typed string and it depends on what type the link is (song or video link).
A song link returns SongUrl formats. And a video link returns VideoUrl formats.

Examples:

Getting a song or a video

	item, err := csn.GetSongVideo(1186398)
	if err != nil {
		dna.Log(err.Error())
	} else {
		switch item.(type) {
		case csn.Song:
			dna.LogStruct(item.(csn.Song))
		case *csn.Song:
			dna.LogStruct(item.(*csn.Song))
		case csn.Video:
			dna.LogStruct(item.(csn.Video))
		case *csn.Video:
			dna.LogStruct(item.(*csn.Video))
		default:
			panic("no type found")
		}
	}

Alternative methods of getting a song or a video using a general type,
initialzing, fetching and saving it to DB

	db, err := sqlpg.Connect(sqlpg.NewSQLConfig("./config/app.ini"))
	dna.PanicError(err)
	sv := csn.NewSongVideo()
	sv.Init(544375)
	err = sv.Fetch()
	if err != nil {
		panic(err.Error())
	}
	err = sv.Save(db)
	if err != nil {
		panic(err.Error())
	}
	db.Close()

*/
package csn
