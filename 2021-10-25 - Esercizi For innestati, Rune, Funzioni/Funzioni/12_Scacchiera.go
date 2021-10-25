package main

/*
Scrivere un programma che legga da standard input un numero intero e, come mostrato nell'Esempio d'esecuzione, stampi a video una schacchiera di dimensione n x n utilizzando i caratteri * (asterisco) e + (più).

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

StampaRigaInizioAsterisco(lunghezza int) che riceve in input la lunghezza della riga da stampare nel parametro lunghezza e stampa in modo alternato i caratteri * e + (partendo dal carattere *);
StampaRigaInizioPiù(lunghezza int) che riceve in input la lunghezza della riga da stampare nel parametro lunghezza e stampa in modo alternato i caratteri + e * (partendo dal carattere +);
StampaScacchiera(dimensione int) che riceve in input la dimensione della scacchiera da stampare nel parametro dimensione e stampa la scacchiera utilizzando le funzioni StampaRigaInizioAsterisco() e StampaRigaInizioPiù(). Se dimensione <= 0, non stampa nulla.
*/

import (
	. "fmt"
)

func main() {
	Print("Inserisci la dimensione: ")
	var x int
	Scan(&x)
	StampaScacchiera(x)
}

func StampaRigaInizioAsterisco(lunghezza int) {
	for i := 0; i < lunghezza; i++ {
		if i%2 == 0 {
			Print("*")
		} else {
			Print("+")
		}
	}

}

func StampaRigaInizioPiù(lunghezza int) {
	for i := 0; i < lunghezza; i++ {
		if i%2 == 0 {
			Print("+")
		} else {
			Print("*")
		}
	}
}

func StampaScacchiera(dimensione int) {
	for i := 0; i < dimensione; i++ {
		if i%2 == 0 {
			StampaRigaInizioAsterisco(dimensione)
		} else {
			StampaRigaInizioPiù(dimensione)
		}
		Println()
	}
}
