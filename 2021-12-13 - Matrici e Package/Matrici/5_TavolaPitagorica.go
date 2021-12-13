package main

/*
Scrivere un programma che legga da riga di comando un numero intero n e, come mostrato nell'Esempio di esecuzione, stampi a video la corrispondente tavola pitagorica n x n.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione CreaTavolaPitagorica(n int) [][]int che riceve in input un valore int nel parametro n e restituisce un valore di tipo [][]int in cui sono memorizzati i valori di una tavola pitagorica n x n;
una funzione StampaTavolaPitagorica(s [][]int) che riceve in input un valore di tipo [][]int nel parametro s e, come mostrato nell'Esempio di esecuzione, stampa la tavola pitagorica corrispondete ai valori memorizzati s.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	n, _ := strconv.Atoi(arg)

	tav := CreaTavolaPitagorica(n)
	StampaTavolaPitagorica(tav)
}

func CreaTavolaPitagorica(n int) [][]int {
	var res [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)

		for j := 0; j < n; j++ {

			res[i][j] = (i + 1) * (j + 1)

		}
	}
	return res
}

func StampaTavolaPitagorica(s [][]int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			Print("\t", s[i][j])
		}
		Println()
	}
}
