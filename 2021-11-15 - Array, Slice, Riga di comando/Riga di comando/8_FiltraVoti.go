package main

/*
Scrivere un programma che:

legga da riga di comando una sequenza di valori (i valori numerici interi che compaiono all'interno della sequenza rappresentano voti in una scala di valutazione tra 0 e 100; i valori numerici interi superiori a 60 corrispondono a voti sufficienti);
stampi a video le due sottosequenze di valori numerici interi che corrisponodo rispettivamente a voti insufficienti e sufficienti.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione LeggiNumeri() (numeri []int) che restituisce un valore numeri di tipo []int in cui sono memorizzati i valori numerici interi specificati a riga di comando;
una funzione FiltraVoti(voti []int) (sufficienti, insufficienti []int) che riceve in input un valore []int nel parametro voti e restituisce due valori di tipo []int, sufficienti e insufficienti, in cui sono memorizzati rispettivamente i voti sufficienti e insufficienti presenti in voti.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	nums := getNums()
	suff, insuff := FiltraVoti(nums)
	Println("Voti sufficienti:", suff)
	Println("Voti insufficienti:", insuff)
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

func FiltraVoti(voti []int) (sufficienti, insufficienti []int) {
	for _, n := range voti {
		if n < 60 {
			insufficienti = append(insufficienti, n)
		} else {
			sufficienti = append(sufficienti, n)
		}
	}
	return
}
