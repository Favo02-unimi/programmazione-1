package main

/*
Scrivere un programma che legga da riga di comando una sequenza di valori interi e stampi a video il valore minimo, massimo e medio tra i valori letti.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Minimo(sl []int) int che riceve in input un valore []int nel parametro sl e restituisce il minimo valore intero presente in sl.
una funzione Massimo(sl []int) int che riceve in input un valore []int nel parametro sl e restituisce il massimo valore intero presente in sl.
una funzione Media(sl []int) float64 che riceve in input un valore []int nel parametro sl e restituisce un valore reale pari alla media aritmetica dei valori interi presenti in sl.
Suggerimento: Il numero di valori interi che il programma deve considerare è pari a len(os.Args)-1.

Nota: Per creare dinamicamente una slice, si utilizzi la funzione make.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args
	var nums []int

	for i := 1; i < len(args); i++ {
		if n, err := strconv.Atoi(args[i]); err == nil {
			nums = append(nums, n)
		}
	}

	Println("Minimo:", Minimo(nums))
	Println("Massimo:", Massimo(nums))
	Printf("Media: %.2f", Media(nums))
}

func Minimo(sl []int) int {
	var min int = sl[0]
	for i := 1; i < len(sl); i++ {
		if sl[i] < min {
			min = sl[i]
		}
	}
	return min
}

func Massimo(sl []int) int {
	var max int = sl[0]
	for i := 1; i < len(sl); i++ {
		if sl[i] > max {
			max = sl[i]
		}
	}
	return max
}

func Media(sl []int) float64 {
	var somma int = 0
	for i := 0; i < len(sl); i++ {
		somma += sl[i]
	}
	return float64(somma) / float64(len(sl))
}
