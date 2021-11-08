package main

/*
Scrivere un programma che legga da standard input il nome di un mese (in minuscolo). Il programma deve stampare a video il numero di giorni di quel mese. Si assume che l'anno non sia bisestile.

In particolare, si scriva una funzione numeroDiGiorni che, dato in input il nome del mese (variabile di tipo stringa) restituisca il numero di giorni del mese (variabile di tipo intera). La funzione deve usare il costrutto switch case.
*/

import (
	. "fmt"
)

func main() {
	Print("Inserire mese: ")
	var mese string
	Scan(&mese)
	Println("Il numero di giorni di", mese, "è", numeroDiGiorni(mese))
}

func numeroDiGiorni(mese string) int {
	switch mese {
	case "gennaio", "marzo", "maggio", "luglio", "agosto", "ottobre", "dicembre":
		return 31
	case "febbraio":
		return 28
	case "aprile", "giugno", "settembre", "novembre":
		return 30
	default:
		return -1
	}
}
