package zi

import (
	. "dna"
)

// Basic video type
type Video struct {
	Key     String
	Id      Int
	Title   String
	Artists StringArray
	Authors StringArray
	Link    String
}

// Video constructor with key
func NewVideo(key String) *Video {
	video := new(Video)
	video.Key = key
	video.Id = Decrypt(key)
	video.Title = ""
	video.Artists = StringArray{}
	video.Authors = StringArray{}
	video.Link = ""
	return video
}

// Getting encoded key used for XML link or getting direct video url
func (v *Video) GetEncodedKey(resolution Resolution) String {
	tailArray := IntArray{10}.Concat(Int(resolution).ToString().Split("").ToIntArray()).Concat(IntArray{10, 2, 0, 1, 0})
	return getCipherText(GetId(v.Key), tailArray)
}

// Getting direct video link from the site with various qualities
func (v *Video) GetDirectLink(resolution Resolution) String {
	return VIDEO_BASE_URL.Concat(v.GetEncodedKey(resolution), "/")
}
