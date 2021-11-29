package main

/*
Scrivere un programma che:

legga da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote (""));
termini la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input l'indicatore End-Of-File (EOF);
stampi a video le seguenti informazioni relative al testo letto:
Il numero di parole distinte presenti nel testo (una parola è una stringa interamente definita da caratteri il cui codice Unicode, se passato come argomento alla funzione func IsLetter(r rune) bool, fa restituire true alla funzione).
La lista di parole distinte presenti nel testo, riportando per ogni parola il relativo numero di occorrenze nel testo (cfr. Esecuzione d'esecuzione).
Oltre alle funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione LeggiTesto() string che legge da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote ("")) e terminato dall'indicatore EOF, restituendo un valore string in cui è memorizzato il testo letto;
una funzione SeparaParole(s string) []string che riceve in input un valore string nel parametro s e restituisce un valore []string in cui sono memorizzate tutte le parole presenti in s;
una funzione ContaRipetizioni(sl []string) map[string]int che riceve in input un valore []string nel parametro sl e restituisce un valore map[string]int in cui, per ogni parola presente in sl, è memorizzato il numero di occorrenze della parola in sl.
*/

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var str string = LeggiTesto()
	parole := SeparaParole(str)
	ripetizioni := ContaRipetizioni(parole)
	Println("Parole distinte:", len(ripetizioni))
	for k, v := range ripetizioni {
		Print(k, ": ", v)
		Println()
	}
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

func SeparaParole(s string) []string {
	righe := strings.Split(s, "\n")
	var parole []string
	for _, v := range righe {
		paroleRiga := strings.Split(v, " ")
		for _, p := range paroleRiga {
			parola, ok := CheckParola(p)
			if ok {
				parole = append(parole, parola)
			}
		}
	}
	return parole
}

func ContaRipetizioni(sl []string) map[string]int {
	var mappa map[string]int = make(map[string]int)
	for _, v := range sl {
		mappa[v]++
	}
	return mappa
}

func CheckParola(s string) (res string, ok bool) {
	var resRune []rune
	for _, v := range s {
		if unicode.IsLetter(v) {
			resRune = append(resRune, v)
		}
	}
	if len(resRune) > 0 {
		return string(resRune), true
	}
	return "", false
}
