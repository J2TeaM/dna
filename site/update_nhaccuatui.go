package site

import (
	"dna"
	"dna/nct"
	"dna/sqlpg"
	"time"
)

// UpdateNhaccuatui gets lastest items from nhaccuatui.com
func UpdateNhaccuatui(sqlConfigPath, siteConfigPath dna.String) {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(sqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("nct", siteConfigPath)
	dna.PanicError(err)
	// Update new songs
	r := NewRange(0, (nct.NewestSongPaths.Length()-1)*nct.TotalSongPages)
	state := NewStateHandlerWithRange(new(nct.SongFinder), r, siteConf, db)
	Update(state)
	dna.Log("The number of songs found:", nct.NewestSongPortions.Length())
	dna.Log("Finding new songs...")
	err = nct.NewestSongPortions.FilterByIds("nctsongs", db)
	dna.PanicError(err)
	dna.Log("The number of NEW songs found:", nct.NewestSongPortions.Length())
	indexRange := NewRange(0, nct.NewestSongPortions.Length()-1)
	state = NewStateHandlerWithRange(new(nct.Song), indexRange, siteConf, db)
	Update(state)
	time.Sleep(3 * time.Second)
	db.Close()
}
