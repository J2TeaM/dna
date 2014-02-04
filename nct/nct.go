package nct

import (
	"dna"
	"dna/sqlpg"
	"dna/utils"
)

const (
	// TotalSongPages defines the number of pages for each path in NewestSongPaths
	TotalSongPages dna.Int = 25
	// TotalAlbumPages defines the number of pages for each path in NewestAlbumPaths
	TotalAlbumPages dna.Int = 28
	// TotalVideoPages defines the number of pages for each path in NewestVideoPaths
	TotalVideoPages dna.Int = 23
)

// Defines resutls of keys, ids when getting latest songs, albums or videos
var (
	// NewestSongPortions defines the total newest song portions getting from topic paths
	NewestSongPortions = dna.StringArray{}
	// NewestSongPortions defines the total newest album portions getting from topic paths
	NewestAlbumPortions = dna.StringArray{}
	// NewestSongPortions defines the total newest video portions getting from topic paths
	NewestVideoPortions = dna.StringArray{}
)

// Defines relevant songs, albums and videos
var (
	EnableRelevantPortionsMode = true
	RelevantSongs              = dna.StringArray{}
	RelevantAlbums             = dna.StringArray{}
	RelevantVideos             = dna.StringArray{}
)

// NewestSongPaths defines the newest song paths
var NewestSongPaths = dna.StringArray{
	"/bai-hat/bai-hat-moi-nhat.html",
	"/bai-hat/nhac-tre-moi-nhat.html",
	"/bai-hat/tru-tinh-moi-nhat.html",
	"/bai-hat/cach-mang-moi-nhat.html",
	"/bai-hat/tien-chien-moi-nhat.html",
	"/bai-hat/nhac-trinh-moi-nhat.html",
	"/bai-hat/thieu-nhi-moi-nhat.html",
	"/bai-hat/rap-viet-moi-nhat.html",
	"/bai-hat/rock-viet-moi-nhat.html",
	"/bai-hat/dance-moi-nhat.html",
	"/bai-hat/rbhip-hoprap-moi-nhat.html",
	"/bai-hat/bluejazz-moi-nhat.html",
	"/bai-hat/country-moi-nhat.html",
	"/bai-hat/latin-moi-nhat.html",
	"/bai-hat/indie-moi-nhat.html",
	"/bai-hat/au-my-khac-moi-nhat.html",
	"/bai-hat/khong-loi-moi-nhat.html",
	"/bai-hat/au-my-moi-nhat.html",
	"/bai-hat/han-quoc-moi-nhat.html",
	"/bai-hat/nhac-hoa-moi-nhat.html",
	"/bai-hat/nhac-nhat-moi-nhat.html",
	"/bai-hat/nhac-phim-moi-nhat.html",
	"/bai-hat/the-loai-khac-moi-nhat.html"}

// NewestAlbumPaths defines the newest album paths
var NewestAlbumPaths = dna.StringArray{
	"/playlist/playlist-moi-nhat.html",
	"/playlist/nhac-tre-moi-nhat.html",
	"/playlist/tru-tinh-moi-nhat.html",
	"/playlist/cach-mang-moi-nhat.html",
	"/playlist/tien-chien-moi-nhat.html",
	"/playlist/nhac-trinh-moi-nhat.html",
	"/playlist/thieu-nhi-moi-nhat.html",
	"/playlist/rap-viet-moi-nhat.html",
	"/playlist/rock-viet-moi-nhat.html",
	"/playlist/au-my-moi-nhat.html",
	"/playlist/han-quoc-moi-nhat.html",
	"/playlist/nhac-hoa-moi-nhat.html",
	"/playlist/nhac-nhat-moi-nhat.html",
	"/playlist/khong-loi-moi-nhat.html",
	"/playlist/nhac-phim-moi-nhat.html",
	"/playlist/the-loai-khac-moi-nhat.html"}

// NewestVideoPaths defines the newest video paths
var NewestVideoPaths = dna.StringArray{
	"/video-am-nhac-viet-nam-nhac-tre-moi-nhat.html",
	"/video-am-nhac-viet-nam-tru-tinh-moi-nhat.html",
	"/video-am-nhac-viet-nam-que-huong-moi-nhat.html",
	"/video-am-nhac-viet-nam-cach-mang-moi-nhat.html",
	"/video-am-nhac-viet-nam-thieu-nhi-moi-nhat.html",
	"/video-am-nhac-viet-nam-nhac-rap-moi-nhat.html",
	"/video-am-nhac-viet-nam-nhac-rock-moi-nhat.html",
	"/video-am-nhac-au-my-pop-moi-nhat.html",
	"/video-am-nhac-au-my-rock-moi-nhat.html",
	"/video-am-nhac-au-my-dance-moi-nhat.html",
	"/video-am-nhac-au-my-r-b-hip-hop-rap-moi-nhat.html",
	"/video-am-nhac-au-my-blue-jazz-moi-nhat.html",
	"/video-am-nhac-au-my-country-moi-nhat.html",
	"/video-am-nhac-au-my-latin-moi-nhat.html",
	"/video-am-nhac-au-my-indie-moi-nhat.html",
	"/video-am-nhac-han-quoc-moi-nhat.html",
	"/video-am-nhac-nhac-hoa-moi-nhat.html",
	"/video-am-nhac-nhac-nhat-moi-nhat.html",
	"/video-am-nhac-the-loai-khac-moi-nhat.html",
	"/video-giai-tri-funny-clip-moi-nhat.html",
	"/video-giai-tri-hai-kich-moi-nhat.html",
	"/video-giai-tri-phim-moi-nhat.html",
	"/video-giai-tri-khac-moi-nhat.html",
	"/video-moi-nhat.html", // Only get first page
}

// ResetRelevantPortions sets RelevantSongs, RelevantAlbums and
// RelevantVideos to &Portions{}
func ResetRelevantPortions() {
	RelevantSongs = dna.StringArray{}
	RelevantAlbums = dna.StringArray{}
	RelevantVideos = dna.StringArray{}
}

func FilterRelevants(db *sqlpg.DB) {
	func() {
		if RelevantSongs.Length() > 100 {
			RelevantSongs = RelevantSongs.Unique()
			FilterKeys(&RelevantSongs, "nctsongs", db)
		}
	}()

	func() {
		if RelevantAlbums.Length() > 100 {
			RelevantAlbums = RelevantAlbums.Unique()
			FilterKeys(&RelevantAlbums, "nctalbums", db)
		}
	}()

	func() {
		if RelevantVideos.Length() > 100 {
			RelevantVideos = RelevantVideos.Unique()
			FilterKeys(&RelevantVideos, "nctvideos", db)
		}
	}()

}

// FilterKeys gets a new portion list that keys are not in a specified table.
func FilterKeys(keys *dna.StringArray, tblName dna.String, db *sqlpg.DB) error {
	// mutex.Lock()
	// defer mutex.Unlock()
	if keys.Length() > 0 {
		missingKeys, err := utils.SelectMissingKeys(tblName, keys, db)
		if err != nil {
			return err
		} else {
			if missingKeys != nil {
				*keys = *missingKeys
				return nil
			}
		}
	} else {
		return nil
	}
	return nil
}
