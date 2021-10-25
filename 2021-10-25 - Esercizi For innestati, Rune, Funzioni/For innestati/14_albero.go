package main

/*
Scrivere un programma che legga da standard input un intero n e, come mostrato nell'Esempio d'esecuzione, stampi a video un albero utilizzando il carattere * (asterisco) per rappresentarne la chioma ed il carattere + (più) per rappresentarne il tronco:

La chioma dell'albero deve essere alta n righe e, nel punto di larghezza massima, larga 2*n-1 colonne.
Il tronco dell'albero deve essere rappresentato con un quadrato di altezza e larghezza pari a 3 caratteri.
Il tronco dell'albero deve essere centrato rispetto alla chioma dell'albero.
Se n<=0, il programma stampa solo il tronco dell'albero.
*/

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	for i := 0; i < x+3; i++ {
		for j := 1; j < x*2; j++ {
			if i >= x {
				if j >= x-1 && j <= x+1 {
					Print("+")
				} else {
					Print(" ")
				}
				continue
			}
			if i+x < j {
				Print(" ")
				continue
			}
			if i+1 > x-j {
				Print("*")
				continue
			}
			Print(" ")
		}
		Println()
	}
}
