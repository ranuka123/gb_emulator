package gameboy

import "gameboyemu/gameboy/cpu"

var paletteOptions = [4][4]uint8{{255, 255, 255, 255}, {192, 192, 192, 255}, {96, 96, 96, 255}, {0, 0, 0, 255}}

//Gpu modes
var modes map[uint8]func(gpu *Gpu) = map[uint8]func(gpu *Gpu){
	0: (*Gpu).hBlank,
	1: (*Gpu).vBlank,
	2: (*Gpu).readOAM,
	3: (*Gpu).readVRAM,
}

type Gpu struct {
	bus         *Bus
	mode        uint8
	clock       uint16
	currentLine uint8
	scrollY     uint8
	scrollX     uint8
	control     uint8
	tiles       [384][8][8]byte
	palette     [4][4]uint8
}

func NewGpu(bus *Bus) (gpu *Gpu) {
	gpu = &Gpu{bus, 0, 0, 0, 0, 0, 0, [384][8][8]byte{}, [4][4]uint8{}}
	return gpu
}

func (gpu *Gpu) Update() {
	gpu.clock += gpu.bus.cpu.Clock()
	modes[gpu.mode](gpu)
}

func (gpu *Gpu) SetTile(addr uint16) {
	addr &= 0x1FFE
	var tile uint16 = (addr >> 4) & 511
	var y uint16 = (addr >> 1) & 7
	var index uint8
	var lowBit, highBit uint8
	for x := 0; x < 8; x++ {
		index = 1 << uint8(7-x) //check this, might be an issue
		lowBit = gpu.bus.memory.ReadByte(addr) & index
		highBit = gpu.bus.memory.ReadByte(addr+1) & index
		if lowBit != 0 {
			lowBit = 1
		}
		if highBit != 0 {
			highBit = 2
		}

		pixelValue := uint8(lowBit + highBit) //pixel value will range from 0-4
		gpu.tiles[tile][y][x] = pixelValue
	}

}

//Register functions

func (gpu *Gpu) SetControl(value uint8) {
	gpu.control = value
}

func (gpu *Gpu) SetScrollY(value uint8) {
	gpu.scrollY = value
}

func (gpu *Gpu) SetScrollX(value uint8) {
	gpu.scrollX = value
}

func (gpu *Gpu) SetPalette(value uint8) {
	for i := 0; i < 4; i++ {
		gpu.palette[i] = paletteOptions[(value>>uint8(i*2))&3]
	}
}

func (gpu *Gpu) GetCurrentLine() uint8 {
	return gpu.currentLine
}

//Gpu modes
func (gpu *Gpu) readOAM() {
	if gpu.clock >= 80 {
		gpu.clock, gpu.mode = 0, 3
	}
}

func (gpu *Gpu) readVRAM() {
	gpu.mode, gpu.clock = 0, 0
	if gpu.clock >= 172 {
		gpu.mode, gpu.clock = 0, 0
		//more tile logic
	}
}

func (gpu *Gpu) hBlank() {
	if gpu.clock >= 204 {
		gpu.currentLine++
		if gpu.currentLine == 143 {
			gpu.mode = 1
			//time to call vblank interrupt
			if interruptsEnabled := gpu.bus.memory.ReadByte(0xFFFF); interruptsEnabled&cpu.VBLANK != 0 {
				interruptFlags := gpu.bus.memory.ReadByte(0xFF0F) | cpu.VBLANK
				gpu.bus.memory.WriteByte(0xFF0F, interruptFlags)
			}
		} else {
			gpu.mode = 2
		}
		gpu.mode, gpu.clock = 0, 0
	}

}

func (gpu *Gpu) vBlank() {
	if gpu.clock >= 456 {
		gpu.clock = 0
		gpu.currentLine++
		if gpu.currentLine > 153 {
			gpu.mode = 2
			gpu.currentLine = 0
			//leftoff here
		}
	}
}
