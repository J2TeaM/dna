package site

import (
	"dna"
	"dna/csn"
	"dna/sqlpg"
	"dna/utils"
	"time"
)

func UpdateChiasenhac(sqlConfigPath, siteConfigPath dna.String) {
	// note: songid 1172662 1172663 1172664 are not continuos
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(sqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("csn", siteConfigPath)
	dna.PanicError(err)

	// Getting LastSongId for SaveNewAlbums func
	csn.LastSongId, err = utils.GetMaxId("csnsongs", db)
	dna.PanicError(err)

	// Getting both new songs and videos and
	// inserting into appropriate tables respectively.
	state := NewStateHandler(new(csn.Song), siteConf, db)
	Update(state)

	// csn.LastSongId = 1172666
	dna.Log("Finding and saving new albums from last songid:", csn.LastSongId)
	nAlbums, err := csn.SaveNewAlbums(db)
	if err != nil {
		dna.Log(err.Error())
	} else {
		dna.Log("New albums inserted:", nAlbums)
	}
	RecoverErrorQueries("./log/sql_error.log", db)
	time.Sleep(2 * time.Second)
	db.Close()
}
