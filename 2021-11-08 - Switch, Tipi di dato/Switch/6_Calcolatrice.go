package main

/*
Scrivere un programma che legga da standard input un numero reale, a, un simbolo di operazione aritmetica (+, -, * o /), ed un secondo numero reale b. Il programma deve calcolare e stampare a video il risultato dell'operazione specificata tra i due numeri. Se l'operatore non è riconosciuto il programma invece deve stampare Operatore non riconosciuto. Si utilizzi il costrutto switch per selezionare l'operazione in base al simbolo inserito.
*/

import (
	. "fmt"
)

func main() {
	Println("inserisci un'operazione aritmetica")
	var a, b float64
	var op rune
	Scanf("%f %c %f", &a, &op, &b)

	switch op {
	case '+':
		Println("risultato:", a+b)
	case '-':
		Println("risultato:", a-b)
	case '*':
		Println("risultato:", a*b)
	case '/':
		Println("risultato:", a/b)
	default:
		Println("Operatore non riconosciuto")
	}

}
