package http

import (
	"compress/gzip"
	"compress/zlib"
	. "dna"
	"io"
	"io/ioutil"
	"net/http"
)

// Get impliments getting site with baisc properties.
// Enable gzip, deflat by default to reduce  network data, redirect to new location from response.
// It returns data (String type) and error
// if err is nil then data is "" (empty).
func Get(url String) (*Result, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url.ToPrimitiveValue(), nil)
	req.Header.Add("Accept-Encoding", "gzip,deflate")
	req.Header.Add("Accept-Language", "en-US,en")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", url.ToPrimitiveValue())
	req.Header.Add("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	req.Header.Add("Cookie", "")

	res, err := client.Do(req)
	if err != nil {
		return new(Result), err
	}

	var data []byte
	var myErr error

	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		var reader io.ReadCloser
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			return new(Result), err
		}
		data, myErr = ioutil.ReadAll(reader)
		reader.Close()
	case "deflate":
		// Logv("sdsafsd")
		reader, err := zlib.NewReader(res.Body)
		if err != nil {
			return new(Result), err
		}
		data, myErr = ioutil.ReadAll(reader)
		reader.Close()
	default:
		data, myErr = ioutil.ReadAll(res.Body)
	}

	if myErr != nil {
		return new(Result), myErr
	}

	res.Body.Close()
	return NewResult(Int(res.StatusCode), String(res.Status), res.Header, String(data)), nil
}
