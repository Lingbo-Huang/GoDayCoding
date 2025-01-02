package mem

import (
	"log"
	"pprof/constants"
)

type Mem struct {
	buffer [][constants.Mi]byte
}

func (m *Mem) Name() string {
	return "mem"
}

func (m *Mem) Run() {
	log.Println(m.Name(), "Run")
	max := constants.Gi
	for len(m.buffer)*constants.Mi < max {
		m.buffer = append(m.buffer, [constants.Mi]byte{})
	}
}
