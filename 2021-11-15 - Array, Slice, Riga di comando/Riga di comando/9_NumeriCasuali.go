package main

/*
Scrivere un programma che:

Legga da riga di comando un numero intero soglia;
Generi in modo casuale una sequenza di lunghezza arbitraria di numeri interi compresi nell'intervallo che va da 0 a 100, estremi inclusi. Il processo di generazione si interrompe quando viene generato un numero inferiore a soglia.
Stampi a video tutti i numeri generati.
Stampi a video tutti i numeri generati superiori a soglia.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Genera(soglia int) []int che riceve in input un valore int nel parametro soglia e restituisce un valore di tipo []int in cui è memorizzata una sequenza di lunghezza arbitraria di numeri interi, generata in base alle specifiche di cui al punto 2.
Suggerimento: per generare in modo casuale un numero intero, potete utilizzare le funzioni dei package math/rand e time come mostrato nel seguente frammento di codice:
*/

import (
	. "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	soglia, _ := strconv.Atoi(os.Args[1])
	nums := Genera(soglia)
	Println("Valori generati:", nums)
	Println("Valori sopra soglia:", nums[:len(nums)-1])
}

func Genera(soglia int) []int {
	var nums []int
	rand.Seed(int64(time.Now().Nanosecond()))
	for {
		num := rand.Intn(100)
		nums = append(nums, num)
		if num < soglia {
			break
		}
	}
	return nums
}
