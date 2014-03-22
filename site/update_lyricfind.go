package site

import (
	"dna"
	"dna/lf"
	"dna/sqlpg"
	"errors"
	"os/exec"
	"time"
)

func testingSongWithLyrics() chan error {
	c := make(chan error)
	go func() {
		testId := dna.Int(2033)
		apiSong, conErr := lf.GetAPIFullSong(testId)
		if conErr != nil {
			c <- conErr
		} else {
			var songErr error
			if apiSong.Track.Title != "Mony Mony" {
				songErr = errors.New("Song with lyric - Wrong Title")
			}
			if apiSong.Track.Instrumental != false {
				songErr = errors.New("Song with lyric - Wrong Instrumental")
			}
			if apiSong.Track.Viewable != true {
				songErr = errors.New("Song with lyric - Wrong Viewable")
			}
			if apiSong.Track.HasLrc != true {
				songErr = errors.New("Song with lyric - Wrong HasLrc")
			}
			if apiSong.Track.LrcVerified != true {
				songErr = errors.New("Song with lyric - Wrong LrcVerified")
			}
			if apiSong.Track.Duration != "2:53" {
				songErr = errors.New("Song with lyric - Wrong Duration")
			}
			if apiSong.Track.Lyricid != 3866124 {
				songErr = errors.New("Song with lyric - Wrong Lyricid")
			}
			if apiSong.Track.Album.Id != 223224 {
				songErr = errors.New("Song with lyric - Wrong Album.Id")
			}
			if apiSong.Track.Album.Year != "1995" {
				songErr = errors.New("Song with lyric - Wrong Album.Year")
			}
			if apiSong.Track.Album.Image != "http://www.lyricfind.com/images/amg/cov75/drf500/f530/f53007zb24v.jpg" {
				songErr = errors.New("Song with lyric - Wrong Album.Image")
			}
			if apiSong.Track.Album.Largeimage != "http://www.lyricfind.com/images/amg/cov200/drf500/f530/f53007zb24v.jpg" {
				songErr = errors.New("Song with lyric - Wrong Album.Largeimage")
			}
			if apiSong.Track.Album.Artist.Id != 0 {
				songErr = errors.New("Song with lyric - Wrong Album.Artist.Id")
			}
			if apiSong.Track.Album.Artist.Image != "http://www.lyricfind.com/images/not_available_pic200.jpg" {
				songErr = errors.New("Song with lyric - Wrong Album.Artist.Image")
			}
			if apiSong.Track.Album.Artist.Name != "Various Artists" {
				songErr = errors.New("Song with lyric - Wrong Album.Artist.Name")
			}
			if apiSong.Track.Artists[0].Id != 4590 {
				songErr = errors.New("Song with lyric - Wrong Artists[0].Id")
			}
			if apiSong.Track.Artists[0].Image != "http://www.lyricfind.com/images/amg/pic200/drp100/p138/p13812erqp5.jpg" {
				songErr = errors.New("Song with lyric - Wrong Artists[0].Image")
			}
			if apiSong.Track.Artists[0].Genre != "Rock" {
				songErr = errors.New("Song with lyric - Wrong Artists[0].Genre")
			}
			if apiSong.Track.Artists[0].Name != "Tommy James & the Shondells" {
				songErr = errors.New("Song with lyric - Wrong Artists[0].Name")
			}
			if apiSong.Track.LastUpdate != "2014-01-03 08:24:58" {
				songErr = errors.New("Song with lyric - Wrong LastUpdate")
			}
			if apiSong.Track.Lyrics != "Here she come down, say Mony Mony\r\nWell, shoot 'em down, turn around come home, honey\r\nHey, she gimme love an' I feel alright now\r\nEverybody! You got me tossin' turnin' in the night\r\nMake me feel alright\r\n\r\nI say yeah, (yeah), yeah, (yeah) \r\nYeah, (yeah), yeah, (yeah), yeah\r\n\r\nWell you make me feel Mony, Mony\r\nSo Mony, Mony\r\nGood Mony, Mony\r\nYeah, Mony, Mony\r\nSo good, Mony, Mony\r\nOh, yeah, Mony, Mony\r\nCome on, Mony, Mony\r\nAll right, baby Mony, Mony\r\nSay yeah, (yeah), yeah, (yeah) \r\nYeah, (yeah), yeah, (yeah) , yeah (yeah), yeah\r\n\r\nBreak 'dis, shake 'dis, Mony, Mony\r\nShot gun, get it done, come on, honey\r\nDon't stop cookin', it feels so good, yeah\r\nHey! well don't stop now, hey, come on Mony, \r\nWell come on, Mony\r\n\r\nI say yeah, (yeah), yeah, (yeah) \r\nYeah, (yeah), yeah, (yeah), yeah, (yeah)\r\n\r\nWell you make me feel Mony, Mony\r\nSo Mony, Mony\r\nGood Mony, Mony\r\nYeah, Mony, Mony\r\nOh, yeah, Mony, Mony\r\nCome on, Mony, Mony\r\nSo good, Mony, Mony\r\nAll right, Mony, Mony\r\n\r\nI say yeah, (yeah), yeah, (yeah) \r\nYeah, (yeah), yeah, (yeah), yeah, (yeah)\r\n\r\nOh, I love your Mony, moan, moan, Mony (so good)\r\nOh, I love your Mony, moan, moan, Mony (so fun)\r\nOh, I love your Mony, moan, moan, Mony \r\nOh, I love your Mony, moan, moan, Mony\r\n\r\nYeah, (yeah), yeah, (yeah) \r\nYeah, (yeah), yeah, (yeah), yeah, (yeah)\r\n\r\nCome on! Mony, Mony\r\nCome on! Mony, Mony\r\nCome on! Mony, Mony\r\nEverybody, Mony, Mony\r\nAll right, Mony, Mony\r\nMony, Mony\r\nMony, Mony" {
				songErr = errors.New("Song with lyric - Wrong Lyrics")
			}
			if apiSong.Track.Copyright != "Lyrics © EMI Music Publishing" {
				songErr = errors.New("Song with lyric - Wrong Copyright")
			}
			if apiSong.Track.Writer != "BLOOM, BOBBY/JAMES, TOMMY/ROSENBLATT, RICHARD/GENTRY, BO" {
				songErr = errors.New("Song with lyric - Wrong Writer")
			}
			c <- songErr
		}

	}()
	return c
}

func testingSongWithMetadata() chan error {
	c := make(chan error)
	go func() {
		testId := dna.Int(30720645)
		apiSong, conErr := lf.GetAPIFullSong(testId)
		if conErr != nil {
			c <- conErr
		} else {
			var songErr error
			// dna.LogStruct(&apiSong.Track)
			if apiSong.Track.Title != "Ourestrou" {
				songErr = errors.New("Song with metadata - Wrong Title")
			}
			if apiSong.Track.Instrumental != false {
				songErr = errors.New("Song with metadata - Wrong Instrumental")
			}
			if apiSong.Track.Viewable != false {
				songErr = errors.New("Song with metadata - Wrong Viewable")
			}
			if apiSong.Track.HasLrc != false {
				songErr = errors.New("Song with metadata - Wrong HasLrc")
			}
			if apiSong.Track.LrcVerified != false {
				songErr = errors.New("Song with metadata - Wrong LrcVerified")
			}
			if apiSong.Track.Duration != "" {
				songErr = errors.New("Song with metadata - Wrong Duration")
			}
			if apiSong.Track.Lyricid != 0 {
				songErr = errors.New("Song with metadata - Wrong Lyricid")
			}
			if apiSong.Track.Album.Id != 2921748 {
				songErr = errors.New("Song with metadata - Wrong Album.Id")
			}
			if apiSong.Track.Album.Year != "2014" {
				songErr = errors.New("Song with metadata - Wrong Album.Year")
			}
			if apiSong.Track.Album.Image != "http://www.lyricfind.com/images/amg/cov75/drv800/v878/v87809r5s65.jpg" {
				songErr = errors.New("Song with metadata - Wrong Album.Image")
			}
			if apiSong.Track.Album.Largeimage != "http://www.lyricfind.com/images/amg/cov200/drv800/v878/v87809r5s65.jpg" {
				songErr = errors.New("Song with metadata - Wrong Album.Largeimage")
			}
			if apiSong.Track.Album.Artist.Id != 0 {
				songErr = errors.New("Song with metadata - Wrong Album.Artist.Id")
			}
			if apiSong.Track.Album.Artist.Image != "http://www.lyricfind.com/images/not_available_pic200.jpg" {
				songErr = errors.New("Song with metadata - Wrong Album.Artist.Image")
			}
			if apiSong.Track.Album.Artist.Name != "Various Artists" {
				songErr = errors.New("Song with metadata - Wrong Album.Artist.Name")
			}
			if apiSong.Track.Artists[0].Id != 289803 {
				songErr = errors.New("Song with metadata - Wrong Artists[0].Id")
			}
			if apiSong.Track.Artists[0].Image != "http://www.lyricfind.com/images/not_available_pic200.jpg" {
				songErr = errors.New("Song with metadata - Wrong Artists[0].Image")
			}
			if apiSong.Track.Artists[0].Genre != "World" {
				songErr = errors.New("Song with metadata - Wrong Artists[0].Genre")
			}
			if apiSong.Track.Artists[0].Name != "Djamel Allam" {
				songErr = errors.New("Song with metadata - Wrong Artists[0].Name")
			}
			if apiSong.Track.LastUpdate != "" {
				songErr = errors.New("Song with metadata - Wrong LastUpdate")
			}
			if apiSong.Track.Lyrics != "" {
				songErr = errors.New("Song with metadata - Wrong Lyrics")
			}
			if apiSong.Track.Copyright != "" {
				songErr = errors.New("Song with metadata - Wrong Copyright")
			}
			if apiSong.Track.Writer != "" {
				songErr = errors.New("Song with metadata - Wrong Writer")
			}
			c <- songErr
		}

	}()
	return c
}

func testingSongs() error {

	metaErr := <-testingSongWithMetadata()
	if metaErr != nil {
		return metaErr
	}

	lyricErr := <-testingSongWithLyrics()
	if lyricErr != nil {
		return lyricErr
	}

	return nil
}

func startVPN(vpnAppName, vpnAppUrl dna.String) chan *exec.Cmd {
	c := make(chan *exec.Cmd)
	dna.Log("Starting", vpnAppName)
	go func() {
		cmd := exec.Command(vpnAppUrl.String())
		c <- cmd
		err := cmd.Run()
		if err != nil {
			dna.Log(err.Error())
		}
	}()

	return c
}

func stopVPN(vpnAppName dna.String, cmd *exec.Cmd) {
	cmdErr := cmd.Process.Kill()
	if cmdErr != nil {
		dna.Log(vpnAppName, "cannot shut it down properly. Error:", cmdErr.Error())
	} else {
		dna.Log(vpnAppName, "was shut down!")
	}
}

func UpdateLyricFind(vpnAppName, vpnAppUrl dna.String, estConDuration time.Duration) {
	cmd := <-startVPN(vpnAppName, vpnAppUrl)
	CountDown(estConDuration*time.Second, "Establishing VPN connection. Estimated time remaining:", "")

	err := testingSongs()
	if err != nil {
		dna.Log("Error occurs: ", err.Error())
		dna.Log("Operation aborted!")
	} else {
		db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
		dna.PanicError(err)
		siteConf, err := LoadSiteConfig("lf", SiteConfigPath)
		dna.PanicError(err)

		state := NewStateHandler(new(lf.Song), siteConf, db)
		Update(state)

		RecoverErrorQueries(SqlErrorLogPath, db)
		CountDown(3*time.Second, QuittingMessage, EndingMessage)
		db.Close()
	}
	stopVPN(vpnAppName, cmd)
}
