package main

/*
Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro righe costituite solo da simboli * (asterisco) intervallati da spazi, righe costituite solo da simboli + (più) intervallati da spazi e righe costituite solo da simboli o (lettera o) intervallati da spazi.
*/

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			if i%3 == 0 {
				Print("* ")
			} else if i%3 == 1 {
				Print("+ ")
			} else {
				Print("o ")
			}
		}
		Println()
	}
}
