package main

/*Definizione: Due numeri primi p e q sono gemelli se p = q + 2.

Scrivere un programma che legga da standard input un numero intero soglia e stampi tutti i numeri primi gemelli tali che p sia inferiore a soglia. Se soglia <= 0 il programma deve stampare La soglia inserita non è positiva.

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione ÈPrimo(n int) bool che riceve in input un valore intero nel parametro n e restituisce true se n è primo e false altrimenti;
una funzione NumeriPrimiGemelli(limite int) che riceve in input un valore intero nel parametro limite e stampa tutte le coppie di numeri primi gemelli tali che p sia inferiore a limite (cfr. Definizione); la funzione deve utilizzare la funzione ÈPrimo().*/

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
	var gemello int = 2
	for i := 2; i < limite; i++ {
		if ÈPrimo(i) {
			if gemello == i-2 {
				Print("(", gemello, ",", i, ")", " ")
			}
			gemello = i
		}
	}
}
