package main

/*
Scrivere un programma che legga da standard input una stringa e, considerando l’insieme delle lettere dell'alfabeto inglese, ristampi a video la stringa due volte: la prima volta in maiuscolo e la seconda volta in minuscolo.

A tal fine definire due funzioni: 'inMaiuscolo(carattere rune) rune' e 'inMinuscolo(carattere rune) rune'

Suggerimento: Sia c una varibile di tipo rune, i seguenti blocchi di codici sono sintatticamente/semanticamente corretti:

if (c >='A' && c <= 'Z') {
	fmt.Println("L’equivalente lettera minuscola è:", string('a' + (c - 'A')))
}
// ... oppure, utilizzando il package "unicode"...
if (c >='A' && c <= 'Z') {
	fmt.Println("L’equivalente lettera minuscola è:", string(unicode.ToLower(c)))
}
*/

import (
	. "fmt"
)

func main() {
	Print("Inserire stringa: ")
	var s string
	Scan(&s)

	var maiuscolo, minuscolo string
	for i := 0; i < len(s); i++ {
		maiuscolo += string(inMaiuscolo(rune(s[i])))
		minuscolo += string(inMinuscolo(rune(s[i])))
	}
	Println("Testo maiuscolo:", maiuscolo)
	Println("Testo minuscolo:", minuscolo)
}

func inMaiuscolo(carattere rune) rune {
	if carattere >= 'a' && carattere <= 'z' {
		return 'A' + carattere - 'a'
	}
	return carattere
}

func inMinuscolo(carattere rune) rune {
	if carattere >= 'A' && carattere <= 'Z' {
		return 'a' + carattere - 'A'
	}
	return carattere
}
