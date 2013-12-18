package site

import (
	"dna"
	"dna/nv"
	"dna/sqlpg"
	"time"
)

func UpdateNhacvui(sqlConfigPath, siteConfigPath dna.String) {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(sqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("nv", siteConfigPath)
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

	RecoverErrorQueries("./log/sql_error.log", db)

	time.Sleep(3 * time.Second)
	db.Close()

}
