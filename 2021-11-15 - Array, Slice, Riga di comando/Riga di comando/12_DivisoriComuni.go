package main

/*
Definizione: I divisori propri di un numero naturale (un numero intero positivo) sono tutti i suoi divisori, tranne il numero stesso.

Definizione: Un numero naturale (un numero intero positivo) è perfetto se è uguale alla somma dei suoi divisori propri (per esempio, 6 è perfetto perché 6 = 1 + 2 + 3 ).

Scrivere un programma che:

legga da riga di comando tre numeri interi positivi, rispettivamente N, DIVISORIMIN, e DIVISORIMAX;
stampi a video tutte le coppie di interi positivi a e b, con a <= N e b <= N, tali che:
a e b abbiano in comune un numero di divisori propri maggiore o uguale a DIVISORIMIN e minore o uguale a DIVISORIMAX;
almeno uno tra a e b sia un numero perfetto.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione DivisoriPropri(n int) []int che riceve in input un valore int nel parametro n e restituisce un valore []int in cui sono memorizzati tutti i divisori propri di n.
Si assuma che:

i valori letti da riga di comando siano specificati nel formato corretto;
N > MAX > MIN > 0.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args
	var n, divisoriMin, divisoriMax int
	n, _ = strconv.Atoi(args[1])
	divisoriMin, _ = strconv.Atoi(args[2])
	divisoriMax, _ = strconv.Atoi(args[3])

	for a := 0; a < n; a++ {
		divA := DivisoriPropri(a)
		for b := a; b < n; b++ {
			divB := DivisoriPropri(b)

			inComune := InComune(divA, divB)

			if inComune >= divisoriMin && inComune <= divisoriMax {
				if IsPerfetto(a) || IsPerfetto(b) {
					Println(a, b)
				}
			}

		}
	}

}

func DivisoriPropri(n int) (div []int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			div = append(div, i)
		}
	}
	return
}

func InComune(a, b []int) (cont int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				cont++
			}
		}
	}
	return cont
}

func IsPerfetto(n int) bool {
	div := DivisoriPropri(n)
	var somma int
	for _, d := range div {
		somma += d
	}
	if somma == n {
		return true
	}
	return false
}
