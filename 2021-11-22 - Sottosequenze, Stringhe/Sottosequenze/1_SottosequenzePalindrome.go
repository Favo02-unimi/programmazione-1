package main

/*
Scrivere un programma che legga da riga di comando una stringa composta di caratteri unicode e stampi a schermo tutte le sottostringhe palindrome (che siano uguali lette da destra e da sinistra) della stringa.
*/

import (
	. "fmt"
	"os"
)

func main() {
	var str string = os.Args[1]
	var s []rune = []rune(str)
	var sub []rune
	var pal []string

	if IsPalindroma(s) {
		pal = append(pal, str)

	}
	for lungh := 2; lungh < len(s); lungh++ {
		for start := 0; start < len(s)-lungh+1; start++ {

			sub = s[start : start+lungh]

			if IsPalindroma(sub) {
				//Println("-->", sub)
				pal = append(pal, string(sub))
			} else {
				//Println(sub)
			}
		}
	}
	Println(pal)
}

func IsPalindroma(s []rune) bool {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
