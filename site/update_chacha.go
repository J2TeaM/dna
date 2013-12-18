package site

import (
	"dna"
	"dna/cc"
	"dna/sqlpg"
	"time"
)

func UpdateChacha(sqlConfigPath, siteConfigPath dna.String) {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(sqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("cc", siteConfigPath)
	dna.PanicError(err)

	state := NewStateHandler(new(cc.Song), siteConf, db)
	Update(state)
	//  update album
	state = NewStateHandler(new(cc.Album), siteConf, db)
	Update(state)

	state = NewStateHandler(new(cc.Video), siteConf, db)
	Update(state)

	RecoverErrorQueries("./log/sql_error.log", db)

	time.Sleep(3 * time.Second)
	db.Close()

}
