package main

/*
Scrivere un programma che legga da standard input un numero reale raggio e stampi i valori di circonferenza e area di un cerchio di raggio raggio.

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

CalcolaArea(raggio float64) float64 che riceve in input il valore reale del raggio di un cerchio nel parametro raggio e restituisce il valore dell'area del cerchio associato;
CalcolaCirconferenza(raggio float64) float64 che riceve in input il valore reale del raggio di un cerchio nel parametro raggio e restituisce il valore della circonferenza del cerchio associato.
*/

import (
	. "fmt"
	"math"
)

func main() {
	Print("Inserisci il raggio del cerchio: ")
	var r float64
	Scan(&r)

	Println("Area del cerchio:", CalcolaArea(r))
	Println("Circonferenza del cerchio:", CalcolaCirconferenza(r))
}

func CalcolaArea(raggio float64) float64 {
	return raggio * raggio * math.Pi
}

func CalcolaCirconferenza(raggio float64) float64 {
	return 2 * raggio * math.Pi
}
