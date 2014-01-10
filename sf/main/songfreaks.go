package main

import (
	"dna"
	"dna/sf"
	"dna/site"
	"dna/sqlpg"
	"dna/terminal"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

const (
	NStep         = 33
	StepAmount    = 1000
	RecoverySteps = 10
)

var (
	StepCount = 0
)

func getErrorIds(inputFile dna.String, mode dna.Int) *dna.IntArray {
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

func RecoverIDS() {
	sf.VideosEnable = true
	sf.CommentsEnable = false
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(site.SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := site.LoadSiteConfig("sf", site.SiteConfigPath)
	siteConf.NConcurrent = 40
	dna.PanicError(err)

	// r := site.NewRange(20987, 30000)
	ids := getErrorIds("./log/http_error.log", 1)
	if ids.Length() > 0 {
		state := site.NewStateHandlerWithExtSlice(new(sf.APISongFreaksTrack), ids, siteConf, db)
		site.Update(state)
	} else {
		dna.Log("No need to recover file")
	}

	// site.RecoverErrorQueries(site.SqlErrorLogPath, db)
	site.CountDown(3*time.Second, site.QuittingMessage, site.EndingMessage)
	db.Close()
}

func main() {
	var first dna.Int
	var filePath = "./log/range_sf.log"
	var console = terminal.NewConsole()

	dna.Log("✦ Total Step:", NStep, "Total records:", NStep*StepAmount)
	dna.Log("✦ The number of steps once recovery mode triggered:", RecoverySteps)

	stopWatch := terminal.NewStopWatch()
	stopWatch.Start()
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	first = dna.String(string(b)).ToInt()

	for StepCount < NStep {
		StepCount += 1
		mes := dna.Sprintf("➤ STEP %v", StepCount)
		dna.Log(terminal.NewColorString(mes).BlueBg().Black().Value())

		// program "spawnsf" is supposed to download tracks
		// and insert raw data into DB
		// Program "spawnsfalbums" is supposed to update songids, review...
		// from an album.
		cmd := exec.Command("./spawnsfalbums", dna.Sprintf("-first=%v", first).String())
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		first += StepAmount
		err = ioutil.WriteFile(filePath, []byte(first.ToString().String()), 0644)
		if err != nil {
			dna.Log("*********************************")
			dna.Log("ERROR WRITTING FILE", err.Error())
			dna.Log("*********************************")
		}
		if StepCount%RecoverySteps == 0 {
			dna.Log(terminal.NewColorString("ENTERING RECOVERY MODE: HTTP & SQL ERRORS").RedBg().Black().Value())
			RecoverIDS()
			db, err := sqlpg.Connect(sqlpg.NewSQLConfig(site.SqlConfigPath))
			dna.PanicError(err)
			site.RecoverErrorQueries(site.SqlErrorLogPath, db)
			dna.Log("\n")
			db.Close()
		}
		if StepCount%6 == 0 {
			console.Erase(terminal.Screen).Move(0, 0)
		}
		stopWatch.Tick()
	}
	stopWatch.Stop()
	stopMes := dna.Sprintf("Total time: %s", stopWatch.Elapsed/time.Millisecond*time.Millisecond)
	dna.Log(stopMes)

}
