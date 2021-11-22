package main

/*
Scrivere un programma che:

legga da standard input un testo su più righe;
termini la lettura quando viene inserita da standard input una riga vuota ( "" ).
Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

aaaa/m/g
aaaa/m/gg
aaaa/mm/g
aaaa/mm/gg
g/m/aaaa
gg/m/aaaa
g/mm/aaaa
gg/mm/aaaa
Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette. In particolare, le date devono essere stampate in ordine cronologico (dalla più antica alla più recente).

Oltre alla funzione main() deve essere definita ed utilizzata la funzione Formatta(data string) string, che data una data in input codificata in uno degli 8 formati descritti sopra, la restituisce formattata aaaa-mm-gg. La funzione deve utilizzare strings.Split per estrarre giorno, mese, e anno dalla stringa di input.

Suggerimento: Una volta codificate nel formato aaaa-mm-gg, le date possono essere ordinate cronologicamente utilizzando la funzione sort.Strings(a []string) del package sort.
*/

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	str := LeggiTesto()
	var date []string
	for _, v := range str {
		date = append(date, Formatta(v))
	}
	sort.Strings(date)
	for _, v := range date {
		Println(v)
	}
}

func LeggiTesto() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var str []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		str = append(str, scanner.Text())
	}
	return str
}

func Formatta(s string) string {
	pezzi := strings.Split(s, "/")
	var anno, mese, giorno int
	if len(pezzi[0]) == 4 {
		anno, _ = strconv.Atoi(pezzi[0])
		mese, _ = strconv.Atoi(pezzi[1])
		giorno, _ = strconv.Atoi(pezzi[2])
	} else {
		anno, _ = strconv.Atoi(pezzi[2])
		mese, _ = strconv.Atoi(pezzi[1])
		giorno, _ = strconv.Atoi(pezzi[0])
	}
	return Sprintf("%d-%02d-%02d", anno, mese, giorno)
}
