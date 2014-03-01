package ke

import (
	"dna"
	"time"
)

func ExampleGetAlbum() {
	_, err := GetAlbum(84935)
	if err == nil {
		panic("Album has to have an error")
	} else {
		if err.Error() != "Keeng - Album ID: 84935 not found" {
			panic("Wrong error message!")
		}
	}
	album, err := GetAlbum(86694)
	dna.PanicError(err)
	album.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
	if album.Plays < 36049 {
		panic("Wrong play")
	}
	album.Plays = 36049
	dna.LogStruct(album)
	// Output:
	// Id : 86694
	// Key : "09BYXAMW"
	// Title : "100 Hit Nhạc Việt 2013 (Part 1)"
	// Artists : dna.StringArray{"Mỹ Tâm", "Thu Minh", "The Men", "Hồ Quang Hiếu", "Khởi My", "Miu Lê", "Hồng Dương M4U", "Bảo Thy", "Noo Phước Thịnh", "Đông Nhi"}
	// Nsongs : 50
	// Plays : 36049
	// Coverart : "http://media3.keeng.vn:8082/medias/images/images_thumb/f_medias_6/album/image/2013/12/23/b6abfc53bcb1f5ecd1c7a2ee7f6f5292a79e12e2_103_103.jpg"
	// Description : ""
	// Songids : dna.IntArray{1771701, 1907293, 1922861, 1786476, 1764324, 1794339, 1944931, 1944729, 1759706, 1787817, 1944090, 1802630, 1907296, 1790298, 1771702, 1776165, 1779905, 1963988, 1962277, 1788903, 1965592, 1782224, 1954418, 1790245, 1788955, 1963617, 1771761, 1965428, 1968399, 1966605, 1821302, 1787318, 1959758, 1747128, 1968036, 1760361, 1967779, 1794784, 1967780, 1919261, 1922843, 1965180, 1785077, 1958497, 1794641, 1966916, 1780127, 1773989, 1967851, 1958909}
	// DateCreated : "2013-12-23 00:00:00"
	// Checktime : "2013-11-21 00:00:00"
}
