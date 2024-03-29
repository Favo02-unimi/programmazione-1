package main

/*
Scrivere un programma che legga da standard input un numero intero e stampi la tabellina corrispondente solo se il numero inserito è compreso tra 1 e 9 (estremi inclusi).

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

Tabellina(numero int) che riceve in input un valore intero nel parametro numero e, se numero è compreso tra 1 e 9 (estremi inclusi), stampa la tabellina associata, altrimenti non stampa nulla.
*/

import (
	. "fmt"
)

func main() {
	Print("Inserisci un numero: ")
	var n int
	Scan(&n)
	Tabellina(n)
}

func Tabellina(n int) {
	if !(n >= 1 && n <= 9) {
		return
	}
	Print("Tabellina del ", n, ": ")
	for i := 0; i <= 10; i++ {
		Print(i*n, " ")
	}
}
