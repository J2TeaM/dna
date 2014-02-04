package am

import (
	"dna"
)

type APICredit struct {
	Id     dna.Int
	Artist dna.String
	Job    dna.String
}

func (apiCredit *APICredit) ToCredit() *Credit {
	credit := NewCredit()
	credit.Id = apiCredit.Id
	credit.Artist = apiCredit.Artist
	credit.Job = apiCredit.Job
	return credit
}

type Credit struct {
	Id      dna.Int
	Artist  dna.String
	Albumid dna.Int
	Job     dna.String
}

func NewCredit() *Credit {
	credit := new(Credit)
	credit.Id = 0
	credit.Albumid = 0
	credit.Artist = ""
	credit.Job = ""
	return credit
}

func (credit *Credit) CSVRecord() []string {
	return []string{
		credit.Id.ToString().String(),
		credit.Artist.String(),
		credit.Albumid.ToString().String(),
		credit.Job.String(),
	}
}
