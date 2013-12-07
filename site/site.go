package site

import (
	. "dna"
	"dna/sqlpg"
	. "dna/terminal"
	"time"
)

// Item defines simple interface for implementation of song,album,video... from different sites
type Item interface {
	New() Item
	Init(interface{})
	Fetch() error
	Save(*sqlpg.DB) error
}

func GetMaxId(tableName String, db *sqlpg.DB) (Int, error) {
	var maxid Int
	err := db.QueryRow("SELECT max(id) FROM " + tableName).Scan(&maxid)
	switch {
	case err == sqlpg.ErrNoRows:
		return 0, err
	case err != nil:
		return 0, err
	default:
		return maxid, nil
	}
}

func atomicUpdate(errChannel chan bool, state *StateHandler) {
	it := state.GetItem().New()
	state.IncreaseCid()
	n := state.GetCid()
	it.Init(n)
	err := it.Fetch()
	if err != nil {
		errChannel <- true
	} else {
		errChannel <- false

		// checking this code.Working only with 1st pattern
		// The goroutine continues to run after DB closed so it will invoke an error
		// state.InsertIgnore(it)
		saveErr := it.Save(state.GetDb())
		if saveErr != nil {
			Log("cannot save item", saveErr.Error())
			LogStruct(it)
			Log("")
		}
	}
	if state.IsComplete() == false {
		go atomicUpdate(errChannel, state)
	}
}

func getUpdateProgressBar(total Int) *ProgressBar {
	var rt String = "$[Running...   $percent%   $current/$total]"
	rt += "\nElapsed $elapsed    Remaining $eta  ($custom)  Speed $speeditems/s"
	var ct String = `$[Done!    t:$elapsed    N:$total  ($custom)  ν:$speeditems/s]`
	upbar := NewProgressBar(total, rt, ct)
	upbar.Width = 70
	upbar.CompleteSymbol = " "
	upbar.IncompleteSymbol = " "
	upbar.CompleteBGColor = Green
	upbar.IncompleteBGColor = White
	upbar.CompleteTextColor = Black
	upbar.IncompleteTextColor = Black
	return upbar
}

// Update gets item from sites and save them to database
func Update(state *StateHandler) {

	CheckStateHandler(state)
	var (
		counter    *Counter = NewCounter(state)
		idcFormat  String
		cData      String
		idc        *Indicator
		bar        *ProgressBar
		errChannel chan bool = make(chan bool)
		tableName  String    = state.GetTableName()
		startupFmt String    = "Update from %v - Cid: %v - Pattern: %v - NCFail: %v - NConcurrent: %v"
	)

	if state.GetPattern() == 1 {
		idcFormat = "  $indicator %v | cid:%v | cf:%v" // cid: current id, cf: continuous failure count
		idc = NewIndicatorWithTheme(ThemeDefault)
		// Getting maxid from an item's table
		id, err := GetMaxId(tableName, state.GetDb())
		PanicError(err)
		state.SetCid(id)
	} else {
		bar = getUpdateProgressBar(counter.Total)
	}

	// 3rd pattern: callind GetCid() wil invoke error
	INFO.Println(Sprintf(startupFmt, tableName, state.Cid, state.GetPattern(), state.GetNCFail(), state.SiteConfig.NConcurrent))

	// Config.NConcurrent
	for i := Int(0); i < state.SiteConfig.NConcurrent; i++ {
		go atomicUpdate(errChannel, state)
	}

	for state.IsComplete() == false {
		hasError := <-errChannel
		counter.Tick(Bool(hasError))
		switch state.GetPattern() {
		case 1:
			if counter.GetCFail() == state.GetNCFail() {
				state.SetCompletion()
			}
			idc.Show(Sprintf(idcFormat, counter, state.GetCid(), counter.GetCFail()))
		case 2:
			if counter.GetCount() == state.GetRange().Total {
				state.SetCompletion()
			}
			cData = Sprintf("%v", counter)
			bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
		case 3:
			if counter.GetCount() == state.GetExtSlice().Length() {
				state.SetCompletion()
			}
			cData = Sprintf("%v", counter)
			bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
		}

	}
	if state.GetPattern() == 1 {
		idc.Close(Sprintf("$indicator Complete updating %v!", tableName))
	}

	INFO.Printf("[%v] %v\n", tableName, counter.FinalString())
	// Delay 2s to ensure all the goroutines left finish it processed before sqlpg.DB closed
	time.Sleep(2 * time.Second)

}
