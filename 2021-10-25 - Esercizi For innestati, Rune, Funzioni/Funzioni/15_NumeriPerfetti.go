package main

/*
Definizione: Un numero naturale è perfetto se è uguale alla somma dei suoi divisori propri (per esempio, 6 è perfetto perché 6 = 1 + 2 + 3).

Scrivere un programma che legga da standard input un numero intero n e stampi se n è perfetto oppure no.

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione SommaDivisoriPropri(n int) int che riceve in input un valore intero nel parametro n e restituisce la somma dei divisori propri di n. Se n non ha divisori propri la funzione restituisce 0;
una funzione ÈPerfetto(n int) bool che riceve in input un valore intero nel parametro n e restituisce true se n è perfetto e false altrimenti; la funzione deve utilizzare la funzione SommaDivisoriPropri().
*/

import (
	. "fmt"
)

func main() {
	Print("Inserisci un numero: ")
	var x int
	Scan(&x)
	if ÈPerfetto(x) {
		Println(x, "è perfetto")
	} else {
		Println(x, "non è perfetto")
	}
}

func SommaDivisoriPropri(n int) int {
	var somma int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			somma += i
		}
	}
	return somma
}

func ÈPerfetto(n int) bool {
	if n == 0 {
		return false
	}
	if SommaDivisoriPropri(n) == n {
		return true
	}
	return false
}
