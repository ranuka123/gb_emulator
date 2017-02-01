package cpu

//import "fmt"

//interrupt bit positions for memory addresses 0xFFFF 0xFF0F
//ranked in order of interrupt priority with vblank being the most important
const (
	VBLANK   uint8 = 1
	LCD_STAT uint8 = 2
	TIMER    uint8 = 3
	SERIAL   uint8 = 4
	JOYPAD   uint8 = 5
)

var interruptCounters [5]uint16 = [5]uint16{0x0040, 0x0048, 0x0050, 0x0058, 0x0060}

func (cpu *Cpu) CheckInterrupts() {
	interruptsEnabled := cpu.bus.ReadByte(0xFFFF)
	interruptFlags := cpu.bus.ReadByte(0xFF0F)
	//if master interrupt is on, there are interrupts enabled and interrupt flags are set
	//run the cpu interrupt handler
	if cpu.masterInterrupt && interruptsEnabled != 0 && interruptFlags != 0 {
		//check which interrupts to handle by masking off the unset or unflagged interrupts
		interruptsToCheck := interruptsEnabled & interruptFlags

		if interruptsToCheck&VBLANK != 0 {
			//disable vblank flag since it's being handled
			interruptFlags &^= VBLANK
			cpu.bus.WriteByte(0xFF0F, interruptFlags)
			//cpu.bus.screen.RenderScreen()
			cpu.runHandler(VBLANK)
		} else if interruptsToCheck&LCD_STAT != 0 {
			interruptFlags &^= LCD_STAT
			cpu.bus.WriteByte(0xFF0F, interruptFlags)
			cpu.runHandler(LCD_STAT)
		} else if interruptsToCheck&TIMER != 0 {
			interruptFlags &^= TIMER
			cpu.bus.WriteByte(0xFF0F, interruptFlags)
			cpu.runHandler(TIMER)
		} else if interruptsToCheck&SERIAL != 0 {
			interruptFlags &^= SERIAL
			cpu.bus.WriteByte(0xFF0F, interruptFlags)
			cpu.runHandler(SERIAL)
		} else if interruptsToCheck&JOYPAD != 0 {
			interruptFlags &^= JOYPAD
			cpu.bus.WriteByte(0xFF0F, interruptFlags)
			cpu.runHandler(JOYPAD)
		}
	}

}

func (cpu *Cpu) runHandler(interrupt uint8) {
	//disable master interrupt
	cpu.masterInterrupt = false
	//decrement stack and save the program counter to the stack
	cpu.stackPointer -= 2
	cpu.bus.WriteWord(cpu.stackPointer, cpu.programCounter)

	cpu.programCounter = interruptCounters[interrupt-1]
	cpu.clock += 12

}

func (cpu *Cpu) returnFromInterrupt() {
	cpu.masterInterrupt = true
	cpu.programCounter = cpu.bus.ReadWord(cpu.stackPointer)
	cpu.stackPointer += 2
	cpu.clock += 12
}
