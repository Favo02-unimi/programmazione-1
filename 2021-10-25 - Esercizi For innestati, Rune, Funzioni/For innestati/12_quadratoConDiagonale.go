package main

/*
Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, stampi a video un quadrato di lato n in cui:

una diagonale è rappresentata utilizzando il carattere o (lettera o);
l'area del quadrato al di sotto della diagonale è rappresentata con il carattere * (asterisco);
l'area del quadrato al di sopra della diagonale è rappresentata con il carattere + (più);
i caratteri di ogni riga del quadrato sono separati tra di loro da un carattere spazio.
*/

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			if i == j {
				Print("o ")
			} else if i > j {
				Print("* ")
			} else {
				Print("+ ")
			}
		}
		Println()
	}
}
