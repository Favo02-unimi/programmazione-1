package main

/*
Definizione: Un numero naturale è abbondante se è inferiore alla somma dei suoi divisori propri (per esempio, 12 è abbondante perché la somma dei suoi divisori propri (1, 2, 3, 4, 6) è 16).

Scrivere un programma che legga da standard input un numero intero soglia e stampi tutti i numeri abbondanti inferiori a soglia. Se soglia <= 0 il programma deve stampare La soglia inserita non è positiva.

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione SommaDivisoriPropri(n int) int che riceve in input un valore intero nel parametro n e restituisce la somma dei divisori propri di n. Se n non ha divisori propri la funzione restituisce 0;
una funzione ÈAbbondante(n int) bool che riceve in input un valore intero nel parametro n e restituisce true se n è abbondante, false altrimenti; la funzione deve utilizzare la funzione SommaDivisoriPropri();
una funzione NumeriAbbondanti(limite int) che riceve in input un valore intero nel parametro limite e stampa tutti i numeri abbondanti inferiori a limite; la funzione deve utilizzare la funzione ÈAbbondante();
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
		NumeriAbbondanti(x)
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

func ÈAbbondante(n int) bool {
	return SommaDivisoriPropri(n) > n
}

func NumeriAbbondanti(limite int) {
	Print("Numeri abbondanti: ")
	for i := 0; i < limite; i++ {
		if ÈAbbondante(i) {
			Print(i, " ")
		}
	}
}
