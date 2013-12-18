package ns

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
)

func getAlbumCategory(albums *[]*Album, genre Genre, page dna.Int) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://nhacso.net/album-theo-the-loai-" + genre.Id.ToString() + "/joke-link-2-" + page.ToString() + ".html"
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			temp := data.FindAllStringSubmatch(`getTotalSongInAlbum\('(.+)', 'album_new_totalsong_'`, -1)
			tmpalbums := &[]*Album{}
			if len(temp) > 0 {
				albumList := temp[0][1].Split(",").ToIntArray()
				for _, albumid := range albumList {
					album := NewAlbum()
					album.Id = albumid
					cats := dna.StringArray{genre.Name}
					album.Category = transformCats(cats)
					*tmpalbums = append(*tmpalbums, album)
				}
			}
			*albums = *tmpalbums

		}
		channel <- true

	}()
	return channel
}

func GetAlbumCategory(genre Genre, page dna.Int) (*AlbumCategory, error) {
	albums := &[]*Album{}
	c := make(chan bool, 1)

	go func() {
		c <- <-getAlbumCategory(albums, genre, page)
	}()

	for i := 0; i < 1; i++ {
		<-c
		// for _, s := range *albums {
		// 	Log(s.Id, "-", s.Category)
		// }
	}

	alca := new(AlbumCategory)
	alca.Genre = genre
	alca.Page = page
	alca.Albums = albums

	return alca, nil

}

type AlbumCategory struct {
	Genre  Genre
	Page   dna.Int
	Albums *[]*Album
}

func NewAlbumCategory() *AlbumCategory {
	alca := new(AlbumCategory)
	alca.Genre = Genre{0, "", 0}
	alca.Page = 0
	alca.Albums = nil
	return alca
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (alca *AlbumCategory) Fetch() error {
	_alca, err := GetAlbumCategory(alca.Genre, alca.Page)
	*alca = *_alca
	// not complete
	return err
}

// GetId implements GetId methods of item.Item interface
// GetId always return zero
func (alca *AlbumCategory) GetId() dna.Int {
	return 0
}

// New implements item.Item interface
// Returns new item.Item interface
func (alca *AlbumCategory) New() item.Item {
	return item.Item(NewAlbumCategory())
}

// Init implements item.Item interface.
func (alca *AlbumCategory) Init(v interface{}) {
	var n dna.Int
	var LastNPages dna.Int = 10                          //  Last 10 pages from each category
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
	alca.Genre = (*SongGenreList)[genreIndex]
	alca.Page = n%LastNPages + 1
}

func (alca *AlbumCategory) Save(db *sqlpg.DB) error {

	var last error
	var aids = dna.IntArray{}
	albums := &[]Album{}

	for _, album := range *(alca.Albums) {
		aids.Push(album.Id)
	}
	query := "SELECT id, topics, genres from nsalbums WHERE id IN (" + aids.Join(",") + ")"
	// dna.Log(query)
	err := db.Select(albums, query)
	dna.PanicError(err)
	for _, album := range *(alca.Albums) {
		foundIndex := 0
		for j, anotherAlbum := range *(albums) {
			if album.Id == anotherAlbum.Id {
				foundIndex = j
			}
		}
		cat := album.Category.Concat((*albums)[foundIndex].Topics).Concat((*albums)[foundIndex].Genres).Unique()
		album.Category = cat.Filter(func(v dna.String, i dna.Int) dna.Bool {
			if v != "" {
				return true
			} else {
				return false
			}
		})
		last = db.Update(album, "id", "category")
	}

	return last
}
