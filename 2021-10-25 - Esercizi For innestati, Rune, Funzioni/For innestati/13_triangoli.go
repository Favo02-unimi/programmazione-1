package main

/*
Scrivere un programma che legga da standard input un intero n > 1 e stampi, utilizzando il carattere *, il perimetro di due triangoli rettangoli con base e altezza di lunghezza n, posizionati come mostrato nell'Esempio d'esecuzione.

$ go run triangoli.go
2
**
 *
 *
 **

$ go run triangoli.go
3
***
 **
  *
  *
  **
  ***
*/

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	for i := 0; i < x*2; i++ {
		for j := 0; j < x*2-1; j++ {
			if i == 0 && j < x { //prima riga metà sinistra
				Print("*")
				continue
			}
			if i == 2*x-1 && j >= x-1 { //ultima riga metà destra
				Print("*")
				continue
			}
			if j == x-1 { //colonna centrale
				Print("*")
				continue
			}

			if i <= j && i < x && j < x { //sopra o uguale alla diagonale, prima di metà, prima della colonna centrale
				Print("*")
				continue
			}

			if i >= j+1 && i >= x && j >= x { //sopra o uguale alla diagonale di quello sotto, dopo la metà, dopo la colonna centrale
				Print("*")
				continue
			}

			Print(" ")
		}
		Println()
	}
}
