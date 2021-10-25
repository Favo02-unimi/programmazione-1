package main

/*
Definizione: Una parola è palindroma se può essere letta normalmente, da sinistra verso destra, sia viceversa, cioè da destra verso sinistra.

Scrivere un programma che:

legga da standard input una stringa senza spazi;
stampi a video il messaggio Palindroma nel caso in cui la stringa letta sia palindroma e Non palindroma altrimenti.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione ÈPalindroma(s string) bool che riceve in input un valore string nel parametro n e restituisce true se s è palindroma e false altrimenti.
*/
//Funziona solo con ASCII

import (
	. "fmt"
)

func main() {
	Print("Inserire stringa: ")
	var s string
	Scan(&s)

	if ÈPalindroma(s) {
		Print("Palindroma")
	} else {
		Print("Non palindroma")
	}
}

func ÈPalindroma(s string) bool {
	var i int = 0
	var j int = len(s) - 1
	for {
		if i >= len(s)/2 {
			break
		}
		Println(string(s[i]), "con", string(s[j]))
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
