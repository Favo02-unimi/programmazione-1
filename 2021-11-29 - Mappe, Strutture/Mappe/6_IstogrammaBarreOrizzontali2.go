package main

/*
Scrivere un programma che:

legga da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote (""));
termini la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input l'indicatore End-Of-File (EOF);
come mostrato nell'Esempio di esecuzione, stampi un istogramma a barre orizzontali per rappresentare il numero di occorrenze di ogni lettera presente nel testo letto:
una lettera è un carattere il cui codice Unicode, se passato come argomento alla funzione func IsLetter(r rune) bool del package unicode, fa restituire true alla funzione;
le lettere minuscole sono da considerarsi diverse dalle lettere maiuscole;
ogni barra viene rappresentata utilizzando il carattere asterisco (*); se il numero di occorrenze della lettera e è per esempio 9, la barra corrispondente sarà formata da 9 caratteri *;
le barre devono essere stampate a partire da quella associata alla lettera con codice Unicode più piccolo fino a quella associata alla lettera con codice Unicode più grande.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione StampaIstogramma(occorrenze map[rune]int) che riceve in input un valore map[rune]int nel parametro occorrenze, in cui ad una data lettera è associato un dato numero di occorrenze, e stampa l'istogramma relativo alle lettere presenti come valori chiave in occorrenze secondo quanto descritto ai punti iii e iv.
*/

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	var str string = LeggiTesto()
	var mappa map[rune]int = Occorrenze(str)
	var chiavi []string
	Println("Occorrenze:")
	for k := range mappa {
		chiavi = append(chiavi, string(k))
	}
	sort.Strings(chiavi)
	for _, v := range chiavi {
		Print(string(v), ": ")
		for i := 0; i < mappa[[]rune(v)[0]]; i++ {
			Print("*")
		}
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

func Occorrenze(s string) map[rune]int {
	var mappa map[rune]int
	mappa = make(map[rune]int)
	for _, v := range s {
		if unicode.IsLetter(v) {
			mappa[v]++
		}
	}
	return mappa
}
