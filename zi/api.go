package zi

import (
	"dna"
)

// APISong maps JSON fields of a song in the API to struct fields.
type APISong struct {
	Id             dna.Int                   `json:"song_id"`
	Key            dna.String                `json:"song_id_encode"`
	Title          dna.String                `json:"title"`
	ArtistId       dna.String                `json:"artist_id"`
	Artists        dna.String                `json:"artist"`
	AlbumId        dna.Int                   `json:"album_id"`
	Album          dna.String                `json:"album"`
	ComposerId     dna.Int                   `json:"composer_id"`
	Composer       dna.String                `json:"composer"`
	GenreId        dna.String                `json:"genre_id"`
	Zaloid         dna.Int                   `json:"zaloid"`
	Username       dna.String                `json:"username"`
	IsHit          dna.Int                   `json:"is_hit"`
	IsOfficial     dna.Int                   `json:"is_official"`
	DownloadStatus dna.Int                   `json:"download_status"`
	Copyright      dna.String                `json:"copyright"`
	Thumbnail      dna.String                `json:"thumbnail"`
	Plays          dna.Int                   `json:"total_play"`
	Link           dna.String                `json:"link"`
	Source         map[dna.String]dna.String `json:"source"`
	LinkDownload   map[dna.String]dna.String `json:"link_download"`
	AlbumCover     dna.String                `json:"album_cover"`
	Likes          dna.Int                   `json:"likes"`
	LikeThis       dna.Bool                  `json:"like_this"`
	Favourites     dna.Int                   `json:"favourites"`
	FavouritesThis dna.Bool                  `json:"favourite_this"`
	Comments       dna.Int                   `json:"comments"`
	GenreName      dna.String                `json:"genre_name"`
	Video          APIVideo                  `json:"video"`
	Response       APIResponse               `json:"response"`
}

// APISongLyric maps JSON fields of a song lyric in the API to struct fields.
type APISongLyric struct {
	Id       dna.String  `json:"id"`
	Content  dna.String  `json:"content"`
	Mark     dna.Int     `json:"mark"`
	Author   dna.String  `json:"author"`
	Response APIResponse `json:"response"`
}

// APIResponse maps JSON fields of a reponse in the API to a struct field.
type APIResponse struct {
	MsgCode dna.Int `json:"msgCode"`
}

// APIAlbum maps JSON fields of an album in the API to struct fields.
type APIAlbum struct {
	Id             dna.Int     `json:"playlist_id"`
	Title          dna.String  `json:"title"`
	ArtistId       dna.String  `json:"artist_id"`
	Artists        dna.String  `json:"artist"`
	GenreId        dna.String  `json:"genre_id"`
	Zaloid         dna.Int     `json:"zaloid"`
	Username       dna.String  `json:"username"`
	Cover          dna.String  `json:"cover"`
	Description    dna.String  `json:"description":`
	IsHit          dna.Int     `json:"is_hit"`
	IsOfficial     dna.Int     `json:"is_official"`
	IsAlbum        dna.Int     `json:"is_album"`
	Year           dna.String  `json:"year"`
	StatusId       dna.Int     `json:"status_id"`
	Link           dna.String  `json:"link"`
	Plays          dna.Int     `json:"total_play"`
	GenreName      dna.String  `json:"genre_name"`
	Likes          dna.Int     `json:"likes"`
	LikeThis       dna.Bool    `json:"like_this"`
	Comments       dna.Int     `json:"comments"`
	Favourites     dna.Int     `json:"favourites"`
	FavouritesThis dna.Int     `json:"favourite_this"`
	Response       APIResponse `json:"response"`
}

// APIVideo maps JSON fields of a video in the API to struct fields.
type APIVideo struct {
	Id             dna.Int                   `json:"video_id"`
	Title          dna.String                `json:"title"`
	ArtistId       dna.String                `json:"artist_id"`
	Artists        dna.String                `json:"artist"`
	GenreId        dna.String                `json:"genre_id"`
	Thumbnail      dna.String                `json:"thumbnail"`
	Duration       dna.Int                   `json:"duration"`
	StatusId       dna.Int                   `json:"status_id"`
	Link           dna.String                `json:"link"`
	Source         map[dna.String]dna.String `json:"source"`
	Plays          dna.Int                   `json:"total_play"`
	Likes          dna.Int                   `json:"likes"`
	LikeThis       dna.Bool                  `json:"like_this"`
	Favourites     dna.Int                   `json:"favourites"`
	FavouritesThis dna.Bool                  `json:"favourite_this"`
	Comments       dna.Int                   `json:"comments"`
	GenreName      dna.String                `json:"genre_name"`
	Response       APIResponse               `json:"response"`
}

// APIVideoLyric maps JSON fields of a video lyric in the API to struct fields.
type APIVideoLyric struct {
	Id          dna.String  `json:"id"`
	Content     dna.String  `json:"content"`
	Mark        dna.Int     `json:"mark"`
	StatusId    dna.Int     `json:"status_id"`
	Author      dna.String  `json:"author"`
	DateCreated dna.Int     `json:"created_date"`
	Response    APIResponse `json:"response"`
}

// APIArtist maps JSON fields of an artist in the API to struct fields.
type APIArtist struct {
	Id           dna.Int     `json:"artist_id"`
	Name         dna.String  `json:"name"`
	Alias        dna.String  `json:"alias"`
	Birthname    dna.String  `json:"birthname`
	Birthday     dna.String  `json:"birthday"`
	Sex          dna.Int     `json:"sex"`
	GenreId      dna.String  `json:"genre_id"`
	Avatar       dna.String  `json:"avatar"`
	Cover        dna.String  `json:"cover"`
	Cover2       dna.String  `json:"cover2"`
	ZmeAcc       dna.String  `json:"zme_acc"`
	Role         dna.String  `json:"role"`
	Website      dna.String  `json:"website"`
	Biography    dna.String  `json:"biography"`
	AgencyName   dna.String  `json:"agency_name"`
	NationalName dna.String  `json:"national_name"`
	IsOfficial   dna.Int     `json:"is_official"`
	YearActive   dna.String  `json:"year_active"`
	StatusId     dna.Int     `json:"status_id"`
	DateCreated  dna.Int     `json:"created_date"`
	Link         dna.String  `json:"link"`
	GenreName    dna.String  `json:"genre_name"`
	Response     APIResponse `json:"response"`
}
