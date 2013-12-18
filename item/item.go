package item

import (
	"dna"
	"dna/sqlpg"
)

// Item defines simple interface for implementation of song,album,video... from different sites
type Item interface {
	New() Item
	Init(interface{})
	Fetch() error
	GetId() dna.Int
	Save(*sqlpg.DB) error
}
