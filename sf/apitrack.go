package sf

import (
	"dna"
)

type APIResponse struct {
	Code       dna.Int    `xml:"code,attr"`
	RenderTime dna.Float  `xml:"renderTime,attr"`
	Value      dna.String `xml:",innerxml"`
}

type APIRating struct {
	AverageRating dna.Float `xml:"averagerating,attr"`
	UserRating    dna.Int   `xml:"userrating,attr"`
	TotalRatings  dna.Int   `xml:"totalratings,attr"`
}

// APITrackAlbum defines a mini album of a track tag.
// It does not have complete fields as APIAlbum's
type APITrackAlbum struct {
	Id            dna.Int    `xml:"id,attr"`
	AMG           dna.Int    `xml:"amg,attr"`
	UrlSlug       dna.String `xml:"urlslug,attr"`
	Year          dna.Int    `xml:"year,attr"`
	Coverart      dna.String `xml:"image,attr"`
	CoverartLarge dna.String `xml:"largeimage,attr"`
	Title         dna.String `xml:"title"`
	Artist        APIArtist  `xml:"artist"`
	Link          dna.String `xml:"link"`
}

type APIArtistBio struct {
	Author dna.String `xml:"author,attr" json:"au"`
	Value  dna.String `xml:",innerxml" json:"value"`
}

type APIArtist struct {
	Id      dna.Int      `xml:"id,attr"`
	AMG     dna.Int      `xml:"amg,attr"`
	UrlSlug dna.String   `xml:"urlslug,attr"`
	Image   dna.String   `xml:"image,attr"`
	Genres  dna.String   `xml:"genre,attr"`
	Name    dna.String   `xml:"name"`
	Link    dna.String   `xml:"link"`
	Rating  APIRating    `xml:"rating"`
	Bio     APIArtistBio `xml:"bio"`
}

type APILrcLine struct {
	Timestamp    dna.String `xml:"lrc_timestamp,attr" json:"ts"`
	Milliseconds dna.Int    `xml:"milliseconds,attr" json:"ms"`
	Line         dna.String `xml:",innerxml" json:"line"`
}

type APILrc struct {
	Lines []APILrcLine `xml:"line"`
}

type APIFoxMobile struct {
	Available dna.Bool `xml:"available,attr"`
}

type APITrack struct {
	// Choose track_group_id as main id
	Id             dna.Int       `xml:"track_group_id,attr"`
	TrackGroupId   dna.Int       `xml:"id,attr"`
	UrlSlug        dna.String    `xml:"urlslug,attr"`
	AMG            dna.Int       `xml:"amg,attr"`
	IsInstrumental dna.Bool      `xml:"instrumental,attr"`
	Viewable       dna.Bool      `xml:"viewable,attr"`
	Duration       dna.String    `xml:"duration,attr"`
	LyricId        dna.Int       `xml:"lyric,attr"`
	HasLrc         dna.Bool      `xml:"has_lrc,attr"`
	TrackNumber    dna.Int       `xml:"tracknumber,attr"`
	DiscNumber     dna.Int       `xml:"discnumber,attr"`
	Title          dna.String    `xml:"title"`
	Rating         APIRating     `xml:"rating"`
	Album          APITrackAlbum `xml:"album"`
	Artists        []APIArtist   `xml:"artists>artist"`
	Lrc            APILrc        `xml:"lrc"`
	Link           dna.String    `xml:"link"`
	Lyrics         dna.String    `xml:"lyrics"`
	Copyright      dna.String    `xml:"copyright"`
	Writer         dna.String    `xml:"writer"`
	SubmittedLyric dna.Bool      `xml:"submittedlyric"`
	FoxMobile      APIFoxMobile  `xml:"foxmobile"`
}

type APIVideo struct {
	YoutubeId   dna.String `xml:"youtubeid,attr"`
	Duration    dna.Int    `xml:"duration,attr"`
	Thumbnail   dna.String `xml:"thumbnail,attr"`
	Title       dna.String `xml:"title"`
	Description dna.String `xml:"description"`
}

type APIContent struct {
	Id   dna.Int    `xml:"contentid,attr"`
	Type dna.String `xml:"contenttype,attr"`
}

type APIUserRating struct {
	Value dna.Int `xml:"value,attr"`
}

type APIUser struct {
	Id          dna.Int    `xml:"id,attr"`
	DisplayName dna.String `xml:"displayname,attr"`
	Avatar      dna.String `xml:"avatar,attr"`
	BadgeScore  dna.Int    `xml:"badgescore,attr"`
}

type APIComment struct {
	Id         dna.Int       `xml:"id,attr"`
	Date       dna.String    `xml:"date,attr"`
	Deleted    dna.Bool      `xml:"deleted,attr"`
	Likes      dna.Int       `xml:"likes,attr"`
	Dislike    dna.Int       `xml:"dislikes,attr"`
	User       APIUser       `xml:"user"`
	UserRating APIUserRating `xml:"userrating"`
	Content    APIContent    `xml:"content"`
	Message    dna.String    `xml:"message"`
	/// Not impplementing "replies" attribute
}

type APIComments struct {
	NResult dna.Int      `xml:"totalresults,attr"`
	NPages  dna.Int      `xml:"totalpages,attr"`
	Values  []APIComment `xml:"comment"`
}
