package nct

import (
	. "dna"
	"dna/http"
)

func ExampleGetRelevantPortions() {
	// UTIL FUNC
	showRelevantPortions := func(portionType String) {
		Log("RELEVANT PORTIONS OF THE " + portionType + ":")
		Log("  Song - Length:", len(*RelevantSongs))
		for _, song := range *RelevantSongs {
			Log("\t"+Sprintf("%d", song.Id), "\t", song.Key, "\t", song.IsOfficial)
		}
		Log("  Album - Length:", len(*RelevantAlbums))
		for _, album := range *RelevantAlbums {
			Log("\t"+Sprintf("%d", album.Id), "\t", album.Key, "\t", album.IsOfficial)
		}
		Log("  Video - Length:", len(*RelevantVideos))
		for _, video := range *RelevantVideos {
			Log("\t"+Sprintf("%d", video.Id), "\t", video.Key, "\t", video.IsOfficial)
		}
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
	//   Song - Length: 20
	// 	2859006 	 osLtFpjkWVvd 	 false
	// 	2806107 	 exOpJcegV6fh 	 true
	// 	2656589 	 nlYH0vDr6CrO 	 true
	// 	2845588 	 KkrxgB3uYPcT 	 false
	// 	2847425 	 MJJGf1ljOt2s 	 false
	// 	2853789 	 L6obM08S8KDN 	 true
	// 	2844828 	 cCb4mCOkn0u8 	 true
	// 	2846976 	 RWoQqEhD1pnP 	 false
	// 	2844817 	 xnfUAYYISpOW 	 false
	// 	2851771 	 fZ6X19zJQFFi 	 false
	// 	2840770 	 en4BhuBUvOHa 	 true
	// 	2825261 	 vyToJbwcCX8s 	 false
	// 	2825255 	 CJUEGOTSQN5B 	 false
	// 	2811987 	 sW9f0HhLsL8G 	 false
	// 	2806107 	 exOpJcegV6fh 	 true
	// 	2770334 	 fnrgn0IbqQck 	 true
	// 	2756676 	 kGTZ70xf1qLF 	 false
	// 	2737320 	 0B3F4bkZjKqH 	 true
	// 	2720484 	 RafDgCQWW9WV 	 true
	// 	2719733 	 FGm2PAG00QhG 	 true
	//   Album - Length: 5
	// 	0 	 7jJ8ENXZaIAj 	 false
	// 	0 	 L3hQwc8CKsp0 	 false
	// 	0 	 Cvs5CotFG3aL 	 false
	// 	0 	 72jnPpGC2B8A 	 false
	// 	0 	 3iwQqVXdrnV8 	 false
	//   Video - Length: 8
	// 	0 	 ADvjp29fT0YCg 	 false
	// 	0 	 k6ajH4XdUbQWf 	 false
	// 	0 	 jlHXXPXcv87EV 	 false
	// 	0 	 t33pNi4Lbsnv6 	 false
	// 	0 	 FYecuJYX1Hjx2 	 false
	// 	0 	 VWcmlySL3RN8o 	 false
	// 	0 	 vjDnW86LDygux 	 false
	// 	0 	 NsXhF3dP5TK4W 	 false
	// RELEVANT PORTIONS OF THE ALBUM:
	//   Song - Length: 10
	// 	2828667 	 1B1TCkm3OfBw 	 false
	// 	2843037 	 lHn3Ds5glpmG 	 true
	// 	2805977 	 7VZkL4MFGP2Y 	 true
	// 	2699925 	 z314XhINPC0m 	 true
	// 	2693887 	 weeqLryzPzQ8 	 true
	// 	2692367 	 TBxX7vbGskLp 	 true
	// 	2696332 	 3iI8wN1Flo7s 	 true
	// 	2701099 	 cbYGMMdxQLhI 	 true
	// 	2693876 	 vDpb0RQHD9R9 	 true
	// 	2699904 	 JmtFIzf59VfB 	 true
	//   Album - Length: 15
	// 	0 	 kmjgSOyykjeR 	 false
	// 	0 	 IGT9vP48AtQk 	 false
	// 	0 	 L3hQwc8CKsp0 	 false
	// 	0 	 uIF7rwd5wo4p 	 false
	// 	0 	 mUClTWrdHZdH 	 false
	// 	0 	 VR4u9fQJ6K4B 	 false
	// 	0 	 wwXt6T4M611v 	 false
	// 	0 	 9Lnf7NQBhzCG 	 false
	// 	0 	 FgqytI2RAfg1 	 false
	// 	0 	 bRg7DLB7VxSe 	 false
	// 	0 	 rPMpl5zsRyBc 	 false
	// 	0 	 0N2DWsjPWTu7 	 false
	// 	0 	 QC8hxITZltYD 	 false
	// 	0 	 VGBQjRaQAqYr 	 false
	// 	0 	 wHh0WvmD0ZZI 	 false
	//   Video - Length: 8
	// 	0 	 aU8vguEHcVfFT 	 false
	// 	0 	 1Xr4XeWluJd5L 	 false
	// 	0 	 ADvjp29fT0YCg 	 false
	// 	0 	 Dw5iteVOEszn3 	 false
	// 	0 	 YCnhJkPLdjkz4 	 false
	// 	0 	 UlkkGWBoCcDNO 	 false
	// 	0 	 8BKRc2CPk10t6 	 false
	// 	0 	 IDYBv1GA8jMVA 	 false
	// RELEVANT PORTIONS OF THE VIDEO:
	//   Song - Length: 9
	// 	2262928 	 dD0w2yC1RoGr 	 true
	// 	2823499 	 ZqTKdh3ZajvG 	 false
	// 	2826603 	 nfjtlx1OR3Gr 	 false
	// 	2807037 	 rA5Hd9xgBeLl 	 true
	// 	2803059 	 h1bdfBMh8ibf 	 true
	// 	2813057 	 WLq9PeL05eat 	 true
	// 	2798028 	 EC4lE0B6SgUw 	 true
	// 	2803062 	 2lxfDRmlQuvk 	 true
	// 	2813059 	 A39xSj3bYtUE 	 true
	//   Album - Length: 5
	// 	0 	 40pYPT4IHqAo 	 false
	// 	0 	 fGItf1521aLi 	 false
	// 	0 	 NDAKbn59eKNw 	 false
	// 	0 	 uIF7rwd5wo4p 	 false
	// 	0 	 LklrneZOZQjO 	 false
	//   Video - Length: 18
	// 	0 	 IfaidDgD27y4E 	 false
	// 	0 	 hl6n6XsajZd2w 	 false
	// 	0 	 Yzc0mxi5mILWZ 	 false
	// 	0 	 YJQAXOP9QhZvM 	 false
	// 	0 	 pyEWT1m9Q2HLo 	 false
	// 	0 	 tld0tCFJ79IkR 	 false
	// 	0 	 wUEgE71BZCfVh 	 false
	// 	0 	 mVOvjfT3MwiVZ 	 false
	// 	0 	 LGutHpSNdF7m 	 false
	// 	0 	 TkcMkoHGz3aj 	 false
	// 	0 	 BKTqWoVypN2R 	 false
	// 	0 	 fjG4FDz7uN4v 	 false
	// 	0 	 eFEKcnBNAztd 	 false
	// 	0 	 sHlBGEE2KOrM 	 false
	// 	0 	 HCmahYMlYHzN 	 false
	// 	0 	 8F8GFdqn0f1G 	 false
	// 	0 	 sZo3wJYHTu5g 	 false
	// 	0 	 OTxrFusHiGRW 	 false
}
