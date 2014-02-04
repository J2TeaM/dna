package site

import (
	"dna"
	"dna/ns"
	"dna/sqlpg"
	"time"
)

// UpdateNhacso gets lastest items from nhacso.com.
// The update process goes through 7 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 2: Updating new songs.
// 	Step 3: Updating new albums.
// 	Step 4: Updating new videos.
// 	Step 5: Updating catagories of new songs.
// 	Step 6: Updating catagories of new albums.
// 	Step 7: Recovering failed sql statements.
func UpdateNhacso() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("ns", SiteConfigPath)
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
	siteConf.NConcurrent = 10
	state = NewStateHandlerWithRange(new(ns.SongCategory), r, siteConf, db)
	Update(state)

	state = NewStateHandlerWithRange(new(ns.AlbumCategory), r, siteConf, db)
	Update(state)

	RecoverErrorQueries(SqlErrorLogPath, db)
	time.Sleep(3 * time.Second)

	CountDown(3*time.Second, QuittingMessage, EndingMessage)

	db.Close()

}
