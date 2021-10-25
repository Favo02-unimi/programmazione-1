package main

/*
Definizione: Se a, b e c sono numeri naturali e a² + b² = c², si dice che la terna di numeri a, b e c è una terna pitagorica.

Scrivere un programma che legga da standard input un intero soglia>0 e stampi a video tutte le terne pitagorighe tali che a<soglia, b<soglia e c<soglia.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

ÈTernaPitagoriga(a int, b int, c int) bool che riceve in input tre valori interi nei parametri a, b e c, e restituisce true se c² è uguale a a² + b² e false altrimenti;
TernePitagoriche(soglia int) che riceve in input un valore intero nel parametro soglia e stampa tutte le terne pitagoriche inferiori a soglia; la funzione deve utilizzare la funzione ÈTernaPitagoriga().
*/

import (
	. "fmt"
)

func main() {
	Print("Inserisci un numero: ")
	var x int
	Scan(&x)
	if x <= 0 {
		Print("La soglia inserita non è positiva")
	} else {
		TernePitagoriche(x)
	}
}

func ÈTernaPitagoriga(a int, b int, c int) bool {
	return c*c == (a*a + b*b)
}

func TernePitagoriche(soglia int) {
	Println("Terne pitagoriche:")
	for a := 1; a < soglia; a++ {
		for b := 1; b < soglia; b++ {
			for c := 1; c < soglia; c++ {
				if ÈTernaPitagoriga(a, b, c) {
					Println("(", a, ",", b, ",", c, ")")
				}
			}
		}
	}
}
