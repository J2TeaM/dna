/*
 This package implements some methods: encryption, decyptions regarding to mp3.zing.vn
*/
package zi

import (
	"dna"
	"math/rand"
	"time"
)

const (
	TV_BASE_URL dna.String = "http://tv.zing.vn/html5/video"
	// SONG_BASE_URL  dna.String = "http://mp3.zing.vn/download/song/joke-link"
	SONG_BASE_URL  dna.String = "http://api.mp3.zing.vn/api/mobile/source/song"
	VIDEO_BASE_URL dna.String = "http://mp3.zing.vn/html5/video"
)

var a = dna.String("0IWOUZ6789ABCDEF").Split("")
var b = dna.String("0123456789abcdef").Split("")
var c = dna.String("GHmn|LZk|DFbv|BVd|ASlz|QWp|ghXC|Nas|Jcx|ERui|Tty|rIU|POwq|efK|Mjo").Split("|")

// Bitrate specifies a bitrate of the output audio
type Bitrate dna.Int

const (
	Lossless   Bitrate = 0
	Bitrate128 Bitrate = 128
	Bitrate256 Bitrate = 256
	Bitrate320 Bitrate = 320
)

// Resolution specifies resolution of the output video
type Resolution dna.Int

const (
	Resolution240p  Resolution = 240
	Resolution360p  Resolution = 360
	Resolution480p  Resolution = 480
	Resolution720p  Resolution = 720
	Resolution1080p Resolution = 1080
)

// Checking if key is valid
func CheckKey(key dna.String) dna.Bool {
	for _, v := range key.Split("") {
		if a.IndexOf(v) == -1 {
			return false
		}
	}
	return true
}

// Encode integer ID into Key
func Encrypt(id dna.Int) dna.String {
	return dna.StringArray(id.ToHex().Split("").Map(
		func(v dna.String, i dna.Int) dna.String {
			return a[b.IndexOf(v)]
		}).([]dna.String)).Join("")
}

// Decode Key into integer ID
func Decrypt(key dna.String) dna.Int {
	return dna.ParseInt(dna.StringArray(key.Split("").Map(func(v dna.String, i dna.Int) dna.String {
		return b[a.IndexOf(v)]
	}).([]dna.String)).Join(""), 16)
}

// Alias of func Encrypt()
func GetKey(id dna.Int) dna.String {
	return Encrypt(id)
}

// Alias of func Decrypt()
func GetId(key dna.String) dna.Int {
	return Decrypt(key)
}

func getCipherText(id dna.Int, tailArray dna.IntArray) dna.String {
	rand.Seed(time.Now().UnixNano())
	return dna.StringArray(dna.IntArray{1, 0, 8, 0, 10}.Concat((id - 307843200).ToString().Split("").ToIntArray()).Concat(tailArray).Map(
		func(v dna.Int, i dna.Int) dna.String {
			return c[v].Split("")[rand.Intn(len(c[v]))]
		}).([]dna.String)).Join("")
}

// Decode the encoded key to key
func DecodeEncodedKey(key dna.String) dna.String {
	var y dna.IntArray = key[5:15].Split("").Map(func(v dna.String, i dna.Int) dna.Int {
		for j, val := range c {
			for _, char := range val.Split("") {
				if char == v {
					return dna.Int(j)
				}
			}
		}
		return -1
	}).([]dna.Int)
	return GetKey(y.ToString().ToInt() + 307843200)
}
