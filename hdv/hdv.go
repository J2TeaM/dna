package hdv

import (
	"dna"
)

// EpisodeKeyList defines a list of episodes
// containing movieid and episode id.
// An episodeKey has formular = movieid*1000 + epid
var EpisodeKeyList = dna.IntArray{}

func ToMovieIdAndEpisodeId(i dna.Int) (movieid, epid dna.Int) {
	movieid = i / 1000
	epid = i % 1000
	return
}

func ToEpisodeKey(movieid, epid dna.Int) dna.Int {
	return movieid*1000 + epid
}
