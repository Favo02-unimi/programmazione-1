package main

/*
Scrivere un programma che:

Legga da standard input un numero reale.
Calcoli la radice quadrata del numero letto (sia x la variabile di tipo float64 in cui è memorizzato il valore; la radice quadrata di del valore reale può essere calcolata utilizzando la funzione math.Sqrt del package math nel seguente modo radiceQuadrata := math.Sqrt(x)).
Stampi a video x, "uguale a", radiceQuadrata, "*", radiceQuadrata, "\n" nel caso in cui radiceQuadrata*radiceQuadrata sia uguale a x e x, "diverso da", radiceQuadrata, "*", radiceQuadrata, "\n" altrimenti.
Supponendo che l'utente inserisca da standard input il valore 9, qual è l'output del programma?

Supponendo che l'utente inserisca da standard input il valore 10, qual è l'output del programma? Perché?
*/

import (
	. "fmt"
	"math"
)

func main() {
	Print("Inserisci un numero: ")
	var x float64
	Scan(&x)

	var sqrt float64 = math.Sqrt(x)

	if sqrt*sqrt == x {
		Println(x, "uguale a", sqrt, "*", sqrt)
	} else {
		Println(x, "diverso da", sqrt, "*", sqrt)
	}
}
