package site

import (
	. "dna"
	"dna/cfg"
)

// SiteConfig describes basic fields for 1st pattern of update proccess
// Example:
// 	sc, err := siteConfig, err := LoadSiteConfig("ns", "./sites.ini")
// 	nda.PanicError(err)
// 	sc.LastSongid = 12121
// 	sc.Save()
type SiteConfig struct {
	NConcurrent Int
	NCSongFail  Int
	NCAlbumFail Int
	NCVideoFail Int
	src         *cfg.ConfigFile // source of the site config
	siteCode    String
}

func NewSiteConfig() *SiteConfig {
	sc := new(SiteConfig)
	sc.NConcurrent = 0
	sc.NCSongFail = 0
	sc.NCAlbumFail = 0
	sc.NCVideoFail = 0
	sc.src = nil
	sc.siteCode = ""
	return sc
}

// LoadSiteConfig loads site config from a path.
//
// 	* siteCode: The code of a site, usually it is the package name of any type implementing dna.Item interface
// 	* filepath: path to config file
// 	* Return an error if occurs
//
// A part from an example config file with a init format:
// 	; nhacso
// 	[ns]
// 	nconcurrent=20
// 	ncsongfail = 0
// 	ncalbumfail = 0
// 	ncvideofail = 0
//
func LoadSiteConfig(siteCode, filepath String) (*SiteConfig, error) {
	cf, err := cfg.LoadConfigFile(filepath)
	if err != nil {
		return nil, err
	}
	section, err := cf.GetSection(siteCode)
	if err != nil {
		return nil, err
	}
	siteConfig := NewSiteConfig()
	siteConfig.NConcurrent = section["nconcurrent"].ToInt()
	siteConfig.NCSongFail = section["ncsongfail"].ToInt()
	siteConfig.NCAlbumFail = section["ncalbumfail"].ToInt()
	siteConfig.NCVideoFail = section["ncvideofail"].ToInt()
	siteConfig.src = cf
	siteConfig.siteCode = siteCode
	return siteConfig, nil
}

// SaveSiteConfig saves site configuration to disk.
func SaveSiteConfig(sc *SiteConfig, filepath String) error {
	sc.src.SetValue(sc.siteCode, "nconcurrent", sc.NConcurrent.ToString())
	sc.src.SetValue(sc.siteCode, "ncsongfail", sc.NCSongFail.ToString())
	sc.src.SetValue(sc.siteCode, "ncalbumfail", sc.NCAlbumFail.ToString())
	sc.src.SetValue(sc.siteCode, "ncvideofail", sc.NCVideoFail.ToString())
	return cfg.SaveConfigFile(sc.src, filepath)
}

// Save stores site configuration on disk.
func (sc *SiteConfig) Save(filepath String) error {
	return SaveSiteConfig(sc, filepath)
}
