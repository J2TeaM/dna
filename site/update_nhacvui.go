package site

import (
	"dna"
	"dna/nv"
	"dna/sqlpg"
	"time"
)

// UpdateNhacvui gets lastest items from nhacvui.vn.
// The update process goes through 5 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 2: Updating new songs.
// 	Step 3: Updating new albums.
// 	Step 4: Updating new videos from FoundVideos var.
// 	Step 5: Recovering failed sql statements.
func UpdateNhacvui() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("nv", SiteConfigPath)
	dna.PanicError(err)

	state := NewStateHandler(new(nv.Song), siteConf, db)
	Update(state)
	//  update album
	state = NewStateHandler(new(nv.Album), siteConf, db)
	Update(state)

	if nv.FoundVideos.Length() > 0 {
		state = NewStateHandlerWithExtSlice(new(nv.Video), nv.FoundVideos, siteConf, db)
		Update(state)
	} else {
		dna.Log("No videos found!")
	}

	RecoverErrorQueries(SqlErrorLogPath, db)

	time.Sleep(3 * time.Second)
	db.Close()

}
