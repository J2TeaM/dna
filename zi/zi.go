/*
 This package implements some methods: encryption, decyptions regarding to mp3.zing.vn
*/
package zi

import (
	. "dna"
	"math/rand"
	"time"
)

const (
	TV_BASE_URL    String = "http://tv.zing.vn/html5/video"
	SONG_BASE_URL  String = "http://mp3.zing.vn/download/song/joke-link"
	VIDEO_BASE_URL String = "http://mp3.zing.vn/html5/video"
)

var a = String("0IWOUZ6789ABCDEF").Split("")
var b = String("0123456789abcdef").Split("")
var c = String("GHmn|LZk|DFbv|BVd|ASlz|QWp|ghXC|Nas|Jcx|ERui|Tty|rIU|POwq|efK|Mjo").Split("|")

// Bitrate specifies a bitrate of the output audio
type Bitrate Int

const (
	Lossless   Bitrate = 0
	Bitrate128 Bitrate = 128
	Bitrate256 Bitrate = 256
	Bitrate320 Bitrate = 320
)

// Resolution specifies resolution of the output video
type Resolution Int

const (
	Resolution240p  Resolution = 240
	Resolution360p  Resolution = 360
	Resolution480p  Resolution = 480
	Resolution720p  Resolution = 720
	Resolution1080p Resolution = 1080
)

// Checking if key is valid
func CheckKey(key String) Bool {
	for _, v := range key.Split("") {
		if a.IndexOf(v) == -1 {
			return false
		}
	}
	return true
}

// Encode integer ID into Key
func Encrypt(id Int) String {
	return StringArray(id.ToHex().Split("").Map(
		func(v String, i Int) String {
			return a[b.IndexOf(v)]
		}).([]String)).Join("")
}

// Decode Key into integer ID
func Decrypt(key String) Int {
	return ParseInt(StringArray(key.Split("").Map(func(v String, i Int) String {
		return b[a.IndexOf(v)]
	}).([]String)).Join(""), 16)
}

// Alias of func Encrypt()
func GetKey(id Int) String {
	return Encrypt(id)
}

// Alias of func Decrypt()
func GetId(key String) Int {
	return Decrypt(key)
}

func getCipherText(id Int, tailArray IntArray) String {
	rand.Seed(time.Now().UnixNano())
	return StringArray(IntArray{1, 0, 8, 0, 10}.Concat((id - 307843200).ToString().Split("").ToIntArray()).Concat(tailArray).Map(
		func(v Int, i Int) String {
			return c[v].Split("")[rand.Intn(len(c[v]))]
		}).([]String)).Join("")
}

// Decode the encoded key to key
func DecodeEncodedKey(key String) String {
	var y IntArray = key[5:15].Split("").Map(func(v String, i Int) Int {
		for j, val := range c {
			for _, char := range val.Split("") {
				if char == v {
					return Int(j)
				}
			}
		}
		return -1
	}).([]Int)
	return GetKey(y.ToString().ToInt() + 307843200)
}
