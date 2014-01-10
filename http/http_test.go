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
	Logv(result.Header.Get("Content-Encoding"))
	Logv(result.Data.Contains("Chỉ Là Em Giấu Đi, Bích Phương"))
	// Log(result.Header)
	//Output:
	// "200 OK"
	// 200
	//"text/html; charset=utf-8"
	//"gzip"
	// true

}
