package gameboy

import (
	"log"
)

type opCode struct {
	name string
	exec func(*Cpu)
}

type Cpu struct {
	debug      bool
	breakpoint uint16
	registers  struct {
		A, B, C, D, E, F, H, L byte
	}
	interrupts struct {
		master bool
	}
	bus            *Bus
	stackPointer   uint16
	programCounter uint16
	mClock         uint16
	tClock         uint16
}

func NewCpu(bus *Bus) *Cpu {
	var cpu *Cpu = &Cpu{debug: false, breakpoint: 0x0297, bus: bus}
	return cpu
}

func (cpu *Cpu) Dump() {
	log.Printf(`af: %x %x
		    bc: %x %x
		    de: %x %x
		    hl: %x %x
		    sp: %x
		    pc: %x
		    `,

		cpu.registers.A, cpu.registers.F, cpu.registers.B, cpu.registers.C,
		cpu.registers.D, cpu.registers.E, cpu.registers.H, cpu.registers.L,
		cpu.stackPointer, cpu.programCounter-1)
}

func (cpu *Cpu) Run() {
	//fetch
	/*
		if cpu.debug && cpu.programCounter == cpu.breakpoint {
			cpu.Dump()
			os.Exit(0)
		}*/
	var instruction byte = cpu.bus.memory.ReadByte(cpu.programCounter)
	//decode
	var op *opCode = baseInstructionSet[instruction]
	cpu.programCounter++
	if op.exec == nil {
		cpu.Dump()
		log.Fatalf("Unable to run instruction: %s %x", op.name, instruction)
	}
	//execute
	op.exec(cpu)

}

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

var baseInstructionSet []*opCode = []*opCode{
	&opCode{"NOP", func(cpu *Cpu) {
		cpu.mClock++
	}},
	&opCode{"LD BC,d16", func(cpu *Cpu) {
		cpu.registers.C = cpu.bus.memory.ReadByte(cpu.programCounter)
		cpu.registers.B = cpu.bus.memory.ReadByte(cpu.programCounter + 1)
		cpu.programCounter += 2
	}},
	&opCode{"LD (BC),A", func(cpu *Cpu) {
		cpu.bus.memory.WriteByte((uint16(cpu.registers.B)<<8)|uint16(cpu.registers.C), cpu.registers.A)
	}},
	&opCode{"INC BC", nil},
	&opCode{"INC B", nil},
	&opCode{"DEC B", func(cpu *Cpu) {
		if (cpu.registers.B & 0x0F) == 0 {
			cpu.setFlags('H')
		} else {
			cpu.clearFlags('H')
		}
		cpu.registers.B--
		cpu.setFlags('N')
		if cpu.registers.B == 0 {
			cpu.setFlags('Z')
		} else {
			cpu.clearFlags('Z')
		}
	}},
	&opCode{"LD B,d8", func(cpu *Cpu) {
		cpu.registers.B = cpu.bus.memory.ReadByte(cpu.programCounter)
		cpu.programCounter++
	}},
	&opCode{"RLCA", nil},
	&opCode{"LD (a16),SP", nil},
	&opCode{"ADD HL,BC", nil},
	&opCode{"LD A,(BC)", nil},
	&opCode{"DEC BC", nil},
	&opCode{"INC C", nil},
	&opCode{"DEC C", func(cpu *Cpu) {
		if (cpu.registers.C & 0x0F) == 0 {
			cpu.setFlags('H')
		} else {
			cpu.clearFlags('H')
		}
		cpu.registers.C--
		cpu.setFlags('N')
		if cpu.registers.C == 0 {
			cpu.setFlags('Z')
		} else {
			cpu.clearFlags('Z')
		}
	}},
	&opCode{"LD C,d8", func(cpu *Cpu) {
		cpu.registers.C = cpu.bus.memory.ReadByte(cpu.programCounter)
		cpu.programCounter++
	}},
	&opCode{"RRCA", nil},
	&opCode{"STOP 0", nil},
	&opCode{"LD DE,d16", nil},
	&opCode{"LD (DE),A", nil},
	&opCode{"INC DE", nil},
	&opCode{"INC D", nil},
	&opCode{"DEC D", nil},
	&opCode{"LD D,d8", nil},
	&opCode{"RLA", nil},
	&opCode{"JR r8", nil},
	&opCode{"ADD HL,DE", nil},
	&opCode{"LD A,(DE)", nil},
	&opCode{"DEC DE", nil},
	&opCode{"INC E", nil},
	&opCode{"DEC E", nil},
	&opCode{"LD E,d8", nil},
	&opCode{"RRA", nil},
	&opCode{"JR NZ,r8", func(cpu *Cpu) {
		if !cpu.isFlagSet('Z') {
			var label uint8 = cpu.bus.memory.ReadByte(cpu.programCounter)
			cpu.programCounter++
			label = ^label + 1
			cpu.programCounter -= uint16(label)
		} else {
			cpu.programCounter++
		}

	}},
	&opCode{"LD HL,d16", func(cpu *Cpu) {
		cpu.registers.L = cpu.bus.memory.ReadByte(cpu.programCounter)
		cpu.registers.H = cpu.bus.memory.ReadByte(cpu.programCounter + 1)
		cpu.programCounter += 2
	}},
	&opCode{"LD (HL+),A", nil},
	&opCode{"INC HL", nil},
	&opCode{"INC H", nil},
	&opCode{"DEC H", nil},
	&opCode{"LD H,d8", nil},
	&opCode{"DAA", nil},
	&opCode{"JR Z,r8", nil},
	&opCode{"ADD HL,HL", nil},
	&opCode{"LD A,(HL+)", nil},
	&opCode{"DEC HL", nil},
	&opCode{"INC L", nil},
	&opCode{"DEC L", nil},
	&opCode{"LD L,d8", nil},
	&opCode{"CPL", nil},
	&opCode{"JR NC,r8", nil},
	&opCode{"LD SP,d16", nil},
	&opCode{"LD (HL-),A", func(cpu *Cpu) {
		cpu.bus.memory.WriteByte((uint16(cpu.registers.H)<<8)+uint16(cpu.registers.L), cpu.registers.A)
		cpu.registers.L--
		//if register L was a 0 then decrementing it would cause it to be 255 so we need to take the one from the H
		if cpu.registers.L == 0xFF {
			cpu.registers.H--
		}
	}},
	&opCode{"INC SP", nil},
	&opCode{"INC (HL)", nil},
	&opCode{"DEC (HL)", nil},
	&opCode{"LD (HL),d8", nil},
	&opCode{"SCF", nil},
	&opCode{"JR C,r8", nil},
	&opCode{"ADD HL,SP", nil},
	&opCode{"LD A,(HL-)", nil},
	&opCode{"DEC SP", nil},
	&opCode{"INC A", nil},
	&opCode{"DEC A", nil},
	&opCode{"LD A,d8", func(cpu *Cpu) {
		cpu.registers.A = cpu.bus.memory.ReadByte(cpu.programCounter)
		cpu.programCounter++
	}},
	&opCode{"CCF", nil},
	&opCode{"LD B,B", nil},
	&opCode{"LD B,C", nil},
	&opCode{"LD B,D", nil},
	&opCode{"LD B,E", nil},
	&opCode{"LD B,H", nil},
	&opCode{"LD B,L", nil},
	&opCode{"LD B,(HL)", nil},
	&opCode{"LD B,A", nil},
	&opCode{"LD C,B", nil},
	&opCode{"LD C,C", nil},
	&opCode{"LD C,D", nil},
	&opCode{"LD C,E", nil},
	&opCode{"LD C,H", nil},
	&opCode{"LD C,L", nil},
	&opCode{"LD C,(HL)", nil},
	&opCode{"LD C,A", nil},
	&opCode{"LD D,B", nil},
	&opCode{"LD D,C", nil},
	&opCode{"LD D,D", nil},
	&opCode{"LD D,E", nil},
	&opCode{"LD D,H", nil},
	&opCode{"LD D,L", nil},
	&opCode{"LD D,(HL)", nil},
	&opCode{"LD D,A", nil},
	&opCode{"LD E,B", nil},
	&opCode{"LD E,C", nil},
	&opCode{"LD E,D", nil},
	&opCode{"LD E,E", nil},
	&opCode{"LD E,H", nil},
	&opCode{"LD E,L", nil},
	&opCode{"LD E,(HL)", nil},
	&opCode{"LD E,A", nil},
	&opCode{"LD H,B", nil},
	&opCode{"LD H,C", nil},
	&opCode{"LD H,D", nil},
	&opCode{"LD H,E", nil},
	&opCode{"LD H,H", nil},
	&opCode{"LD H,L", nil},
	&opCode{"LD H,(HL)", nil},
	&opCode{"LD H,A", nil},
	&opCode{"LD L,B", nil},
	&opCode{"LD L,C", nil},
	&opCode{"LD L,D", nil},
	&opCode{"LD L,E", nil},
	&opCode{"LD L,H", nil},
	&opCode{"LD L,L", nil},
	&opCode{"LD L,(HL)", nil},
	&opCode{"LD L,A", nil},
	&opCode{"LD (HL),B", nil},
	&opCode{"LD (HL),C", nil},
	&opCode{"LD (HL),D", nil},
	&opCode{"LD (HL),E", nil},
	&opCode{"LD (HL),H", nil},
	&opCode{"LD (HL),L", nil},
	&opCode{"HALT", nil},
	&opCode{"LD (HL),A", nil},
	&opCode{"LD A,B", nil},
	&opCode{"LD A,C", nil},
	&opCode{"LD A,D", nil},
	&opCode{"LD A,E", nil},
	&opCode{"LD A,H", nil},
	&opCode{"LD A,L", nil},
	&opCode{"LD A,(HL)", nil},
	&opCode{"LD A,A", nil},
	&opCode{"ADD A,B", nil},
	&opCode{"ADD A,C", nil},
	&opCode{"ADD A,D", nil},
	&opCode{"ADD A,E", nil},
	&opCode{"ADD A,H", nil},
	&opCode{"ADD A,L", nil},
	&opCode{"ADD A,(HL)", nil},
	&opCode{"ADD A,A", nil},
	&opCode{"ADC A,B", nil},
	&opCode{"ADC A,C", nil},
	&opCode{"ADC A,D", nil},
	&opCode{"ADC A,E", nil},
	&opCode{"ADC A,H", nil},
	&opCode{"ADC A,L", nil},
	&opCode{"ADC A,(HL)", nil},
	&opCode{"ADC A,A", nil},
	&opCode{"SUB B", nil},
	&opCode{"SUB C", nil},
	&opCode{"SUB D", nil},
	&opCode{"SUB E", nil},
	&opCode{"SUB H", nil},
	&opCode{"SUB L", nil},
	&opCode{"SUB (HL)", nil},
	&opCode{"SUB A", nil},
	&opCode{"SBC A,B", nil},
	&opCode{"SBC A,C", nil},
	&opCode{"SBC A,D", nil},
	&opCode{"SBC A,E", nil},
	&opCode{"SBC A,H", nil},
	&opCode{"SBC A,L", nil},
	&opCode{"SBC A,(HL)", nil},
	&opCode{"SBC A,A", nil},
	&opCode{"AND B", nil},
	&opCode{"AND C", nil},
	&opCode{"AND D", nil},
	&opCode{"AND E", nil},
	&opCode{"AND H", nil},
	&opCode{"AND L", nil},
	&opCode{"AND (HL)", nil},
	&opCode{"AND A", nil},
	&opCode{"XOR B", nil},
	&opCode{"XOR C", nil},
	&opCode{"XOR D", nil},
	&opCode{"XOR E", nil},
	&opCode{"XOR H", nil},
	&opCode{"XOR L", nil},
	&opCode{"XOR (HL)", nil},
	&opCode{"XOR A", func(cpu *Cpu) {
		cpu.registers.A ^= cpu.registers.A
		if cpu.registers.A == 0 {
			cpu.setFlags('Z')
		} else {
			cpu.clearFlags('Z')
		}
	}},
	&opCode{"OR B", nil},
	&opCode{"OR C", nil},
	&opCode{"OR D", nil},
	&opCode{"OR E", nil},
	&opCode{"OR H", nil},
	&opCode{"OR L", nil},
	&opCode{"OR (HL)", nil},
	&opCode{"OR A", nil},
	&opCode{"CP B", nil},
	&opCode{"CP C", nil},
	&opCode{"CP D", nil},
	&opCode{"CP E", nil},
	&opCode{"CP H", nil},
	&opCode{"CP L", nil},
	&opCode{"CP (HL)", nil},
	&opCode{"CP A", nil},
	&opCode{"RET NZ", nil},
	&opCode{"POP BC", nil},
	&opCode{"JP NZ,a16", nil},
	&opCode{"JP a16", func(cpu *Cpu) {
		cpu.programCounter = cpu.bus.memory.ReadWord(cpu.programCounter)
	}},
	&opCode{"CALL NZ,a16", nil},
	&opCode{"PUSH BC", nil},
	&opCode{"ADD A,d8", nil},
	&opCode{"RST 00H", nil},
	&opCode{"RET Z", nil},
	&opCode{"RET", nil},
	&opCode{"JP Z,a16", nil},
	&opCode{"PREFIX CB", nil},
	&opCode{"CALL Z,a16", nil},
	&opCode{"CALL a16", nil},
	&opCode{"ADC A,d8", nil},
	&opCode{"RST 08H", nil},
	&opCode{"RET NC", nil},
	&opCode{"POP DE", nil},
	&opCode{"JP NC,a16", nil},
	&opCode{" ", nil},
	&opCode{"CALL NC,a16", nil},
	&opCode{"PUSH DE", nil},
	&opCode{"SUB d8", nil},
	&opCode{"RST 10H", nil},
	&opCode{"RET C", nil},
	&opCode{"RETI", nil},
	&opCode{"JP C,a16", nil},
	&opCode{" ", nil},
	&opCode{"CALL C,a16", nil},
	&opCode{" ", nil},
	&opCode{"SBC A,d8", nil},
	&opCode{"RST 18H", nil},
	&opCode{"LDH (a8),A", func(cpu *Cpu) {
		cpu.bus.memory.WriteByte(0xFF00+uint16(cpu.bus.memory.ReadByte(cpu.programCounter)), cpu.registers.A)
		cpu.programCounter++
	}},
	&opCode{"POP HL", nil},
	&opCode{"LD (C),A", nil},
	&opCode{" ", nil},
	&opCode{" ", nil},
	&opCode{"PUSH HL", nil},
	&opCode{"AND d8", nil},
	&opCode{"RST 20H", nil},
	&opCode{"ADD SP,r8", nil},
	&opCode{"JP (HL)", nil},
	&opCode{"LD (a16),A", nil},
	&opCode{" ", nil},
	&opCode{" ", nil},
	&opCode{" ", nil},
	&opCode{"XOR d8", nil},
	&opCode{"RST 28H", nil},
	&opCode{"LDH A,(a8)", func(cpu *Cpu) {
		cpu.registers.A = cpu.bus.memory.ReadByte(0xFF00 + uint16(cpu.bus.memory.ReadByte(cpu.programCounter)))
		cpu.programCounter++
	}},
	&opCode{"POP AF", nil},
	&opCode{"LD A,(C)", nil},
	&opCode{"DI", func(cpu *Cpu) {
		cpu.interrupts.master = false
	}},
	&opCode{" ", nil},
	&opCode{"PUSH AF", nil},
	&opCode{"OR d8", nil},
	&opCode{"RST 30H", nil},
	&opCode{"LD HL,SP+r8", nil},
	&opCode{"LD SP,HL", nil},
	&opCode{"LD A,(a16)", nil},
	&opCode{"EI", nil},
	&opCode{" ", nil},
	&opCode{" ", nil},
	&opCode{"CP d8", nil},
	&opCode{"RST 38H", nil},
}

/*
func (cpu *Cpu) RLCA() {
	//http://karma.ticalc.org/guide/lesson14.html
	var carryBit byte = (cpu.registers.A & 0x80) >> 7
	cpu.flags.C = carryBit == 1
	cpu.registers.A <<= 1
	cpu.registers.A += carryBit
	cpu.flags.N, cpu.flags.Z, cpu.flags.H = false, false, false
}
*/

/*used with any instruction that has 0xCB. Another byte is read and decoded and run in the extended
instruction set.*/
var extendedInstructionSet []*opCode = []*opCode{}
