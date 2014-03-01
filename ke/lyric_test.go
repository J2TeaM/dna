package ke

import (
	"dna"
)

func ExampleGetLyric() {
	_, err := GetLyric(1966248)
	if err == nil {
		panic("Lyric has to have an error")
	} else {
		if err.Error() != "Keeng - Lyric ID: 1966248 not found" {
			panic("Wrong error message!")
		}
	}
	album, err := GetLyric(1966613)
	dna.PanicError(err)
	dna.LogStruct(album)
	// Output:
	// Id : 1966613
	// Content : "<p>\r\n\tNgày ấy ta yêu nhau,em đã biết mang thương đau<br />\r\n\tNhưng trái tim em lỡ yêu anh rồi<br />\r\n\tAnh chỉ muốn lấp khoảng trống trong lòng<br />\r\n\tNào có yêu thương em gì đâu.<br />\r\n\t<br />\r\n\tGiờ đây anh ra đi, không một câu chia ly<br />\r\n\tAnh chỉ xem em giống như nhân tình<br />\r\n\tKhông thể ở bên nhau suốt đời<br />\r\n\tChỉ tìm vui nhau trong phút giây.<br />\r\n\t<br />\r\n\tNgười yêu ơi, hãy nói em nghe<br />\r\n\tSao anh vô tình lặng thinh quay bước<br />\r\n\tNgười yêu ơi, phút chốc cô đơn bơ vơ<br />\r\n\tĐôi chân lạc loài về đâu?<br />\r\n\t<br />\r\n\t[ĐK:]<br />\r\n\tLòng em tự hỏi quá yếu mềm hay em đã yêu cuồng si<br />\r\n\tĐã trao cho anh tất cả tình yêu<br />\r\n\t<span style=\"font-size: 10%; line-height: 1px; color: #EEFFFF;\">Xin Một Lần Được Yêu (Ballad Version) lyrics on ChiaSeNhac.com</span><br />\r\n\tGiấc mơ đôi ta xây nơi thiên đường<br />\r\n\tGiờ mình em nghe mưa mang anh rời xa.<br />\r\n\t<br />\r\n\tNhiều khi muốn níu lấy tình, nhưng đã quá xa tầm tay<br />\r\n\tGió ơi hãy mang hết đi buồn đau<br />\r\n\tNếu em được quay về phút ban đầu<br />\r\n\tEm xin chấp nhận lấy thêm một niềm đau.</p>"
}
