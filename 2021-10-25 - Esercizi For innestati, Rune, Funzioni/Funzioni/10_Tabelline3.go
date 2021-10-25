package main

/*
Estendete il programma che risolve l'esercizio Tabelline 2 in modo tale che soddisfi la richiesta seguente.

Scrivere un programma che legga da standard input una sequenza di numeri interi compresi tra 1 e 9 (estremi inclusi) e stampi per ognuno di essi la tabellina corrispondente. Il programma si interrompe quando viene inserito in input un numero non valido (inferiore a 1 o superiore a 9) stampando Programma terminato..

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

Tabellina(numero int) bool che riceve in input un valore intero nel parametro numero. Se numero è compreso tra 1 e 9 (estremi inclusi), la funzione stampa la tabellina associata e restituisce un valore logico true. Altrimenti, la funzione non stampa nulla e restituisce un valore logico false.
*/

import (
	. "fmt"
)

func main() {
	var n int
	for {
		Print("Inserisci un numero: ")
		Scan(&n)
		if !Tabellina(n) {
			Print("Programma terminato.")
			break
		}
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
	Println()
	return true
}
