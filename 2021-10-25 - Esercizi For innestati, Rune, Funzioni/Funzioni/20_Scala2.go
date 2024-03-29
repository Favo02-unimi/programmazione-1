package main

/*
Scrivere un programma che legga da standard input un numero intero n e, come mostrato nell'Esempio d'esecuzione, stampi a video una scala utilizzando il carattere '*' (asterisco):

La scala è formata da n gradini.
Ciascun gradino è profondo 3 caratteri e alto 2.
Il gradino più in basso deve essere stampato senza alcuna rientranza verso destra; considerando poi i successivi gradini dal basso verso l'alto, ciascun di essi rientra (è traslato) verso destra rispetto al precedente di due caratteri ' ' (spazio) .
Se n è negativo o nullo, anziché stampare la scala il programma deve stampare un messaggio d'errore.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

StampaGradino(gradino int) che riceve in input un valore intero nel parametro gradino e stampa a video un singolo gradino della scala, opportunamente traslato verso destra. Se gradino < 0 non stampa nulla. Se gradino = 0 non stampa la rientranza. Se gradino > 0 stampa il gradino traslato di gradino * 2 spazi verso destra;
StampaScala(gradini int) che riceve in input un intero nel parametro gradini e stampa una scala composta da gradini gradini.
*/

import (
	. "fmt"
)

func main() {
	Print("Inserisci il numero dei gradini: ")
	var x int
	Scan(&x)
	if x <= 0 {
		Print("Dimensione non sufficiente")
	} else {
		StampaScala(x)
	}
}

func StampaGradino(gradino int) {
	for i := 0; i < gradino*2; i++ {
		Print(" ")
	}
	Println("***")
	for i := 0; i < gradino*2; i++ {
		Print(" ")
	}
	Println("*")
}

func StampaScala(gradini int) {
	for i := gradini - 1; i >= 0; i-- {
		StampaGradino(i)
	}
}
