/*
This package implements some functionalities relating to updating items from different sites.

It supports different update methods: range, external slice or stop once reaching n continuous failed items. Read more at type StateHandler

Example:
	// main.go
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig("./config/app.ini"))
	dna.PanicError(err)
	siteConf, err := site.LoadSiteConfig("ns", "./config/sites.ini")
	dna.PanicError(err)

	// first pattern
	state := site.NewStateHandler(new(ns.Song), siteConf, db)
	site.Update(state)

	// second pattern
	r := site.NewRange(1315976, 1316276)
	state = site.NewStateHandlerWithRange(new(ns.Song), r, siteConf, db)
	site.Update(state)

	// third pattern
	slice := &dna.IntArray{}
	for i := 0; i < 300; i++ {
		slice.Push(dna.Int(1315976 + i))
	}
	state = site.NewStateHandlerWithExtSlice(new(ns.Song), slice, siteConf, db)
	site.Update(state)
	db.Close()
*/
package site
