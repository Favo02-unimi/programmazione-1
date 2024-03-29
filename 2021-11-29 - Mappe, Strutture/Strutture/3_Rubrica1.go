package main

/*
Si consideri una rubrica in cui:

Ogni contatto deve avere un cognome, un nome, un indirizzo ed un numero di telefono.
Ogni indirizzo deve contenere le seguenti informazioni: via, numero civico, CAP e città.
Non possono esistere due contatti con lo stesso cognome e lo stesso nome.
Strutture dati da definire
Definire i seguenti tipi di dati:

Indirizzo: una struttura che memorizzi un indirizzo nei seguenti campi: via, cap e città di tipo string e numeroCivico di tipo uint;

Contatto: una struttura che memorizzi i dati di un contatto nei seguenti campi: cognome, nome e telefono di tipo string e indirizzo di tipo Indirizzo;

Rubrica: una slice dove ogni elemento è di tipo Contatto. Ogni elemento della slice rappresenta un contatto inserito nella rubrica. Una slice vuota rappresenta una rubrica vuota.

Funzioni da implementare
Implementare le funzioni:

NuovaRubrica() Rubrica che restituisce un valore Rubrica rappresentante una rubrica senza alcun contatto inserito;

InserisciContatto(r Rubrica, c Contatto) Rubrica che data la rubrica r restituisce una nuova rubrica in cui è inserito il nuovo contatto creato con i dati passati come argomento. Se la rubrica r contiene già un contatto con identici nome e cognome non avviene nessuna modifica e la rubrica restituita è r stessa;

Esempio d'esecuzione:
La stampa della rubrica dopo due operazioni di inserimento dovrebbe risultare simile al seguente:

$ go run rubrica.go
[{nome cognome 212312 {via x 20138 Milano 5}} {altronome altrocognome 4484848 {via y 20099 Sesto San Giovanni 8}}]
*/

import (
	. "fmt"
)

type Indirizzo struct {
	via, cap, citta string
	numeroCivico    uint
}

type Contatto struct {
	cognome, nome, telefono string
	indirizzo               Indirizzo
}

type Rubrica struct {
	contatti []Contatto
}

func main() {
	var r Rubrica = NuovaRubrica()
	var contatto = Contatto{"Favini", "Luca", "3333333333", Indirizzo{"Via puffo", "10021", "Mylano", 25}}
	r = InserisciContatto(r, contatto)
	contatto = Contatto{"Cognome", "Nome", "1111111111", Indirizzo{"Via asdsad", "545834", "Rowam", 54}}
	r = InserisciContatto(r, contatto)

	Println(r)
}

func NuovaRubrica() Rubrica {
	var rubrica Rubrica
	return rubrica
}

func InserisciContatto(r Rubrica, c Contatto) Rubrica {
	if ControllaDuplicato(r, c) {
		return r
	}
	r.contatti = append(r.contatti, c)
	return r
}

func ControllaDuplicato(r Rubrica, c Contatto) bool {
	for _, cont := range r.contatti {
		if cont.cognome == c.cognome && cont.nome == c.nome {
			return true
		}
	}
	return false
}
