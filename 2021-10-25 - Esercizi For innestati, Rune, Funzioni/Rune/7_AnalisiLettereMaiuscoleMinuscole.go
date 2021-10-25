package main

/*
Scrivere un programma che legga da standard input una stringa senza spazi e, considerando solamente l’insieme delle lettere dell'alfabeto inglese, stampi

il numero di lettere maiuscole;
il numero di lettere minuscole;
il numero di consonanti;
il numero di vocali.
A tal fine definire le seguenti funzioni: 'èLetteraValida(l rune) bool', 'èMaiuscola(l rune) bool', 'èVocale(l rune) bool'.

Suggerimento: Le lettere dell'alfabeto inglese sono considerate nello standard US-ASCII (integrato nello standard Unicode). Per i caratteri considerati nello standard US-ASCII, il codice Unicode varia tra 0 e 127. In particolare, per le lettere dell’alfabeto inglese il codice Unicode varia negli intervalli:

[65, 90] per le lettere MAIUSCOLE (con ‘A’=65 e ‘Z’=90)
[97, 122] per le lettere minuscole (con ‘a’=97 e ‘z’=122)
Sia c una varibile di tipo rune, i seguenti blocchi di codici sono sintatticamente/semanticamente corretti:
*/

import (
	. "fmt"
)

func main() {
	Print("Inserire stringa: ")
	var s string
	Scan(&s)

	var maiuscole, minuscole, vocali, consonanti int

	for _, l := range s {
		if !èLetteraValida(l) {
			Println(string(l), "non valida")
			continue
		}
		if èMaiuscola(l) {
			maiuscole++
		} else {
			minuscole++
		}
		if èVocale(l) {
			vocali++
		} else {
			consonanti++
		}
	}
	Println("Maiuscole:", maiuscole)
	Println("Minuscole:", minuscole)
	Println("Vocali:", vocali)
	Println("Consonanti:", consonanti)
}

func èLetteraValida(l rune) bool {
	return l >= 65 && l < 122
}

func èMaiuscola(l rune) bool {
	return l >= 65 && l <= 90
}

func èVocale(l rune) bool {
	return l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U' || l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u'
}
