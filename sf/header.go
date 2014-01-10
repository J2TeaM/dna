package sf

import (
	"dna"
	"dna/http"
	dhttp "net/http"
	"sync"
)

// SET COOKIE HTTP REQUEST
var Cookie = "JSESSIONID=352FFB19E8C5EBA662C490C8401D42B3"
var mutex = &sync.Mutex{}

var Header = dhttp.Header{
	"Accept-Encoding": []string{"gzip"},
	"Content-Type":    []string{"application/x-www-form-urlencoded; charset=utf-8"},
	"Connection":      []string{"keep-alive"},
	"Content-Length":  []string{"0"},
	"User-Agent":      []string{"SongFreaks 2.1.1 (iPad; iPhone OS 7.0.4; en_US)"},
	"Cookie":          []string{""},
}

func Post(url, bodyStr dna.String) (*http.Result, error) {
	mutex.Lock()
	http.DefaulHeader = Header
	http.DefaulHeader.Set("Cookie", Cookie)
	http.DefaulHeader.Set("Content-Length", string(bodyStr.Length().ToString()))
	mutex.Unlock()
	// dna.Log(http.DefaulHeader)
	return http.Post(url, bodyStr)
}
