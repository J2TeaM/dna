package main

import (
	"dna"
	"dna/sf"
	"dna/site"
	"dna/sqlpg"
	"flag"
	"time"
)

const (
	StepAmount = 1000
)

func spawn(first dna.Int) {
	sf.VideosEnable = true
	sf.CommentsEnable = false
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(site.SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := site.LoadSiteConfig("sf", site.SiteConfigPath)
	siteConf.NConcurrent = 40
	dna.PanicError(err)

	// r := site.NewRange(20987, 30000)

	siteRange := site.NewRange(first, first+StepAmount-1)
	dna.Log(dna.Sprintf("Range from %v", siteRange))
	state := site.NewStateHandlerWithRange(new(sf.APISongFreaksTrack), siteRange, siteConf, db)
	site.Update(state)

	// site.RecoverErrorQueries(site.SqlErrorLogPath, db)
	site.CountDown(2*time.Second, site.QuittingMessage, "")
	db.Close()

}

func main() {
	var first int
	flag.IntVar(&first, "first", 0, "first element of Range")
	flag.Parse()
	if first == 0 {
		dna.Log("Cannot run because first ele of Range is 0")
	} else {
		spawn(dna.Int(first) + 1)
	}

}
