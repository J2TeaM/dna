package hdv

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"errors"
	"time"
)

type Movie struct {
	Id            dna.Int
	Title         dna.String
	AnotherTitle  dna.String
	ForeignTitle  dna.String
	VnTitle       dna.String
	Topics        dna.StringArray
	Actors        dna.StringArray
	Directors     dna.StringArray
	Countries     dna.StringArray
	Description   dna.String
	YearReleased  dna.Int
	IMDBRating    dna.IntArray
	Similars      dna.IntArray
	Thumbnail     dna.String
	MaxResolution dna.Int
	IsSeries      dna.Bool
	SeasonId      dna.Int
	Seasons       dna.IntArray
	Epid          dna.Int
	CurrentEps    dna.Int
	MaxEp         dna.Int
	Checktime     time.Time
}

func NewMovie() *Movie {
	movie := new(Movie)
	movie.Id = 0
	movie.Title = ""
	movie.AnotherTitle = ""
	movie.ForeignTitle = ""
	movie.VnTitle = ""
	movie.Topics = dna.StringArray{}
	movie.Actors = dna.StringArray{}
	movie.Directors = dna.StringArray{}
	movie.Countries = dna.StringArray{}
	movie.Description = ""
	movie.YearReleased = 0
	movie.IMDBRating = dna.IntArray{}
	movie.Similars = dna.IntArray{}
	movie.Thumbnail = ""
	movie.MaxResolution = 0
	movie.IsSeries = false
	movie.SeasonId = 0
	movie.Seasons = dna.IntArray{}
	movie.Epid = 0
	movie.CurrentEps = 0
	movie.MaxEp = 0
	movie.Checktime = time.Now()
	return movie
}

func getAnchorTagsData(data dna.String) dna.StringArray {
	return dna.StringArray(data.Split("|").Map(func(val dna.String, idx dna.Int) dna.String {
		return val.RemoveHtmlTags("").Trim()
	}).([]dna.String)).Filter(func(val dna.String, idx dna.Int) dna.Bool {
		if val != "" {
			return true
		} else {
			return false
		}
	})
}

func getNames(data, from, to dna.String) dna.StringArray {
	pattern := "(?mis)" + from + "(.+?)" + to
	arr := data.FindAllStringSubmatch(pattern, 1)
	if len(arr) > 0 {
		return getAnchorTagsData(arr[0][1])
	} else {
		return dna.StringArray{}
	}
}
func getMovieFromPage(movie *Movie) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://mmovie.hdviet.com/hdviet." + movie.Id.ToString() + ".html"
		result, err := http.Get(link)
		if err == nil {
			// Renew ACCESS_TOKEN_KEY
			if AccessTokenKeyRenewable == false {
				keyArr := result.Data.FindAllString(`_strLinkPlay.+`, 1)
				if keyArr.Length() > 0 {
					ACCESS_TOKEN_KEY = keyArr[0].Replace(`';`, "").ReplaceWithRegexp(`^.+;`, "").String()
					if len(ACCESS_TOKEN_KEY) != 32 {
						panic("Wrong ACCESS_TOKEN_KEY: Length has to be 32")
					}
				}
				AccessTokenKeyRenewable = true
			}

			// Finding main data
			mainDat := result.Data.FindAllString(`(?mis)<div class="main_decs">.+?<div class="clear">`, 1)
			if len(mainDat) > 0 {
				main := mainDat[0]
				if main.GetTags("h4").Length() > 0 {
					movie.Title = main.GetTags("h4")[0].RemoveHtmlTags("").DecodeHTML().Trim()
					switch movie.Title.Split(" - ").Length() {
					case 2:
						movie.ForeignTitle = movie.Title.Split(" - ")[0]
						movie.VnTitle = movie.Title.Split(" - ")[1]
					case 3:
						movie.ForeignTitle = movie.Title.Split(" - ")[0]
						movie.VnTitle = movie.Title.Split(" - ")[1]
						third := movie.Title.Split(" - ")[2]
						if third.Match(`Tập`) == true {
							movie.IsSeries = true
							movie.CurrentEps = third.Replace("Tập", "").Trim().Split("/")[0].ToInt()
							movie.MaxEp = third.Replace("Tập", "").Trim().Split("/")[1].ToInt()
							seasonArr := movie.VnTitle.FindAllStringSubmatch(`Phần ([0-9]+)`, 1)
							if len(seasonArr) > 0 {
								movie.SeasonId = seasonArr[0][1].ToInt()
							}
						}
					}

					movie.Topics = getNames(main, "Danh mục:", "Diễn viên")
					movie.Actors = getNames(main, "Diễn viên:", "Đạo diễn")
					movie.Directors = getNames(main, "Đạo diễn:", "<p>Quốc gia")
					movie.Countries = getNames(main, "<p>Quốc gia:", "Năm khởi chiếu")

					descArr := main.FindAllStringSubmatch(`(?mis)</h4>(.+?)Danh mục:`, 1)
					if len(descArr) > 0 {
						movie.Description = descArr[0][1].RemoveHtmlTags("").DecodeHTML().Trim()
					}

					yearArr := main.FindAllStringSubmatch(`(?mis)Năm khởi chiếu:(.+?)</p>`, 1)
					if len(yearArr) > 0 {
						movie.YearReleased = yearArr[0][1].RemoveHtmlTags("").DecodeHTML().Trim().ToInt()
					}

					ratingArr := main.FindAllStringSubmatch(`(?mis)Đánh giá IMDB:(.+?)<p>`, 1)
					if len(ratingArr) > 0 {
						if ratingArr[0][1].GetTags("span").Length() > 0 {
							rating := ratingArr[0][1].GetTags("span")[0].RemoveHtmlTags("").Replace(".", "").ToInt()
							movie.IMDBRating.Push(rating)
						}
						nvotes := ratingArr[0][1].FindAllStringSubmatch(`\((.+) phiếu bình chọn`, 1)
						if len(nvotes) > 0 {
							movie.IMDBRating.Push(nvotes[0][1].Replace(",", "").ToInt())
						}
					}
				}
			}

			seasonsArr := result.Data.FindAllStringSubmatch(`(?mis)Season:(.+?)<div class="main_decs">`, 1)

			if len(seasonsArr) > 0 {
				movie.Seasons = dna.IntArray(seasonsArr[0][1].Split(`</li>`).Map(func(val dna.String, idx dna.Int) dna.Int {
					idArr := val.GetTagAttributes("href").FindAllStringSubmatch(`/hdviet\.(.+)\.html`, 1)
					if len(idArr) > 0 {
						return idArr[0][1].ToInt()
					} else {
						return 0
					}
				}).([]dna.Int)).Filter(func(val dna.Int, idx dna.Int) dna.Bool {
					if val != 0 {
						return true
					} else {
						return false
					}
				})
			}

			anotherTitleArr := result.Data.FindAllStringSubmatch(`var _moviename = '(.+)'`, 1)
			if len(anotherTitleArr) > 0 {
				movie.AnotherTitle = anotherTitleArr[0][1].Trim()
			}

			similarDat := result.Data.FindAllStringSubmatch(`(?mis)<div class="other-films">(.+?)</div>`, 1)
			if len(similarDat) > 0 {
				movie.Similars = dna.IntArray(similarDat[0][1].FindAllString(`<a.+</a>`, -1).Map(func(val dna.String, idx dna.Int) dna.Int {
					href := val.GetTagAttributes("href")
					midArr := href.FindAllStringSubmatch(`\.([0-9]+)\.html`, 1)
					if len(midArr) > 0 {
						return midArr[0][1].ToInt()
					} else {
						return 0
					}
				}).([]dna.Int))
			}

			thumbArr := result.Data.FindAllString(`<div class="fd-poster">[\n\t\r]+.+`, 1)
			if thumbArr.Length() > 0 {
				movie.Thumbnail = thumbArr[0].GetTagAttributes("src")
			}

			if result.Data.Match(`http://mmovie.hdviet.com/images/720.png`) == true {
				movie.MaxResolution = 720
			}
			if result.Data.Match(`http://mmovie.hdviet.com/images/1080.png`) == true {
				movie.MaxResolution = 1080
			}

		}
		channel <- true

	}()
	return channel
}

// GetMovie returns a movie or an error.
func GetMovie(id dna.Int) (*Movie, error) {
	var movie *Movie = NewMovie()
	movie.Id = id
	c := make(chan bool, 1)

	go func() {
		c <- <-getMovieFromPage(movie)
	}()

	for i := 0; i < 1; i++ {
		<-c
	}

	if movie.Title == "" && movie.AnotherTitle == "" {
		return nil, errors.New(dna.Sprintf("Hdviet - Movie %v: Not available", movie.Id).String())
	} else {
		// Generating EpisodeKeyList
		if movie.IsSeries == false {
			EpisodeKeyList.Push(movie.Id * 1000)
		} else {
			for i := dna.Int(1); i <= movie.CurrentEps; i++ {
				EpisodeKeyList.Push(ToEpisodeKey(movie.Id, i))
			}
		}
		return movie, nil
	}

	// if movie.Link == "" || movie.Link == "/" {
	// 	return nil, errors.New(fmt.Sprintf("Nhacso - Movie %v: Mp3 link not found", movie.Id))
	// } else {
	// 	movie.Checktime = time.Now()
	// 	return movie, nil
	// }
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (movie *Movie) Fetch() error {
	_movie, err := GetMovie(movie.Id)
	if err != nil {
		return err
	} else {
		*movie = *_movie
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (movie *Movie) GetId() dna.Int {
	return movie.Id
}

// New implements item.Item interface
// Returns new item.Item interface
func (movie *Movie) New() item.Item {
	return item.Item(NewMovie())
}

// Init implements item.Item interface.
// It sets Id or key.
// dna.Interface v has type int or dna.Int, it calls Id field.
// Otherwise if v has type string or dna.String, it calls Key field.
func (movie *Movie) Init(v interface{}) {
	switch v.(type) {
	case int:
		movie.Id = dna.Int(v.(int))
	case dna.Int:
		movie.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (movie *Movie) Save(db *sqlpg.DB) error {
	insertStmt := getInsertStmt(movie, dna.Sprintf("WHERE NOT EXISTS (SELECT 1 FROM %v WHERE id=%v)", getTableName(movie), movie.Id))
	_, err := db.Exec(insertStmt.String())
	if err != nil {
		err = errors.New(err.Error() + " $$$error$$$" + insertStmt.String() + "$$$error$$$")
	}
	return err
}
