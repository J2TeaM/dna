package nct

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"dna/utils"
	"errors"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

// SongsInAlbums stores all song portions found in albums
var SongsInAlbums = &Portions{}

type Album struct {
	Id          dna.Int
	Key         dna.String
	Title       dna.String
	Artists     dna.StringArray
	Topics      dna.StringArray
	Plays       dna.Int
	Songids     dna.IntArray
	Nsongs      dna.Int
	Description dna.String
	Coverart    dna.String
	LinkKey     dna.String
	DateCreated time.Time
	Checktime   time.Time
}

func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Key = ""
	album.Title = ""
	album.Artists = dna.StringArray{}
	album.Topics = dna.StringArray{}
	album.Plays = 0
	album.Songids = dna.IntArray{}
	album.Nsongs = 0
	album.Description = ""
	album.Coverart = ""
	album.LinkKey = ""
	album.DateCreated = time.Time{}
	album.Checktime = time.Time{}
	return album
}

// getAlbumPlays returns song plays
func getAlbumPlays(album *Album, body dna.String) {
	link := "http://www.nhaccuatui.com/interaction/api/hit-counter?jsoncallback=nct"
	http.DefaulHeader.Set("Content-Type", "application/x-www-form-urlencoded ")
	result, err := http.Post(dna.String(link), body)
	// Log(link)
	if err == nil {
		data := &result.Data
		tpl := dna.String(`{"counter":([0-9]+)}`)
		playsArr := data.FindAllStringSubmatch(tpl, -1)
		if len(playsArr) > 0 {
			album.Plays = playsArr[0][1].ToInt()
		}
	}
}

// getAlbumFromMainPage returns album from main page
func getAlbumFromMainPage(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.nhaccuatui.com/playlist/google-bot." + album.Key + ".html"
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data

			idArr := data.FindAllStringSubmatch(`value="(.+)" id="inpHiddenId"`, 1)
			if len(idArr) > 0 {
				album.Id = idArr[0][1].ToInt()
			}

			topicsArr := data.FindAllStringSubmatch(`<strong>Thể loại</strong></p>[\n\t\r]+(.+)`, 1)
			if len(topicsArr) > 0 {
				album.Topics = topicsArr[0][1].RemoveHtmlTags("").Trim().Split(", ")
			}

			nsongArr := data.FindAllStringSubmatch(`<span class="tag black">(.+)bài hát</span>`, 1)
			if len(nsongArr) > 0 {
				album.Nsongs = nsongArr[0][1].Trim().ToInt()
			}

			linkkeyArr := data.FindAllStringSubmatch(`"flashPlayer", "playlist", "(.+?)"`, 1)
			if len(linkkeyArr) > 0 {
				album.LinkKey = linkkeyArr[0][1].Trim()
			}

			titleArr := data.FindAllStringSubmatch(`<h1>(.+?)</h1>`, 1)
			if len(titleArr) > 0 {
				album.Title = titleArr[0][1].Trim().SplitWithRegexp(" - ", 2)[0].Trim()
			}

			artistsArr := data.FindAllStringSubmatch(`<h1>(.+?)</h1>`, 1)
			if len(artistsArr) > 0 {
				artists := artistsArr[0][1].RemoveHtmlTags("").SplitWithRegexp(" - ", 2)
				if artists.Length() == 2 {
					album.Artists = artists[1].Split(", ").Filter(func(v dna.String, i dna.Int) dna.Bool {
						if v != "" {
							return true
						} else {
							return false
						}
					})
				}
			}

			descArr := data.FindAllString(`(?mis)<p id="shortPlDesc">.+?</p>`, 1)
			if descArr.Length() > 0 {
				album.Description = descArr[0].RemoveHtmlTags("").Trim()
			}

			coverartArr := data.FindAllString(`<meta property="og:image".+`, 1)
			if coverartArr.Length() > 0 {
				album.Coverart = coverartArr[0].GetTagAttributes("content")
				datecreatedArr := album.Coverart.FindAllStringSubmatch(`/([0-9]+)\..+$`, -1)
				if len(datecreatedArr) > 0 {
					// Log(int64(datecreatedArr[0][1].ToInt()))
					album.DateCreated = time.Unix(int64(datecreatedArr[0][1].ToInt()/1000), 0)
				} else {
					dateCreatedArr := album.Coverart.FindAllStringSubmatch(`/?(\d{4}/\d{2}/\d{2})`, -1)
					if len(dateCreatedArr) > 0 {
						year := dateCreatedArr[0][1].FindAllStringSubmatch(`(\d{4})/\d{2}/\d{2}`, -1)[0][1].ToInt()
						month := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/(\d{2})/\d{2}`, -1)[0][1].ToInt()
						day := dateCreatedArr[0][1].FindAllStringSubmatch(`\d{4}/\d{2}/(\d{2})`, -1)[0][1].ToInt()
						album.DateCreated = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

					}
				}
			}

			// Find params for the number of album plays
			itemIdArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('(.+?)'.+`, 1)
			timeArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('.+?', '(.+?)'.+\);`, 1)
			signArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('.+?', '.+?', '(.+?)'.+;`, 1)
			typeArr := data.FindAllStringSubmatch(`NCTWidget.hitCounter\('.+?', '.+?', '.+?', "(.+?)"\);`, 1)
			if len(itemIdArr) > 0 && len(timeArr) > 0 && len(signArr) > 0 && len(typeArr) > 0 {
				// boday has post form:
				// item_id=2870710&time=1389009424631&sign=2499ab08f6662842a02b06aad603d8ab&type=playlist
				body := dna.Sprintf(`item_id=%v&time=%v&sign=%v&type=%v`, itemIdArr[0][1], timeArr[0][1], signArr[0][1], typeArr[0][1])
				getAlbumPlays(album, body)
				album.Id = itemIdArr[0][1].ToInt()
			}

			songidsArr := data.FindAllString(`<a href="javascript:;" class="button_download".+`, -1)
			songkeysArr := data.FindAllString(`<a href="javascript:;" class="button_add_playlist".+`, -1)
			if songidsArr.Length() == songkeysArr.Length() {
				// Getting error when album key is "rwwC6U1wow3X"
				// panic(`Song ids and keys do not match. Album key: ` + album.Key.String())
				songidsArr.ForEach(func(val dna.String, idx dna.Int) {
					tmp := val.FindAllStringSubmatch(`_([0-9]+)`, 1)
					if len(tmp) > 0 {
						id := tmp[0][1].Trim().ToInt()
						album.Songids.Push(id)
						portion := NewPortion()
						portion.Id = int32(id)
						tmpKeys := songkeysArr[idx].FindAllStringSubmatch(`btnShowBoxPlaylist_([a-zA-Z0-9]+)"`, 1)
						if len(tmpKeys) > 0 {
							portion.Key = string(tmpKeys[0][1])
							SongsInAlbums.Push(portion)
						}

					}
				})
			}

			GetRelevantPortions(&result.Data)
		}
		channel <- true
	}()
	return channel
}

// GetAlbum returns a album or an error
// 	* key: A unique key of a album
// 	* Official : 0 or 1, if its value is unknown, set to 0
// 	* Returns a found album or an error
func GetAlbum(key dna.String) (*Album, error) {
	var album *Album = NewAlbum()
	album.Key = key
	c := make(chan bool, 1)
	go func() {
		c <- <-getAlbumFromMainPage(album)
	}()
	for i := 0; i < 1; i++ {
		<-c
	}
	if album.Nsongs != album.Songids.Length() {
		return nil, errors.New(dna.Sprintf("Nhaccuatui - Album %v: - Key: %v Songids and Nsongs do not match", album.Id, album.Key).String())
	} else if album.Nsongs == 0 && album.Songids.Length() == 0 {
		return nil, errors.New(dna.Sprintf("Nhaccuatui - Album %v: - Key: %v No song found", album.Id, album.Key).String())
	} else {
		album.Checktime = time.Now()
		return album, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (album *Album) Fetch() error {
	_album, err := GetAlbum(album.Key)
	if err != nil {
		return err
	} else {
		*album = *_album
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (album *Album) GetId() dna.Int {
	return album.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (album *Album) New() item.Item {
	return item.Item(NewAlbum())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (album *Album) Init(v interface{}) {
	switch v.(type) {
	case int:
		idx := dna.Int(v.(int))
		length := (*NewestAlbumPortions).Length()
		if idx >= length {
			idx = length - 1
		}
		if length > 0 {
			album.Key = dna.String((*NewestAlbumPortions)[idx].Key)
		}
	case dna.Int:
		idx := v.(dna.Int)
		length := (*NewestAlbumPortions).Length()
		if idx >= length {
			idx = length - 1
		}
		if length > 0 {
			album.Key = dna.String((*NewestAlbumPortions)[idx].Key)
		}

	default:
		panic("Interface v has to be int")
	}
}

func (album *Album) Save(db *sqlpg.DB) error {
	filterRelevants(db)
	ids, err := utils.SelectMissingIds("nctsongs", &album.Songids, db)
	portions := &Portions{}
	// potential data race
	mutex.Lock()
	if err == nil {
		for _, portion := range *SongsInAlbums {
			// 1st condition: if portion id is in ids result. It means that portion is new
			// then it has to be added to SongsInAlbums
			//
			// 2nd condition: if portion id is not in abum songids,
			// it means it has no relation to select statement, add it back
			if ids.IndexOf(dna.Int(portion.Id)) > -1 || album.Songids.IndexOf(dna.Int(portion.Id)) == -1 {
				portions.Push(&portion)
			}
		}
	}
	*SongsInAlbums = *portions
	mutex.Unlock()
	return db.InsertIgnore(album)
}
