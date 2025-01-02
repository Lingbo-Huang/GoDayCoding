package goroutine

import (
	"log"
	"time"
)

type Goroutine struct {
}

func (g *Goroutine) Name() string {
	return "goroutine"
}

func (g *Goroutine) Run() {
	log.Println(g.Name(), "Run")
	for i := 0; i < 10; i++ {
		time.Sleep(30 * time.Millisecond)
	}
}
