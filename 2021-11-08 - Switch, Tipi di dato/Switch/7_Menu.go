package main

/*
Scrivere un programma che permetta di ordinare al fast food con consegna a domicilio. Il programma deve chiedere iterativamente da standard input il numero corrispondente al tipo di articolo (1 patatine,2 hamburger o 3 cocacola) e la quantità richiesta. Con 0 viene terminata la lettura e viene stampato il costo dell'ordine. I prezzi sono 2€ per le patatine 5€ per gli hamburger e 2€ per la cocacola. In più ci sono 2€ di spesa per la consegna.

Si utilizzi il costrutto switch per selezionare l'articolo ed aggiornare il totale in base al numero inserito.
*/

import (
	. "fmt"
)

func main() {
	var s, q int
ciclo:
	for {
		Println(`Cosa vuoi ordinare?
			1. patatine 2€
			2. hamburger 5€
			3. cocacola 2€
			0. termina`)
		var sel int
		Scan(&sel)
		switch sel {
		case 1:
			Println("patatine? ottimo, quante?")
			Scan(&q)
			s += q * 2
		case 2:
			Println("hamburger? ottimo, quanti?")
			Scan(&q)
			s += q * 5
		case 3:
			Println("coca cola? ottimo, quante?")
			Scan(&q)
			s += q * 2
		case 0:
			break ciclo
		}
	}
	Println("Sono", s, " euro + 2 di consegna. Totale:", s+2)
}
