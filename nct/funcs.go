package nct

import (
	"dna"
	"dna/http"
	// "dna/item"
	// "dna/sqlpg"
	// "encoding/csv"
	"encoding/json"
	// "os"
)

func isSongFormat(str dna.String) dna.Bool {
	if str.EndsWith("mp3") == true || str.EndsWith("m4a") == true {
		return true
	} else {
		return false
	}
}

func isVideoFormat(str dna.String) dna.Bool {
	if str.EndsWith("mp4") == true || str.EndsWith("mpg") == true || str.EndsWith("flv") == true {
		return true
	} else {
		return false
	}
}

type apiSongList struct {
	Result  dna.Bool  `json:"Result"`
	HasMore dna.Bool  `json:"IsMore"`
	Data    []APISong `json:"Data"`
}

// GetAPISongsByArtist returns a list of songs, HasMore songs form an artist
// from artistid, pageIdex and pageSize.
// If artist id is not available, then the hot song lists will be returned.
func GetAPISongsByArtist(id, pageIndex, pageSize dna.Int) ([]APISong, dna.Bool, error) {
	urlb := NewURLBuilder()
	link := urlb.GetSongsByArtist(id, pageIndex, pageSize)
	result, err := http.Get(link)
	if err == nil {
		var asongList = &apiSongList{}
		errd := json.Unmarshal(result.Data.ToBytes(), asongList)
		if errd == nil {
			return asongList.Data, asongList.HasMore, nil
		} else {
			return nil, false, errd
		}
	} else {
		return nil, false, err
	}
}

// All the part below is for testing
//
//	psql -c "COPY nctsongs (id,key,title,artists,topics,link_key,type,bitrate,official,likes,plays,link_share,stream_url,image,coverart,duration,linkdown,linkdown_hq,lyricid,has_lyric,lyric,lyric_status,has_lrc,lrc,lrc_url,username_created,checktime) FROM '$GOPATH/nctsongs.csv' DELIMITER ',' CSV"
// var (
// 	SongWriter *csv.Writer
// 	SongFile   *os.File

// 	ArtistWriter *csv.Writer
// 	ArtistFile   *os.File
// 	PageSize     dna.Int = 30
// )

// func InitWriters() {
// 	SongFile, err := os.Create("./nctsongs.csv")
// 	if err != nil {
// 		panic("cannot create songFile")
// 	}
// 	SongWriter = csv.NewWriter(SongFile)

// 	ArtistFile, err = os.Create("./nctartists.csv")
// 	if err != nil {
// 		panic("cannot create songFile")
// 	}
// 	ArtistWriter = csv.NewWriter(ArtistFile)
// }

// func TerminateWriters() {
// 	SongWriter.Flush()
// 	SongFile.Close()

// 	ArtistWriter.Flush()
// 	ArtistFile.Close()
// }

// type APISongByArtist struct {
// 	Id     dna.Int // artist id
// 	Nsongs dna.Int
// }

// func NewAPISongByArtist() *APISongByArtist {
// 	sba := new(APISongByArtist)
// 	sba.Id = 0
// 	sba.Nsongs = 0
// 	return sba
// }

// func getAllAPISongsByArtist(apiSongBA *APISongByArtist, pageIndex dna.Int) {
// 	apiSongs, hasMore, err := GetAPISongsByArtist(apiSongBA.Id, pageIndex, PageSize)
// 	if err == nil {
// 		mutex.Lock()
// 		for _, apisong := range apiSongs {
// 			song := NewSong()
// 			apisong.FillSong(song)
// 			SongWriter.Write(song.CSVRecord())
// 		}
// 		apiSongBA.Nsongs += dna.Int(len(apiSongs))
// 		SongWriter.Flush()
// 		mutex.Unlock()
// 	} else {
// 		dna.Log("ERROR FROM FUNC:getAllAPISongsByArtist")
// 		dna.Log(err.Error())
// 	}
// 	if hasMore == true {
// 		getAllAPISongsByArtist(apiSongBA, pageIndex+1)
// 	}

// }

// //psql -c "COPY nctartists (id,name,avatar,nsongs,nalbums,nvideos,obj_type) FROM '$GOPATH/nctartists.csv' DELIMITER ',' CSV"
// func (apiSongBA *APISongByArtist) Fetch() error {
// 	// If artist id is not available, then the hot song lists will be returned.
// 	apiArtist, err := GetAPIArtist(apiSongBA.Id)
// 	if err == nil {
// 		getAllAPISongsByArtist(apiSongBA, 1)
// 		mutex.Lock()
// 		apiArtist.NSongs = apiSongBA.Nsongs
// 		ArtistWriter.Write(apiArtist.CSVRecord())
// 		ArtistWriter.Flush()
// 		mutex.Unlock()
// 	}

// 	return nil
// }

// func (apiSongBA *APISongByArtist) GetId() dna.Int {
// 	return apiSongBA.Id
// }

// func (apiSongBA *APISongByArtist) New() item.Item {
// 	return item.Item(NewAPISongByArtist())
// }

// func (apiSongBA *APISongByArtist) Init(v interface{}) {
// 	switch v.(type) {
// 	case int:
// 		apiSongBA.Id = dna.Int(v.(int))
// 	case dna.Int:
// 		apiSongBA.Id = v.(dna.Int)
// 	default:
// 		panic("Interface v has to be int")
// 	}
// }

// func (apiSongBA *APISongByArtist) Save(db *sqlpg.DB) error {
// 	return nil
// }
