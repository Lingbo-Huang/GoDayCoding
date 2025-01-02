package mutex

import (
	"log"
	"sync"
	"time"
)

type Mutex struct{}

func (m *Mutex) Name() string {
	return "mutex"
}

func (m *Mutex) Run() {
	log.Println(m.Name(), "Run")
	mutex := &sync.Mutex{}
	mutex.Lock()
	go func() {
		time.Sleep(time.Second)
		mutex.Unlock()
	}()
	mutex.Lock()
}
