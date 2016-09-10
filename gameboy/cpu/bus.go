package cpu

type Bus interface {
	Memory
}

type Memory interface {
	ReadByte(uint16) byte
	ReadWord(uint16) uint16
	WriteByte(uint16, byte)
	WriteWord(uint16, uint16)
}
