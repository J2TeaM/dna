package ns

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
)

// transformCats turns genre name into multiple genres if applicable
func transformCats(cats dna.StringArray) dna.StringArray {
	ret := dna.StringArray{}
	vnSongs := dna.StringArray{"Nhạc Trẻ", "Nhạc Trữ Tình", "Nhạc Cách Mạng", "Nhạc Trịnh", "Nhạc Tiền Chiến", "Nhạc Dân Tộc", "Nhạc Thiếu Nhi", "Rock Việt", "Nhạc Hải Ngoại", "Nhạc Quê Hương", "Rap Việt - Hiphop"}
	for _, cat := range cats {
		if vnSongs.IndexOf(cat) > -1 {
			ret.Push("Nhạc Việt Nam")
		}
		switch cat {
		case "Pop/Ballad":
			ret.Push("Pop")
			ret.Push("Ballad")
		case "Dance/Electronic":
			ret.Push("Dance")
			ret.Push("Electronic")
		case "Nhạc Spa | Thư Giãn":
			ret.Push("Nhạc Spa")
			ret.Push("Thư Giãn")
		case "Hiphop/Rap":
			ret.Push("Hiphop")
			ret.Push("Rap")
		case "Nhạc Bà Bầu & Baby":
			ret.Push("Nhạc Bà Bầu")
			ret.Push("Nhạc Baby")
		case "Rap Việt - Hiphop":
			ret.Push("Rap Việt")
			ret.Push("Hiphop")
		case "Radio - Cảm Xúc":
			ret.Push("Radio")
			ret.Push("Cảm Xúc")
		default:
			ret.Push(cat)
		}
	}
	return ret
}

func getCategory(songs *[]*Song, genre Genre, page dna.Int) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/bai-hat-theo-the-loai-" + genre.Id.ToString() + "/joke-link-2-" + page.ToString() + ".html"
		// dna.Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data

			// transform string {"2":[0,3,5,7,9,11,13,15,29],"10":[1,2,4,6,8,]}
			// to  map[dna.Int]dna.Int{20:2, 28:2, 4:10, 12:10} Ex: map[29] = 2
			temp := data.FindAllStringSubmatch(`getCategory.+'(\{.+\})'`, -1)
			mapping := map[dna.Int]dna.Int{}
			if len(temp) > 0 && temp[0].Length() > 0 {
				vals := temp[0][1].FindAllString(`"[0-9]+":\[[0-9,]+?\]`, -1)
				if vals.Length() > 0 {
					for _, val := range vals {
						target := val.FindAllStringSubmatch(`"(.+)"`, -1)[0][1].ToInt()
						arr := val.FindAllStringSubmatch(`\[(.+)\]`, -1)[0][1].Split(",").ToIntArray()
						for _, itm := range arr {
							mapping[itm] = target
						}
					}
				}
			}
			// Finding cat id for each song. cats var is 2-dimentional array.
			// Each index of it represents the correspondent song, its value is the categories the song belongs to
			catStrings := data.FindAllString(`Thể loại :.+`, -1)
			cats := []dna.IntArray{}
			for _, val := range catStrings {
				tagids := dna.IntArray{}
				tmp := val.FindAllStringSubmatch(`cate_tag_song_([0-9]+)`, -1)
				if len(tmp) > 0 {
					for _, el := range tmp {
						tagids.Push(el[1].ToInt())
					}
				}
				cats = append(cats, tagids)
			}
			// Log(cats)

			// get songids
			temps := data.FindAllStringSubmatch(`play" id="blocksongtag_([0-9]+)`, -1)
			songids := dna.IntArray{}
			if len(temps) > 0 {
				for _, val := range temps {
					songids.Push(val[1].ToInt())
				}
			}

			tmpsongs := &[]*Song{}
			for idx, songid := range songids {
				song := NewSong()
				song.Id = songid
				category := dna.StringArray{}
				for _, val := range cats[idx] {
					if mapping[val] > 0 && mapping[val] < CatTags.Length() {
						if CatTags[mapping[val]] != "" {
							category.Push(CatTags[mapping[val]])
						}
					} else {
						mess := dna.Sprintf("WRONG INDEX AT CATTAGS: %v %v %v - %v", mapping[val], genre, page, link)
						panic(mess.String())
					}

				}
				category.Push(genre.Name)
				song.Category = transformCats(category.Unique()).Unique()
				*tmpsongs = append(*tmpsongs, song)
			}
			*songs = *tmpsongs

		}
		channel <- true

	}()
	return channel
}

func GetSongCategory(genre Genre, page dna.Int) (*SongCategory, error) {
	songs := &[]*Song{}
	c := make(chan bool, 1)

	go func() {
		c <- <-getCategory(songs, genre, page)
	}()

	for i := 0; i < 1; i++ {
		<-c
		// for _, s := range *songs {
		// 	Log(s.Id, "-", s.Category)
		// }
	}

	soca := new(SongCategory)
	soca.Genre = genre
	soca.Page = page
	soca.Songs = songs

	return soca, nil

}

type SongCategory struct {
	Genre Genre
	Page  dna.Int
	Songs *[]*Song
}

func NewSongCategory() *SongCategory {
	soca := new(SongCategory)
	soca.Genre = Genre{0, "", 0}
	soca.Page = 0
	soca.Songs = nil
	return soca
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (soca *SongCategory) Fetch() error {
	_soca, err := GetSongCategory(soca.Genre, soca.Page)
	*soca = *_soca
	// not complete
	return err
}

// GetId implements GetId methods of item.Item interface
// GetId always return zero
func (soca *SongCategory) GetId() dna.Int {
	return 0
}

// New implements item.Item interface
// Returns new item.Item interface
func (soca *SongCategory) New() item.Item {
	return item.Item(NewSongCategory())
}

// Init implements item.Item interface.
func (soca *SongCategory) Init(v interface{}) {
	var n dna.Int

	var NGenres dna.Int = dna.Int(len((*SongGenreList))) // The total of genres
	switch v.(type) {
	case int:
		n = dna.Int(v.(int))
	case dna.Int:
		n = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
	genreIndex := dna.Int(n / LastNPages)
	if genreIndex >= NGenres {
		genreIndex = NGenres - 1
	}
	soca.Genre = (*SongGenreList)[genreIndex]
	soca.Page = n%LastNPages + 1
}

func (soca *SongCategory) Save(db *sqlpg.DB) error {

	var last error
	for _, song := range *(soca.Songs) {
		// dna.Log(song)
		last = db.Update(song, "id", "category")
	}
	return last
}
