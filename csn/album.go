package csn

import (
	"dna"
	"dna/sqlpg"
	"errors"
	"time"
)

type Album struct {
	Id           dna.Int
	Title        dna.String
	Artists      dna.StringArray
	Topics       dna.StringArray
	Href         dna.String
	Nsongs       dna.Int
	Songids      dna.IntArray
	Coverart     dna.String
	Producer     dna.String
	Downloads    dna.Int
	Plays        dna.Int
	DateReleased dna.String
	DateCreated  time.Time
	Checktime    time.Time
}

// SaveNewAlbums finds all new albums available in new found songs.
// New albums depends on aggregate funcs from the last songid.
// It returns an error or nil and the number of new albums inserted into DB
func SaveNewAlbums(db *sqlpg.DB) (nAlbumsInserted dna.Int, err error) {
	if LastSongId <= 0 {
		return 0, errors.New("Last song id has to be greater than Zero (0)")
	}
	query := dna.Sprintf(`insert into %v 
		(title,artists,topics,href,nsongs,coverart,producer,downloads,plays,date_released,date_created,songids) 
		(
		select title,array_agg_csn_artist_cat(artists) as artists,array_agg_csn_cat(topics) as topics, min(link) as href ,
		sum(nsongs)::int8 as nsongs,max(coverart) as coverart, max(producer) as producer,
		floor(avg(downloads))::int8 as downloads, floor(avg(plays))::int8 as plays,
		max(date_released) as date_released , max(date_created) as date_created, array_agg_cat_unique(songids) as songids  
		from 
		(
		select max(album_title) as title,array_agg_csn_artist_cat(artists) as artists,
		array_agg_csn_cat(topics) as topics , min(album_href) as link, 
		count(*) as nsongs, album_coverart as coverart, max(producer) as producer, 
		floor(avg(downloads)) as downloads, floor(avg(plays)) as plays, 
		max(date_released) as date_released,max(date_created) as date_created, array_agg(id) as songids 
		from %v
		where id > %v 
		and album_title <> '' 
		and album_href <> '' 
		group by album_coverart 
		) as anbinh
		group by title
		)`, "csnalbums", "csnsongs", LastSongId)
	// dna.Log(query)
	ret, err := db.Exec(query.String())
	if err != nil {
		return 0, err
	} else {
		nAl, err := ret.RowsAffected()
		nAlbumsInserted = dna.Int(nAl)
		if err != nil {
			return 0, nil
		} else {
			return nAlbumsInserted, nil
		}

	}
}
