package hdv

import (
	"dna"
)

func ToMovieIdAndEpisodeId(i dna.Int) (movieid, epid dna.Int) {
	movieid = i / 1000
	epid = i % 1000
	return
}

func ToEpisodeKey(movieid, epid dna.Int) dna.Int {
	return movieid*1000 + epid
}
