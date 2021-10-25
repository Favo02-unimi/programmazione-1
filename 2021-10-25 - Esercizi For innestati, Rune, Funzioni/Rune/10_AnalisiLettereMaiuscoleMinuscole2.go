package main

/*
Scrivere un programma che legga da standard input una stringa senza spazi e, considerando l’insieme delle lettere dell'alfabeto inglese, stampi

il numero di vocali maiuscole;
il numero di vocali minuscole;
il numero di consonanti maiuscole;
il numero di consonanti minuscole.
A tal fine definire le seguenti funzioni: 'èLetteraValida(l rune) bool', 'èMaiuscola(l rune) bool', 'èVocale(l rune) bool'.

Alternativamente, è possibile anche utilizzare le funzioni 'unicode.IsLetter' e 'unicode.IsUpper' del package 'unicode' al posto di 'èLetteraValida' e 'èMaiuscola' rispettivamente. Che differenze ci sono?
*/

import (
	. "fmt"
	"unicode"
)

func main() {
	Print("Inserire stringa: ")
	var s string
	Scan(&s)

	var vocaliMaiuscole, vocaliMinuscole, consonantiMaiuscole, consonantiMinuscole int

	for _, r := range s {
		if unicode.IsUpper(r) {
			if èVocale(r) {
				vocaliMaiuscole++
			} else {
				consonantiMaiuscole++
			}
		}
		if unicode.IsLower(r) {
			if èVocale(r) {
				vocaliMinuscole++
			} else {
				consonantiMinuscole++
			}
		}
	}
	Println("Vocali maiuscole:", vocaliMaiuscole)
	Println("Consonanti maiuscole:", consonantiMaiuscole)
	Println("Vocali minuscole:", vocaliMinuscole)
	Println("Consonanti minuscole:", consonantiMinuscole)
}

func èVocale(l rune) bool {
	return l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U' || l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u'
}
