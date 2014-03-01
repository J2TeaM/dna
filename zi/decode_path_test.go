package zi

import (
	"dna"
)

func ExampleDecodePath() {
	val := DecodePath("MjAxMyUyRjExJTJGMDUlMkYwJTJGMiUyRjAyN2UzN2M4NDUwMWFlOTEwNGNkZjgyMDZjYWE4OTkzLm1wMyU3QzI=")
	dna.Logv(val)
	// Output:
	// "2013/11/05/0/2/027e37c84501ae9104cdf8206caa8993.mp3|2"
}
