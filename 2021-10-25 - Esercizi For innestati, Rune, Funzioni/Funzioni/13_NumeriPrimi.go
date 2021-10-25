package main

/*
Definizione: Un numero naturale è primo se è divisibile solo per se stesso e per 1.

Scrivere un programma che legga da standard input un numero intero soglia e stampi tutti i numeri primi inferiori a soglia. Se soglia <= 0 il programma deve stampare La soglia inserita non è positiva.

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione ÈPrimo(n int) bool che riceve in input un valore intero nel parametro n e restituisce true se n è primo e false altrimenti;
una funzione NumeriPrimi(limite int) che riceve in input un valore intero nel parametro limite e stampa tutti i numeri primi inferiori a limite; la funzione deve utilizzare la funzione ÈPrimo().
*/

import (
	. "fmt"
)

func main() {
	Print("Inserisci un numero: ")
	var x int
	Scan(&x)
	if x <= 0 {
		Print("La soglia inserita non è positiva")
	} else {
		NumeriPrimi(x)
	}
}

func ÈPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NumeriPrimi(limite int) {
	Println("Numeri primi inferiori a", limite)
	for i := 2; i < limite; i++ {
		if ÈPrimo(i) {
			Print(i, " ")
		}
	}
}
