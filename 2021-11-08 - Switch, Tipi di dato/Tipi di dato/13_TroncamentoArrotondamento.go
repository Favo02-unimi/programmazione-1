package main

/*
Scrivere una versione generalizzata dei programmi Troncamento e Arrotondamento che legga da standard input un intero n oltre al numero reale. L'intero n specifica che il troncamento e l'arrotondamento devono avvenire alla n-esima cifra decimale.
*/

import (
	. "fmt"
	"math"
)

func main() {
	var n float64
	Print("Inserisci il valore: ")
	Scan(&n)

	var dec int
	Print("Inserisci il numero di cifre dopo la virgola: ")
	Scan(&dec)

	var molt float64 = 1
	for i := 0; i < dec; i++ {
		molt *= 10
	}

	Println("Valore troncato =", math.Trunc(n*molt)/molt)
	Println("Valore arrotondato =", math.Round(n*molt)/molt)
}
