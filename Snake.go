package main

import (
	. "fmt"
	"time"
)

const dimX int = 10
const dimY int = 20

var field [dimX][dimY]bool
var isHead [dimX][dimY]bool
var isFood [dimX][dimY]bool

func main() {

	field[0][0] = true
	for {
		Print("\033[H\033[2J")
		Println("Benvenuto nello Snake di Favo :=)")
		printField()
		moveHead('d')
		time.Sleep(1 * time.Second)
	}
}

func printField() {
	for x := -1; x < dimX+1; x++ {
		for y := -1; y < dimY+1; y++ {
			if x == -1 {
				if y == -1 {
					Print("┌-")
				} else if y == dimY {
					Print("-┐")
				} else {
					Print("--")
				}
			} else if x == dimX {
				if y == -1 {
					Print("└-")
				} else if y == dimY {
					Print("-┘")
				} else {
					Print("--")
				}
			} else if y == -1 {
				Print("│ ")
			} else if y == dimY {
				Print(" │")
			} else {

				if field[x][y] {
					Print("⬜")
				} else {
					Print("  ")
				}
			}
		}
		Println()
	}
}

func moveHead(direction rune) {
	var modifierX, modifierY int
	switch direction {
	case 'w':
		modifierY, modifierX = 0, -1
	case 'a':
		modifierY, modifierX = -1, 0
	case 's':
		modifierY, modifierX = 0, +1
	case 'd':
		modifierY, modifierX = +1, 0
	}
	Println(modifierX, modifierY)
	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			if field[x][y] {
				if x+modifierX < 0 || x+modifierX >= dimX {
					Println("Perso", "x", x, "modx", modifierX, "y", y, "mody", modifierY)
				} else if y+modifierY < 0 || y+modifierY >= dimY {
					Println("Perso", "x", x, "modx", modifierX, "y", y, "mody", modifierY)
				} else {
					field[x][y] = false
					field[x+modifierX][y+modifierY] = true
				}
			}
		}
	}
}
