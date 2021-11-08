package main

/*
Scrivere un programma che legga da standard input un numero reale e ne stampi il valore troncato alla seconda cifra decimale.

Suggerimento: Il valore troncato alla seconda cifra decimale di un numero reale può essere ottenuto effettuando le seguenti operazioni:

Moltiplicare il numero reale per 100
Convertire il valore ottenuto al passo 1) in un valore di tipo int
Riconvertire il valore ottenuto al passo 2) in un valore di tipo float64
Dividere il valore ottenuto al passo 3) per 100
*/

import (
	. "fmt"
	"math"
)

func main() {
	var n float64
	Print("Inserisci il valore da troncare: ")
	Scan(&n)

	Println("Valore troncato =", math.Trunc(n*100)/100)
}
