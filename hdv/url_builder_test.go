package hdv

import (
	"dna"
)

func ExampleURLBuilder() {
	urlb := NewURLBuilder()
	dna.Log("MOVIE  :", urlb.GetMovie(5585))
	dna.Log("EPISOLE:", urlb.GetEpisole(4807, 5))
	dna.Log("CHANNEL:", urlb.GetChannel(8))
	// Output:
	// MOVIE  : https://api.hdviet.com/movie/play?movieid=5585&accesstokenkey=1c3102056acd3c12440bd05af8b9c560&sign=849ac3a1bac4ad70ef424c63d8285d7e
	// EPISOLE: https://api.hdviet.com/movie/play?movieid=4807&accesstokenkey=1c3102056acd3c12440bd05af8b9c560&ep=5&sign=d2bb08aaa5b11ecbe4d58b45b5705890
	// CHANNEL: https://api.hdviet.com/channel/play?channelid=8&accesstokenkey=1c3102056acd3c12440bd05af8b9c560&sign=9f08c9663c064bd6713ded7f8bd4cfd6
}
