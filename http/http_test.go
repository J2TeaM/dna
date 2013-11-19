package http

import (
	. "dna"
)

func ExampleGet() {
	result, err := Get("http://mp3.zing.vn/album/Chi-La-Em-Giau-Di-Bich-Phuong/ZWZB0I67.html")
	if err != nil {
		panic("ERROR OCCURS")
	}
	Logv(result.Status)
	Logv(result.StatusCode)
	Logv(result.Header.Get("Content-Type"))
	Logv(result.Data.Contains("Chỉ Là Em Giấu Đi, Bích Phương"))
	//Output:
	// "200 OK"
	// 200
	//"text/html; charset=utf-8"
	// true

}
