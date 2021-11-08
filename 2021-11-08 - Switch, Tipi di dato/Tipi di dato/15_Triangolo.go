package main

/*
Scrivere un programma che legga da standard input un valore intero e sei valori reali:

il primo valore è il seme (seed) s da utilizzare per inizializzare il generatore di numeri casuali;
il secondo ed il terzo valore sono il coefficiente angolare m1 e il termine noto q1 di una retta r1: y = m1*x + q1 sul piano cartesiano su cui giace la base AB di un triangolo ABC;
il quarto ed il quinto valore sono il coefficiente angolare m2 e il termine noto q2 di una retta r2: y = m2*x + q2 sul piano cartesiano su cui giace il lato BC di un triangolo ABC;
il sesto ed il settimo valore sono il coefficiente angolare m3 e il termine noto q3 di una retta r3: y = m3*x + q3 sul piano cartesiano su cui giace il lato AC di un triangolo ABC.
Una volta terminata la fase di lettura, il programma deve generare per 10 volte una coppia di valori reali px e py che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, per ciascun punto:

determinare se, a meno di una costante EPSILON = 1.0e-9, il punto sta all'interno o all'esterno del triangolo ABC;
stampare a video il relativo messaggio (come mostrato nell'Esempio di esecuzione).
I valori px e py devono essere compresi nell'intervallo [0, 10.0).

Si assuma che:

i valori introdotti dall'utente che definiscono le rette r1, r2, ed r3 permettano di caratterizzare un triangolo ABC in cui la base AB sia parallela all'asse delle ascisse, ed il vertice C, opposto alla base, sia al di sopra della retta r1 su cui giace la base del triangolo:
       C
       /\
      /  \
     /____\
    A      B
i valori introdotti dall'utente che definiscono le rette r1, r2, ed r3 permettano di definire un triangolo ABC che contenga almeno un punto con coordinare px e py compresi nell'intervallo [0, 10.0).
Suggerimento: Un punto non appartiene al triangolo se sta al di sotto della retta r1, al di sopra della retta r2, o al di sopra della retta r3.
*/

import (
	. "fmt"
	"math/rand"
)

const EPSILON = 1.0e-9

func main() {

	var seed int64
	Print("Inserisci s: ")
	Scan(&seed)
	rand.Seed(seed)

	var m1, q1 float64
	Print("Inserisci m1 e q1: ")
	Scan(&m1, &q1)

	var m2, q2 float64
	Print("Inserisci m2 e q2: ")
	Scan(&m2, &q2)

	var m3, q3 float64
	Print("Inserisci m3 e q3: ")
	Scan(&m3, &q3)

	for i := 0; i < 10; i++ {
		px := rand.Float64() * 10
		py := rand.Float64() * 10

		if ÈXMinoreDiY(py, m1*px+q1) {
			Print("Punto (", px, ",", py, ") - Il punto sta all'esterno del triangolo.\n")
		} else if ÈXMaggioreDiY(py, m2*px+q2) {
			Print("Punto (", px, ",", py, ") - Il punto sta all'esterno del triangolo.\n")
		} else if ÈXMaggioreDiY(py, m3*px+q3) {
			Print("Punto (", px, ",", py, ") - Il punto sta all'esterno del triangolo.\n")
		} else {
			Print("Punto (", px, ",", py, ") - Il punto sta all'interno del triangolo.\n")
		}
	}

}

func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}
