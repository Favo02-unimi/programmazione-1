package main

/*
Scrivere un programma che, dopo aver letto da standard input un numero intero n, chiede all'utente di inserire n numeri interi (sempre da standard input).

Il programma deve stampare gli n numeri interi in ordine inverso rispetto a quello di inserimento.

Suggerimento: Per creare dinamicamente una slice si utilizzi la funzione make().
*/

import (
	. "fmt"
)

func main() {
	var n int
	Scan(&n)

	s := make([]int, n)

	Println("Inserisci", n, "numeri:")
	for i := 0; i < n; i++ {
		Scan(&s[i])
	}
	for i := n - 1; i >= 0; i-- {
		Print(s[i], " ")
	}
}
