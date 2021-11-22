package main

/*
Scrivere un programma che:

legga da standard input un testo su più righe;
termini la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input l'indicatore End-Of-File (EOF);
ristampi il testo letto con spaziatura, ovvero inserendo uno spazio ' ' tra ogni coppia di caratteri che non sono spazi.
Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione LeggiTesto() string che legge da standard input un testo su più righe e terminato dall'indicatore EOF, restituendo un valore string in cui è memorizzato il testo letto;
una funzione Spazia(s string) string che riceve in input un valore string nel parametro s e restituisce un valore string che rappresenta la versione spaziata di s.
Suggerimento: per sapere se un carattere è uno spazio utilizzate la funzione unicode.IsSpace del package unicode.
*/

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

func main() {
	Println("Inserisci un testo su più righe (termina con CTRL+D):")
	var str string = LeggiTesto()
	Println("Risultato:")
	Println(Spazia(str))
}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		str += scanner.Text()
		str += "\n"
	}
	return str
}

func Spazia(s string) string {
	var res []rune
	var str []rune = []rune(s)
	for i, v := range str[:len(str)-1] {
		if !unicode.IsSpace(str[i+1]) && !unicode.IsSpace(str[i]) {
			res = append(res, v, ' ')
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}
