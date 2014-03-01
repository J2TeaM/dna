package hdv

import (
	"dna"
	"dna/item"
	"dna/sqlpg"
	"errors"
)

// EpUpdater finds new episodes.
type EpUpdater struct {
	EpId       dna.Int
	CurrentEps dna.Int
}

func NewEpUpdater() *EpUpdater {
	movieUpdater := new(EpUpdater)
	movieUpdater.EpId = 0
	movieUpdater.CurrentEps = 0
	return movieUpdater
}

func GetNewEpisodeKeys(movieid, currentEps dna.Int) (dna.IntArray, error) {
	movie, err := GetMovie(movieid)
	// Reset hdv.EpisodeKeyList
	EpisodeKeyList = dna.IntArray{}
	if err == nil {
		if movie.CurrentEps > currentEps {
			var ret = dna.IntArray{}
			LastestMovieCurrentEps[movieid] = movie.CurrentEps
			for i := currentEps + 1; i <= movie.CurrentEps; i++ {
				ret.Push(ToEpisodeKey(movieid, i))
			}
			return ret, nil
		} else {
			return nil, errors.New(dna.Sprintf("Ep ID: %v has to updated episode", movieid).String())
		}
	} else {
		return nil, err
	}
}

// Fetch implements item.Item interface.
// Returns error if can not get item
func (eu *EpUpdater) Fetch() error {
	epKeys, err := GetNewEpisodeKeys(eu.EpId, eu.CurrentEps)
	if err != nil {
		return err
	} else {
		LastestEpisodeKeyList = LastestEpisodeKeyList.Concat(epKeys)
		return nil
	}
}

// GetId implements GetId methods of item.Item interface
func (eu *EpUpdater) GetId() dna.Int {
	return ToEpisodeKey(eu.EpId, eu.CurrentEps)
}

// New implements item.Item interface
// Returns new item.Item interface
func (eu *EpUpdater) New() item.Item {
	return item.Item(NewEpUpdater())
}

// Init implements item.Item interface.
// It sets EpisodeKey
func (eu *EpUpdater) Init(v interface{}) {
	switch v.(type) {
	case int:
		eu.EpId, eu.CurrentEps = ToMovieIdAndEpisodeId(dna.Int(v.(int)))
	case dna.Int:
		eu.EpId, eu.CurrentEps = ToMovieIdAndEpisodeId(v.(dna.Int))
	default:
		panic("Interface v has to be int")
	}
}

// Save does not run.
func (eu *EpUpdater) Save(db *sqlpg.DB) error {
	return nil
}
