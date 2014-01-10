// This program gets records from sfsongfreaks
// and save them to sfsongs.csv, sfartists.csv,
// sfvideos.csv,sfalbums.csv
package main

import (
	"dna"
	"dna/sf"
	"dna/site"
	"dna/sqlpg"
	"dna/terminal"
	"encoding/csv"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"time"
)

const N_RECS_PER_SELECT = 1000

var SQLERROR = terminal.NewLogger(terminal.Magenta, ioutil.Discard, "", "./log/sql_error.log", 0)

var counter = site.NewCounter(1000)
var bar *terminal.ProgressBar

var (
	songWriter   *csv.Writer
	albumWriter  *csv.Writer
	artistWriter *csv.Writer
	videoWriter  *csv.Writer
)

var (
	songFile   *os.File
	artistFile *os.File
	albumFile  *os.File
	videoFile  *os.File
)

type SongFreak struct {
	Id     dna.Int
	Track  dna.String
	Videos dna.String
}

type Videos struct {
	APIVideo []sf.APIVideo
}

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

func decodeSongfreakTrack(songfreak *SongFreak, id dna.Int) (*sf.APISongFreaksTrack, error) {
	var merr error
	sfTrack := sf.NewAPISongFreaksTrack()
	sfTrack.Response.Code = 101
	merr = xml.Unmarshal([]byte(songfreak.Track.String()), &sfTrack.Track)
	if merr != nil {
		SQLERROR.Println(getErrMes(merr, id, " while decoding Track"))
	}

	var videos Videos
	if songfreak.Videos.String() != "" {
		merr = xml.Unmarshal([]byte("<Videos>"+songfreak.Videos.String()+"</Videos>"), &videos)
		if merr != nil {
			SQLERROR.Println(getErrMes(merr, id, "while decoding Videos"))
		}
	}
	sfTrack.Videos = videos.APIVideo

	return sfTrack, merr
}

func getComponents(apisf *sf.APISongFreaksTrack, id dna.Int) (*sf.Artist, *sf.Album, *sf.Song, []*sf.Video, error) {
	artist, err := apisf.ToAlbumArtist()
	if err != nil {
		if err.Error() != "Cannot convert artist" {
			SQLERROR.Println(getErrMes(err, id, "while tranforming to Artist"))
		}

	}
	album, err := apisf.ToAlbum()
	if err != nil {
		SQLERROR.Println(getErrMes(err, id, "while tranforming to Album"))
	}
	song, err := apisf.ToSong()
	if err != nil {
		SQLERROR.Println(getErrMes(err, id, "while tranforming to Song"))
	}
	videos, err := apisf.ToVideos()
	if err != nil {
		SQLERROR.Println(getErrMes(err, id, "while tranforming to Video"))
	}

	return artist, album, song, videos, err

}

func printComponent(artist *sf.Artist, album *sf.Album, song *sf.Song, videos []*sf.Video, dumpMode dna.Bool) {
	if dumpMode == false {
		dna.Log("----------------------ARTIST----------------------")
		dna.LogStruct(artist)
		dna.Log("----------------------ALBUM----------------------")
		dna.LogStruct(album)
		dna.Log("----------------------SONG----------------------")
		dna.LogStruct(song)
		dna.Log("----------------------VIDEOS----------------------")
		for _, video := range videos {
			dna.LogStruct(video)
		}
	} else {
		// do nothing
	}

}

func saveComponents(id dna.Int, artist *sf.Artist, album *sf.Album, song *sf.Song, videos []*sf.Video) error {

	var err error = nil

	// songWriter.Write(song.CSVRecord())

	// if artist != nil {
	// 	artistWriter.Write(artist.CSVRecord())
	// }

	// if album != nil {
	// 	albumWriter.Write(album.CSVRecord())
	// }

	if videos != nil {
		for _, video := range videos {
			videoWriter.Write(video.CSVRecord())
		}
	}

	return err
}

func selectSQL(db *sqlpg.DB, id dna.Int) {
	songFreaks := &[]SongFreak{}
	intA := dna.IntArray{}
	for j := id; j < id+N_RECS_PER_SELECT; j++ {
		intA.Push(j)
	}
	query := dna.Sprintf("select * from sfsongfreaks where id in (%v)", intA.Join(","))
	// dna.Log(query)
	err := db.Select(songFreaks, query)
	if err != nil {
		SQLERROR.Println(getErrMes(err, id, " while selecting a record"))
	} else {
		// change counter
		length := dna.Int(len(*songFreaks))
		counter.Pass += length
		counter.Fail += N_RECS_PER_SELECT - length
		counter.Count += N_RECS_PER_SELECT
		counter.ElapsedTime = time.Since(counter.GetStartingTime())
		if counter.ElapsedTime/time.Second > 0 {
			counter.Speed = dna.Int(int64(counter.Count) / int64(counter.ElapsedTime/time.Second))
		}

		for _, songfreak := range *songFreaks {
			apisf, err := decodeSongfreakTrack(&songfreak, id)
			if err != nil {
				SQLERROR.Println(getErrMes(err, id, "while decoding SongFreakTrack"))
			} else {
				artist, album, song, videos, err := getComponents(apisf, id)
				if err != nil {
					panic("Error occurs at id:" + dna.Int(id).ToString().String())
				}
				if song.Lrc != "{}" {
					song.HasLrc = true
				}
				saveComponents(id, artist, album, song, videos)
				// printComponent(artist, album, song, videos, true)
			}
		}

	}
	cData := dna.Sprintf("%v | Cid:%v", counter, id)
	bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
}

func displayCSV() {
	file, err := os.Open("./song.csv")
	if err != nil {
		panic("cannot open")
	}

	var count = 0
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			dna.Log("Error:", err)
			return
		}
		count += 1
		dna.Log(count, len(record)) // record has the type []string
	}
}

func initWriters() {
	var err error
	// create files
	songFile, err = os.Create("./sfsongs.csv")
	if err != nil {
		panic("cannot create songFile")
	}
	artistFile, err = os.Create("./sfartists.csv")
	if err != nil {
		panic("cannot create artistFile")
	}
	albumFile, err = os.Create("./sfalbums.csv")
	if err != nil {
		panic("cannot create albumFile")
	}
	videoFile, err = os.Create("./sfvideos.csv")
	if err != nil {
		panic("cannot create videoFile")
	}

	// init writers
	songWriter = csv.NewWriter(songFile)
	artistWriter = csv.NewWriter(artistFile)
	albumWriter = csv.NewWriter(albumFile)
	videoWriter = csv.NewWriter(videoFile)
}

func terminateWriters() {
	// flush writers
	songWriter.Flush()
	artistWriter.Flush()
	albumWriter.Flush()
	videoWriter.Flush()

	// close writers
	artistFile.Close()
	songFile.Close()
	albumFile.Close()
	videoFile.Close()
}

func main() {
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(site.SqlConfigPath))
	dna.PanicError(err)
	bar = getUpdateProgressBar(4715353, "ALL")
	initWriters()
	for i := 1; i <= 4715353; i += N_RECS_PER_SELECT {
		selectSQL(db, dna.Int(i))
	}
	terminateWriters()

	// close DB
	db.Close()

}
