package gameboy

import (
	"gameboyemu/gameboy/cpu"
	"io/ioutil"
	"log"
)

type Gameboy struct {
	bus *Bus
}

//allows the components in the gameboy to communicate
//with other components
type Bus struct {
	cpu    *cpu.Cpu
	memory *Memory
	gpu    *Gpu
}

type CpuBus struct {
	*Memory
}

const MAX_CYCLES_PER_SECOND int = 69905

func New(bpAddr uint16) (gameboy *Gameboy) {
	var bus *Bus = &Bus{nil, nil, nil}
	memory := NewMemory(bus)
	var cpuBus CpuBus = CpuBus{memory}
	cpu, gpu := cpu.NewCpu(cpuBus, 0x100, bpAddr), NewGpu(bus)
	bus.memory, bus.cpu, bus.gpu = memory, cpu, gpu
	gameboy = &Gameboy{bus}
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
		gameboy.bus.memory.WriteByte(i, file[i])
	}
	log.Println("ROM loaded.")
}

func (gameboy *Gameboy) Run() {

	for {
		gameboy.bus.cpu.Update()
		gameboy.bus.gpu.Update()
		gameboy.bus.cpu.CheckInterrupts()
	}
}
