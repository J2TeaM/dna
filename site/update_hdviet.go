package site

import (
	"dna"
	"dna/hdv"
	"dna/sqlpg"
	"time"
)

// UpdateHDViet gets lastest items from hdviet.com.
//
// The update process goes through 8 steps and 2 parts:
//
// 	PART I: FINDING AND UPDATING NEW EPS OF MOVIES IN DB
// 	Step 1: Initalizing db connection, loading site config and state handler.
// 	Step 2: Finding all movies possibly having new episodes in DB
// 	Step 3: Finding all movies found in Step 2 actually having new episodes available from source website.
// 	Step 4: Checking consitency of movies' new found episodes from the source.
// 	Step 5: Getting new episode data from the source, save them to DB and update current_eps field of movies having new eps.
//
// 	PART II: GETTING NEW MOVIES FROM HDVIET SITE.
// 	Step 6: Find newest movies from the source.
// 	Step 7: Updating newest episodes from newest found movies in Step 6.
// 	Step 8: Recovering failed sql statements.
func UpdateHDViet() {
	var mvTable = dna.String("hdvmovies")
	db, err := sqlpg.Connect(sqlpg.NewSQLConfig(SqlConfigPath))
	dna.PanicError(err)
	siteConf, err := LoadSiteConfig("hdv", SiteConfigPath)
	dna.PanicError(err)
	TIMEOUT_SECS = 100

	// PART 1: UPDATING NEW EPISODES OF MOVIES IN DB
	// STEP 2: Finding all movies possibly having new episodes in DB.
	movieCurrentEps, err := hdv.GetMoviesCurrentEps(db, mvTable)
	dna.PanicError(err)
	newEpisodeKeys := dna.IntArray{}
	for movieid, currentEps := range movieCurrentEps {
		newEpisodeKeys.Push(hdv.ToEpisodeKey(movieid, currentEps))
	}
	newEpisodeKeys.Sort()

	// STEP 3: Checking and getting new episodes if available from source website.
	state := NewStateHandlerWithExtSlice(new(hdv.EpUpdater), &newEpisodeKeys, siteConf, db)
	Update(state)

	// STEP 4:Checking consitency of new found episodes of movies from the source.
	movieIdList := dna.IntArray{}
	for _, epKey := range hdv.LastestEpisodeKeyList {
		mvid, _ := hdv.ToMovieIdAndEpisodeId(epKey)
		movieIdList.Push(mvid)
	}
	movieIdList = movieIdList.Unique()
	dna.Log("\nNumber of movies: ", movieIdList.Length(), "having new episodes:", hdv.LastestEpisodeKeyList.Length())

	if movieIdList.Length() != dna.Int(len(hdv.LastestMovieCurrentEps)) {
		dna.Log("LastestEpisodeKeyList & LastestMovieCurrentEps do not match! GOT:", movieIdList.Length(), len(hdv.LastestMovieCurrentEps))
	}

	// STEP 5: Getting new episode data from the source, save them to DB and update current_eps field of movies having new eps.
	if hdv.LastestEpisodeKeyList.Length() > 0 {
		state = NewStateHandlerWithExtSlice(new(hdv.Episode), &hdv.LastestEpisodeKeyList, siteConf, db)
		Update(state)
		RecoverErrorQueries(SqlErrorLogPath, db)
		// dna.Log(hdv.LastestEpisodeKeyList)

		hdv.SaveLastestMovieCurrentEps(db, mvTable, SQLERROR)
		RecoverErrorQueries(SqlErrorLogPath, db)
	} else {
		dna.Log("No new episodes found. Update operation has been aborted!")
	}

	// PART 2: UPDATING NEW MOVIES FROM hdv site.
	// STEP 6: Find newest movies from sources.
	db.Ping()
	state = NewStateHandler(new(hdv.Movie), siteConf, db)
	Update(state)

	// STEP 7: Updating newest episodes from newest found movies in Step 6.
	hdv.EpisodeKeyList = hdv.EpisodeKeyList.Unique()
	if hdv.EpisodeKeyList.Length() > 0 {
		state = NewStateHandlerWithExtSlice(new(hdv.Episode), &hdv.EpisodeKeyList, siteConf, db)
		Update(state)
	} else {
		dna.Log("No new movies found")
	}

	// STEP 8: Recovering failed sql statements.
	RecoverErrorQueries(SqlErrorLogPath, db)

	CountDown(3*time.Second, QuittingMessage, EndingMessage)
	db.Close()
}
