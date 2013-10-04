package zi

import (
	. "dna"
)

func ExampleTV_GetEncodedKey() {
	var x String = "ZW67FWWF"
	y := DecodeEncodedKey(NewTV(x).GetEncodedKey())
	Logv(y)
	// Output: "ZW67FWWF"
}

func ExampleTV_GetDirectLink() {
	var x String = "IWZA0O0O"
	// Get the same result with different bitrates
	x = NewTV(x).GetDirectLink()
	// x has a form of "http://tv.zing.vn/html5/video/LmcntlQhEitDHLn"
}
