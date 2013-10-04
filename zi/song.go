package zi

import (
	. "dna"
)

// The basic song type
type Song struct {
	Id          Int
	Key         String
	Title       String
	Artists     StringArray
	Authors     StringArray
	Plays       Int
	Topics      StringArray
	Link        String
	Path        String
	Lyric       String
	DateCreated String
}

// Song constructor with key
func NewSong(key String) *Song {
	song := new(Song)
	song.Key = key
	song.Id = Decrypt(key)
	song.Title = ""
	song.Artists = StringArray{}
	song.Authors = StringArray{}
	song.Plays = 0
	song.Link = ""
	song.Path = ""
	song.Lyric = ""
	song.DateCreated = ""
	return song
}

// Song constructor with ID
func NewSongWithId(id Int) *Song {
	return NewSong(GetKey(id))
}

// Getting encoded key used for XML file or direct link
func (s *Song) GetEncodedKey(bitrate Bitrate) String {
	var temp IntArray
	if bitrate == Lossless {
		temp = IntArray{11, 12, 13, 13, 11, 14, 13, 13}
	} else {
		temp = Int(bitrate).ToString().Split("").ToIntArray()
	}
	tailArray := IntArray{10}.Concat(temp).Concat(IntArray{10, 2, 0, 1, 0})
	return getCipherText(GetId(s.Key), tailArray)

}

// Notice: The interface of getting mp3 direct link with high quality( 320kbps or lossless) has been deprecated. Need to check it out later!!!
func (s *Song) GetDirectLink(bitrate Bitrate) String {
	return SONG_BASE_URL.Concat(s.GetEncodedKey(bitrate), "/")
}
