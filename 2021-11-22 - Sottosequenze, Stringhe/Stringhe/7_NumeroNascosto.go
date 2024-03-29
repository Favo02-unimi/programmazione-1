package main

/*
Scrivere un programma che legga da standard input una riga di testo e stampi in output il doppio del numero nascosto all'interno della riga di testo, ovvero il doppio del numero che si ottiene concatenando le cifre presenti all'interno della riga di testo. Il programma non stampa nulla se non è presente alcun numero nascosto.

Suggerimento: Per convertire una stringa in un numero intero utilizzate la funzione strconv.Atoi() del package strconv. Invece, per sapere se un carattere è una cifra utilizzate la funzione unicode.IsDigit() del package unicode.

Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione LeggiTesto() string che legge da standard input una riga di testo, restituendo un valore string in cui è memorizzato il testo letto;
una funzione NumeroNascosto(testo string) (int, error) che riceve in input un valore string nel parametro testo e restituisce due valori:
il primo valore è un numero intero che rappresenta il numero nascosto all'interno del testo. Se il testo in input non contiene alcun numero il valore restituito deve essere 0;
il secondo valore è l'eventuale errore restituito dalla funzione strconv.Atoi().
*/

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var str string = LeggiTesto()
	num, err := NumeroNascosto(str)
	if err != nil {
		Println("Nessun numero nascosto")
	} else {
		Print("Doppio del numero nascosto: ", num*2, " (", num, " * 2)")
	}
}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func NumeroNascosto(testo string) (int, error) {
	var num []rune
	for _, v := range testo {
		if unicode.IsDigit(v) {
			num = append(num, v)
		}
	}
	n, err := strconv.Atoi(string(num))
	return n, err
}
