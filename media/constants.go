package media

import (
	"dna"
)

const (
	NhacSo dna.Int = 1 << iota
	Zing
	NhacCuaTui
	ChaCha
	NhacVui
	ChiaSeNhac
	HDViet
	Keeng
	SongFreaks
	AllMusic
	LyricWiki
	MetroLyrics
	LyricFind
	VietGiaiTri
	MusicVNN
)

// ToShortForm returns a short form string representing a site name.
// For example: NhacSo is for "ns"
func ToShortForm(siteid dna.Int) dna.String {
	switch siteid {
	case NhacSo:
		return "ns"
	case Zing:
		return "zi"
	case NhacCuaTui:
		return "nct"
	case ChaCha:
		return "cc"
	case NhacVui:
		return "nv"
	case ChiaSeNhac:
		return "csn"
	case HDViet:
		return "hdv"
	case Keeng:
		return "ke"
	case SongFreaks:
		return "sf"
	case AllMusic:
		return "am"
	case LyricWiki:
		return "lw"
	case MetroLyrics:
		return "ml"
	case LyricFind:
		return "lf"
	case VietGiaiTri:
		return "vg"
	case MusicVNN:
		return "mv"
	default:
		panic("Cannot convert siteid to short form - GOT:" + siteid.ToString().String())
	}
}

// ToSiteid returns a siteid id from a short form.
// For example: "ns" is for NhacSo
func ToSiteid(shortForm dna.String) dna.Int {
	switch shortForm {
	case "ns":
		return NhacSo
	case "zi":
		return Zing
	case "nct":
		return NhacCuaTui
	case "cc":
		return ChaCha
	case "nv":
		return NhacVui
	case "csn":
		return ChiaSeNhac
	case "hdv":
		return HDViet
	case "ke":
		return Keeng
	case "sf":
		return SongFreaks
	case "am":
		return AllMusic
	case "lw":
		return LyricWiki
	case "ml":
		return MetroLyrics
	case "lf":
		return LyricFind
	case "vg":
		return VietGiaiTri
	case "mv":
		return MusicVNN
	default:
		panic("Cannot convert shortform to siteid - GOT:" + shortForm.String())
	}
}
