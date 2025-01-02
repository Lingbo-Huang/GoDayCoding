package block

import (
	"log"
	"time"
)

type Block struct {
}

func (b *Block) Name() string {
	return "block"
}

func (b *Block) Run() {
	log.Println(b.Name(), "Run")
	<-time.After(time.Second)
}
