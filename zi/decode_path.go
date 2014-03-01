package zi

import (
	"dna"
	"encoding/base64"
	"net/url"
)

// DecodePath decodes encoded string such as "MjAxMyUyRjExJTJGMDUlMkYwJTJGMiUyRjAyN2UzN2M4NDUwMWFlOTEwNGNkZjgyMDZjYWE4OTkzLm1wMyU3QzI="
// into its real path on server such as "/2013/11/05/0/2/027e37c84501ae9104cdf8206caa8993.mp3|2"
func DecodePath(encodedPath dna.String) dna.String {
	ret, err := base64.StdEncoding.DecodeString(encodedPath.String())
	if err == nil {
		escape, err := url.QueryUnescape(string(ret))
		if err == nil {
			return dna.String(escape)
		} else {
			return ""
		}
	} else {
		return ""
	}
}

// EncodePath encodes "/2013/11/05/0/2/027e37c84501ae9104cdf8206caa8993.mp3|2"
// to base64 string.
func EncodePath(path dna.String) dna.String {
	escape := url.QueryEscape(path.String())
	return dna.String(base64.StdEncoding.EncodeToString([]byte(escape)))
}
