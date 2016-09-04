package gameboy

func (cpu *Cpu) setFlags(flags ...byte) {
	for _, flag := range flags {
		switch flag {
		case 'Z':
			cpu.registers.F |= 0x80
		case 'N':
			cpu.registers.F |= 0x40
		case 'H':
			cpu.registers.F |= 0x20
		case 'C':
			cpu.registers.F |= 0x10
		}
	}
}

func (cpu *Cpu) clearFlags(flags ...byte) {
	for _, flag := range flags {
		switch flag {
		case 'Z':
			cpu.registers.F &= ^uint8(0x80)
		case 'N':
			cpu.registers.F &= ^uint8(0x40)
		case 'H':
			cpu.registers.F &= ^uint8(0x20)
		case 'C':
			cpu.registers.F &= ^uint8(0x10)
		}
	}
}

func (cpu *Cpu) isFlagSet(flag byte) bool {
	if flag == 'Z' {
		return (cpu.registers.F & 0x80) != 0
	} else if flag == 'N' {
		return (cpu.registers.F & 0x40) != 0
	} else if flag == 'H' {
		return (cpu.registers.F & 0x20) != 0
	} else {
		return (cpu.registers.F & 0x10) != 0
	}

}
