package site

import (
	. "dna"
	"testing"
)

func TestSiteConfig(t *testing.T) {
	siteConfig, err := LoadSiteConfig("ns", "./sites.ini")
	PanicError(err)
	siteConfig.NConcurrent = 1
	siteConfig.NCSongFail = 2
	siteConfig.NCAlbumFail = 3
	siteConfig.NCVideoFail = 4
	err = siteConfig.Save("./sites.ini")
	PanicError(err)
	newConfig, err := LoadSiteConfig("ns", "./sites.ini")
	PanicError(err)
	if newConfig.NConcurrent != 1 {
		t.Error("Wrong nconcurrent")
	}
	if newConfig.NCSongFail != 2 {
		t.Error("Wrong ncfail")
	}
	if newConfig.NCAlbumFail != 3 {
		t.Error("Wrong ncfail")
	}
	if newConfig.NCVideoFail != 4 {
		t.Error("Wrong ncfail")
	}
}
