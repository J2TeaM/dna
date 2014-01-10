package am

import (
	"dna"
	"dna/http"
	"dna/item"
	"dna/sqlpg"
	"encoding/json"
	"errors"
	"time"
)

type Album struct {
	Id            dna.Int
	Title         dna.String
	Artistids     dna.IntArray
	Artists       dna.StringArray
	Discographies dna.String // Json decoded string from []Discography
	Review        dna.String
	Coverart      dna.String
	Duration      dna.Int
	// Ratings contains 3 elements:
	// 	1st is site rating (0 : no rating, 1 the worst, 10 the best)
	// 	2nd is average rating
	// 	3rd is average count
	Ratings      dna.IntArray
	Similars     dna.IntArray
	Genres       dna.String // Json decoded string from []Category
	Styles       dna.String // Json decoded string from []Category
	Moods        dna.String // Json decoded string from []Category
	Themes       dna.String // Json decoded string from []Category
	Songs        dna.String // Json decoded string from []Song
	Releases     dna.String // Json decoded string from []Released
	Awards       dna.String // Json decoded string from []Award
	DateReleased time.Time
	Credits      dna.String // Json decoded string from []Credit
	Checktime    time.Time
}

func NewAlbum() *Album {
	album := new(Album)
	album.Id = 0
	album.Title = ""
	album.Artistids = dna.IntArray{}
	album.Artists = dna.StringArray{}
	album.Discographies = "[]"
	album.Review = ""
	album.Coverart = ""
	album.Duration = 0
	album.Ratings = dna.IntArray{0, 0, 0}
	album.Similars = dna.IntArray{}
	album.Genres = "[]"
	album.Styles = "[]"
	album.Moods = "[]"
	album.Themes = "[]"
	album.Songs = "[]"
	album.Releases = "[]"
	album.DateReleased = time.Time{}
	album.Credits = "[]"
	album.Checktime = time.Time{}
	return album
}

// ToSeconds returns total seconds from the time format "01:02:03"
func ToSeconds(str dna.String) dna.Int {
	if str == "" {
		return 0
	} else {
		intervals := dna.IntArray(str.Split(":").Map(func(val dna.String, idx dna.Int) dna.Int {
			return val.ToInt()
		}).([]dna.Int))
		switch intervals.Length() {
		case 3:
			return intervals[0]*3600 + intervals[1]*60 + intervals[2]
		case 2:
			return intervals[0]*60 + intervals[1]
		case 1:
			return intervals[0]
		default:
			return 0
		}
	}
}

func getTSGM(data *dna.String, kind dna.String) dna.String {
	var itemArr dna.StringArray
	switch kind {
	case "genres":
		itemArr = data.FindAllString(`(?mis)<h4>Genre</h4>(.+?)</div>`, 1)
	case "styles":
		itemArr = data.FindAllString(`(?mis)<h4>Styles</h4>(.+?)</div>`, 1)
	case "moods":
		itemArr = data.FindAllString(`(?mis)<h4>Album Moods</h4>(.+?)</div>`, 1)
	case "themes":
		itemArr = data.FindAllString(`(?mis)<h4>Themes</h4>(.+?)</div>`, 1)
	default:
		panic("Wrong kind!!!")
	}
	if itemArr.Length() > 0 {
		catArr := itemArr[0].FindAllString(`<a href=.+?</a>`, -1)
		categories := catArr.Map(func(val dna.String, idx dna.Int) Category {
			var idArr []dna.StringArray
			var id dna.Int = 0
			name := val.RemoveHtmlTags("")
			if kind == "moods" {
				idArr = val.FindAllStringSubmatch(`xa([0-9]+)`, 1)
			} else {
				idArr = val.FindAllStringSubmatch(`ma([0-9]+)`, 1)
			}
			if len(idArr) > 0 {
				id = idArr[0][1].ToInt()
			}
			return Category{id, name}
		}).([]Category)

		if len(categories) > 0 {
			bCat, merr := json.Marshal(categories)
			if merr != nil {
				return "[]"
			} else {
				return dna.String(string(bCat))
			}
		} else {
			return "[]"
		}
	} else {
		return "[]"
	}

}

// getAlbumCredit fetches album's credits
func getAlbumCredit(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.allmusic.com/album/album/google-mw" + album.Id.ToFormattedString(10, true) + "/credits/mobile"
		result, err := http.Get(link)
		if err == nil {

			data := &result.Data
			artistArr := data.FindAllString(`(?mis)<li>.+?</li>`, -1)
			credits := artistArr.Map(func(val dna.String, idx dna.Int) Credit {
				var credit dna.String = ""
				var id dna.Int = 0
				name := val.GetTags("a")[0].RemoveHtmlTags("")
				artistIdArr := val.FindAllStringSubmatch(`mn([0-9]+)`, 1)
				if len(artistIdArr) > 0 {
					id = artistIdArr[0][1].ToInt()
				}
				creditArr := val.FindAllString(`(?mis)<div class="credit">.+</div>`, 1)
				if creditArr.Length() > 0 {
					credit = creditArr[0].RemoveHtmlTags("").Trim()
				}
				return Credit{Id: id, Artist: name, Job: credit}
			}).([]Credit)

			if len(credits) > 0 {
				bCredits, derr := json.Marshal(credits)
				if derr == nil {
					album.Credits = dna.String(string(bCredits))
				}
			}
		}

		channel <- true
	}()
	return channel

}

// getAlbumAverageRating fetches average ratings has the following format
// [{"average":81.428571428571,"count":7,"itemId":"MW0002585207"}]
func getAlbumAverageRating(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.allmusic.com/rating/average/MW" + album.Id.ToFormattedString(10, true)
		result, err := http.Get(link)
		if err == nil {
			var avgRatings []AverageRating
			data := &result.Data
			umerr := json.Unmarshal([]byte(data.String()), &avgRatings)
			if umerr == nil {
				album.Ratings[1] = dna.Int(avgRatings[0].Average / 10)
				album.Ratings[2] = avgRatings[0].Count
			}
		}

		channel <- true
	}()
	return channel

}

// getAlbumSimilars fetches album's similars
// with the following url format:
// http://www.allmusic.com/album/google-bot-mw0002585207/similar/mobile
func getAlbumSimilars(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.allmusic.com/album/google-bot-mw" + album.Id.ToFormattedString(10, true) + "/similar/mobile"
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			idsArr := data.FindAllString(`<a href=".+`, -1)
			ids := dna.IntArray(idsArr.Map(func(val dna.String, idx dna.Int) dna.Int {
				idArr := val.FindAllStringSubmatch(`mw([0-9]+)`, -1)
				if len(idArr) > 0 {
					return idArr[0][1].ToInt()
				} else {
					return 0
				}
			}).([]dna.Int)).Filter(func(val dna.Int, idx dna.Int) dna.Bool {
				if val > 0 {
					return true
				} else {
					return false
				}
			})

			if ids.Length() > 0 {
				album.Similars = ids
			}
		}

		channel <- true
	}()
	return channel

}

// getAlbumReleases fetches album's releases
// with the following url format:
// http://www.allmusic.com/album/google-bot-mw0002585207/similar/mobile
func getAlbumReleases(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.allmusic.com/album/google-bot-mw" + album.Id.ToFormattedString(10, true) + "/releases/mobile"
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			rows := data.FindAllString(`(?mis)<tr>.+?</tr>`, -1)
			releases := rows.Map(func(row dna.String, idx dna.Int) Release {
				var (
					id     dna.Int    = 0
					year   dna.Int    = 0
					format dna.String = ""
					title  dna.String = ""
					label  dna.String = ""
				)

				formatArr := row.FindAllString(`(?mis)<div class="format">.+?</div>`, 1)
				if formatArr.Length() > 0 {
					format = formatArr[0].RemoveHtmlTags("").Trim()
				}

				yearArr := row.FindAllString(`(?mis)<div class="year">.+?</div>`, 1)
				if yearArr.Length() > 0 {
					year = yearArr[0].RemoveHtmlTags("").Trim().ToInt()
				}

				labelArr := row.FindAllString(`(?mis)<div class="label">.+?</div>`, 1)
				if labelArr.Length() > 0 {
					label = labelArr[0].RemoveHtmlTags("").DecodeHTML().Trim()
				}

				titleArr := row.FindAllString(`(?mis)<div class="title">.+?</div>`, 1)
				if titleArr.Length() > 0 {
					title = titleArr[0].RemoveHtmlTags("").DecodeHTML().Trim()
				}

				idArr := row.FindAllStringSubmatch(`<a href=.+mr([0-9]+)"`, 1)
				if len(idArr) > 0 {
					id = idArr[0][1].ToInt()
				}

				return Release{
					Id:     id,
					Title:  title,
					Format: format,
					Year:   year,
					Label:  label,
				}
			}).([]Release)

			if len(releases) > 0 {
				bRelease, derr := json.Marshal(releases)
				if derr == nil {
					album.Releases = dna.String(string(bRelease))
				}
			}

		}

		channel <- true
	}()
	return channel

}

func getGrammyAwardSection(section dna.String) AwardSection {
	var awards = []Award{}
	tbody := section.FindAllString(`(?mis)<tbody>.+?</tbody>`, -1)
	if tbody.Length() > 0 {
		rows := tbody[0].FindAllString(`(?mis)<tr>.+?</tr>`, -1)
		awards = rows.Map(func(row dna.String, idx dna.Int) Award {
			var (
				id      dna.Int    = 0
				year    dna.Int    = 0
				chart   dna.String = ""
				title   dna.String = ""
				peak    dna.Int    = 0
				winners            = []Person{}
				award   dna.String = ""
				atype   dna.Int    = 0
			)

			yearArr := row.FindAllString(`(?mis)<div class="year">.+?</div>`, 1)
			if yearArr.Length() > 0 {
				year = yearArr[0].RemoveHtmlTags("").Trim().ToInt()
			}

			atypeArr := row.FindAllString(`(?mis)<div class="type">.+?</div>`, 1)
			if atypeArr.Length() > 0 {
				switch atypeArr[0].RemoveHtmlTags("").Trim() {
				case "T":
					atype = 2
				case "A":
					atype = 1
				default:
					atype = 0
				}
			}

			awardArr := row.FindAllString(`(?mis)<div class="award-name">.+?</div>`, 1)
			if awardArr.Length() > 0 {
				award = awardArr[0].RemoveHtmlTags("").Trim()
			}

			titleArr := row.FindAllString(`(?mis)<td class="title".+?</div>`, 1)
			if titleArr.Length() > 0 {
				title = titleArr[0].RemoveHtmlTags("").Trim()
			}

			winnerArr := row.FindAllString(`<a href=".+?</a>`, -1)
			winners = winnerArr.Map(func(winner dna.String, idx dna.Int) Person {
				var winnerName dna.String = ""
				var winnerId dna.Int = 0
				winnerName = winner.RemoveHtmlTags("").Trim()
				winnerIdArr := winner.FindAllStringSubmatch(`<a href=.+mn([0-9]+)`, 1)
				if len(winnerIdArr) > 0 {
					winnerId = winnerIdArr[0][1].ToInt()
				}
				return Person{Id: winnerId, Name: winnerName}
			}).([]Person)

			return Award{
				Id:      id,
				Title:   title,
				Year:    year,
				Chart:   chart,
				Peak:    peak,
				Type:    atype,
				Winners: winners,
				Award:   award,
			}
		}).([]Award)
		return AwardSection{Name: "Grammy Awards", Type: "ALBUMs & SONGS", Awards: awards}
	} else {
		return AwardSection{}
	}

}

func getSection(section dna.String) AwardSection {
	var awards = []Award{}
	var name dna.String = ""
	var sectionType dna.String = ""
	var awardType dna.Int = 0

	nameArr := section.FindAllString(`(?mis)<h2 class="headline">.+?</h2>`, -1)
	if nameArr.Length() > 0 {
		name = nameArr[0].RemoveHtmlTags("").Trim()
		switch {
		case name.Match("Singles") == true:
			sectionType = "SINGLE"
			awardType = 2
		case name.Match("Albums") == true:
			sectionType = "ALBUM"
			awardType = 1

		}
	}

	if name.Match("Grammy Awards") == true {
		return getGrammyAwardSection(section)
	}

	tbody := section.FindAllString(`(?mis)<tbody>.+?</tbody>`, -1)
	if tbody.Length() > 0 {
		rows := tbody[0].FindAllString(`(?mis)<tr>.+?</tr>`, -1)
		awards = rows.Map(func(row dna.String, idx dna.Int) Award {
			var (
				id      dna.Int    = 0
				year    dna.Int    = 0
				chart   dna.String = ""
				title   dna.String = ""
				peak    dna.Int    = 0
				winners            = []Person{}
				award   dna.String = ""
			)

			yearArr := row.FindAllString(`(?mis)<td class="year".+?</div>`, 1)
			if yearArr.Length() > 0 {
				year = yearArr[0].RemoveHtmlTags("").Trim().ToInt()
			}

			peakArr := row.FindAllString(`(?mis)<td class="peak".+?</td>`, 1)
			if peakArr.Length() > 0 {
				peak = peakArr[0].RemoveHtmlTags("").Trim().ToInt()
			}

			charArr := row.FindAllString(`(?mis)<div class="chart-name">.+?</div>`, 1)
			if charArr.Length() > 0 {
				chart = charArr[0].RemoveHtmlTags("").DecodeHTML().Trim()
			}

			titleArr := row.FindAllString(`<a href=".+</a>`, 1)
			if titleArr.Length() > 0 {
				title = titleArr[0].RemoveHtmlTags("").DecodeHTML().Trim()
			}

			var match dna.String = ""
			switch sectionType {
			case "SINGLE":
				match = `<a href=.+mt([0-9]+)"`
			case "ALBUM":
				match = `<a href=.+mw([0-9]+)"`
			default:
				match = `<a href=.+mw([0-9]+)"`
			}
			idArr := row.FindAllStringSubmatch(match, 1)
			if len(idArr) > 0 {
				id = idArr[0][1].ToInt()
			}

			return Award{
				Id:      id,
				Title:   title,
				Year:    year,
				Chart:   chart,
				Peak:    peak,
				Type:    awardType,
				Winners: winners,
				Award:   award,
			}
		}).([]Award)
	}
	return AwardSection{Name: name, Type: sectionType, Awards: awards}
}

// getAlbumAwards fetches album's awards
// with the following url format:
// http://www.allmusic.com/album/google-bot-mw0002585207/similar/mobile
func getAlbumAwards(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.allmusic.com/album/google-bot-mw" + album.Id.ToFormattedString(10, true) + "/awards/mobile"
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			var awardSections = []AwardSection{}

			sectionsArr := data.FindAllString(`(?mis)<section class=.+?</section>`, -1)
			sectionsArr.ForEach(func(section dna.String, idx dna.Int) {
				awardSections = append(awardSections, getSection(section))
			})

			if len(awardSections) > 0 {
				bAwards, derr := json.Marshal(awardSections)
				if derr == nil {
					album.Awards = dna.String(string(bAwards))
				}
			}
		}

		channel <- true
	}()
	return channel

}

// getAlbumFromMainPage returns album from main page
func getAlbumFromMainPage(album *Album) <-chan bool {
	channel := make(chan bool, 1)
	go func() {
		link := "http://www.allmusic.com/album/google-bot-mw" + album.Id.ToFormattedString(10, true)
		// dna.Log(link)
		result, err := http.Get(link)
		if err == nil {
			data := &result.Data
			artistsArr := data.FindAllString(`(?mis)<h3 class="album-artist".+?</h3>`, 1)
			if artistsArr.Length() > 0 {
				// Getting Artists
				album.Artists = artistsArr[0].RemoveHtmlTags("").Trim().Split(" / ")

				// Getting Artistids
				idsArr := artistsArr[0].FindAllString(`mn[0-9]+`, -1)
				album.Artistids = dna.IntArray(idsArr.Map(func(val dna.String, idx dna.Int) dna.Int {
					idArr := val.FindAllStringSubmatch(`mn([0-9]+)`, -1)
					if len(idArr) > 0 {
						return idArr[0][1].ToInt()
					} else {
						return 0
					}
				}).([]dna.Int))
			}

			// Getting Title
			titleArr := data.FindAllString(`(?mis)<h2 class="album-title".+?</h2>`, 1)
			if titleArr.Length() > 0 {
				album.Title = titleArr[0].RemoveHtmlTags("").Trim()
			}

			// Getting Review
			reviewArr := data.FindAllStringSubmatch(`(?mis)<div class="text" itemprop="reviewBody">(.+?)</div>`, 1)
			if len(reviewArr) > 0 {
				album.Review = reviewArr[0][1].Trim().ReplaceWithRegexp(`^<p>`, ``).ReplaceWithRegexp(`</p>$`, ``).Trim()
			}

			// Getting site rating
			ratingArr := data.FindAllStringSubmatch(`<div class="allmusic-rating.+([0-9]+)"`, 1)
			if len(ratingArr) > 0 {
				siteRating := ratingArr[0][1].ToInt()
				if siteRating > 0 {
					album.Ratings[0] = siteRating + 1
				}
			}

			// Getting Duration
			durationArr := data.FindAllStringSubmatch(`(?mis)<h4>Duration</h4>.+?<span>(.+?)</span>`, 1)
			if len(durationArr) > 0 {
				album.Duration = ToSeconds(durationArr[0][1])
			}

			// Getting DateReleased
			dateReleasedArr := data.FindAllStringSubmatch(`(?mis)<h4>Release Date</h4>.+?<span>(.+?)</span>`, 1)
			if len(dateReleasedArr) > 0 {
				// dna.Log(dateReleasedArr[0][1].String())
				album.DateReleased, _ = time.Parse(`January 02, 2006`, dateReleasedArr[0][1].String())
			}

			// Getting Discographies
			discoArr := data.FindAllString(`(?mis)<li class="album">.+?</li>`, -1)
			discos := discoArr.Map(func(val dna.String, idx dna.Int) Discography {
				var id dna.Int
				title := val.GetTagAttributes("title")
				href := val.GetTagAttributes("href")
				coverart := val.GetTagAttributes("src")
				idArr := href.FindAllStringSubmatch(`mw([0-9]+)`, 1)
				if len(idArr) > 0 {
					id = idArr[0][1].ToInt()
				} else {
					id = 0
				}
				return Discography{id, title, coverart}
			}).([]Discography)
			if len(discos) > 0 {
				bDisco, err := json.Marshal(discos)
				if err == nil {
					album.Discographies = dna.String(string(bDisco))
				}
			}

			// Getting Genres, Moods, Styles and Themes
			album.Genres = getTSGM(data, "genres")
			album.Moods = getTSGM(data, "moods")
			album.Styles = getTSGM(data, "styles")
			album.Themes = getTSGM(data, "themes")

			// Getting Songs
			songTitleArr := data.FindAllString(`(?mis)<tr class="track.+?</tr>`, -1)
			songs := songTitleArr.Map(func(track dna.String, idx dna.Int) Song {
				var id, songDuration dna.Int = 0, 0
				var title dna.String = ""
				var composers, performers = []Person{}, []Person{}

				// Getting song's title and id
				titleArr := track.FindAllString(`(?mis)<div class="title" itemprop="name">.+?</div>`, 1)
				if titleArr.Length() > 0 {
					title = titleArr[0].RemoveHtmlTags("").Trim()
					idArr := titleArr[0].FindAllStringSubmatch(`mt([0-9]+)`, 1)
					if len(idArr) > 0 {
						id = idArr[0][1].ToInt()
					}
				}

				// Getting song's duration
				durationArr := track.FindAllString(`(?mis)<td class="time">.+?</td>`, 1)
				if durationArr.Length() > 0 {
					songDuration = ToSeconds(durationArr[0].RemoveHtmlTags("").Trim())
				}

				// Getting composers
				composerArr := track.FindAllString(`(?mis)<div class="composer">.+?</div>`, 1)
				if composerArr.Length() > 0 {
					composers = composerArr[0].Split(" / ").Map(func(val dna.String, idx dna.Int) Person {
						var cid dna.Int = 0
						name := val.RemoveHtmlTags("").Trim()
						performerIdArr := val.FindAllStringSubmatch(`mn([0-9]+)`, 1)
						if len(performerIdArr) > 0 {
							cid = performerIdArr[0][1].ToInt()
						}
						return Person{cid, name}
					}).([]Person)
				}

				// Getting artists
				performerArr := track.FindAllString(`(?mis)<td class="performer".+?</td>`, 1)
				if performerArr.Length() > 0 {
					perList := performerArr[0].FindAllString(`<a href=.+?</a>`, -1)
					if perList.Length() > 0 {
						performers = perList.Map(func(val dna.String, idx dna.Int) Person {
							var cid dna.Int = 0
							name := val.RemoveHtmlTags("").Trim()
							artistIdArr := val.FindAllStringSubmatch(`mn([0-9]+)`, 1)
							if len(artistIdArr) > 0 {
								cid = artistIdArr[0][1].ToInt()
							}
							return Person{cid, name}
						}).([]Person)
					}

				}

				return Song{id, title, performers, composers, songDuration}
			}).([]Song)

			if len(songs) > 0 {
				bSongs, derr := json.Marshal(songs)
				if derr == nil {
					album.Songs = dna.String(string(bSongs))
				}
			}

			// Getting Ratings

		}
		channel <- true
	}()
	return channel
}

// GetAlbum returns a album or an error
// 	* key: A unique key of a album
// 	* Official : 0 or 1, if its value is unknown, set to 0
// 	* Returns a found album or an error
func GetAlbum(id dna.Int) (*Album, error) {
	var album *Album = NewAlbum()
	album.Id = id
	c := make(chan bool, 6)
	go func() {
		c <- <-getAlbumFromMainPage(album)
	}()

	go func() {
		c <- <-getAlbumAverageRating(album)
	}()

	go func() {
		c <- <-getAlbumCredit(album)
	}()

	go func() {
		c <- <-getAlbumSimilars(album)
	}()

	go func() {
		c <- <-getAlbumReleases(album)
	}()

	go func() {
		c <- <-getAlbumAwards(album)
	}()

	for i := 0; i < 6; i++ {
		<-c
	}
	return album, nil
	if album.Songs == "[]" {
		return nil, errors.New(dna.Sprintf("Allmusic - Album %v: No song found", album.Id).String())
	} else {
		album.Checktime = time.Now()
		return album, nil
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (album *Album) Fetch() error {
	_album, err := GetAlbum(album.Id)
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
		album.Id = dna.Int(v.(int))
	case dna.Int:
		album.Id = v.(dna.Int)
	default:
		panic("Interface v has to be int")
	}
}

func (album *Album) Save(db *sqlpg.DB) error {
	return db.InsertIgnore(album)
}
