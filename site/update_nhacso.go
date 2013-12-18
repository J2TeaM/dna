package site

import (
	"dna"
	"dna/ns"
	"dna/sqlpg"
	"time"
)

// UpdateNhacso gets lastest items from nhacso.com
func UpdateNhacso(sqlConfigPath, siteConfigPath dna.String) {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(sqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("ns", siteConfigPath)
	dna.PanicError(err)
	// update song
	state := NewStateHandler(new(ns.Song), siteConf, db)
	Update(state)
	//  update album
	state = NewStateHandler(new(ns.Album), siteConf, db)
	Update(state)
	// update video
	state = NewStateHandler(new(ns.Video), siteConf, db)
	Update(state)

	r := NewRange(0, 209)
	state = NewStateHandlerWithRange(new(ns.SongCategory), r, siteConf, db)
	Update(state)

	state = NewStateHandlerWithRange(new(ns.AlbumCategory), r, siteConf, db)
	Update(state)

	RecoverErrorQueries("./log/sql_error.log", db)
	time.Sleep(3 * time.Second)
	db.Close()

}
