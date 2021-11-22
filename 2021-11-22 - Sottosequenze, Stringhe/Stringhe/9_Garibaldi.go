package main

/*
Scrivere un programma che:

legga da standard input un testo su più righe;
termini la lettura quando viene inserita da standard input una riga vuota ("");
ristampi il testo letto (riga vuota esclusa) sostituendo tutte le vocali con delle u.
Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione LeggiTesto() string che legge da standard input un testo su più righe e terminato da una riga vuota (""), restituendo un valore string in cui è memorizzato il testo letto (riga vuota esclusa);
una funzione TrasformaCarattere(c rune) rune che riceve in input un valore rune nel parametro c e restituisce un valore rune uguale a u se c è una vocale, e uguale a c altrimenti.
una funzione Garibaldi(s string) string che riceve in input un valore string nel parametro s e restituisce un valore string in cui ogni carattere è trasformato utilizzando la funzione TrasformaCarattere().
*/

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	Println("Inserisci un testo su più righe (termina con riga vuota):")
	var str = LeggiTesto()
	Println("Risultato trasformazione:")
	Println(Garibaldi(str))
}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		str += scanner.Text()
		str += "\n"
	}
	return str[:len(str)-1]
}

func TrasformaCarattere(c rune) rune {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return 'u'
	case 'A', 'E', 'I', 'O', 'U':
		return 'U'
	default:
		return c
	}
}

func Garibaldi(s string) string {
	var res []rune
	for _, v := range s {
		res = append(res, TrasformaCarattere(v))
	}
	return string(res)
}
