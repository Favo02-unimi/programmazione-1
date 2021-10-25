package main

/*
Scrivere un programma che legga da standard input due stringhe senza spazi s1 e s2 e stampi a video la stringa creata alternando i caratteri delle stringhe s1 e s2.

A tal fine utilizzare una funzione 'StringheAlternate(s1, s2 string) (risultato string)'

Esempio: Se "ciao!" e "MONDO" sono le stringhe lette, allora la stringa stampata video deve essere "cMiOaNoD!O".

Si assuma che le stringhe lette siano interamente definite da caratteri considerati nello standard US-ASCII.

Se le stringhe lette non sono definite dallo stesso numero di caratteri, si deve utilizzare il carattere '-' come carattere di riempimento:

Esempio: Se "ciao" e "mondo!", sono le stringhe lette, allora la stringa stampata video deve essere "cmioanod-o-!".
*/

import (
	. "fmt"
	"math"
)

func main() {
	Print("Inserire due stringhe: ")
	var s1, s2 string
	Scan(&s1, &s2)

	Print(StringheAlternate(s1, s2))
}

func StringheAlternate(s1, s2 string) (risultato string) {
	var lenght = int(math.Max(float64(len(s1)), float64(len(s2))))
	for i := 0; i < lenght; i++ {
		if i >= len(s1) {
			risultato += "-"
		} else {
			risultato += string(s1[i])
		}
		if i >= len(s2) {
			risultato += "-"
		} else {
			risultato += string(s2[i])
		}
	}
	return
}
