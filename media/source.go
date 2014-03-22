package media

import (
	"dna"
)

type Source struct {
	Id     dna.Int
	Siteid dna.Int
}

// To Source converts string format such as "(12,123123)" to Source Type
func ToSource(sourceStr dna.String) Source {
	if sourceStr.Match(`\([0-9]+,[0-9]+\)`) == false {
		panic("Wrong Source format")
	}
	sourArr := sourceStr.ReplaceWithRegexp(`^\(|\)$`, "").Split(",")
	return Source{sourArr[0].ToInt(), sourArr[1].ToInt()}
}

func (s Source) SQLValue() dna.String {
	return dna.Sprintf("(%v,%v)", s.Id, s.Siteid)
}

func (s Source) String() string {
	return s.SQLValue().String()
}
