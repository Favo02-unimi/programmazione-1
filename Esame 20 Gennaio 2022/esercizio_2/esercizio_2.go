package main

import (
	. "fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	arg := os.Args[1]
	sotto := sottostringhe(arg)
	var sottoSlice []int
	for k, _ := range sotto {
		if k[:1] == "0" { //per evitare duplicati dovuti a 3011 e 03011
			continue
		}
		n, _ := strconv.Atoi(k)
		if n < 100 {
			continue
		}
		sottoSlice = append(sottoSlice, n)
	}
	sort.Ints(sottoSlice)
	//Println(sottoSlice)

	var primi []int
	for _, v := range sottoSlice {
		if èPrimo(v) {
			primi = append(primi, v)
		}
	}
	formatPrint(primi)
}

func sottostringhe(s string) (res map[string]int) {
	res = make(map[string]int)
	for lungh := 2; lungh <= len(s); lungh++ {
		for inizio := 0; inizio < len(s)-lungh+1; inizio++ {
			res[s[inizio:inizio+lungh]]++
		}
	}
	return res
}

func èPrimo(n int) bool {
	for i := 2; i < ((n / 2) + 1); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func formatPrint(primi []int) {
	for i, v := range primi {
		Print(v)
		if i != len(primi)-1 {
			Print(", ")
		}
	}
}
