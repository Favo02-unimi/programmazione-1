package main

/*
Scrivere un programma che legga da riga di comando una sequenza di numeri interi di lunghezza pari. Data la sequenza, il programma deve moltiplicare ciascun numero in una posizione di indice pari per il successivo numero in posizione di indice dispari e sommare i prodotti ottenuti.

Esempio: Se 10 2 3 4 5 6 è la sequenza letta, allora la somma calcolata deve essere 10*2 + 3*4 + 5*6 = 62.

Il programma deve infine stampare a video il valore della somma calcolata.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Calcola(sl []int) int che riceve in input un valore []int nel parametro sl e restituisce un valore di tipo int pari alla somma dei prodotti ottenuti moltiplicando ciascun numero in una posizione di indice pari di sl per il successivo numero in posizione di indice dispari.
Nota: Per creare dinamicamente una slice, si utilizzi la funzione make.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	nums := getNums()
	if len(nums)%2 == 1 {
		Println("Sequenza di lunghezza dispari")
		return
	}
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

func Calcola(sl []int) (somma int) {
	for i := 0; i < len(sl); i += 2 {
		somma += sl[i] * sl[i+1]
	}
	return somma
}
