package main

/*
Scrivere un programma che:

legga da riga di comando una stringa s costituita da cifre decimali;
stampi a schermo tutte le sottosequenze (anche ripetute) della stringa s nelle quali le cifre sono in ordine crescente (si considerino solamente sottosequenze di almeno 2 cifre).
Se la stringa s letta da riga di comando non è costituita solamente da cifre decimali, il programma termina senza stampare nulla.

Oltre alla funzione main(), deve essere definita ed utilizzata almeno la funzione Sottostringhe(s string) []string, che riceve in input un valore string nel parametro s, e restituisce un valore []string che contiene tutte le sottosequenze desiderate
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var str string = os.Args[1]
	if _, err := strconv.Atoi(str); err != nil {
		return
	}

	var sott []string = sott(str)
	var ordinate []string

	for _, v := range sott {
		if ordine(v) {
			ordinate = append(ordinate, v)
		}
	}
	Println(ordinate)
}

func sott(s string) (sott []string) {
	var sub string
	for lungh := 2; lungh <= len(s); lungh++ {
		for start := 0; start < len(s)-lungh+1; start++ {
			sub = s[start : start+lungh]
			sott = append(sott, sub)
		}
	}
	return
}

func ordine(s string) bool {
	for i := 1; i < len(s); i++ {
		v, _ := strconv.Atoi(string(s[i]))
		vindietro, _ := strconv.Atoi(string(s[i-1]))
		if v <= vindietro {
			return false
		}
	}
	return true
}
