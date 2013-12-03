package zi

import (
	. "dna"
)

func ExampleTV_GetEncodedKey() {
	var x String = "ZW67FWWF"
	tv := NewTV()
	tv.Key = x
	tv.Id = GetId(x)
	y := DecodeEncodedKey(tv.GetEncodedKey())
	Logv(y)
	// Output: "ZW67FWWF"
}

func ExampleTV_GetDirectLink() {
	var x String = "IWZA0O0O"
	tv := NewTV()
	tv.Key = x
	tv.Id = GetId(x)
	// Get the same result with different bitrates
	x = tv.GetDirectLink()
	// x has a form of "http://tv.zing.vn/html5/video/LmcntlQhEitDHLn"
}
