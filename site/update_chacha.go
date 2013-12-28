package site

import (
	"dna"
	"dna/cc"
	"dna/sqlpg"
	"time"
)

// UpdateChacha updates chacha.vn.
// The update process goes through 5 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 1: Updating new songs.
// 	Step 3: Updating new albums.
// 	Step 4: Updating new videos.
// 	Step 5: Recovering failed sql statements.
func UpdateChacha() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("cc", SiteConfigPath)
	dna.PanicError(err)

	state := NewStateHandler(new(cc.Song), siteConf, db)
	Update(state)
	//  update album
	state = NewStateHandler(new(cc.Album), siteConf, db)
	Update(state)

	state = NewStateHandler(new(cc.Video), siteConf, db)
	Update(state)

	RecoverErrorQueries(SqlErrorLogPath, db)

	time.Sleep(3 * time.Second)
	db.Close()

}
