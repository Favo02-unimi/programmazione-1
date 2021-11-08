package main

/*
Scrivere un programma che legga da standard input un valore intero e due valori reali:

il primo valore è il seme (seed) s da utilizzare per inizializzare il generatore di numeri casuali;
il secondo ed il terzo valore sono il coefficiente angolare m e il termine noto q di una retta r: y = m*x + q sul piano cartesiano.
Una volta terminata la fase di lettura, il programma deve generare per 10 volte una coppia di valori reali px e py che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, per ciascun punto:

determinare se, a meno di una costante EPSILON = 1.0e-9, il punto sta sopra, sotto, o appartiene alla retta r;
stampare a video il relativo messaggio (come mostrato nell'Esempio di esecuzione).
I valori px e py devono essere compresi nell'intervallo [0, 10.0).
*/

import (
	. "fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-9

func main() {
	var seed int64
	Print("Inserire seed: ")
	Scan(&seed)

	var m, q float64
	Print("Inserire q ed m della retta: ")
	Scan(&m, &q)

	rand.Seed(seed)

	for i := 0; i < 10; i++ {
		x := rand.Float64() * 10
		y := rand.Float64() * 10

		if ÈXUgualeAY(y, m*x+q) {
			Print("Punto (", x, ",", y, ") - Il punto appartiene alla retta\n")
		} else if ÈXMaggioreDiY(y, m*x+q) {
			Print("Punto (", x, ",", y, ") - Il punto sta sopra la retta\n")
		} else {
			Print("Punto (", x, ",", y, ") - Il punto sta sotto la retta\n")
		}
	}

}

func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}
