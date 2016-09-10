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
	debug      bool
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

func NewCpu(bus Bus, counter uint16) *Cpu {
	var cpu *Cpu = &Cpu{debug: false, breakpoint: 0x02e3, bus: bus, programCounter: counter}
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
	if cpu.debug {
		log.Printf("Running instruction %s PC: %x: \n", op.name, cpu.programCounter)
	}
	if cpu.debug && cpu.programCounter == cpu.breakpoint {
		log.Printf("Stopped at instruction %s: \n", op.name)
		cpu.Dump()
		os.Exit(0)
	}
	cpu.clock += uint16(ticksPerInstruction[instruction])
	if op.exec == nil {
		cpu.Dump()
		log.Fatalf("Unable to run instruction: %s %x", op.name, instruction)
	}
	cpu.programCounter++
	//execute
	op.exec(cpu)

}

func (cpu *Cpu) Clock() uint16 {
	return cpu.clock
}
