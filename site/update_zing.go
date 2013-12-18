package site

import (
	"dna"
	"dna/sqlpg"
	"dna/utils"
	"dna/zi"
	"time"
)

// UpdateZing gets lastest items from mp3.zing.vn
func UpdateZing(sqlConfigPath, siteConfigPath dna.String) {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(sqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("zi", siteConfigPath)
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
	nids, err := utils.SelectNonExistedIds("zisongs", ids, db)
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

	RecoverErrorQueries("./log/sql_error.log", db)

	time.Sleep(3 * time.Second)
	db.Close()
}
