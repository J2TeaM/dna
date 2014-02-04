package nct

import (
	"dna"
)

func ExampleSongFinder() {
	songFinder := NewSongFinder()
	songFinder.Init(1)
	songFinder.Fetch()
	if songFinder.SongPortions.Length() >= 40 {
		dna.Log("Checked!")
	}
	//Output:
	//Checked!

}

func ExampleAlbumFinder() {
	albumFinder := NewAlbumFinder()
	albumFinder.Init(1)
	albumFinder.Fetch()
	if albumFinder.AlbumPortions.Length() >= 36 {
		dna.Log("Checked!")
	}
	//Output:
	//Checked!
}

func ExampleVideoFinder() {
	videoFinder := NewVideoFinder()
	videoFinder.Init(1)
	videoFinder.Fetch()
	if videoFinder.VideoPortions.Length() >= 44 {
		dna.Log("Checked!")
	}
	//Output:
	//Checked!

}
