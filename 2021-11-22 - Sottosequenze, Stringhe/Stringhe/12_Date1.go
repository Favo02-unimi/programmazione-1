package main

/*
Scrivere un programma che:

legga da standard input un testo su più righe;
termini la lettura quando viene inserita da standard input una riga vuota ("").
Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

aaaa/m/g
aaaa/m/gg
aaaa/mm/g
aaaa/mm/gg
Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette.

Oltre alla funzione main() devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione LeggiTesto() []string che legge da standard input un testo su più righe e terminato da una riga vuota (""), restituendo un valore []string in cui sono memorizzate le righe del testo letto;
una funzione Formatta(s string) string che riceve in input un valore string nel parametro s in cui è codificata una data nel formato aaaa/m/g, aaaa/m/gg, aaaa/mm/g o aaaa/mm/gg, e restituisce un valore string in cui la data in input è codificata nel formato aaaa-mm-gg.
*/

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := LeggiTesto()
	for _, v := range str {
		Println(Formatta(v))
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
	anno, _ := strconv.Atoi(pezzi[0])
	mese, _ := strconv.Atoi(pezzi[1])
	giorno, _ := strconv.Atoi(pezzi[2])
	return Sprintf("%d-%02d-%02d", anno, mese, giorno)
}
