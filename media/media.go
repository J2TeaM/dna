package media

import (
	"dna"
)

type Media struct {
	Id      dna.Int
	Title   dna.String
	Artists dna.StringArray
	Sources dna.StringArray
}
