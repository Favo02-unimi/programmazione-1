package main

/*
Scrivere un programma che legga da riga di comando una sequenza di valori e stampi a video la somma dei valori letti che rappresentano numeri interi eche compaiono nella sequenza una sola volta.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione LeggiNumeri() (numeri []int) che restituisce un valore numeri di tipo []int in cui sono memorizzati i valori numerici interi specificati a riga di comando;
una funzione Occorrenze(numeri []int, n int) int che riceve in input un valore []int nel parametro numeri e restituisce un valore int pari al numero di occorrenze di n in numeri.
Esempio:
Supponendo di leggere da riga di comando la sequenza 1 2 a 4 ciao 3 2 1 5, il programma deve stampare 12, ovvero la somma dei numeri 4, 3 e 5. Se la sequenza fosse 4 3 5 non_conto 4 2 2 3 2, l'output sarebbe invece 5.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	nums := getNums()
	Println(somma(nums))
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

func somma(nums []int) int {
	var somma int
	for _, n := range nums {
		if !checkDoppio(nums, n) {
			somma += n
		}
	}
	return somma
}

func checkDoppio(nums []int, n int) bool {
	var cont int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == n {
			cont++
		}
	}
	return cont > 1
}
