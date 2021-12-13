package main

/*
Scrivere un programma che:

legga da riga di comando un numero intero n;
utilizzi le funzioni CreaSliceBidimensionale(l int) [][]int e InizializzaSliceBidimensionale([][]int) descritte nell'Esercizio 2 (Laboratorio 9 - Array e slice III) per inizializzare una variabile s di tipo [][]int con lunghezza/capacità pari a n in cui ogni elemento s[i], con 0 <= i < l, è un valore di tipo []int con lunghezza/capacità pari a n;
stampi a video tutte le coppie di indici (i, j), con 0 <= i < l e 0 <= j < l, tali che s[i][j] è uguale a 1.
Oltre alle funzioni main(), CreaSliceBidimensionale(), e InizializzaSliceBidimensionale() devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Coppie(s [][]int) (coppie [][]int) che riceve in input un valore [][]int nel parametro s e restituisce il valore di tipo [][]int nella variabile coppie in cui sono memorizzate tutte le coppie di indici (i, j), con 0 <= i < l e 0 <= j < l, tali che s[i][j] è uguale a 1 (coppie[k], con 0 <= k < len(coppie), è un valore di tipo []int di lunghezza 2).
*/

import (
	. "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	arg := os.Args[1]
	n, _ := strconv.Atoi(arg)

	s := CreaSliceBidimensionale(n)
	AssegnaSliceBidimensionale(s)
	Stampa(s)
	coppie := Coppie(s)
	for _, v := range coppie {
		Println(v)
	}
}

func CreaSliceBidimensionale(l int) [][]int {
	var s [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		s[i] = make([]int, l)
	}
	return s
}

func AssegnaSliceBidimensionale(s [][]int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}
}

func Stampa(s [][]int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			Print(s[i][j])
		}
		Println()
	}
}

func Coppie(s [][]int) (coppie [][]int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == 1 {
				coppie = append(coppie, []int{i, j})
			}
		}
	}
	return
}
