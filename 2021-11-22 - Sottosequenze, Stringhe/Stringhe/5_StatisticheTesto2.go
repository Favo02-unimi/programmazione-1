package main

/*
Estendere il programma precedente in modo tale che il conteggio della lunghezza media delle parole non prenda in considerazione numeri e caratteri di separazione. Si implementi una funzione contaLettere(string) int che, data in input una parola derivata dal testo, restituisca il numero effettivo di lettere.

Suggerimento: per sapere se un carattere rappresenta una lettera utilizza la funzione unicode.IsLetter() del package unicode. Utilizzare go doc unicode.IsLetter per scoprire il suo funzionamento.
*/

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	Println("Inserisci un testo su più righe (termina con Ctrl+D):")
	var str string = LeggiTesto()
	linee, parole, lunghezzaTot := StatisticheParole(str)
	Println("Statistiche:")
	Println("Numero linee:", linee)
	Println("Numero parole:", parole)
	Println("Lunghezza media:", float64(lunghezzaTot)/float64(parole))
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

func StatisticheParole(s string) (int, int, int) {
	lines := strings.Split(s, "\n")
	var words []string
	words = make([]string, 0)
	for _, v := range lines {
		words = append(words, strings.Split(v, " ")...)
	}
	var chars int
	for _, v := range words {
		for _, p := range v {
			if unicode.IsLetter(p) {
				chars++
			}
		}
	}
	return len(lines) - 1, len(words) - 1, chars
}
