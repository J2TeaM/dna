package nct

import (
	. "dna"
	"dna/http"
)

func ExampleGetRelevantPortions() {
	// UTIL FUNC
	showRelevantPortions := func(portionType String) {
		Log("RELEVANT PORTIONS OF THE " + portionType + ":")
		Log("Song - Length:", len(RelevantSongs))
		Log(RelevantSongs)

		Log("Album - Length:", len(RelevantAlbums))
		Log(RelevantAlbums)

		Log("Video - Length:", len(RelevantVideos))
		Log(RelevantVideos)
		ResetRelevantPortions()
	}
	// GET FROM SONG URL
	var link String = "http://www.nhaccuatui.com/bai-hat/google-bot.PdVrkoUSKvuR.html"
	result, err := http.Get(link)
	if err == nil {
		GetRelevantPortions(&result.Data)
	}
	showRelevantPortions("SONG")

	// GET FROM ALBUM URL
	link = "http://www.nhaccuatui.com/playlist/google-bot.Cvs5CotFG3aL.html"
	result, err = http.Get(link)
	if err == nil {
		GetRelevantPortions(&result.Data)
	}
	showRelevantPortions("ALBUM")

	// GET FROM VIDEO URL
	link = "http://www.nhaccuatui.com/video/google-bot.8V46a78ezO1Yu.html"
	result, err = http.Get(link)
	if err == nil {
		GetRelevantPortions(&result.Data)
	}
	showRelevantPortions("VIDEO")

	// Output:
	// RELEVANT PORTIONS OF THE SONG:
	// Song - Length: 20
	// dna.StringArray{"en4BhuBUvOHa", "vyToJbwcCX8s", "CJUEGOTSQN5B", "sW9f0HhLsL8G", "exOpJcegV6fh", "fnrgn0IbqQck", "kGTZ70xf1qLF", "0B3F4bkZjKqH", "RafDgCQWW9WV", "FGm2PAG00QhG", "Rw3XKvgMrSlS", "iE3mwAlhOk9v", "KNwtyyzfUQRt", "uPA2X5F5Hjvm", "8DhMQyHGw7UN", "KkrxgB3uYPcT", "MJJGf1ljOt2s", "L6obM08S8KDN", "xnfUAYYISpOW", "RWoQqEhD1pnP"}
	// Album - Length: 10
	// dna.StringArray{"Qror45VOSaE0", "7jJ8ENXZaIAj", "L3hQwc8CKsp0", "uIF7rwd5wo4p", "zOxw7Ayj0ToA", "qG3V4fq8zMbA", "XiSkxJkknVXe", "uhyJJvkle0vK", "KxlCrvAdoK9d", "3Bt7GLRy46MD"}
	// Video - Length: 9
	// dna.StringArray{"I5auynRU3g7A2", "CEJbI0MI20yWQ", "ADvjp29fT0YCg", "t33pNi4Lbsnv6", "FYecuJYX1Hjx2", "VWcmlySL3RN8o", "vjDnW86LDygux", "NsXhF3dP5TK4W", "JkjveHzMmkiG"}
	// RELEVANT PORTIONS OF THE ALBUM:
	// Song - Length: 10
	// dna.StringArray{"LeAxuYRoFHZE", "VbrExEGOyS2o", "V4I4sQgwJlq8", "weeqLryzPzQ8", "z314XhINPC0m", "TBxX7vbGskLp", "3iI8wN1Flo7s", "cbYGMMdxQLhI", "JmtFIzf59VfB", "vDpb0RQHD9R9"}
	// Album - Length: 18
	// dna.StringArray{"S8DBBanHq11Z", "Qror45VOSaE0", "7jJ8ENXZaIAj", "uIF7rwd5wo4p", "zOxw7Ayj0ToA", "uhyJJvkle0vK", "FQai2P1Ugqyo", "YImzKQJh0MwT", "VR4u9fQJ6K4B", "wwXt6T4M611v", "9Lnf7NQBhzCG", "FgqytI2RAfg1", "bRg7DLB7VxSe", "rPMpl5zsRyBc", "0N2DWsjPWTu7", "QC8hxITZltYD", "VGBQjRaQAqYr", "wHh0WvmD0ZZI"}
	// Video - Length: 9
	// dna.StringArray{"b5ecEOtf4EXbw", "mWfhLIXcVo82H", "CEJbI0MI20yWQ", "Dw5iteVOEszn3", "YCnhJkPLdjkz4", "UlkkGWBoCcDNO", "8BKRc2CPk10t6", "IDYBv1GA8jMVA", "JkjveHzMmkiG"}
	// RELEVANT PORTIONS OF THE VIDEO:
	// Song - Length: 9
	// dna.StringArray{"dD0w2yC1RoGr", "ZqTKdh3ZajvG", "nfjtlx1OR3Gr", "rA5Hd9xgBeLl", "h1bdfBMh8ibf", "WLq9PeL05eat", "EC4lE0B6SgUw", "2lxfDRmlQuvk", "A39xSj3bYtUE"}
	// Album - Length: 9
	// dna.StringArray{"40pYPT4IHqAo", "fGItf1521aLi", "NDAKbn59eKNw", "uIF7rwd5wo4p", "LklrneZOZQjO", "mUClTWrdHZdH", "Kveukbhre1ry", "O9u6hddTKoud", "3Bt7GLRy46MD"}
	// Video - Length: 60
	// dna.StringArray{"0Jk6r66tyMEX", "2FWQzXUPflOb", "2VarrzE0UoQy", "2wssWCVxjiYz", "5U7O7E7HbEYA", "60pv4sLVOer5", "8F8GFdqn0f1G", "8V46a78ezO1Yu", "b7XNCPGnNiku", "BKTqWoVypN2R", "COPjMCHn4EIT", "Cz24ilc7PLKa", "DEnvLAgIIoNn", "Ek6cKnO0qgXY", "fd1IcXvjmfrL", "ffDXOWePpeQM", "FVDIEJsdZqUF", "hAYqA5mGmcp8", "JkXCJseuA6ZJ", "JnGNt8jJkORP", "Jzptw837KVEs", "KeOeYEDRGOTG", "Ki1rVpxIit2e", "lBkpaEsFsvmN", "Lgan2iGxcdo5", "lsqh3f3cbDId", "m2vPXF1P4vKa", "oDxvLOl14kwA", "OFArT2eLNWrq", "OTxrFusHiGRW", "PhiWMfBPGqDK", "pV2LxcTWFEsx", "QvR8bCuUtWYN", "rxvWvU2xOc7F", "s0YsKyZ06Z4F", "spk4F5oWY88e", "sZo3wJYHTu5g", "tKWBrFMofIDp", "tPFBevjceIPV", "uhyZAZJgDSrI", "Uyq3LihV1ELO", "VgHZqIXsC0Ez", "vuR4HzERRRn1Q", "VXeneMf4fxdY", "Whx6lAepDmsE", "zOXnuQNNfNDC", "IfaidDgD27y4E", "hl6n6XsajZd2w", "Yzc0mxi5mILWZ", "YJQAXOP9QhZvM", "pyEWT1m9Q2HLo", "tld0tCFJ79IkR", "wUEgE71BZCfVh", "mVOvjfT3MwiVZ", "LGutHpSNdF7m", "TkcMkoHGz3aj", "fjG4FDz7uN4v", "eFEKcnBNAztd", "sHlBGEE2KOrM", "HCmahYMlYHzN"}
}
