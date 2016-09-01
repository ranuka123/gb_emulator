package gameboy

import (
	"io/ioutil"
	"log"
)

type Gameboy struct {
	cpu    *Cpu
	memory *Memory
}

func New() (gameboy *Gameboy) {
	var memory *Memory = NewMemory()
	gameboy = &Gameboy{NewCpu(memory), memory}
	gameboy.cpu.programCounter = 0x100
	return gameboy
}

func (gameboy *Gameboy) LoadRom(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if len(file) > 0x7FFF {
		//if it exceeds need to do dump rest in the cartridge ram
		log.Printf("File exceeded ROM bank. Size %d", len(file))
	}

	var i uint16
	for i = 0; i < 0x7FFF; i++ {
		gameboy.memory.WriteByte(i, file[i])
	}
	log.Println("ROM loaded.")
}

func (gameboy *Gameboy) Run() {
	for {
		gameboy.cpu.FetchDecodeExecute()
	}
}
