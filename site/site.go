package site

import (
	"dna"
	"dna/terminal"
	"dna/utils"
	"errors"
	"time"
)

// TIMEOUT_SECS defines the total seconds before timeout.
var TIMEOUT_SECS int64 = 8

func atomicUpdate(errChannel chan bool, state *StateHandler) {
	var err error
	var c1 = make(chan error, 1)

	it := state.GetItem().New()
	state.IncreaseCid()
	n := state.GetCid()
	it.Init(n)
	go func() {
		c1 <- it.Fetch()
	}()

	select {
	case err = <-c1:
		// do nothing
	case <-time.After(time.Second * time.Duration(TIMEOUT_SECS)):
		strErr := dna.Sprintf("%v - Timeout after %vs at id :%v", state.GetTableName(), TIMEOUT_SECS, it.GetId())
		err = errors.New(strErr.String())
	}
	state.AddNcCount(1)
	// if it != nil {
	// 	dna.LogStruct(it)
	// }
	if err != nil {
		// dna.Log(err.Error())
		HTTPERROR.Println(it.GetId(), err.Error())
		errChannel <- true
	} else {
		// checking this code.Working only with 1st pattern
		// The goroutine continues to run after DB closed so it will invoke an error
		// state.InsertIgnore(it)
		saveErr := it.Save(state.GetDb())
		if saveErr != nil {
			SQLERROR.Println(dna.String(saveErr.Error()))
		}
		errChannel <- false
	}
	state.AddNcCount(-1)
	if state.IsComplete() == false {
		switch state.GetPattern() {
		case 2:
			if state.GetCid() < state.GetRange().Last {
				go atomicUpdate(errChannel, state)
			}
		default:
			go atomicUpdate(errChannel, state)
		}

	}
}

func getUpdateProgressBar(total dna.Int, tableName dna.String) *terminal.ProgressBar {
	var rt dna.String = "$[ " + tableName + " $percent% $current/$total]"
	rt += "\nElapsed: $elapsed    ETA: $eta   Speed: $speeditems/s"
	rt += "\nStats: $custom"
	var ct dna.String = "$[  " + tableName + "  t:$elapsed    N:$total  ($custom)  ν:$speeditems/s]"
	upbar := terminal.NewProgressBar(total, rt, ct)
	upbar.Width = 70
	upbar.CompleteSymbol = " "
	upbar.IncompleteSymbol = " "
	upbar.CompleteBGColor = terminal.Green
	upbar.IncompleteBGColor = terminal.White
	upbar.CompleteTextColor = terminal.Black
	upbar.IncompleteTextColor = terminal.Black
	return upbar
}

// Update gets item from sites and save them to database
func Update(state *StateHandler) *Counter {

	CheckStateHandler(state)
	var (
		counter    *Counter = NewCounterWithStateHandler(state)
		idcFormat  dna.String
		cData      dna.String
		idc        *terminal.Indicator
		bar        *terminal.ProgressBar
		errChannel chan bool  = make(chan bool)
		tableName  dna.String = state.GetTableName()
		startupFmt dna.String = "Update %v - Cid:%v - Pat:%v - Ncf:%v - NCon:%v"
	)

	// 3rd pattern: callind GetCid() wil invoke error
	INFO.Println(dna.Sprintf(startupFmt, tableName, state.Cid, state.GetPattern(), state.GetNCFail(), state.SiteConfig.NConcurrent))

	if utils.IsValidTable(tableName, state.GetDb()) == false {
		tableName = ""
	}

	if state.GetPattern() == 1 {
		idcFormat = " $indicator %v|cid:%v|cf:%v" // cid: current id, cf: continuous failure count
		idc = terminal.NewIndicatorWithTheme(terminal.ThemeDefault)
		// Getting maxid from an item's table
		id, err := utils.GetMaxId(tableName, state.GetDb())
		dna.PanicError(err)
		state.SetCid(id)
	} else {
		bar = getUpdateProgressBar(counter.Total, tableName)
	}

	// Config.NConcurrent
	for i := dna.Int(0); i < state.SiteConfig.NConcurrent; i++ {
		go atomicUpdate(errChannel, state)
	}

	for state.IsComplete() == false {
		hasError := <-errChannel
		counter.Tick(dna.Bool(hasError))
		switch state.GetPattern() {
		case 1:
			if counter.GetCFail() == state.GetNCFail() {
				state.SetCompletion()
			}
			idc.Show(dna.Sprintf(idcFormat, counter, state.GetCid(), counter.GetCFail()))
		case 2:
			if counter.GetCount() == state.GetRange().Total {
				state.SetCompletion()
			}
			cData = dna.Sprintf("%v | Ncc:%v | Cid:%v", counter, state.GetNcCount(), state.GetCid())
			bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
		case 3:
			if counter.GetCount() == state.GetExtSlice().Length() {
				state.SetCompletion()
			}
			cData = dna.Sprintf("%v | Ncc:%v | Cid:%v", counter, state.GetNcCount(), state.GetCid())
			bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
		}

	}
	if state.GetPattern() == 1 {
		idc.Close(dna.Sprintf("$indicator Complete updating %v!", tableName))
	}

	INFO.Printf("[%v] %v\n", tableName, counter.FinalString())
	// Delay 2s to ensure all the goroutines left finish it processed before sqlpg.DB closed
	// time.Sleep(2 * time.Second)
	return counter
}
