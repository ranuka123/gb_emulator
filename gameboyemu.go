package main

import (
	"gameboyemu/gameboy"
	"os"
)

func main() {
	var gameboy *gameboy.Gameboy = gameboy.New()
	gameboy.LoadRom(os.Args[1])
	gameboy.Run()
}
