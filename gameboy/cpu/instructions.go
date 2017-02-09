package cpu

import (
	"fmt"
	"os"
)

var baseInstructionSet []*opCode = []*opCode{
	&opCode{"NOP", func(cpu *Cpu) {
	}},
	&opCode{"LD BC,d16", func(cpu *Cpu) {
		cpu.registers.C = cpu.bus.ReadByte(cpu.programCounter)
		cpu.registers.B = cpu.bus.ReadByte(cpu.programCounter + 1)
		cpu.programCounter += 2
	}},
	&opCode{"LD (BC),A", func(cpu *Cpu) {
		cpu.bus.WriteByte((uint16(cpu.registers.B)<<8)|uint16(cpu.registers.C), cpu.registers.A)
	}},
	&opCode{"INC BC", func(cpu *Cpu) {
		cpu.incrementRegisters(&cpu.registers.B, &cpu.registers.C)
	}},
	&opCode{"INC B", func(cpu *Cpu) {
		cpu.incrementRegister(&cpu.registers.B)
	}},
	&opCode{"DEC B", func(cpu *Cpu) {
		cpu.decrementRegister(&cpu.registers.B)
	}},
	&opCode{"LD B,d8", func(cpu *Cpu) {
		cpu.registers.B = cpu.bus.ReadByte(cpu.programCounter)
		cpu.programCounter++
	}},
	&opCode{"RLCA", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (a16),SP", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD HL,BC", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,(BC)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"DEC BC", func(cpu *Cpu) {
		cpu.decrementRegisters(&cpu.registers.B, &cpu.registers.C)
	}},
	&opCode{"INC C", func(cpu *Cpu) {
		cpu.incrementRegister(&cpu.registers.C)
	}},
	&opCode{"DEC C", func(cpu *Cpu) {
		cpu.decrementRegister(&cpu.registers.C)
	}},
	&opCode{"LD C,d8", func(cpu *Cpu) {
		cpu.registers.C = cpu.bus.ReadByte(cpu.programCounter)
		cpu.programCounter++
	}},
	&opCode{"RRCA", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"STOP 0", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD DE,d16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (DE),A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"INC DE", func(cpu *Cpu) {
		cpu.incrementRegisters(&cpu.registers.D, &cpu.registers.E)
	}},
	&opCode{"INC D", func(cpu *Cpu) {
		cpu.incrementRegister(&cpu.registers.D)
	}},
	&opCode{"DEC D", func(cpu *Cpu) {
		cpu.decrementRegister(&cpu.registers.D)
	}},
	&opCode{"LD D,d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RLA", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JR r8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD HL,DE", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,(DE)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"DEC DE", func(cpu *Cpu) {
		cpu.decrementRegisters(&cpu.registers.D, &cpu.registers.E)
	}},
	&opCode{"INC E", func(cpu *Cpu) {
		cpu.incrementRegister(&cpu.registers.E)
	}},
	&opCode{"DEC E", func(cpu *Cpu) {
		cpu.decrementRegister(&cpu.registers.E)
	}},
	&opCode{"LD E,d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RRA", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JR NZ,r8", func(cpu *Cpu) {
		if !cpu.isFlagSet('Z') {
			var label uint8 = cpu.bus.ReadByte(cpu.programCounter)
			cpu.programCounter++
			if label > 127 {
				label = (^label + 1)
				cpu.programCounter -= uint16(label)
			} else {

				cpu.programCounter += uint16(label)
			}
			cpu.clock += 4
		} else {
			cpu.programCounter++
		}

	}},
	&opCode{"LD HL,d16", func(cpu *Cpu) {
		cpu.registers.L = cpu.bus.ReadByte(cpu.programCounter)
		cpu.registers.H = cpu.bus.ReadByte(cpu.programCounter + 1)
		cpu.programCounter += 2
	}},
	&opCode{"LD (HL+),A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"INC HL", func(cpu *Cpu) {
		cpu.incrementRegisters(&cpu.registers.H, &cpu.registers.L)
	}},
	&opCode{"INC H", func(cpu *Cpu) {
		cpu.incrementRegister(&cpu.registers.H)
	}},
	&opCode{"DEC H", func(cpu *Cpu) {
		cpu.decrementRegister(&cpu.registers.H)
	}},
	&opCode{"LD H,d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"DAA", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JR Z,r8", func(cpu *Cpu) {
		if cpu.isFlagSet('Z') {
			var label uint8 = cpu.bus.ReadByte(cpu.programCounter)
			cpu.programCounter++
			if label > 127 {
				label = (^label + 1)
				cpu.programCounter -= uint16(label)
			} else {

				cpu.programCounter += uint16(label)
			}
			cpu.clock += 4
		} else {
			cpu.programCounter++
		}
	}},
	&opCode{"ADD HL,HL", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,(HL+)", func(cpu *Cpu) {
		cpu.registers.A = cpu.bus.ReadByte((uint16(cpu.registers.H) << 8) + uint16(cpu.registers.L))
		cpu.incrementRegisters(&cpu.registers.H, &cpu.registers.L)
	}},
	&opCode{"DEC HL", func(cpu *Cpu) {
		cpu.decrementRegisters(&cpu.registers.H, &cpu.registers.L)
	}},
	&opCode{"INC L", func(cpu *Cpu) {
		cpu.incrementRegister(&cpu.registers.L)
	}},
	&opCode{"DEC L", func(cpu *Cpu) {
		cpu.decrementRegister(&cpu.registers.L)
	}},
	&opCode{"LD L,d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CPL", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JR NC,r8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD SP,d16", func(cpu *Cpu) {
		cpu.stackPointer = cpu.bus.ReadWord(cpu.programCounter)
		cpu.programCounter += 2
	}},
	&opCode{"LD (HL-),A", func(cpu *Cpu) {
		cpu.bus.WriteByte((uint16(cpu.registers.H)<<8)+uint16(cpu.registers.L), cpu.registers.A)
		cpu.decrementRegisters(&cpu.registers.H, &cpu.registers.L)
	}},
	&opCode{"INC SP", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"INC (HL)", func(cpu *Cpu) {
		b := cpu.bus.ReadByte((uint16(cpu.registers.H) << 8) + uint16(cpu.registers.L))
		cpu.incrementRegister(&b)
		cpu.bus.WriteByte((uint16(cpu.registers.H)<<8)+uint16(cpu.registers.L), b)

	}},
	&opCode{"DEC (HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),d8", func(cpu *Cpu) {
		cpu.bus.WriteByte(
			(uint16(cpu.registers.H)<<8)+uint16(cpu.registers.L),
			cpu.bus.ReadByte(cpu.programCounter),
		)
		cpu.programCounter++
	}},
	&opCode{"SCF", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JR C,r8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD HL,SP", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,(HL-)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"DEC SP", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"INC A", func(cpu *Cpu) {
		cpu.incrementRegister(&cpu.registers.A)
	}},
	&opCode{"DEC A", func(cpu *Cpu) {
		cpu.decrementRegister(&cpu.registers.A)
	}},
	&opCode{"LD A,d8", func(cpu *Cpu) {
		cpu.registers.A = cpu.bus.ReadByte(cpu.programCounter)
		cpu.programCounter++
	}},
	&opCode{"CCF", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD B,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD C,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD D,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD E,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD H,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD L,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"HALT", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (HL),A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,B", func(cpu *Cpu) {
		cpu.registers.A = cpu.registers.B
	}},
	&opCode{"LD A,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD A,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADC A,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB (HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SUB A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,(HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"AND B", func(cpu *Cpu) {
		cpu.andRegister(cpu.registers.B)
	}},
	&opCode{"AND C", func(cpu *Cpu) {
		cpu.andRegister(cpu.registers.C)
	}},
	&opCode{"AND D", func(cpu *Cpu) {
		cpu.andRegister(cpu.registers.D)
	}},
	&opCode{"AND E", func(cpu *Cpu) {
		cpu.andRegister(cpu.registers.E)
	}},
	&opCode{"AND H", func(cpu *Cpu) {
		cpu.andRegister(cpu.registers.H)
	}},
	&opCode{"AND L", func(cpu *Cpu) {
		cpu.andRegister(cpu.registers.L)
	}},
	&opCode{"AND (HL)", func(cpu *Cpu) {
		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"AND A", func(cpu *Cpu) {
		cpu.andRegister(cpu.registers.A)
	}},
	&opCode{"XOR B", func(cpu *Cpu) {
		cpu.xorRegister(cpu.registers.B)
	}},
	&opCode{"XOR C", func(cpu *Cpu) {
		cpu.xorRegister(cpu.registers.C)
	}},
	&opCode{"XOR D", func(cpu *Cpu) {
		cpu.xorRegister(cpu.registers.D)
	}},
	&opCode{"XOR E", func(cpu *Cpu) {
		cpu.xorRegister(cpu.registers.E)
	}},
	&opCode{"XOR H", func(cpu *Cpu) {
		cpu.xorRegister(cpu.registers.H)
	}},
	&opCode{"XOR L", func(cpu *Cpu) {
		cpu.xorRegister(cpu.registers.L)
	}},
	&opCode{"XOR (HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"XOR A", func(cpu *Cpu) {
		cpu.xorRegister(cpu.registers.A)
	}},
	&opCode{"OR B", func(cpu *Cpu) {
		cpu.orRegister(cpu.registers.B)
	}},
	&opCode{"OR C", func(cpu *Cpu) {
		cpu.orRegister(cpu.registers.C)
	}},
	&opCode{"OR D", func(cpu *Cpu) {
		cpu.orRegister(cpu.registers.D)
	}},
	&opCode{"OR E", func(cpu *Cpu) {
		cpu.orRegister(cpu.registers.E)
	}},
	&opCode{"OR H", func(cpu *Cpu) {
		cpu.orRegister(cpu.registers.H)
	}},
	&opCode{"OR L", func(cpu *Cpu) {
		cpu.orRegister(cpu.registers.L)
	}},
	&opCode{"OR (HL)", func(cpu *Cpu) {
		cpu.orRegister(cpu.bus.ReadByte((uint16(cpu.registers.H) << 8) | uint16(cpu.registers.L)))
	}},
	&opCode{"OR A", func(cpu *Cpu) {
		cpu.orRegister(cpu.registers.A)
	}},
	&opCode{"CP B", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP D", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP E", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP L", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP (HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP A", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RET NZ", func(cpu *Cpu) {
		if !cpu.isFlagSet('Z') {
			cpu.programCounter = cpu.fromStack()
			cpu.clock += 12
		}
	}},
	&opCode{"POP BC", func(cpu *Cpu) {
		word := cpu.fromStack()
		cpu.registers.C = byte((0xFF00 & word) >> 8)
		cpu.registers.B = byte(0x00FF & word)
	}},
	&opCode{"JP NZ,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JP a16", func(cpu *Cpu) {
		cpu.programCounter = cpu.bus.ReadWord(cpu.programCounter)
	}},
	&opCode{"CALL NZ,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"PUSH BC", func(cpu *Cpu) {
		cpu.toStack((uint16(cpu.registers.B) << 8) | uint16(cpu.registers.C))
	}},
	&opCode{"ADD A,d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RST 00H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RET Z", func(cpu *Cpu) {
		if cpu.isFlagSet('Z') {
			cpu.programCounter = cpu.fromStack()
			cpu.clock += 12
		}
	}},
	&opCode{"RET", func(cpu *Cpu) {
		cpu.programCounter = cpu.fromStack()
	}},
	&opCode{"JP Z,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"PREFIX CB", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CALL Z,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CALL a16", func(cpu *Cpu) {
		//save  current program counter to stack and prepare to jump
		jumpAddress := cpu.bus.ReadWord(cpu.programCounter)
		cpu.programCounter += 2
		cpu.toStack(cpu.programCounter)
		cpu.programCounter = jumpAddress
	}},
	&opCode{"ADC A,d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RST 08H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RET NC", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"POP DE", func(cpu *Cpu) {
		word := cpu.fromStack()
		cpu.registers.E = byte((0xFF00 & word) >> 8)
		cpu.registers.D = byte(0x00FF & word)
	}},
	&opCode{"JP NC,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CALL NC,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"PUSH DE", func(cpu *Cpu) {
		cpu.toStack((uint16(cpu.registers.D) << 8) | uint16(cpu.registers.E))
	}},
	&opCode{"SUB d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RST 10H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RET C", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RETI", func(cpu *Cpu) {
		cpu.masterInterrupt = true
		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JP C,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CALL C,a16", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"SBC A,d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RST 18H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LDH (a8),A", func(cpu *Cpu) {
		cpu.bus.WriteByte(0xFF00+uint16(cpu.bus.ReadByte(cpu.programCounter)), cpu.registers.A)
		cpu.programCounter++
	}},
	&opCode{"POP HL", func(cpu *Cpu) {
		word := cpu.fromStack()
		cpu.registers.L = byte((0xFF00 & word) >> 8)
		cpu.registers.H = byte(0x00FF & word)
	}},
	&opCode{"LD (C),A", func(cpu *Cpu) {
		cpu.bus.WriteByte(0xFF00+uint16(cpu.registers.C), cpu.registers.A)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"PUSH HL", func(cpu *Cpu) {
		cpu.toStack((uint16(cpu.registers.L) << 8) | uint16(cpu.registers.H))
	}},
	&opCode{"AND d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RST 20H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"ADD SP,r8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"JP (HL)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD (a16),A", func(cpu *Cpu) {
		cpu.bus.WriteByte(
			cpu.bus.ReadWord(cpu.programCounter),
			cpu.registers.A,
		)
		cpu.programCounter += 2
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"XOR d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RST 28H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LDH A,(a8)", func(cpu *Cpu) {
		location := 0xFF00 + uint16(cpu.bus.ReadByte(cpu.programCounter))
		cpu.registers.A = cpu.bus.ReadByte(location)
		cpu.programCounter++
	}},
	&opCode{"POP AF", func(cpu *Cpu) {
		word := cpu.fromStack()
		cpu.registers.F = byte((0xFF00 & word) >> 8)
		cpu.registers.A = byte(0x00FF & word)
	}},
	&opCode{"LD A,(C)", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"DI", func(cpu *Cpu) {
		cpu.masterInterrupt = false
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"PUSH AF", func(cpu *Cpu) {
		cpu.toStack((uint16(cpu.registers.A) << 8) | uint16(cpu.registers.F))
	}},
	&opCode{"OR d8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"RST 30H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD HL,SP+r8", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD SP,HL", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"LD A,(a16)", func(cpu *Cpu) {
		cpu.registers.A = cpu.bus.ReadByte(cpu.bus.ReadWord(cpu.programCounter))
		cpu.programCounter += 2
	}},
	&opCode{"EI", func(cpu *Cpu) {
		cpu.masterInterrupt = true
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{" ", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
	&opCode{"CP d8", func(cpu *Cpu) {
		value := cpu.bus.ReadByte(cpu.programCounter)
		aCopy := int(cpu.registers.A) - int(value)
		cpu.clearFlags('N', 'Z', 'C', 'H')
		fmt.Println(aCopy, value)
		cpu.setFlags('N')
		if aCopy < 0 {
			cpu.setFlags('C')
		} else if aCopy == 0 {
			cpu.setFlags('Z')
		}

		if (value & 0x0F) > (cpu.registers.A & 0x0F) {
			cpu.setFlags('H')
		}
		cpu.programCounter++
	}},
	&opCode{"RST 38H", func(cpu *Cpu) {

		cpu.Dump()
		os.Exit(1)
	}},
}

var ticksPerInstruction []uint8 = []uint8{
	4, 12, 8, 8, 4, 4, 8, 8, 20, 8, 8, 8, 4, 4, 8, 8,
	4, 12, 8, 8, 4, 4, 8, 8, 8, 8, 8, 8, 4, 4, 8, 8,
	8, 12, 8, 8, 4, 4, 8, 4, 8, 8, 8, 8, 4, 4, 8, 4,
	8, 12, 8, 8, 12, 12, 12, 4, 0, 8, 8, 8, 4, 4, 8, 4,
	4, 4, 4, 4, 4, 4, 8, 4, 4, 4, 4, 4, 4, 4, 8, 4,
	4, 4, 4, 4, 4, 4, 8, 4, 4, 4, 4, 4, 4, 4, 8, 4,
	4, 4, 4, 4, 4, 4, 8, 4, 4, 4, 4, 4, 4, 4, 8, 4,
	8, 8, 8, 8, 8, 8, 4, 8, 4, 4, 4, 4, 4, 4, 8, 4,
	4, 4, 4, 4, 4, 4, 8, 4, 4, 4, 4, 4, 4, 4, 8, 4,
	4, 4, 4, 4, 4, 4, 8, 4, 4, 4, 4, 4, 4, 4, 8, 4,
	4, 4, 4, 4, 4, 4, 8, 4, 4, 4, 4, 4, 4, 4, 8, 4,
	4, 4, 4, 4, 4, 4, 8, 4, 4, 4, 4, 4, 4, 4, 8, 4,
	8, 12, 0, 12, 0, 16, 8, 16, 8, 4, 0, 0, 0, 12, 8, 16,
	0, 12, 0, 0, 0, 16, 8, 16, 0, 16, 0, 0, 0, 0, 8, 16,
	12, 12, 8, 0, 0, 16, 8, 16, 16, 4, 16, 0, 0, 0, 8, 16,
	12, 12, 8, 4, 0, 16, 8, 16, 12, 8, 16, 4, 0, 0, 8, 16,
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
