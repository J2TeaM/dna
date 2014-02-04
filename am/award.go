package am

import (
	"dna"
)

type APIAward struct {
	Id    dna.Int
	Title dna.String
	Year  dna.Int
	Chart dna.String
	Peak  dna.Int
	Type  dna.Int // 0: undefined, 1 : album, 2 song, 3 song & album
	Award dna.String
	// If Chart is empty, then Award is the name if the award
	// artists reveive
	Winners []Person
}

func (apiaward *APIAward) ToAward() *Award {
	award := NewAward()
	award.Id = apiaward.Id
	award.Title = apiaward.Title
	award.Section = ""
	award.Year = apiaward.Year
	award.Chart = apiaward.Chart
	award.Peak = apiaward.Peak
	award.Type = apiaward.Type
	award.Prize = apiaward.Award
	winnerids := dna.IntArray{}
	winners := dna.StringArray{}
	for _, winner := range apiaward.Winners {
		winnerids.Push(winner.Id)
		winners.Push(winner.Name)
	}
	award.Winnerids = winnerids
	award.Winners = winners
	return award
}

type Award struct {
	Id      dna.Int
	Title   dna.String
	Albumid dna.Int
	Section dna.String
	Year    dna.Int
	Chart   dna.String
	Peak    dna.Int
	Type    dna.Int // 0: undefined, 1 : album, 2 song, 3 song & album
	Prize   dna.String
	// If Chart is empty, then Award is the name if the award
	// artists reveive
	Winnerids dna.IntArray
	Winners   dna.StringArray
}

func NewAward() *Award {
	award := new(Award)
	award.Id = 0
	award.Albumid = 0
	award.Title = ""
	award.Section = ""
	award.Year = 0
	award.Chart = ""
	award.Peak = 0
	award.Type = 0
	award.Prize = ""
	award.Winnerids = dna.IntArray{}
	award.Winners = dna.StringArray{}
	return award
}

func (award *Award) CSVRecord() []string {
	return []string{
		award.Id.ToString().String(),
		award.Title.String(),
		award.Albumid.ToString().String(),
		award.Section.String(),
		award.Year.ToString().String(),
		award.Chart.String(),
		award.Peak.ToString().String(),
		award.Type.ToString().String(),
		award.Prize.String(),
		dna.Sprintf("%#v", award.Winnerids).Replace("dna.IntArray", "").String(),
		dna.Sprintf("%#v", award.Winners).Replace("dna.StringArray", "").String(),
	}
}
