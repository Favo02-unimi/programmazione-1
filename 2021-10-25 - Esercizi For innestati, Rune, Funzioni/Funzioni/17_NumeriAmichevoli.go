package main

/*
Definizione: Due numeri naturali x e y, con x < y, sono detti amichevoli se la somma dei divisori propri di ciascuno è uguale all’altro; ad esempio (220, 284) è una coppia di amichevoli, essendo 284 = 1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110 (dove 1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110 sono i divisori di 220) e 220 = 1 + 2 + 4 + 71 + 142 (dove 1, 2, 4, 71, 142 sono i divisori di 284).

Scrivere un programma che legga da standard input un numero intero soglia e stampi tutte le coppie di numeri amichevoli tali che y sia inferiore a soglia. Se soglia <= 0 il programma deve stampare La soglia inserita non è positiva.

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione SommaDivisoriPropri(n int) int che riceve in input un valore intero nel parametro n e restituisce la somma dei divisori propri di n. Se n non ha divisori propri la funzione restituisce 0;
una funzione SonoAmichevoli(n, m int) bool che riceve in input due valori interi nei parametri n ed m e restituisce true se n ed m sono amichevoli e false altrimenti (utilizzando la funzione SommaDivisoriPropri());
una funzione NumeriAmichevoli(limite int) che riceve in input un valore intero nel parametro limite e stampa tutte le coppie di numeri amichevoli tali che y sia inferiore a limite (cfr. Definizione); la funzione deve utilizzare la funzione SonoAmichevoli().
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
		NumeriAmichevoli(x)
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

func SonoAmichevoli(n, m int) bool {
	return SommaDivisoriPropri(n) == m && SommaDivisoriPropri(m) == n
}

func NumeriAmichevoli(limite int) {
	Println("Numeri amichevoli inferiori a", limite)
	for x := 0; x < limite; x++ {
		for y := x + 1; y < limite; y++ {
			if x == y {
				continue
			}
			if SonoAmichevoli(x, y) {
				Print("(", x, ",", y, ")", " ")
			}
		}
	}
}
