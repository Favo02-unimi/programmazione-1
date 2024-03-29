package main

/*
Nel linguaggio farfallino ciascuna vocale non accentata (vocale) viene sostituita da una sequenza di tre caratteri vocale-f-vocale. Per esempio, la vocale a viene sostituita dalla sequenza afa, la vocale e dalla sequenza efe e così via. Se una vocale è maiuscola, anche la sequenza di tre caratteri che sostituisce la vocale deve essere definita da caratteri maiuscoli (ad esempio, la vocale A viene sostituita dalla sequenza AFA).

Scrivere un programma che:

legga da standard input un testo su più righe;
termini la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input l'indicatore End-Of-File (EOF);
ristampi il testo letto dopo averlo tradotto in linguaggio farfallino.
Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione LeggiTesto() string che legge da standard input un testo su più righe e terminato dall'indicatore EOF, restituendo un valore string in cui è memorizzato il testo letto;
una funzione TraduciCarattereInFarfallino(c rune) string che riceve in input un valore rune nel parametro c e restituisce un valore string che rappresenta la traduzione in linguaggio farfallino di c;
una funzione TraduciTestoInFarfallino(s string) string che riceve in input un valore string nel parametro s e restituisce un valore string che rappresenta la traduzione in linguaggio farfallino di s.
Suggerimento: per verificare se un carattere è una vocale potete utilizzare i costrutti if o switch, oppure utilizzare la funzione strings.ContainsRune() del package strings.
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
	Println(TraduciTestoInFarfallino(str))
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

func TraduciCarattereInFarfallino(c rune) string {
	if unicode.IsUpper(c) {
		return string(c) + string('F') + string(c)
	}
	return string(c) + string('f') + string(c)
}

func TraduciTestoInFarfallino(s string) string {
	var res string
	for _, v := range s {
		switch v {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			res += TraduciCarattereInFarfallino(v)
		default:
			res += string(v)
		}
	}
	return res
}
