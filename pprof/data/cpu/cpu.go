package cpu

import "log"

type Cpu struct {
}

func (c *Cpu) Name() string {
	return "cpu"
}

func (c *Cpu) Run() {
	log.Println(c.Name(), "Run")
	for i := 0; i < 1000000000; i++ {

	}
}
