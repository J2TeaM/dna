package nct

import (
	"dna"
	"time"
)

func ExampleURLBuilder() {
	logURL := func(descStr, url dna.String) {
		dna.Log(descStr+dna.String(" ").Repeat(23-descStr.Length()), url)
	}
	urlb := NewURLBuilder()
	urlb.SetTimestamp(time.Date(2014, time.February, 4, 12, 12, 12, 12, time.UTC))
	logURL("ARTIST INFO:", urlb.GetArtistInfo(49674))
	logURL("CHECKVERSION:", urlb.GetCheckVersion())
	dna.Log("--------------------------------------------------------")
	logURL("LIST OF GENRES:", urlb.GetListOfGenres("song"))
	logURL("LIST OF TOPICS:", urlb.GetListOfTopics(1, 30))
	dna.Log("--------------------------------------------------------")
	logURL("PLAYLIST INFO:", urlb.GetPlaylistInfo(12255234))
	logURL("PLAYLISTS BY ARTIST:", urlb.GetPlaylistsByArtist(49674, 1, 30))
	logURL("PLAYLISTS BY GENRE:", urlb.GetPlaylistsByGenre(0, 1, 30))
	logURL("PLAYLISTS BY SEARCHING:", urlb.GetPlaylistsBySearching("tell", 1, 30))
	logURL("PLAYLISTS BY TOPIC:", urlb.GetPlaylistsByTopic(51))
	logURL("PLAYLISTS RELATED:", urlb.GetPlaylistsRelated(12336694))
	dna.Log("--------------------------------------------------------")
	logURL("SONG INFO:", urlb.GetSongInfo(2854574))
	logURL("SONG LYRIC:", urlb.GetSongLyric(2909727))
	logURL("SONGS BY ARTIST:", urlb.GetSongsByArtist(49674, 50, 30))
	logURL("SONGS BY GENRE:", urlb.GetSongsByGenre(5015, 34, 30))
	logURL("SONGS BY SEARCHING:", urlb.GetSongsBySearching("man", 1, 30))
	logURL("SONGS OF PLAYLIST", urlb.GetSongsOfPlaylist(12095591))
	dna.Log("--------------------------------------------------------")
	logURL("VIDEO INFO:", urlb.GetVideoInfo(2876055))
	logURL("VIDEOS BY ARTIST:", urlb.GetVideosByArist(49674, 1, 30))
	logURL("VIDEOS BY GENRE:", urlb.GetVideosByGenre(5142, 1, 30))
	logURL("VIDEOS BY SEARCHING:", urlb.GetVideosBySearching("man", 1, 30))
	logURL("VIDEOS RELATED:", urlb.GetVideosRelated(12336694))
	// Output:
	// ARTIST INFO:            http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-artist-detail&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&artistid=49674&time=1391515932000&token=a6ffbbb6da2b822f647829ff6215c210
	// CHECKVERSION:           http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=check-version&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&time=1391515932000&token=31842a81519047f5dd477b54608eb2d7
	// --------------------------------------------------------
	// LIST OF GENRES:         http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-list-genre&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&type=song&time=1391515932000&token=c1b31f50c1ea19250898a1a2e14ca60b
	// LIST OF TOPICS:         http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-list-topic&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&pageindex=1&pagesize=30&time=1391515932000&token=d7abea98a130db44bce09600f25cb316
	// --------------------------------------------------------
	// PLAYLIST INFO:          http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-playlist-info&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&playlistid=12255234&time=1391515932000&token=f1f0b8e4827086b37c3c3accee79017f
	// PLAYLISTS BY ARTIST:    http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-playlist-by-artist&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&artistid=49674&pageindex=1&pagesize=30&time=1391515932000&token=2be9e5d0c3fd1325d2a7ee2426922d5b
	// PLAYLISTS BY GENRE:     http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-playlist-by-genre&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&genreid=0&pageindex=1&pagesize=30&time=1391515932000&token=ec761c1865e50560e61c5036fca893ce
	// PLAYLISTS BY SEARCHING: http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=search-playlist&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&keyword=tell&pageindex=1&pagesize=30&time=1391515932000&token=5370cbfb8040183ad5570b999c813c83
	// PLAYLISTS BY TOPIC:     http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-playlist-by-topic&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&topicid=51&time=1391515932000&token=ff8f012db0d705b3cdf2698ec16536a7
	// PLAYLISTS RELATED:      http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-playlist-related&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&playlistid=12336694&time=1391515932000&token=5c8efd0df7e283a58c28b943b85e1262
	// --------------------------------------------------------
	// SONG INFO:              http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-song-info&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&songid=2854574&time=1391515932000&token=b1db67e99010e541ea910e997350a75d
	// SONG LYRIC:             http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-lyric&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&songid=2909727&time=1391515932000&token=c752ab44e4f60177474b3ff554c380f3
	// SONGS BY ARTIST:        http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-song-by-artist&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&artistid=49674&pageindex=50&pagesize=30&time=1391515932000&token=cb00833e14a158c8eae098be2834e166
	// SONGS BY GENRE:         http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-song-by-genre&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&genreid=5015&pageindex=34&pagesize=30&time=1391515932000&token=21bd96e68c4179d0459f5692e0144ca0
	// SONGS BY SEARCHING:     http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=search-song&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&keyword=man&pageindex=1&pagesize=30&time=1391515932000&token=3ba6c692cd09ef54ee196bc8e3798078
	// SONGS OF PLAYLIST       http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-song-by-playlistid&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&playlistid=12095591&time=1391515932000&token=9e461944fa434c9e197b5bae1ad9cc3f
	// --------------------------------------------------------
	// VIDEO INFO:             http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-video-detail&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&videoid=2876055&time=1391515932000&token=ad13b3cee157e1ab1dfa337ec47ab20e
	// VIDEOS BY ARTIST:       http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-video-by-artist&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&artistid=49674&pageindex=1&pagesize=30&time=1391515932000&token=c1641290c36d3a36590df55cbad083dc
	// VIDEOS BY GENRE:        http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-video-by-genre&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&genreid=5142&pageindex=1&pagesize=30&time=1391515932000&token=270c039a73d299c688d8b3b086ae1e7e
	// VIDEOS BY SEARCHING:    http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=search-video&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&keyword=man&pageindex=1&pagesize=30&time=1391515932000&token=4bce7e9c8333ba31076c3cb1a2b85544
	// VIDEOS RELATED:         http://api.m.nhaccuatui.com/mobile/v5.0/api?secretkey=nct@mobile_service&action=get-video-related&deviceinfo={"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}&videoid=12336694&time=1391515932000&token=57aa7164a27ac3dacb25526d7f4a3cbc
}
