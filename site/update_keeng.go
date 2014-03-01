package site

import (
	"dna"
	"dna/ke"
	"dna/sqlpg"
	"time"
)

// UpdateKeeng gets lastest items from keeng.com.
// The update process goes through 6 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 2: Updating new artists.
// 	Step 3: Updating new songs.
// 	Step 4: Updating new albums.
// 	Step 5: Updating new videos.
// 	Step 6: Recovering failed sql statements.
func UpdateKeeng() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("ke", SiteConfigPath)
	dna.PanicError(err)

	state := NewStateHandler(new(ke.Artist), siteConf, db)
	Update(state)

	state = NewStateHandler(new(ke.Song), siteConf, db)
	Update(state)

	state = NewStateHandler(new(ke.Album), siteConf, db)
	Update(state)

	state = NewStateHandler(new(ke.Video), siteConf, db)
	Update(state)

	RecoverErrorQueries(SqlErrorLogPath, db)
	CountDown(3*time.Second, QuittingMessage, EndingMessage)
	db.Close()
}
