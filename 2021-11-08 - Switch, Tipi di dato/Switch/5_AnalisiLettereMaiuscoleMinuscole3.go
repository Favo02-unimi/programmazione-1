package main

/*
Scrivere un programma che legga da standard input una stringa senza spazi e, considerando l’insieme delle lettere dell'alfabeto inglese, stampi

il numero di vocali maiuscole;
il numero di vocali minuscole;
il numero di consonanti maiuscole;
il numero di consonanti minuscole.
A tal fine definire le seguenti funzioni: 'èLetteraValida(l rune) bool', 'èMaiuscola(l rune) bool', 'èVocale(l rune) bool'.

Modificare il programma (già svolto la scorsa volta) per

Utilizzare il costrutto switch per il controllo delle vocali nella funzione 'èVocale'
Utilizzare le funzioni 'unicode.IsLetter' e 'unicode.IsUpper' del package 'unicode' al posto di 'èLetteraValida' e 'èMaiuscola' rispettivamente.
*/

import (
	. "fmt"
	"unicode"
)

func main() {
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
	switch l {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
