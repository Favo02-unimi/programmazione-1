package main

/*
Definizione: Si definisce fattoriale di un numero intero positivo, il prodotto dei numeri interi positivi minori o uguali a tale numero. Il fattoriale di k è uguale a 1*2*3*...*(k-3)*(k-2)*(k-1)*k.

Scrivere un programma che legga da riga di comando un numero intero n e stampi a video il fattoriale di tutti i numeri compresi tra 1 e n (estremi inclusi).

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Fattoriali(n int) (f []int) che riceve in input un valore int nel parametro n e restituisce il valore f di tipo []int in cui in f[0] è memorizzato il fattoriale di 1, in f[1] è memorizzato il fattoriale di 2, ..., in f[n-1] è memorizzato il fattoriale di n.
Nota: Per creare dinamicamente una slice, si utilizzi la funzione make.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	Println(Fattoriali(n))
}

func Fattoriali(n int) (f []int) {
	f = make([]int, n)
	for i := 1; i <= n; i++ {
		f[i-1] = fattoriale(i)
	}
	return f
}

func fattoriale(n int) int {
	var fact int = 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}
