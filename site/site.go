package site

import (
	. "dna"
	"dna/sqlpg"
	. "dna/terminal"
)

// Item defines simple interface for implementation of song,album,video from different sites
type Item interface {
	New() Item
	Fetch() error
	SetPrimaryColumn(interface{})
}

func atomicUpdate(item Item, errChannel chan bool, stateHandler *StateHandler) {
	it := item.New()
	it.SetPrimaryColumn(stateHandler.IncreaseId(1))
	err := it.Fetch()
	if err != nil {
		errChannel <- true
	} else {
		errChannel <- false
		stateHandler.InsertIgnore(it)
	}
	if stateHandler.Done == false {
		go atomicUpdate(item, errChannel, stateHandler)
	}
}

// Update gets item from sites and save them to database
func Update(lastId Int, item Item) {
	var (
		idcFormat  String     = "  $indicator %v | cid:%v | cf:%v" // cid: current id, cf: continuous failure count
		errChannel chan bool  = make(chan bool)
		counter    *Counter   = NewCounter()
		idc        *Indicator = NewIndicatorWithTheme(ThemeDefault)
		tableName  String     = sqlpg.GetTableName(item)
	)
	db, err := sqlpg.Connect(sqlpg.DefaultConfig)
	if err != nil {
		ERROR.Panic("Cannot connect to database")
	}
	stateHandler := NewStateHandler(lastId-1, db)

	INFO.Println("Starting updating items from", tableName)

	for i := 0; i < 20; i++ {
		go atomicUpdate(item, errChannel, stateHandler)
	}

	for stateHandler.Done == false {
		hasError := <-errChannel
		stateHandler.UpdateFailure(Bool(hasError))
		counter.Tick(Bool(hasError))
		if stateHandler.FailureCount == 100 {
			stateHandler.SetDone(true)
		}
		idc.Show(Sprintf(idcFormat, counter, stateHandler.CurrentId, stateHandler.FailureCount))

	}
	idc.Close(Sprintf("$indicator Complete updating %v!", tableName))
	INFO.Printf("[%v] %v\n", tableName, counter.FinalString())
}
