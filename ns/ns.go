package ns

import (
	"dna"
)

var ns = []dna.StringArray{
	dna.String("bw bg bQ bA aw ag aQ aA Zw Zg").Split(" "),
	dna.String("fedcbaZYXW").Split(""),
	dna.String("NJFBdZVRtp").Split(""),
	dna.String("U0 Uk UU UE V0 Vk VU VE W0 Wk").Split(" "),
	dna.String("RQTSVUXWZY").Split(""),
	dna.String("hlptx159BF").Split(""),
	dna.String(" X1 XF XV Wl W1 WF WV Vl V1").Split(" "),
}

var LastNPages dna.Int = 10 //  Last 10 pages from each category

// This function will encode the Id of nhacso into cipher text.
func Encrypt(id dna.Int) dna.String {
	return dna.StringArray(id.ToString().Split("").Map(
		func(v dna.String, i dna.Int) dna.String {
			return ns[6-i][v.ToInt()]
		}).([]dna.String)).Join("")
}

// This function will decode a cipher string into id.
func Decrypt(cipher dna.String) dna.Int {
	arr := dna.StringArray{cipher[0:2], cipher[2:3], cipher[3:4], cipher[4:6], cipher[6:7], cipher[7:8], cipher[8:10]}.Filter(
		func(value dna.String, index dna.Int) dna.Bool {
			return value != ""
		})
	return dna.IntArray(arr.Map(func(v dna.String, i dna.Int) dna.Int {
		return ns[6-i].IndexOf(v)
	}).([]dna.Int)).Join("").ToInt()
}

// The shortcut to encode ID
func GetKey(id dna.Int) dna.String {
	return Encrypt(id)
}

// The shortcut to decode cipher string
func GetId(key dna.String) dna.Int {
	return Decrypt(key)
}
