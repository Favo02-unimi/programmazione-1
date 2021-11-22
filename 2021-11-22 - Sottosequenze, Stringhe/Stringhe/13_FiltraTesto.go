package main

/*
Scrivere un programma che:

legga da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote (""));
termini la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input l'indicatore End-Of-File (EOF);
ristampi le righe del testo letto in cui compaiono cifre.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione LeggiTesto() []string che legge da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote ("")) e terminato dall'indicatore EOF, restituendo un valore []string in cui sono memorizzate le stringhe corrispondenti alle righe del testo letto;
una funzione ContieneCifre(s string) bool che riceve in input un valore string nel parametro s e restituisce true se almeno un carattere in s è una cifra, false altrimenti;
una funzione FiltraTesto(sl []string) []string che riceve in input un valore []string nel parametro sl e restituisce un valore []string in cui sono memorizzate le stringhe di sl in cui compaiono cifre; la funzione deve utilizzare la funzione ContieneCifre().
Suggerimento: per sapere se un carattere è una cifra si può far riferimento alla documentazione della funzione unicode.IsDigit del package unicode.
*/

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

func main() {
	str := LeggiTesto()
	Println("Testo filtrato:")
	for _, v := range FiltraTesto(str) {
		Println(v)
	}
}

func LeggiTesto() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var str []string
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	return str
}

func ContieneCifre(s string) bool {
	for _, v := range s {
		if unicode.IsDigit(v) {
			return true
		}
	}
	return false
}

func FiltraTesto(sl []string) []string {
	var res []string
	for _, l := range sl {
		if ContieneCifre(l) {
			res = append(res, l)
		}
	}
	return res
}
