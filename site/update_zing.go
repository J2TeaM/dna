package site

import (
	"dna"
	"dna/sqlpg"
	"dna/utils"
	"dna/zi"
	"time"
)

// UpdateZing gets lastest items from mp3.zing.vn.
// The update process goes through 8 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler
// 	Step 1: Updating new songs
// 	Step 3: Updating new albums
// 	Step 4: Updating new videos
// 	Step 5: Updating new artists
// 	Step 6: Updating new songids found in new albums
// 	Step 7: Updating new tvs
// 	Step 8: Recovering failed sql statements
func UpdateZing() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("zi", SiteConfigPath)
	dna.PanicError(err)
	// update song
	state := NewStateHandler(new(zi.Song), siteConf, db)
	Update(state)
	// update album
	state = NewStateHandler(new(zi.Album), siteConf, db)
	Update(state)
	// update video
	state = NewStateHandler(new(zi.Video), siteConf, db)
	Update(state)
	// update artist
	state = NewStateHandler(new(zi.Artist), siteConf, db)
	Update(state)

	// update new songids found in albums
	dna.Log("Update new songs from albums")
	ids := utils.SelectNewSidsFromAlbums("zialbums", time.Now(), db)
	nids, err := utils.SelectMissingIds("zisongs", ids, db)
	if err != nil {
		dna.Log(err.Error())
	} else {
		if nids != nil && nids.Length() > 0 {
			state = NewStateHandlerWithExtSlice(new(zi.Song), nids, siteConf, db)
			Update(state)
		} else {
			dna.Log("No new songs found")
		}

	}

	state = NewStateHandler(new(zi.TV), siteConf, db)
	Update(state)

	RecoverErrorQueries(SqlErrorLogPath, db)

	time.Sleep(3 * time.Second)
	db.Close()
}
