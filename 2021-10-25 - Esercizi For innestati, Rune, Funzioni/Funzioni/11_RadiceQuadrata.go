package main

/*
Scrivere un programma che legga da standard input un numero reale n e stampi a video il valore della radice quadrata di n solo se n >= 0. Se n < 0 il programma deve stampare Il valore inserito non appartiene al dominio della funzione.

Suggerimento: Per calcolare la radice quadrata di un numero relae n potete usare la funzione Sqrt del package math:

radiceQuadrata = math.Sqrt(n)
Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

RadiceQuadrata(numero float64) (float64, bool) che riceve in input un valore reale nel parametro numero. Se numero >= 0 la funzione restituisce il valore della radice quadrata di numero e un valore logico true come certificato della corretta esecuzione del calcolo. Altrimenti, la funzione restituisce un valore reale 0 e un valore logico false, per segnalare che non è stato possibile calcolare la radice quadrata di numero nel campo dei reali.
*/

import (
	. "fmt"
	"math"
)

func main() {
	Print("Inserisci un numero: ")
	var x float64
	Scan(&x)
	sqrt, ok := RadiceQuadrata(x)
	if !ok {
		Print("Il valore inserito non appartiene al dominio della funzione")
	} else {
		Print("Radice quadrata: ", sqrt)
	}
}

func RadiceQuadrata(numero float64) (float64, bool) {
	if numero >= 0 {
		return math.Sqrt(numero), true
	}
	return 0, false
}
