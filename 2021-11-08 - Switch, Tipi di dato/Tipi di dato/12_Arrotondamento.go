package main

/*
Scrivere un programma che legga da standard input un numero reale e ne stampi il valore arrotondato alla seconda cifra decimale.

Suggerimento: Il valore arrotondato alla seconda cifra decimale di un numero reale può essere ottenuto effettuando le seguenti operazioni:

Moltiplicare il numero reale per 100
Sommare 0.5 al valore ottenuto al passo 1)
Convertire il valore ottenuto al passo 2) in un valore di tipo int
Riconvertire il valore ottenuto al passo 3) in un valore di tipo float64
Dividere il valore ottenuto al passo 4) per 100
oppure:

Moltiplicare il numero reale per 100
Utilizzare la funzione math.Round del package math per arrontondare all'intero valore ottenuto al passo 1) (si utilizzi il comando go doc math.Round per capire come utilizzare la funzione)
Dividere il valore ottenuto al passo 2) per 100
*/

import (
	. "fmt"
	"math"
)

func main() {
	var n float64
	Print("Inserisci il valore da arrotondare: ")
	Scan(&n)

	Println("Valore arrotondato =", math.Round(n*100)/100)
}
