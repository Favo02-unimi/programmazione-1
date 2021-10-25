package main

/*
Scrivere un programma che legga legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, stampi a video un quadrato di lato n in cui:

il perimetro è rappresentato con il carattere * (asterisco);
l'area interna è rappresentata con il carattere + (più);
i caratteri di ogni riga del quadrato sono separati tra di loro da un carattere spazio.
*/

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == x-1 || j == 0 || j == x-1 {
				Print("* ")
			} else {
				Print("+ ")
			}
		}
		Println()
	}
}
