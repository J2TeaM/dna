package site

import (
	"dna"
	"dna/terminal"
	"dna/utils"
	"io/ioutil"
	"os"
	// "time"
)

var (
	// SQLERROR is a default logger to print warning message
	HTTPERROR = terminal.NewLogger(terminal.Magenta, ioutil.Discard, "", "./log/http_error.log", 0)
	// SQLERROR is a default logger to print warning message
	SQLERROR = terminal.NewLogger(terminal.Magenta, ioutil.Discard, "", "./log/sql_error.log", 0)
	// INFO is a default logger to print info message
	INFO = terminal.NewLogger(terminal.White, os.Stdout, "INFO:", "./log/std.log", terminal.Ldate|terminal.Ltime)
	// // WARNING is a default logger to print warning message
	// WARNING = terminal.NewLogger(terminal.Magenta, os.Stdout, "WARNING:", "./log/std.log", terminal.Ldate|terminal.Ltime|terminal.Lshortfile)
	// // ERROR is a default logger to print error message
	// ERROR = terminal.NewLogger(terminal.Red, os.Stderr, "ERROR:", "./log/std.log", terminal.Ldate|terminal.Ltime|terminal.Lshortfile)
)

func atomicUpdate(errChannel chan bool, state *StateHandler) {
	it := state.GetItem().New()
	state.IncreaseCid()
	n := state.GetCid()
	it.Init(n)
	err := it.Fetch()
	// if it != nil {
	// 	dna.LogStruct(it)
	// }
	if err != nil {
		// dna.Log(err.Error())
		HTTPERROR.Println(err.Error())
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
	if state.IsComplete() == false {
		go atomicUpdate(errChannel, state)
	}
}

func getUpdateProgressBar(total dna.Int, tableName dna.String) *terminal.ProgressBar {
	var rt dna.String = "$[  " + tableName + "   $percent%   $current/$total]"
	rt += "\nElapsed $elapsed    Remaining $eta  ($custom)  Speed $speeditems/s"
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
func Update(state *StateHandler) {

	CheckStateHandler(state)
	var (
		counter    *Counter = NewCounter(state)
		idcFormat  dna.String
		cData      dna.String
		idc        *terminal.Indicator
		bar        *terminal.ProgressBar
		errChannel chan bool  = make(chan bool)
		tableName  dna.String = state.GetTableName()
		startupFmt dna.String = "Update %v - Cid:%v - Pat:%v - Ncf:%v - NCon:%v"
	)

	if state.GetPattern() == 1 {
		idcFormat = "  $indicator %v | cid:%v | cf:%v" // cid: current id, cf: continuous failure count
		idc = terminal.NewIndicatorWithTheme(terminal.ThemeDefault)
		// Getting maxid from an item's table
		id, err := utils.GetMaxId(tableName, state.GetDb())
		dna.PanicError(err)
		state.SetCid(id)
	} else {
		bar = getUpdateProgressBar(counter.Total, tableName)
	}

	// 3rd pattern: callind GetCid() wil invoke error
	INFO.Println(dna.Sprintf(startupFmt, tableName, state.Cid, state.GetPattern(), state.GetNCFail(), state.SiteConfig.NConcurrent))

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
			cData = dna.Sprintf("%v", counter)
			bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
		case 3:
			if counter.GetCount() == state.GetExtSlice().Length() {
				state.SetCompletion()
			}
			cData = dna.Sprintf("%v", counter)
			bar.Show(counter.GetCount(), cData, cData.Replace("|", "-"))
		}

	}
	if state.GetPattern() == 1 {
		idc.Close(dna.Sprintf("$indicator Complete updating %v!", tableName))
	}

	INFO.Printf("[%v] %v\n", tableName, counter.FinalString())
	// Delay 2s to ensure all the goroutines left finish it processed before sqlpg.DB closed
	// time.Sleep(2 * time.Second)

}
