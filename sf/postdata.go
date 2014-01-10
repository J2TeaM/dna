package sf

import (
	"dna"
	"net/url"
	"sync"
)

type QueryValues struct {
	Values         url.Values
	LrcEnable      dna.Bool
	VideosEnable   dna.Bool
	CommentsEnable dna.Bool
	mutex          *sync.Mutex
}

// PostData defines data to post when request resource.
// It has a form:
// sfid=3921858&appname=sf_iphone&comments=true&territory=US&lrc=lrc&videos=true&version=2.1.1&apikey=6c7a95aa8238f3b437b1db24e644c2ee
var PostData = QueryValues{
	Values: url.Values{
		"sfid":      []string{"0"},
		"appname":   []string{"sf_iphone"},
		"territory": []string{"US"},
		"version":   []string{"2.1.1"},
		"apikey":    []string{"6c7a95aa8238f3b437b1db24e644c2ee"},
	},
	LrcEnable:      true,
	VideosEnable:   false,
	CommentsEnable: false,
	mutex:          &sync.Mutex{},
}

func (qv *QueryValues) SetIdKey(id dna.Int) {
	qv.mutex.Lock()
	qv.Values.Set("sfid", id.ToString().String())
	qv.mutex.Unlock()
}

func (qv *QueryValues) AddVideosKey() {
	qv.Values.Add("videos", "true")
}

func (qv *QueryValues) RemoveVideosKey() {
	qv.Values.Del("videos")
}

func (qv *QueryValues) AddLrcKey() {
	qv.Values.Add("lrc", "lrc")
}

func (qv *QueryValues) RemoveLrcKey() {
	qv.Values.Del("lrc")
}

func (qv *QueryValues) AddCommentsKey() {
	qv.Values.Add("comments", "true")
}

func (qv *QueryValues) RemoveCommentsKey() {
	qv.Values.Del("comments")
}

func (qv *QueryValues) Encode() dna.String {
	qv.mutex.Lock()
	defer qv.mutex.Unlock()
	if qv.LrcEnable == true {
		qv.AddLrcKey()
	} else {
		qv.RemoveLrcKey()
	}
	if qv.VideosEnable == true {
		qv.AddVideosKey()
	} else {
		qv.RemoveVideosKey()
	}
	if qv.CommentsEnable == true {
		qv.AddCommentsKey()
	} else {
		qv.RemoveCommentsKey()
	}
	return dna.String(qv.Values.Encode())
}
