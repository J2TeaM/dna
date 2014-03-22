package media

import (
	"dna"
	"dna/cfg"
	"dna/sqlpg"
	"dna/utils"
	"time"
)

const (
	FArtist = 1 << iota
	FAlbum
	FSong
	FVideo
)

type RefTableMap map[dna.String]dna.Int

func (rtm RefTableMap) HasArtist(key dna.String) dna.Bool {
	val, ok := rtm[key]
	if ok {
		if val&1 == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func (rtm RefTableMap) HasAlbum(key dna.String) dna.Bool {
	val, ok := rtm[key]
	if ok {
		if (val>>1)&1 == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func (rtm RefTableMap) HasSong(key dna.String) dna.Bool {
	val, ok := rtm[key]
	if ok {
		if (val>>2)&1 == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func (rtm RefTableMap) HasVideo(key dna.String) dna.Bool {
	val, ok := rtm[key]
	if ok {
		if (val>>3)&1 == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

var RefTables = RefTableMap{
	"cc":  FAlbum | FSong | FVideo,
	"csn": FSong | FVideo, // Disable FAlbum
	"ke":  FArtist | FAlbum | FSong | FVideo,
	"mv":  FSong | FVideo,
	"nct": FArtist | FAlbum | FSong | FVideo,
	"ns":  FAlbum | FSong | FVideo,
	"nv":  FAlbum | FSong | FVideo,
	"vg":  FSong,
	"zi":  FArtist | FAlbum | FSong | FVideo,
}

func GetArtistTables() dna.StringArray {
	var tblArr = dna.StringArray{}
	for key, _ := range RefTables {
		if RefTables.HasArtist(key) {
			tblArr.Push(key + "artists")
		}
	}
	return tblArr
}

func GetAlbumTables() dna.StringArray {
	var tblArr = dna.StringArray{}
	for key, _ := range RefTables {
		if RefTables.HasAlbum(key) {
			tblArr.Push(key + "albums")
		}
	}
	return tblArr
}

func GetSongTables() dna.StringArray {
	var tblArr = dna.StringArray{}
	for key, _ := range RefTables {
		if RefTables.HasSong(key) {
			tblArr.Push(key + "songs")
		}
	}
	return tblArr
}

func GetVideoTables() dna.StringArray {
	var tblArr = dna.StringArray{}
	for key, _ := range RefTables {
		if RefTables.HasVideo(key) {
			tblArr.Push(key + "videos")
		}
	}
	return tblArr
}

// DumpHashTables uses psql command. Set its path to ENV
func DumpHashTables() {
	var commands = dna.StringArray{}
	for index, table := range GetAlbumTables().Concat(GetSongTables()).Concat(GetVideoTables()) {
		shortForm := table.Replace("songs", "").Replace("albums", "").Replace("videos", "")
		stmt := dna.Sprintf(`SELECT dna_hash(title,artists), ROW(%v,id) from %v`, ToSiteid(shortForm), table)
		command := dna.Sprintf(`psql -c 'COPY (%v) TO STDOUT'`, stmt)
		switch {
		case table.Match("albums") == true:
			command += " >> data/hash_albums.log"
		case table.Match("songs") == true:
			command += " >> data/hash_songs.log"
		case table.Match("videos") == true:
			command += " >> data/hash_videos.log"
		}
		commands.Push(`INTERNAL_TIME=$(date +%s)`)
		commands.Push(dna.Sprintf(`printf "%-3v:Extracting hashids from %-12v "`, index+1, table+"..."))
		commands.Push(command)
		commands.Push(`echo "Completed in $(($(date +%s) - $INTERNAL_TIME)) seconds!"`)
		commands.Push("#--------------------------")
	}
	dna.Log(commands.Join("\n"))
}

func DumpFiles() {
	var ret = dna.StringArray{}
	for _, table := range GetAlbumTables().Concat(GetSongTables()).Concat(GetVideoTables()) {
		ret.Push(`"` + table + `"`)
	}
	dna.Log(ret.Join(","))
}

type site_checktime struct {
	Site      dna.String
	Checktime time.Time
}

// GetLastedChecktime returns a map which maps site to lasted checktime.
// Ex: "nssongs" => 2014-03-17 12:09:37
func GetLastedChecktime(db *sqlpg.DB) (map[dna.String]time.Time, error) {
	siteCts := &[]site_checktime{}
	ret := make(map[dna.String]time.Time)
	err := db.Select(siteCts, "select * from get_lasted_checktime();")
	if err == nil {
		for _, sitect := range *siteCts {
			ret[sitect.Site] = sitect.Checktime
			// dna.Log(sitect.Site, sitect.Checktime.Format(utils.DefaultTimeLayout))

		}
		return ret, nil
	} else {
		return nil, err
	}
}

func SaveLastedChecktime(db *sqlpg.DB, filePath dna.String) error {
	cf, err := cfg.LoadConfigFile(filePath)
	if err != nil {
		return err
	}
	siteCts, err := GetLastedChecktime(db)
	if err != nil {
		return err
	}
	for site, checktime := range siteCts {
		var key dna.String
		var section dna.String
		switch {
		case site.Match("songs") == true:
			key = "songs"
			section = site.Replace("songs", "")
		case site.Match("albums") == true:
			key = "albums"
			section = site.Replace("albums", "")
		case site.Match("videos") == true:
			key = "videos"
			section = site.Replace("videos", "")
		default:
			panic("site is not valid")
		}
		cf.SetValue(section, key, dna.String(checktime.Format(utils.DefaultTimeLayout)))
	}
	return cfg.SaveConfigFile(cf, filePath)
}

func LoadLastedChecktime(filePath dna.String) (map[dna.String]time.Time, error) {
	ret := make(map[dna.String]time.Time)
	cf, err := cfg.LoadConfigFile(filePath)
	if err != nil {
		return nil, err
	}
	for site, _ := range RefTables {
		section, err := cf.GetSection(site)
		if err != nil {
			dna.PanicError(err)
		}
		for postfix, timeStr := range section {
			if timeStr != "" {
				t, err := time.Parse(utils.DefaultTimeLayout, timeStr.String())
				if err == nil {
					ret[site+postfix] = t
				} else {
					dna.PanicError(err)
				}
			}
		}
	}
	return ret, nil
}
