package main

/*
Scrivere un programma che legga da standard input un valore intero e tre valori reali:

il primo valore è il seme (seed) s da utilizzare per inizializzare il generatore di numeri casuali;
il secondo ed il terzo valore, xC e yC, rappresentano rispettivamente l'ascissa e l'ordinata di un punto C sul piano cartesiano: il centro di un cerchio;
il quarto valore, raggio, è il raggio del cerchio con centro C.
Una volta terminata la fase di lettura, il programma deve generare per 1.000.000 di volte una coppia di valori reali px e py che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, come mostrato nell'Esempio di esecuzione:

per ogni punto che, a meno di una costante EPSILON = 1.0e-6, giace sulla circoferenza del cerchio con centro C, deve stampare a video il relativo messaggio;

deve stampare l'ascissa e l'ordinata del punto che, di almeno una costante EPSILON = 1.0e-6:

a) è all'esterno del cerchio;

b) ha distanza massima dal centro C;

deve stampare l'ascissa e l'ordinata del punto che, di almeno una costante EPSILON = 1.0e-6:

a) è all'interno al cerchio;

b) ha distanza minima dal centro C;

I valori px e py devono essere compresi rispettivamente negli intervalli [xC-raggio, xC+raggio) e [yC-raggio, yC+raggio).
*/

import (
	. "fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-6

func main() {

	var seed int64
	Print("Inserisci s: ")
	Scan(&seed)
	rand.Seed(seed)

	var xC, yC float64
	Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C (il centro del cerchio): ")
	Scan(&xC, &yC)

	var raggio float64
	Print("Inserisci il valore del raggio: ")
	Scan(&raggio)

	max := xC + raggio
	min := xC - raggio

	var xMaxEsterno, yMaxEsterno, xMinInterno, yMinInterno float64

	for i := 0; i < 1_000_000; i++ {

		px := min + rand.Float64()*(max-min)
		py := min + rand.Float64()*(max-min)

		if i == 0 { //inizializzazione punti massimo e minimo
			xMaxEsterno = px
			yMaxEsterno = py
			xMinInterno = px
			yMinInterno = py
		}

		if ÈXUgualeAY(raggio*raggio, (math.Pow(px-xC, 2) + (math.Pow(py-yC, 2)))) {
			Print("Il punto (", px, ",", py, ") giace sulla circonferenza del cerchio\n")
		}
		if Distanza(px, py, xC, yC) > Distanza(xMaxEsterno, yMaxEsterno, xC, yC) {
			xMaxEsterno = px
			yMaxEsterno = py
		}
		if Distanza(px, py, xC, yC) < Distanza(xMinInterno, yMinInterno, xC, yC) {
			xMinInterno = px
			yMinInterno = py
		}

	}
	Println()
	Print("Il punto (", xMaxEsterno, ",", yMaxEsterno, ") è quello all'esterno del cerchio che ha distanza massima dal centro C\n")
	Println()
	Print("Il punto (", xMinInterno, ",", yMinInterno, ") è quello all'esterno del cerchio che ha distanza massima dal centro C\n")
	Println()

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

func Distanza(x1, y1, x2, y2 float64) float64 {
	n := math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
	return math.Pow(n, 0.5)
}
