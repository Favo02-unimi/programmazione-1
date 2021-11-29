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

InserisciContatto(r Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica che data la rubrica r restituisce una nuova rubrica in cui è inserito il nuovo contatto creato con i dati passati come argomento. Se la rubrica r contiene già un contatto con identici nome e cognome non avviene nessuna modifica e la rubrica restituita è r stessa;

EliminaContatto(r Rubrica, cognome, nome string) Rubrica che restituisce una rubrica in cui è eliminato il contatto avente nome e cognome uguali a quelli passati in input. Se tale contatto non esiste, la rubrica restituita è r stessa;

StampaContatto(c Contatto) che stampa a video i dettagli del contatto c nel seguente formato: nome cognome: via numeroCivico, città, cap - Tel. telefono\n;

StampaRubrica(r Rubrica) che stampa a video tutti i contatti presenti nella rubrica utilizzando per ogni contatto la funzione StampaContatto(). La stampa deve rispettare il seguente formato:

Rubrica:\n
[1] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
[2] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
[3] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
AggiornaContatto(rubrica Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica che aggiorna i dettagli del contatto identificato da nome e cognome e restituisce la rubrica con il contatto aggiornato. Se il contatto non esiste, viene restituita la rubrica r stessa.

Esempio d'esecuzione:
Il formato dell'output dopo alcune operazioni di inserimento, cancellazione e modifica della rubrica dovrebbe risultare simile al seguente:

$ go run rubrica.go
Rubrica:
[1] - Mario Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
[2] - Anna Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
[3] - Carlo Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Anna Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
[2] - Carlo Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Anna Rossi: Via S. Sofia 25, Milano, 20122 - Tel.
[2] - Carlo Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
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
	contatto = Contatto{"Grande", "Puffo", "05435124", Indirizzo{"Via grande", "8732", "Puffolandia", 2}}
	r = InserisciContatto(r, contatto)

	Println("Inseriti 3 contatti:")
	StampaRubrica(r)

	r = EliminaContatto(r, "Grande", "Puffo")

	Println("Eliminato contatto Puffo:")
	StampaRubrica(r)

	r = AggiornaContatto(r, "Cognome", "Nome", "Via test2", 12, "1983", "Napoli", "73545343")

	Println("Aggiornato contatto cognome nome:")
	StampaRubrica(r)
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

func EliminaContatto(r Rubrica, cognome, nome string) Rubrica {
	var newRubrica Rubrica = NuovaRubrica()
	for _, cont := range r.contatti {
		if !(cont.cognome == cognome && cont.nome == nome) {
			newRubrica = InserisciContatto(newRubrica, cont)
		}
	}
	return newRubrica
}

func StampaContatto(c Contatto) {
	Println(c.nome, c.cognome, c.indirizzo.via, c.indirizzo.numeroCivico, c.indirizzo.citta, c.indirizzo.cap, "- Tel.", c.telefono)
}

func StampaRubrica(r Rubrica) {
	Println("Rubrica:")
	for i, v := range r.contatti {
		Print("[", i+1, "] - ")
		StampaContatto(v)
	}
}

func AggiornaContatto(rubrica Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica {
	var newRubrica Rubrica = NuovaRubrica()
	for _, cont := range rubrica.contatti {
		if cont.cognome == cognome && cont.nome == nome {
			cont.cognome = cognome
			cont.nome = nome
			cont.indirizzo.via = via
			cont.indirizzo.numeroCivico = numero
			cont.indirizzo.cap = cap
			cont.indirizzo.citta = città
			cont.telefono = telefono
		}
		newRubrica = InserisciContatto(newRubrica, cont)
	}
	return newRubrica
}
