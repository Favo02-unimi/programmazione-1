package main

/*
Giulio Cesare usava per le sue corrispondenze riservate un codice di sostituzione molto semplice, nel quale la lettera chiara viene sostituita dalla lettera che la segue di tre posti nell’alfabeto: la lettera A è sostituita dalla D, la B dalla E, e così via fino alle ultime lettere che sono cifrate con le prime.

Chiaro: A B C D E F G H I J K L M N O P Q R S T U V W X Y Z

Cifrato: D E F G H I J K L M N O P Q R S T U V W X Y Z A B C

Più in generale si dice cifrario di Cesare un codice nel quale ogni lettera del messaggio chiaro viene spostata di un numero fisso k di posti (non necessariamente tre), dove k è detto chiave di cifratura.

Scrivere un programma che:

legga da standard input un numero intero k (la chiave di cifratura);
legga da standard input un messaggio in chiaro su più righe, terminando la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input l'indicatore End-Of-File (EOF);
stampi il messaggio cifrato (ottenuto con chiave di cifratura k) corrispondente al messaggio in chiaro letto.
Oltre alla funzione main(), il programma deve definire e utilizzare le seguenti funzioni:

una funzione LeggiNumero() int che legge da standard input un numero intero e ne restituisce il valore;
una funzione LeggiTesto() string che legge da standard input un testo su più righe e terminato dall'indicatore EOF, restituendo un valore string in cui è memorizzato il testo letto;
una funzione CifraCarattere(c rune, chiave int) rune che riceve in input un valore rune nel parametro c ed un valore int nel parametro chiave, e restituisce un valore rune uguale a c nel caso in cui c non sia una lettera dell'alfabeto inglese, uguale al valore cifrato corrispondente a c (ottenuto con chiave di cifratura chiave) altrimenti. In particolare, il valore cifrato deve essere minuscolo se c è minuscolo e maiuscolo se c è maiuscolo;
una funzione CifraTesto(s string, chiave int) string che riceve in input un valore string nel parametro s ed un valore int nel parametro chiave, e restituisce un valore string ottenuto cifrando ogni carattere di s tramite la funzione CifraCarattere().
*/

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	Println("Inserisci un numero:")
	var k int = LeggiNumero()
	Println("Inserisci un testo su più righe (termina con CTRL D):")
	var str string = LeggiTesto()
	Println("Testo cifrato:")
	Println(CifraTesto(str, k))
}

func LeggiNumero() int {
	var k int
	Scan(&k)
	return k
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

func CifraCarattere(c rune, chiave int) rune {
	var numa, numz, numA, numZ int
	numa = 'a'
	numz = 'z'
	numA = 'A'
	numZ = 'Z'

	var numC int = int(c)
	var newC int

	if numC >= numa && numC <= numz {
		newC = numC + chiave
		if newC > numz {
			newC = (numa - 1) + (newC - numz)
		}
		if newC < numa {
			newC = (numz + 1) - (numa - newC)
		}
		return rune(newC)
	}

	if numC >= numA && numC <= numZ {
		newC = numC + chiave
		if newC > numZ {
			newC = (numA - 1) + (newC - numZ)
		}
		if newC < numA {
			newC = (numZ + 1) - (numA - newC)
		}
		return rune(newC)
	}

	return c
}

func CifraTesto(s string, chiave int) string {
	var res []rune
	for _, v := range s {
		res = append(res, CifraCarattere(v, chiave))
	}
	return string(res)
}
