package hdv

import (
	"crypto/md5"
	"dna"
	"encoding/base64"
	"io"
)

func getMD5(s dna.String) dna.String {
	h := md5.New()
	io.WriteString(h, s.String())
	return dna.Sprintf("%x", h.Sum(nil))
}

type URLBuilder struct {
}

func NewURLBuilder() *URLBuilder {
	return new(URLBuilder)
}

// GetMovie returns a movie URL.
func (urlb *URLBuilder) GetMovie(movieid dna.Int) dna.String {
	str := dna.Sprintf("movieid=%v&accesstokenkey=%v", movieid, ACCESS_TOKEN_KEY)
	data := []byte(str.String())
	strBase64 := base64.StdEncoding.EncodeToString(data)
	sign := getMD5(dna.String(strBase64) + SECRET_KEY)
	return dna.Sprintf("%vmovieid=%v&accesstokenkey=%v&sign=%v", BASE_URL, movieid, ACCESS_TOKEN_KEY, sign)
}

// GetEpisole returns an episole URL of a series.
func (urlb *URLBuilder) GetEpisole(movieid, ep dna.Int) dna.String {
	str := dna.Sprintf("movieid=%v&accesstokenkey=%v&ep=%v", movieid, ACCESS_TOKEN_KEY, ep)
	data := []byte(str.String())
	strBase64 := base64.StdEncoding.EncodeToString(data)
	sign := getMD5(dna.String(strBase64) + SECRET_KEY)
	return dna.Sprintf("%vmovieid=%v&accesstokenkey=%v&ep=%v&sign=%v", BASE_URL, movieid, ACCESS_TOKEN_KEY, ep, sign)
}

// GetChannel returns TV channel URL.
func (urlb *URLBuilder) GetChannel(channelid dna.Int) dna.String {
	str := dna.Sprintf("channelid=%v&accesstokenkey=%v", channelid, ACCESS_TOKEN_KEY)
	data := []byte(str.String())
	strBase64 := base64.StdEncoding.EncodeToString(data)
	sign := getMD5(dna.String(strBase64) + SECRET_KEY)
	return dna.Sprintf("%vchannelid=%v&accesstokenkey=%v&sign=%v", CHANNEL_BASE_URL, channelid, ACCESS_TOKEN_KEY, sign)
}
