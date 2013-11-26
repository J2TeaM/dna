package site

import (
	. "dna"
	. "dna/terminal"
	"sync"
)

type Item interface {
	New() Item
	Fetch() error
	SetPrimaryCol(interface{})
}

type StateHandler struct {
	mu           sync.Mutex
	CurrentId    Int
	Done         Bool
	FailureCount Int
}

func NewStateHandler(id Int) *StateHandler {
	return &StateHandler{
		CurrentId:    id,
		Done:         false,
		FailureCount: 0,
	}
}

func (sh *StateHandler) Inc(amount Int) Int {
	sh.mu.Lock()
	sh.CurrentId += amount
	sh.mu.Unlock()
	return sh.CurrentId
}

func (sh *StateHandler) FailureUpdate(hasError Bool) Int {
	sh.mu.Lock()
	if hasError == true {
		sh.FailureCount += 1
	} else {
		sh.FailureCount = 0
	}

	sh.mu.Unlock()
	return sh.FailureCount
}

func (sh *StateHandler) SetDone(v Bool) {
	sh.mu.Lock()
	sh.Done = v
	sh.mu.Unlock()
}

func atomicUpdate(item Item, errChannel chan bool, stateHandler *StateHandler) {

	// Log(*CurrentIdPtr)
	// song := ns.NewSong()
	// song.Id = *CurrentIdPtr
	// err := song.Fetch()
	// LogStruct(song)
	it := item.New()
	it.SetPrimaryCol(stateHandler.Inc(1))
	// Log(stateHandler.CurrentId)
	err := it.Fetch()
	if err != nil {
		errChannel <- true
	} else {

		errChannel <- false
	}

	if stateHandler.Done == false {
		go atomicUpdate(item, errChannel, stateHandler)
	}

}

func Update(lastId Int, item Item) {
	var (
		stateHandler *StateHandler = NewStateHandler(lastId - 1)
		idcFormat    String        = ""
	)

	INFO.Println("Starting updating items...")

	errChannel := make(chan bool)
	counter := NewCounter()

	for i := 0; i < 20; i++ {
		go atomicUpdate(item, errChannel, stateHandler)
	}

	idc := NewIndicatorWithTheme(ThemeDefault)
	for stateHandler.Done == false {
		hasError := <-errChannel
		stateHandler.FailureUpdate(Bool(hasError))
		counter.Tick(Bool(hasError))
		if stateHandler.FailureCount == 100 {
			stateHandler.Done = true
		}
		idcFormat = "  $indicator  %v CurrentId:%v con_failures:%v"
		idc.Show(Sprintf(idcFormat, counter, stateHandler.CurrentId, stateHandler.FailureCount))

	}
	Log("")
	INFO.Println("DonePtr")

}
