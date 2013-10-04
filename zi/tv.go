package zi

import (
	. "dna"
)

// Basic TV type
type TV struct {
	Key     String
	Id      Int
	Title   String
	Artists StringArray
	Authors StringArray
	Link    String
}

// TV constructor with key
func NewTV(key String) *TV {
	tv := new(TV)
	tv.Key = key
	tv.Id = Decrypt(key)
	tv.Title = ""
	tv.Artists = StringArray{}
	tv.Authors = StringArray{}
	tv.Link = ""
	return tv
}

// Getting encoded key of video
func (t *TV) GetEncodedKey() String {
	return getCipherText(GetId(t.Key), IntArray{10, 2, 0, 1, 0})
}

// Getting direct url for specific episode
func (t *TV) GetDirectLink() String {
	return TV_BASE_URL.Concat(t.GetEncodedKey(), "/")
}
