package hdv

import "dna"

// ACCESS_TOKEN_KEY is produced by server and based on a
// HEADER of http request.
//
// It's available in movie main page (Ex:http://mmovie.hdviet.com/hdviet.5571.html in form of `var _strLinkPlay = 'hdviet://5571;9a83c5ed2e92466ed7ae459aceea3274';`)
//
// It is gonna be expired after a periof of time.
var (
	// AccessTokenKeyRenewable is to checked whether the constant ACCESS_TOKEN_KEY has been
	// re-fetched.
	// It is used when GetMovie is called.
	AccessTokenKeyRenewable = false
	ACCESS_TOKEN_KEY        = "69641ca27dacc33ff9564ad789ad6bea" // It is expired.
)

// EpisodeKeyList defines a list of episodes
// containing movieid and episode id.
// An episodeKey has formular = movieid*1000 + epid
var EpisodeKeyList = dna.IntArray{}

// LastedEpisodeKeyList defines the lastest list of episode keys
// found while udpating.
var LastestEpisodeKeyList = dna.IntArray{}

// LastestMovieCurrentEps defines last movie current eps
// if found.
var LastestMovieCurrentEps = make(map[dna.Int]dna.Int)
