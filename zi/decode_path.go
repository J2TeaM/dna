package zi

import (
	"dna"
	"net/url"
)

func decodePathComponent(str dna.String) dna.String {
	s := dna.StringArray{"IJKLMNOPQRSTUVWXYZabcdef", "CDEFGHSTUVWXijklmnyz0123"}
	var base24 dna.String = "0123456789abcdefghijklmn"
	s1 := dna.String("AEIMQUYcgkosw048").Split("")
	category := dna.String("")
	var additionalFactor dna.Int

	x1x2 := dna.StringArray(str.Substring(0, 2).Split("").Map(func(val dna.String, idx dna.Int) dna.String {
		return base24.CharAt(s[idx].IndexOf(val))
	}).([]dna.String)).Join("")
	n := dna.ParseInt(x1x2, 24)

	x3 := str.CharAt(2)
	x4 := str.CharAt(3)
	c_x3 := x3.CharCodeAt(0)
	c_x4 := x4.CharCodeAt(0)

Loop:
	for i := dna.Int(0); i <= s1.Length()-2; i++ {
		if 56 <= c_x3 && c_x3 <= 57 {
			if c_x3 == 56 {
				category = "typeA"
			} else {
				category = "typeB"
			}
			additionalFactor = s1.IndexOf("8")
			break Loop
		} else {
			if s1[i].CharCodeAt(0) <= c_x3 && c_x3 < s1[i+1].CharCodeAt(0) {
				additionalFactor = s1.IndexOf(s1[i])
				if c_x3 == s1[i].CharCodeAt(0) {
					category = "typeA"
				} else {
					if c_x3 == s1[i].CharCodeAt(0)+1 {
						category = "typeB"
					} else {
						category = "typeC"
					}
				}

			}
		}
	}
	c_y2 := (n%6+2)*16 + additionalFactor
	c_y1 := dna.Float((dna.Float(n) / 6)).Floor() + 32
	c_y3 := dna.Int(0)

	switch category {
	case "typeA":
		if 103 <= c_x4 && c_x4 <= 122 {
			c_y3 = c_x4 - 103 + 32
		}
		if 48 <= c_x4 && c_x4 <= 57 {
			c_y3 = c_x4 + 4
		}
	case "typeB":
		if 65 <= c_x4 && c_x4 <= 90 {
			c_y3 = c_x4 - 1
		} else if 97 <= c_x4 && c_x4 <= 122 {
			c_y3 = c_x4 - 7
		} else if 48 <= c_x4 && c_x4 <= 57 {
			c_y3 = c_x4 + 68
		}
	default:
		c_y3 = 63
	}
	return dna.FromCharCode(c_y1) + dna.FromCharCode(c_y2) + dna.FromCharCode(c_y3)

}

// DecodePath decodes encoded string such as "MjAxMyUyRjExJTJGMDUlMkYwJTJGMiUyRjAyN2UzN2M4NDUwMWFlOTEwNGNkZjgyMDZjYWE4OTkzLm1wMyU3QzI="
// into its real path on server such as "/2013/11/05/0/2/027e37c84501ae9104cdf8206caa8993.mp3"
func DecodePath(str dna.String) dna.String {
	ret := dna.StringArray{}
	for i := dna.Int(0); i < str.Length(); i = i + 4 {
		// dna.Log(str[i : i+4])
		ret.Push(decodePathComponent(str[i : i+4]))
	}
	u, err := url.Parse(ret.Join("").Replace(`%7C2`, "").String())
	if err != nil {
		panic(err.Error())
	}
	return dna.Sprintf("%v", u).ReplaceWithRegexp(`%00$`, "")
}
