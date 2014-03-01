package site

import (
	"dna"
	"dna/nct"
	"dna/sqlpg"
	"time"
)

// UpdateNhaccuatui gets lastest items from nhaccuatui.com.
//
// The update process goes through 6 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 2: Updating new songs.
// 	Step 3: Updating new albums.
// 	Step 4: Updating new videos.
// 	Step 5: Updating new artists.
// 	Step 6: Recovering failed sql statements.
func UpdateNhaccuatui() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("nct", SiteConfigPath)
	dna.PanicError(err)
	TIMEOUT_SECS = 60 // for safe
	// siteConf.NConcurrent = 15

	state := NewStateHandler(new(nct.Song), siteConf, db)
	Update(state)

	state = NewStateHandler(new(nct.Album), siteConf, db)
	Update(state)

	state = NewStateHandler(new(nct.Video), siteConf, db)
	Update(state)

	state = NewStateHandler(new(nct.Artist), siteConf, db)
	// Because there is no setting for total concurrent artists failed,
	// the default value NCSongFail will be returned.
	// Now, it is set to new value
	state.SiteConfig.NCSongFail = 1000 // Set concurrency for artists
	Update(state)
	RecoverErrorQueries(SqlErrorLogPath, db)
	CountDown(3*time.Second, QuittingMessage, EndingMessage)
	db.Close()
}
