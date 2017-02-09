package cpu

import "fmt"

//common operations used in instructions defined here
func (cpu *Cpu) andRegister(register byte) {
	cpu.registers.A &= register
	if cpu.registers.A == 0 {
		cpu.setFlags('Z')
	} else {
		cpu.clearFlags('Z')
	}
	cpu.clearFlags('C', 'N', 'H')

}

func (cpu *Cpu) orRegister(register byte) {
	cpu.registers.A |= register
	if cpu.registers.A == 0 {
		cpu.setFlags('Z')
	} else {
		cpu.clearFlags('Z')
	}
	cpu.clearFlags('C', 'N', 'H')

}

func (cpu *Cpu) incrementRegister(register *byte) {
	cpu.clearFlags('N')
	if (*register & 0x0F) == 0x0F {
		cpu.setFlags('H')
	} else {
		cpu.clearFlags('H')
	}
	*register++
	if *register == 0 {
		cpu.setFlags('Z')
	} else {
		cpu.clearFlags('Z')
	}
}

func (cpu *Cpu) incrementRegisters(registerA *byte, registerB *byte) {
	*registerB++
	if *registerB == 0 {
		*registerA++
	}
}

func (cpu *Cpu) decrementRegister(register *byte) {
	if (*register & 0x0F) == 0 {
		cpu.setFlags('H')
	} else {
		cpu.clearFlags('H')
	}
	*register--
	cpu.setFlags('N')
	if *register == 0 {
		cpu.setFlags('Z')
	} else {
		cpu.clearFlags('Z')
	}
}

func (cpu *Cpu) decrementRegisters(registerA *byte, registerB *byte) {
	*registerB--
	//if register B was a 0 then decrementing it would cause it to be 255 so we need to take the one from the A
	if *registerB == 0xFF {
		*registerA--
	}
}

func (cpu *Cpu) xorRegister(register byte) {
	cpu.registers.A ^= register
	if cpu.registers.A == 0 {
		cpu.setFlags('Z')
	} else {
		cpu.clearFlags('Z')
	}
	cpu.clearFlags('C', 'N', 'H')
}

func (cpu *Cpu) toStack(word uint16) {
	cpu.stackPointer -= 2
	fmt.Printf("Word %d written at %d", word, cpu.stackPointer)
	cpu.bus.WriteWord(cpu.stackPointer, word)
}

func (cpu *Cpu) fromStack() uint16 {
	word := cpu.bus.ReadWord(cpu.stackPointer)
	fmt.Printf("Word %d fetched from %d\n", word, cpu.stackPointer)
	cpu.stackPointer += 2
	return word
}
