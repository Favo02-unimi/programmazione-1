package main

import (
	. "fmt"
	"os"
)

func main() {
	arg := os.Args[1]

	dritta := []rune(arg)
	inverso := inverso(arg)

	lungh := len(dritta)

	spazi := lungh / 2

	for i := 0; i < spazi; i++ {
		inserisciSpazi(spazi)
		Print(string(inverso[i]))
		inserisciSpazi(spazi)
		Println()
	}
	Println(string(inverso))
	for i := spazi + 1; i < lungh; i++ {
		inserisciSpazi(spazi)
		Print(string(inverso[i]))
		inserisciSpazi(spazi)
		Println()
	}

}

func inverso(str string) (res []rune) {
	runeConv := []rune(str)
	res = make([]rune, len(runeConv))

	for i := 0; i < len(runeConv); i++ {
		res[len(runeConv)-1-i] = runeConv[i]
	}
	return res
}

func inserisciSpazi(n int) {
	for i := 0; i < n; i++ {
		Print(" ")
	}
}
