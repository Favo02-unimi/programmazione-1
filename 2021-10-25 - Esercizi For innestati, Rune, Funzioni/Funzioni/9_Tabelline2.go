package main

/*
Estendete il programma che risolve l'esercizio Tabelline in modo tale che soddisfi la richiesta seguente.

Scrivere un programma che legga da standard input un numero intero e stampi la tabellina corrispondente solo se il numero inserito è compreso tra 1 e 9 (estremi inclusi). Se il numero inserito non è valido (inferiore di 1 o superiore a 9), il programma deve stampare Numero non valido..

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

Tabellina(numero int) bool che riceve in input un valore intero nel parametro numero. Se numero è compreso tra 1 e 9 (estremi inclusi), la funzione stampa la tabellina associata e restituisce un valore logico true. Altrimenti, la funzione non stampa nulla e restituisce un valore logico false.
*/

import (
	. "fmt"
)

func main() {
	Print("Inserisci un numero: ")
	var n int
	Scan(&n)
	if !Tabellina(n) {
		Print("Numero non valido.")
	}
}

func Tabellina(n int) bool {
	if !(n >= 1 && n <= 9) {
		return false
	}
	Print("Tabellina del ", n, ": ")
	for i := 0; i <= 10; i++ {
		Print(i*n, " ")
	}
	return true
}
