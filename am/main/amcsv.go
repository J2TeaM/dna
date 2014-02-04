// This program gets records from sfsongfreaks
// and save them to sfsongs.csv, sfartists.csv,
// sfvideos.csv,sfalbums.csv
package main

import (
	"dna"
	"dna/am"
	"dna/site"
	"dna/sqlpg"
	"dna/terminal"
	"encoding/csv"
	"io/ioutil"
	"os"
	"time"
)

var N_RECS_PER_SELECT dna.Int

var SQLERROR = terminal.NewLogger(terminal.Magenta, ioutil.Discard, "", "./log/sql_error.log", 0)

var counter *site.Counter
var bar *terminal.ProgressBar

var (
	awardWriter   *csv.Writer
	creditWriter  *csv.Writer
	discoWriter   *csv.Writer
	releaseWriter *csv.Writer
	songWriter    *csv.Writer
	albumWriter   *csv.Writer
)

var (
	awardFile   *os.File
	discoFile   *os.File
	creditFile  *os.File
	releaseFile *os.File
	songFile    *os.File
	albumFile   *os.File
)

func getUpdateProgressBar(total dna.Int, tableName dna.String) *terminal.ProgressBar {
	var rt dna.String = "$[ " + tableName + " $percent% $current/$total]"
	rt += "\nElapsed: $elapsed    ETA: $eta   Speed: $speeditems/s"
	rt += "\nStats: $custom"
	var ct dna.String = "$[  " + tableName + "  t:$elapsed    N:$total  ($custom)  ν:$speeditems/s]"
	upbar := terminal.NewProgressBar(total, rt, ct)
	upbar.Width = 70
	upbar.CompleteSymbol = " "
	upbar.IncompleteSymbol = " "
	upbar.CompleteBGColor = terminal.Green
	upbar.IncompleteBGColor = terminal.White
	upbar.CompleteTextColor = terminal.Black
	upbar.IncompleteTextColor = terminal.Black
	return upbar
}

func getErrMes(err error, id dna.Int, custom dna.String) dna.String {
	errMess := err.Error() + " at the record:" + id.ToString().String() + " " + custom.String()
	return dna.String(errMess)
}

func saveApiAlbum(apiAlbums *[]am.APIAlbum) {
	for _, apiAlbum := range *apiAlbums {
		awards, credits, discos, releases, songs, album := apiAlbum.Convert()
		for _, award := range awards {
			awardWriter.Write(award.CSVRecord())
		}

		for _, credit := range credits {
			creditWriter.Write(credit.CSVRecord())
		}

		for _, disco := range discos {
			discoWriter.Write(disco.CSVRecord())
		}

		for _, release := range releases {
			releaseWriter.Write(release.CSVRecord())
		}

		for _, song := range songs {
			songWriter.Write(song.CSVRecord())
		}

		albumWriter.Write(album.CSVRecord())

	}
}

func selectSQL(db *sqlpg.DB, id dna.Int) {
	apiAlbums := &[]am.APIAlbum{}
	intA := dna.IntArray{}
	for j := id; j < id+N_RECS_PER_SELECT; j++ {
		intA.Push(j)
	}
	query := dna.Sprintf("select * from amapialbums where id in (%v)", intA.Join(","))
	// dna.Log(query)
	err := db.Select(apiAlbums, query)
	if err != nil {
		SQLERROR.Println(getErrMes(err, id, " while selecting a record"))
	} else {
		// change counter
		length := dna.Int(len(*apiAlbums))
		counter.Pass += length
		counter.Fail += N_RECS_PER_SELECT - length
		counter.Count += N_RECS_PER_SELECT
		counter.ElapsedTime = time.Since(counter.GetStartingTime())
		if counter.ElapsedTime/time.Second > 0 {
			counter.Speed = dna.Int(int64(counter.Count) / int64(counter.ElapsedTime/time.Second))
		}
		saveApiAlbum(apiAlbums)

	}
	cData := dna.Sprintf("%v | Cid:%v", counter, id)
	bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
}

func initWriters() {
	var err error

	// create files
	awardFile, err = os.Create("./amawards.csv")
	if err != nil {
		panic("cannot create amawards.csv")
	}
	discoFile, err = os.Create("./amdiscographies.csv")
	if err != nil {
		panic("cannot create amdiscos.csv")
	}
	creditFile, err = os.Create("./amcredits.csv")
	if err != nil {
		panic("cannot create amcredits.csv")
	}
	releaseFile, err = os.Create("./amreleases.csv")
	if err != nil {
		panic("cannot create amreleases.csv")
	}
	songFile, err = os.Create("./amsongs.csv")
	if err != nil {
		panic("cannot create amsongs.csv")
	}
	albumFile, err = os.Create("./amalbums.csv")
	if err != nil {
		panic("cannot create amalbums.csv")
	}

	// init writers
	awardWriter = csv.NewWriter(awardFile)
	discoWriter = csv.NewWriter(discoFile)
	creditWriter = csv.NewWriter(creditFile)
	releaseWriter = csv.NewWriter(releaseFile)
	songWriter = csv.NewWriter(songFile)
	albumWriter = csv.NewWriter(albumFile)
}

func terminateWriters() {
	// flush writers
	awardWriter.Flush()
	discoWriter.Flush()
	creditWriter.Flush()
	releaseWriter.Flush()
	songWriter.Flush()
	albumWriter.Flush()

	// close files
	awardFile.Close()
	discoFile.Close()
	creditFile.Close()
	releaseFile.Close()
	songFile.Close()
	albumFile.Close()
}

func main() {
	N_RECS_PER_SELECT = 200
	//243930, 244930
	rge := site.NewRange(1, 2595000)
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(site.SqlConfigPath))
	dna.PanicError(err)
	bar = getUpdateProgressBar(rge.Total, "ALL")
	counter = site.NewCounter(rge.Total)
	initWriters()
	for i := rge.First; i <= rge.Last; i += N_RECS_PER_SELECT {
		selectSQL(db, dna.Int(i))
	}
	terminateWriters()

	// close DB
	db.Close()

}
