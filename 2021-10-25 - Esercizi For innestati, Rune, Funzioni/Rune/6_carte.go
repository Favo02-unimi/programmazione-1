package main

import . "fmt"

func main() {
	var carta int = '\U0001F0B1'
	for i := 0; i <= 10; i++ {
		Printf("Simbolo: %c - Codice numerico in base 10: %d", carta, carta)
		println()
		carta++
	}

}
