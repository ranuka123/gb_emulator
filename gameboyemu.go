package main

import (
	"gameboyemu/gameboy"
	"os"
	"strconv"
)

func main() {
	var gb *gameboy.Gameboy
	if len(os.Args) < 3 {
		gb = gameboy.New(0xFFFF)
	} else {
		bp, _ := strconv.ParseInt(os.Args[2], 16, 16)
		gb = gameboy.New(uint16(bp))
	}
	gb.LoadRom(os.Args[1])
	gb.Run()
}
