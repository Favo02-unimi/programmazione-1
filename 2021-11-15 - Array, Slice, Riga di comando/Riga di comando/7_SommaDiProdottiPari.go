package main

/*
Scrivere un programma che:

legga da riga di comando una sequenza di numeri interi;
stampi a video il risultato della somma dei prodotti pari associati alle coppie non ordinate di numeri che si possono definire a partire dai numeri letti (data la coppia non ordinata di numeri (numero_1, numero_2), il valore del prodotto associato è numero_1 * numero_2).
Esempio: Se 10 1 31 4 è la sequenza letta, le coppie non ordinate di numeri che si possono definire a partire dai numeri letti sono: (10, 1); (10, 31); (10, 4); (1, 31); (1, 4); (31, 4). Di queste, quelle il cui prodotto è pari sono: (10, 1); (10, 31); (10, 4); (1, 4); (31, 4). La somma dei prodotti pari è 488.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Calcola(sl []int) int che riceve in input un valore []int nel parametro sl e restituisce un valore di tipo int pari alla somma dei prodotti pari associati alle coppie non ordinate di numeri che si possono definire a partire dai numeri presenti in sl.
Nota: Per creare dinamicamente una slice, si utilizzi la funzione make.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	nums := getNums()
	Println("La somma è", Calcola(nums))
}

func getNums() (nums []int) {
	var args []string = os.Args
	for i := 1; i < len(args); i++ {
		if n, err := strconv.Atoi(args[i]); err == nil {
			nums = append(nums, n)
		}
	}
	return nums
}

func Calcola(sl []int) int {
	var somma int
	for i := 0; i < len(sl); i++ {
		for j := i; j < len(sl); j++ {
			if i == j {
				continue
			}
			if sl[i]*sl[j]%2 == 0 {
				somma += sl[i] * sl[j]
			}
		}
	}
	return somma
}
