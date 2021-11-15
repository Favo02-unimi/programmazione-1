package main

/*
Scrivere un programma che legga da riga di comando una sequenza di valori e stampi a video il risultato della moltiplicazione tra i valori che rappresentano numeri interi.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args
	var nums []int
	var molt int = 1

	for i, s := range args {
		if i == 0 {
			continue
		}

		if n, ok := getNum(s); ok {
			nums = append(nums, n)
		}
	}

	for _, n := range nums {
		molt *= n
	}

	Println("Il risultato della moltiplicazione tra i numeri interi è", molt)

}

func getNum(s string) (int, bool) {
	if n, ok := strconv.Atoi(s); ok != nil {
		return 0, false
	} else {
		return n, true
	}
}
