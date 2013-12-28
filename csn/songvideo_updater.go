package csn

import (
	"dna/item"
	"dna/sqlpg"
)

// SongVideoUpdater re-fetches and updates some missing fields
// from Song or Video.
type SongVideoUpdater struct {
	SongVideo
}

// NewSongVideoUpdater returns new SongVideoUpdater.
func NewSongVideoUpdater() *SongVideoUpdater {
	return &SongVideoUpdater{*NewSongVideo()}
}

// New returns warpper of item.Item interface
func (svu *SongVideoUpdater) New() item.Item {
	return item.Item(NewSongVideoUpdater())
}

// Save stores some fields to DB.
func (svu *SongVideoUpdater) Save(db *sqlpg.DB) error {
	switch svu.Item.(type) {
	case *Song:
		return db.Update(svu.Item, "id", "title", "artists", "authors", "topics", "album_title", "album_href", "album_coverart", "producer", "lyric", "date_released", "date_created", "is_lyric")
	case *Video:
		return db.Update(svu.Item, "id", "title", "artists", "authors", "topics", "producer", "lyric", "date_released", "date_created", "is_lyric")
	default:
		panic("Invalid type of SongVideoUpdater.Item")
	}

}
