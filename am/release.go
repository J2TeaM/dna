package am

import (
	"dna"
)

type APIRelease struct {
	Id     dna.Int
	Title  dna.String
	Format dna.String
	Year   dna.Int
	Label  dna.String
}

func (apiRelease *APIRelease) ToRelease() *Release {
	release := NewRelease()
	release.Id = apiRelease.Id
	release.Title = apiRelease.Title
	release.Format = apiRelease.Format
	release.Year = apiRelease.Year
	release.Label = apiRelease.Label
	return release
}

type Release struct {
	Id      dna.Int
	Title   dna.String
	Albumid dna.Int
	Format  dna.String
	Year    dna.Int
	Label   dna.String
}

func NewRelease() *Release {
	release := new(Release)
	release.Id = 0
	release.Albumid = 0
	release.Title = ""
	release.Format = ""
	release.Year = 0
	release.Label = ""
	return release
}

func (release *Release) CSVRecord() []string {
	return []string{
		release.Id.ToString().String(),
		release.Title.String(),
		release.Albumid.ToString().String(),
		release.Format.String(),
		release.Year.ToString().String(),
		release.Label.String(),
	}
}
