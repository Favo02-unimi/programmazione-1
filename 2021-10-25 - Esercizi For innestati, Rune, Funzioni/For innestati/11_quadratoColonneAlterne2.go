package main

/*
Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro 2 colonne costituite da simboli * (asterisco) a 2 colonne costituite da simboli + (più). In particolare, se n è dispari solo una delle due colonne più a destra sarà stampata.
*/

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	for i := 0; i < x; i++ {
		var asterischi bool = true
		var count int
		for j := 0; j < x; j++ {
			if count == 2 {
				asterischi = !asterischi
				count = 0
			}
			count++

			if asterischi {
				Print("* ")
			} else {
				Print("+ ")
			}
		}
		Println()
	}
}
