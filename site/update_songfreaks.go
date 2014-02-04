package site

import (
	"dna"
	"dna/sf"
	"dna/sqlpg"
	"time"
)

// UpdateSongFreaks gets lastest songs,albums,artists and videos from songfreaks.com
// The update process goes through 4 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 2: Finding new songs, insert new albums,artists and videos if found.
// 	Step 3: Updating found new albums in Step 2.
// 	Step 4: Recovering failed sql statements in Step 2.
func UpdateSongFreaks() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("sf", SiteConfigPath)
	siteConf.NConcurrent = 20
	dna.PanicError(err)

	// Update new songs
	state := NewStateHandler(new(sf.APISongFreaksTrack), siteConf, db)
	state.TableName = "sfsongs"
	Update(state)

	// Update "ratings", "songids", "review_author", "review" of song
	ids := &[]dna.Int{}
	query := dna.Sprintf("SELECT id FROM sfalbums where checktime > '%v' AND array_length(songids, 1) is NULL", time.Now().Format("2006-01-02"))
	// dna.Log(query)
	err = db.Select(ids, query)
	if err != nil {
		dna.PanicError(err)
	}
	idsSlice := dna.IntArray(*ids)
	if idsSlice.Length() > 0 {
		state = NewStateHandlerWithExtSlice(new(sf.APISongFreaksAlbum), &idsSlice, siteConf, db)
		Update(state)

	} else {
		dna.Log("No new albums found")
	}

	// Recover failed sql statements
	RecoverErrorQueries(SqlErrorLogPath, db)

	CountDown(3*time.Second, QuittingMessage, EndingMessage)
	db.Close()
}
