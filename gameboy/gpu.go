package gameboy

type Gpu struct {
	bus         *Bus
	mode        uint8
	clock       uint16
	currentLine uint8
	tiles       [384][8][8]byte
}

func NewGpu(bus *Bus) (gpu *Gpu) {
	gpu = &Gpu{bus, 0, 0, 0, [384][8][8]byte{}}
	return gpu
}

func (gpu *Gpu) Run() {

}

func (gpu *Gpu) UpdateTile(addr uint16) {
	addr &= 0x1FFE
	var tile uint16 = (addr >> 4) & 511
	var y uint16 = (addr >> 1) & 7
	var index uint16
	var lowBit, highBit uint8
	for x := 0; x < 8; x++ {
		index = 1 << uint16(7-x) //check this, might be an issue
		lowBit = gpu.bus.memory.ram[addr] & index
		highBit = gpu.bus.memory.ram[addr+1] & index
		if lowBit != 0 {
			lowBit = 1
		}
		if highBit != 0 {
			highBit = 2
		}

		pixelValue := lowBit + highBit
		gpu.tiles[tile][y][x] = pixelValue
	}

}

//Gpu modes
func (gpu *Gpu) readOAM() {
	gpu.mode, gpu.clock = 0, 0
}

func (gpu *Gpu) readVRAM() {
	gpu.mode, gpu.clock = 0, 0
	//gpu.Scan()
}

func (gpu *Gpu) hBlank() {
	gpu.mode, gpu.clock = 0, 0
	gpu.currentLine++
	if gpu.currentLine == 143 {
		gpu.mode = 1
		//screen.Write(gpu screen, 0, 0)
	} else {
		gpu.mode = 2
	}
}

func (gpu *Gpu) vBlank() {
	if gpu.clock >= 456 {
		gpu.clock = 0
		gpu.currentLine++
		if gpu.currentLine > 153 {
			gpu.mode = 2
			gpu.currentLine = 0
		}
	}
}
