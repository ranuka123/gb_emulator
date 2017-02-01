package cpu

import (
	"log"
	"os"
)

type opCode struct {
	name string
	exec func(*Cpu)
}

type Cpu struct {
	debug      uint
	breakpoint uint16
	registers  struct {
		A, B, C, D, E, F, H, L byte
	}
	masterInterrupt bool
	bus             Bus
	stackPointer    uint16
	programCounter  uint16
	clock           uint16
}

func NewCpu(bus Bus, counter uint16, breakpoint uint16) *Cpu {
	var cpu *Cpu = &Cpu{debug: 1, breakpoint: breakpoint, bus: bus, programCounter: counter}
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
		cpu.stackPointer, cpu.programCounter)
}

func (cpu *Cpu) Update() {
	//fetch
	var instruction byte = cpu.bus.ReadByte(cpu.programCounter)
	//decode
	var op *opCode = baseInstructionSet[instruction]
	if cpu.debug == 1 {
		log.Printf("%s %x \n", op.name, cpu.programCounter)
	}
	if cpu.debug == 1 && cpu.programCounter == cpu.breakpoint {
		log.Printf("Stopped at instruction %s: \n", op.name)
		cpu.Dump()
		os.Exit(0)
	}
	if ticksPerInstruction[instruction] == 0 {
		cpu.Dump()
		log.Fatalf("0 tick instruction found %s", op.name)
	}
	cpu.clock += uint16(ticksPerInstruction[instruction])
	cpu.programCounter++
	//execute
	op.exec(cpu)

}

func (cpu *Cpu) Clock() uint16 {
	return cpu.clock
}
