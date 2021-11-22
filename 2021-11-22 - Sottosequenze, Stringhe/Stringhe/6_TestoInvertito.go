package main

/*
Scrivere un programma che legga da standard input un testo formato da un numero variabile di righe, terminando la lettura quando viene inserita da standard input una riga vuota (""), e lo ristampi dall’ultimo carattere al primo.

Scrivere un programma che:

legga da standard input un testo su più righe;
termini la lettura quando viene inserita da standard input una riga vuota ("");
ristampi il testo letto (riga vuota esclusa) dall’ultimo carattere al primo.
Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione LeggiTesto() string che legge da standard input un testo su più righe e terminato da una riga vuota (""), restituendo un valore string in cui è memorizzato il testo letto (riga vuota esclusa);
una funzione InvertiStringa(s string) string che riceve in input un valore string nel parametro s e ne inverte l'ordine dei caratteri, ovvero restituisce un valore string in cui il primo carattere è l'ultimo che definisce s, il secondo carattere è il penultimo che definisce s, ... e l'ultimo carattere è il primo che definisce s.
*/

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	var str string = LeggiTesto()
	var inv string = InvertiStringa(str)
	Println("Testo invertito:")
	Println(inv)
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

func InvertiStringa(s string) string {
	var str []rune = []rune(s)
	var len int = len(str)
	var res []rune

	res = make([]rune, len)
	for i := 0; i < len; i++ {
		res[len-1-i] = str[i]
	}
	return string(res)
}
