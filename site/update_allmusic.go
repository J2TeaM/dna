package site

import (
	"dna"
	"dna/am"
	"dna/sqlpg"
	"io/ioutil"
	"time"
)

func amGetErrIds(inputFile dna.String, mode dna.Int) *dna.IntArray {
	var ret = dna.IntArray{}
	b, err := ioutil.ReadFile(inputFile.String())
	if err != nil {
		panic(err)
	}
	data := dna.String(string(b))
	lines := data.Split("\n")
	for _, line := range lines {
		switch mode {
		case 1:
			idArr := line.FindAllStringSubmatch(`([0-9]+) Post.+no such host`, 1)
			if len(idArr) > 0 {
				ret.Push(idArr[0][1].ToInt())
			}
			idArr = line.FindAllStringSubmatch(`Timeout.+at id :([0-9]+)`, 1)
			if len(idArr) > 0 {
				ret.Push(idArr[0][1].ToInt())
			}
		case 2:
			ret.Push(line.ToInt())
		}
	}
	if mode == 1 {
		err = ioutil.WriteFile(inputFile.String(), []byte{}, 0644)
		if err != nil {
			dna.Log("Cannot write to file1:", err.Error())
		}

	}
	ret = ret.Unique()
	return &ret
}

func amRecovery() {
	// Recovering failed ambum ids
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("am", SiteConfigPath)
	siteConf.NConcurrent = 30
	dna.PanicError(err)

	// r := NewRange(20987, 30000)
	ids := amGetErrIds("./log/http_error.log", 1)
	tmp := ids.Unique()
	ids = &tmp
	if ids.Length() > 0 {
		state := NewStateHandlerWithExtSlice(new(am.APIAlbum), ids, siteConf, db)
		Update(state)
	} else {
		dna.Log("No need to recover file")
	}

	// Recover failed SQL statements
	RecoverErrorQueries(SqlErrorLogPath, db)

	CountDown(3*time.Second, QuittingMessage, EndingMessage)
	db.Close()
}

func amUpdateNewAlbums() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("am", SiteConfigPath)
	siteConf.NConcurrent = 10
	dna.PanicError(err)

	state := NewStateHandler(new(am.APIAlbum), siteConf, db)
	state.TableName = "amalbums"
	Update(state)
	CountDown(5*time.Second, QuittingMessage, EndingMessage)
	db.Close()
}

// UpdateAllmusic gets lastest albums from allmusic.com.
// The update process goes through 3 steps:
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 2: Updating new albums.
// 	Step 3: Refeching timeout album ids and recovering failed sql statements.
func UpdateAllmusic() {
	amUpdateNewAlbums()
	amRecovery()
}
