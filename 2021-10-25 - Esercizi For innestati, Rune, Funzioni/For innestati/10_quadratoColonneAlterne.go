package main

/*
Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro colonne costituite da simboli * (asterisco) a colonne costituite da simboli + (più).
*/

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			if j%2 == 0 {
				Print("* ")
			} else {
				Print("+ ")
			}
		}
		Println()
	}
}
