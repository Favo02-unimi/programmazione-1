package main

/*
Scrivere un programma che:

legga da standard input una stringa senza spazi ed interamente definita da lettere dell'alfabeto inglese;
stampi la stessa stringa in modo tale che ogni lettera sia separata da quella successiva da uno spazio.
*/

import (
	. "fmt"
)

func main() {
	Print("Inserire stringa: ")
	var s string
	Scan(&s)

	for i := 0; i < len(s); i++ {
		Print(string(s[i]), " ")
	}
}
