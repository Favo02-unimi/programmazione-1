package main

import (
	. "fmt"
	"time"
)

func main() {

	for i := 0; true; i++ {
		Print("\r")
		switch i {
		case 0:
			Print("|")
		case 1:
			Print("/")
		case 2:
			Print("-")
		case 3:
			Print("\\")
			i = -1
		}
		time.Sleep(10000)
	}
}
