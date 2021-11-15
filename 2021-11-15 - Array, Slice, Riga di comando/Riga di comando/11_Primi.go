package main

/*
Definizione: Un numero naturale è primo se è divisibile solo per se stesso e per 1.

Scrivere un programma che legga da riga di comando un numero intero numero e stampi tutti i numeri primi ottenibili rimuovendo al più 3 cifre consecutive tra quelle che definiscono numero.

In particolare, i numeri primi devono essere stampate in ordine crescente (cioè dal più piccolo al più grande).

Ad esempio, se il numero intero letto da riga di comando fosse:

5899

i numeri ottenibili rimuovendo al più 3 cifre consecutive tra quelle che definiscono 5899 sarebbero:

5899
899
599
589
589
99
59
58
9
5
Si assuma che il valore specificato a riga di comando sia nel formato corretto e, in particolare, sia un intero maggiore o uguale a 1000.

Oltre alle funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione ÈPrimo(n int) bool che riceve in input un valore (di tipo) int nel parametro n e restituisce true se n è primo (i.e., se il valore di n rappresenta un numero primo) e false altrimenti.
Suggerimento: I numeri primi possono essere ordinati in senso crescente utilizzando la funzione sort.Ints(a []int) del package sort.
*/

import (
	. "fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var args []string = os.Args
	var n int
	if num, err := strconv.Atoi(args[1]); err == nil && num >= 1000 {
		n = num
	} else {
		Println("Numero invalido")
		return
	}

	var primi []int = generatePrime(n)
	sort.Ints(primi)
	for _, val := range primi {
		Println(val)
	}
}

func generatePrime(n int) (nums []int) {

	var snum string = strconv.Itoa(n)
	var sgen string

	for rimossi := 0; rimossi <= 3; rimossi++ {
		if rimossi == 0 {
			if isPrime(convInt(snum)) {
				nums = append(nums, convInt(snum))
			}
			continue
		}
		for i := 0; i < len(snum)-(rimossi-1); i++ {
			sgen = snum[:i] + snum[i+rimossi:]
			if isPrime(convInt(sgen)) {
				nums = append(nums, convInt(sgen))
			}
		}

	}
	return
}

func convInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
