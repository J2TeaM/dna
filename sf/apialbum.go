package sf

import (
	"dna"
)

// APIAlbum defines aln album.
type APIAlbum struct {
	Id            dna.Int         `xml:"id,attr"`
	AMG           dna.Int         `xml:"amg,attr"`
	UrlSlug       dna.String      `xml:"urlslug,attr"`
	Year          dna.Int         `xml:"year,attr"`
	Coverart      dna.String      `xml:"image,attr"`
	CoverartLarge dna.String      `xml:"largeimage,attr"`
	Title         dna.String      `xml:"title"`
	Rating        APIRating       `xml:"rating"` // Added field
	Artist        APIArtist       `xml:"artist"`
	Link          dna.String      `xml:"link"`
	Tracks        []APIAlbumTrack `xml:"tracks>track"`
	Review        APIAlbumReview  `xml:"review"`
}

// APIAlbumTrack defines a mini track of an album.
// It does not have complete fields as APITrack's.
// Choose track_group_id as main id
type APIAlbumTrack struct {
	Id             dna.Int    `xml:"track_group_id,attr"`
	TrackGroupId   dna.Int    `xml:"id,attr"`
	UrlSlug        dna.String `xml:"urlslug,attr"`
	AMG            dna.Int    `xml:"amg,attr"`
	IsInstrumental dna.Bool   `xml:"instrumental,attr"`
	Viewable       dna.Bool   `xml:"viewable,attr"`
	Duration       dna.String `xml:"duration,attr"`
	LyricId        dna.Int    `xml:"lyric,attr"`
	HasLrc         dna.Bool   `xml:"has_lrc,attr"`
	TrackNumber    dna.Int    `xml:"tracknumber,attr"`
	DiscNumber     dna.Int    `xml:"discnumber,attr"`
	Title          dna.String `xml:"title"`
	Link           dna.String `xml:"link"`
}

type APIAlbumReview struct {
	Author dna.String `xml:"author,attr"`
	Value  dna.String `xml:",innerxml"`
}
