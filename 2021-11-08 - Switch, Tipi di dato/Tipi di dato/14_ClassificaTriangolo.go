package main

/*
Scrivere un programma che legga da standard input sei valori reali:

il primo ed il secondo valore, xA e yA, rappresentano rispettivamente l'ascissa e l'ordinata di un punto A sul piano cartesiano;
il terzo ed il quarto valore, xB e yB, rappresentano rispettivamente l'ascissa e l'ordinata di un punto B sul piano cartesiano;
il quinto ed il sesto valore, xC e yC, rappresentano rispettivamente l'ascissa e l'ordinata di un punto C sul piano cartesiano.
Una volta terminata la fase di lettura, il programma deve stampare a video (come mostrato nell'Esempio di esecuzione), se il triangolo ABC definito dai segmenti/lati AB, BC e AC è equilatero, iscoscele o scaleno.

Un triangolo è equilatero se ha tutti e tre i lati di lunghezza uguale.

Un triangolo è isoscele se ha solo due lati di lunghezza uguale.

Un triangolo è scaleno se ha tutti e tre i lati di lunghezza diversa.

La lunghezza di ciascun lato di un triangolo è pari alla distanza euclidea tra gli estremi del lato.

Per esempio, la lunghezza del lato AB del triangolo ABC è pari alla distanza euclidea tra i punti A e B: ((xA-xB)2 + (yA-yB)2)1/2.

Le lunghezze dei lati del triangolo vanno confrontate a meno di una costante EPSILON = 1.0e-9.

Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione Distanza(x1, y1, x2, y2 float64) float64 che riceve in input:

due valori float64 nei parametri x1 e y1 che rappresentano rispettivamente l'ascissa e l'ordinata di un punto P1 sul piano cartesiano;

due valori float64 nei parametri x2 e y2 che rappresentano rispettivamente l'ascissa e l'ordinata di un punto P2 sul piano cartesiano;

e restituisce un valore float64 pari alla distanza euclidea tra i punti P1 e P2.
*/

import (
	. "fmt"
	"math"
)

const EPSILON = 1.0e-9

func main() {
	var xA, yA float64
	Print("Inserisci i valori dell'ascissa e dell'ordinata del punto A: ")
	Scan(&xA, &yA)

	var xB, yB float64
	Print("Inserisci i valori dell'ascissa e dell'ordinata del punto B: ")
	Scan(&xB, &yB)

	var xC, yC float64
	Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C: ")
	Scan(&xC, &yC)

	var ab, bc, ac float64

	ab = Distanza(xA, yA, xB, yB)
	bc = Distanza(xB, yB, xC, yC)
	ac = Distanza(xA, yA, xC, yC)

	if math.Abs(ab-bc) < EPSILON && math.Abs(bc-ac) < EPSILON {
		Println("Il triangolo ABC è equilatero.")
		Println("Lunghezza del lato:", ab)
		return
	}

	if math.Abs(ab-bc) < EPSILON {
		Println("Il triangolo ABC è isoscele.")
		Println("I lati di lunghezza uguale sono AB e BC.")
		Println("Lunghezza dei lati AB e BC:", ab)
		Println("Lunghezza del lato AC:", ac)
		return
	}

	if math.Abs(ab-ac) < EPSILON {
		Println("Il triangolo ABC è isoscele.")
		Println("I lati di lunghezza uguale sono AB e AC.")
		Println("Lunghezza dei lati AB e AC:", ab)
		Println("Lunghezza del lato BC:", bc)
		return
	}

	if math.Abs(bc-ac) < EPSILON {
		Println("Il triangolo ABC è isoscele.")
		Println("I lati di lunghezza uguale sono BC e AC.")
		Println("Lunghezza dei lati BC e AC:", bc)
		Println("Lunghezza del lato AB:", ab)
		return
	}

	Println("Il triangolo ABC è scaleno.")
	Println("Lunghezza del lato AB:", ab)
	Println("Lunghezza del lato BC:", bc)
	Println("Lunghezza del lato AC:", ac)

}

func Distanza(x1, y1, x2, y2 float64) float64 {
	n := math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
	return math.Pow(n, 0.5)
}
