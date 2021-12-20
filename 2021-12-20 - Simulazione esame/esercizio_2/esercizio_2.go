package main

import (
	. "fmt"
	"sort"
	"unicode"
)

func main() {
	var str string
	Scan(&str)

	if len(str) != len([]rune(str)) { //tutte ASCII
		return
	}
	for _, v := range str {
		if !(unicode.IsLetter(v) && unicode.IsLower(v)) {
			return
		}
	}
	//Println("Controlli fatti")

	sottostringhe := sottostringhe(str)
	var slice []string = make([]string, 0)

	for sotto, _ := range sottostringhe {
		if !controllaOrdine(sotto) {
			delete(sottostringhe, sotto)
		} else {
			slice = append(slice, sotto)
		}
	}
	sort.Strings(slice)

	output(sottostringhe, slice)
}

func sottostringhe(s string) map[string]int {
	var res map[string]int = make(map[string]int)
	for lungh := 2; lungh <= len(s); lungh++ {
		for iniz := 0; iniz <= len(s)-lungh; iniz++ {
			sott := s[iniz : iniz+lungh]
			res[sott]++
		}
	}
	return res
}

func controllaOrdine(str string) bool {
	s := []rune(str)
	for i, v := range s[:len(s)-1] {
		if !(v < s[i+1]) {
			return false
		}
	}
	return true
}

func output(s map[string]int, order []string) {
	Println("output:")
	for _, v := range order {
		Println(v, s[v])
	}
}
