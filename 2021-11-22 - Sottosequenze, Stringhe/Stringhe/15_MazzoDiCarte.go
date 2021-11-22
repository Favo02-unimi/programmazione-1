package main

/*
Scrivere un programma che:

Legga da riga di comando un numero intero n tale che 0 < n < 10.
Inizializzi una stringa che rappresenti un mazzo di carte formato dalle sole carte di cuori. i) Le carte di cuori corrispondono ai caratteri con codice Unicode compreso nell'intervallo che va da '\U0001F0B1' a '\U0001F0BA', estremi inclusi. ii) Le carte del mazzo non sono mischiate: la prima carta del mazzo è l'asse di cuori; la seconda carta del mazzo è il due di cuori;... l'ultima carta del mazzo è il dieci di cuori.
Simuli l'estrazione casuale (ed in sequenza) di n carte dal mazzo, stampando a video le carte estratte e quelle rimaste nel mazzo dopo ogni estrazione.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione LeggiNumero() int che legge da riga di comando un numero intero e ne restituisce il valore;
una funzione GeneraMazzo() string che restituisce un valore string che rappresenta un mazzo di carte formato dalle sole carte di cuori;
una funzione EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) che riceve in input un valore string nel parametro mazzo e restituisce un valore di tipo rune nella variabile cartaEstratta ed un valore di tipo string nella variabile mazzoResiduo, che rappresentano rispettivamente la carta estratta in modo casuale dal mazzo di carte rappresentato da mazzo ed il mazzo residuo di carte dopo l'estrazione;
una funzione EstraiCarte(mazzo string, n int) che riceve in input un valore string nel parametro mazzo ed un valore int nel parametro n e simula l'estrazione casuale (ed in sequenza) di n carte dal mazzo di carte rappresentato da mazzo, stampando a video le carte estratte e quelle rimaste nel mazzo dopo ogni estrazione; la funzione deve utilizzare la funzione EstraiCarta().
Suggerimento: per generare in modo casuale un numero intero, potete utilizzare le funzioni dei package math/rand e time come mostrato nel seguente frammento di codice:

inizializzazione del generatore di numeri casuali
rand.Seed(int64(time.Now().Nanosecond()))
generazione di un numero casuale compreso nell'intervallo
   che va da 0 a 99 (estremi inclusi)
numeroGenerato := rand.Intn(100)
*/

import (
	. "fmt"
	"math/rand"
	"time"
)

func main() {
	n := LeggiNumero()
	mazzo := GeneraMazzo()
	EstraiCarte(mazzo, n)
}

func LeggiNumero() int {
	var k int
	Scan(&k)
	return k
}

func GeneraMazzo() string {
	var mazzo string = "1234567890" //ogni carta corrisponde al suo numero, 10 = 0
	return mazzo
}

func EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) {
	rand.Seed(int64(time.Now().Nanosecond()))
	n := rand.Intn(len(mazzo))
	carta := mazzo[n : n+1]
	runaCarta := CodificaCarta(carta)
	return runaCarta, mazzo[:n] + mazzo[n+1:]
}

func EstraiCarte(mazzo string, n int) {
	var carta rune
	for i := 0; i < n; i++ {
		carta, mazzo = EstraiCarta(mazzo)
		Print("Estratta la carta ", string(carta), " - Carte rimaste nel mazzo:")
		PrintMazzo(mazzo)
		Println()
	}
}

func CodificaCarta(s string) rune {
	switch s {
	case "0":
		return '\U0001F0BA'
	case "1":
		return '\U0001F0B1'
	case "2":
		return '\U0001F0B2'
	case "3":
		return '\U0001F0B3'
	case "4":
		return '\U0001F0B4'
	case "5":
		return '\U0001F0B5'
	case "6":
		return '\U0001F0B6'
	case "7":
		return '\U0001F0B7'
	case "8":
		return '\U0001F0B8'
	case "9":
		return '\U0001F0B9'
	}
	return '\U0001F0B1'
}

func PrintMazzo(mazzo string) {
	for _, v := range mazzo {
		Print(string(CodificaCarta(string(v))))
	}
}
