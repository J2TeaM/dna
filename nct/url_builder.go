package nct

import (
	"crypto/md5"
	"dna"
	"io"
	"time"
)

const (
	TOKEN_KEY    = "nct@asdgvhfhyth"
	SECRET_KEY   = "nct@mobile_service"
	DEVICE_INFOR = `{"DeviceID":"90c18c4cb3c37d442e8386631d46b46f","OsName":"ANDROID","OsVersion":"10","AppName":"NhacCuaTui","AppVersion":"5.0.1","UserInfo":"","LocationInfo":""}`
	BASE_URL     = "http://api.m.nhaccuatui.com/mobile/v5.0/api"
)

func getTimestampStr() dna.String {
	return dna.Sprintf("%v", time.Now().UnixNano()/1000000)
}

func getMD5(s dna.String) dna.String {
	h := md5.New()
	io.WriteString(h, s.String())
	return dna.Sprintf("%x", h.Sum(nil))
}

type URLBuilder struct {
	tsStr dna.String // A string containing timestamp in milisecs
}

func NewURLBuilder() *URLBuilder {
	urlb := new(URLBuilder)
	urlb.tsStr = getTimestampStr()
	return urlb
}

func (urlb *URLBuilder) SetTimestamp(ts time.Time) {
	urlb.tsStr = dna.Sprintf("%v", ts.UnixNano()/1000000)
}

func (urlb *URLBuilder) GetSongsByGenre(genreId, pageIndex, pageSize dna.Int) dna.String {
	s1 := "get-song-by-genre" + genreId.ToString() + pageIndex.ToString() + pageSize.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-song-by-genre&deviceinfo=%v&genreid=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, genreId, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetSongsOfPlaylist(albumid dna.Int) dna.String {
	s1 := "get-song-by-playlistid" + albumid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-song-by-playlistid&deviceinfo=%v&playlistid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, albumid, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetVideoInfo(videoid dna.Int) dna.String {

	//NOT COMPLETE
	s1 := "get-video-detail" + videoid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-video-detail&deviceinfo=%v&videoid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, videoid, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetSongLyric(songid dna.Int) dna.String {
	s1 := "get-lyric" + songid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-lyric&deviceinfo=%v&songid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, songid, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetPlaylistInfo(albumid dna.Int) dna.String {
	//NOT COMPLETE
	s1 := "get-playlist-info" + albumid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-playlist-info&deviceinfo=%v&playlistid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, albumid, urlb.tsStr, getMD5(s1))
}

// GetListOfGenres return a list of genres
// genreType is an one of ["song","video","playlist"]
func (urlb *URLBuilder) GetListOfGenres(genreType dna.String) dna.String {
	s1 := "get-list-genre" + genreType + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-list-genre&deviceinfo=%v&type=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, genreType, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetSongsByArtist(artistid, pageIndex, pageSize dna.Int) dna.String {
	s1 := "get-song-by-artist" + dna.Sprintf("%v%v%v", artistid, pageIndex, pageSize) + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-song-by-artist&deviceinfo=%v&artistid=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, artistid, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetPlaylistsByArtist(artistid, pageIndex, pageSize dna.Int) dna.String {
	s1 := "get-playlist-by-artist" + dna.Sprintf("%v%v%v", artistid, pageIndex, pageSize) + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-playlist-by-artist&deviceinfo=%v&artistid=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, artistid, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetVideosByArist(artistid, pageIndex, pageSize dna.Int) dna.String {
	s1 := "get-video-by-artist" + dna.Sprintf("%v%v%v", artistid, pageIndex, pageSize) + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-video-by-artist&deviceinfo=%v&artistid=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, artistid, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetListOfTopics(pageIndex, pageSize dna.Int) dna.String {
	s1 := "get-list-topic" + dna.Sprintf("%v%v", pageIndex, pageSize) + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-list-topic&deviceinfo=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetPlaylistsByTopic(topicId dna.Int) dna.String {
	s1 := "get-playlist-by-topic" + topicId.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-playlist-by-topic&deviceinfo=%v&topicid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, topicId, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetPlaylistsRelated(playlistid dna.Int) dna.String {
	s1 := "get-playlist-related" + playlistid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-playlist-related&deviceinfo=%v&playlistid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, playlistid, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetVideosRelated(videoid dna.Int) dna.String {
	s1 := "get-video-related" + videoid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-video-related&deviceinfo=%v&videoid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, videoid, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetSongsBySearching(keyword dna.String, pageIndex, pageSize dna.Int) dna.String {
	s1 := "search-song" + dna.Sprintf("%v%v%v", keyword, pageIndex, pageSize) + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=search-song&deviceinfo=%v&keyword=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, keyword, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetPlaylistsBySearching(keyword dna.String, pageIndex, pageSize dna.Int) dna.String {
	s1 := "search-playlist" + dna.Sprintf("%v%v%v", keyword, pageIndex, pageSize) + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=search-playlist&deviceinfo=%v&keyword=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, keyword, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetVideosBySearching(keyword dna.String, pageIndex, pageSize dna.Int) dna.String {
	s1 := "search-video" + dna.Sprintf("%v%v%v", keyword, pageIndex, pageSize) + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=search-video&deviceinfo=%v&keyword=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, keyword, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetPlaylistsByGenre(genreId, pageIndex, pageSize dna.Int) dna.String {
	s1 := "get-playlist-by-genre" + genreId.ToString() + pageIndex.ToString() + pageSize.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-playlist-by-genre&deviceinfo=%v&genreid=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, genreId, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetVideosByGenre(genreId, pageIndex, pageSize dna.Int) dna.String {
	s1 := "get-video-by-genre" + genreId.ToString() + pageIndex.ToString() + pageSize.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-video-by-genre&deviceinfo=%v&genreid=%v&pageindex=%v&pagesize=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, genreId, pageIndex, pageSize, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetArtistInfo(artistid dna.Int) dna.String {
	s1 := "get-artist-detail" + artistid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-artist-detail&deviceinfo=%v&artistid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, artistid, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetSongInfo(songid dna.Int) dna.String {
	//NOT COMPLETE
	s1 := "get-song-info" + songid.ToString() + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=get-song-info&deviceinfo=%v&songid=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, songid, urlb.tsStr, getMD5(s1))
}

func (urlb *URLBuilder) GetCheckVersion() dna.String {
	s1 := "check-version" + TOKEN_KEY + urlb.tsStr
	return dna.Sprintf("%v?secretkey=%v&action=check-version&deviceinfo=%v&time=%v&token=%v", BASE_URL, SECRET_KEY, DEVICE_INFOR, urlb.tsStr, getMD5(s1))
}
