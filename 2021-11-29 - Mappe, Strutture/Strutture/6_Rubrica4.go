package main

/*
Si consideri l'esercizio 4 Rubrica (2). Modificare il programma in modo tale che il tipo Rubrica sia implementato come una map[string]Contatto invece che come una slice []Contatto.

Nella nuova implementazione la chiave di tipo string della mappa sarà la stringa ottenuta concatenando nome e cognome del Contatto associato.
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
	contatti map[string]Contatto
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
	Println()

	r = EliminaContatto(r, "Grande", "Puffo")

	Println("Eliminato contatto Puffo:")
	StampaRubrica(r)
	Println()

	r = AggiornaContatto(r, "Cognome", "Nome", "Via test2", 12, "1983", "Napoli", "73545343")

	Println("Aggiornato contatto cognome nome:")
	StampaRubrica(r)
	Println()
}

func NuovaRubrica() Rubrica {
	var rubrica Rubrica
	rubrica.contatti = make(map[string]Contatto)
	return rubrica
}

func InserisciContatto(r Rubrica, c Contatto) Rubrica {
	if ControllaDuplicato(r, c) {
		return r
	}
	r.contatti[c.cognome+c.nome] = c
	return r
}

func ControllaDuplicato(r Rubrica, c Contatto) bool {
	_, ok := r.contatti[c.cognome+c.nome]
	return ok
}

func EliminaContatto(r Rubrica, cognome, nome string) Rubrica {
	_, ok := r.contatti[cognome+nome]
	if ok {
		delete(r.contatti, cognome+nome)
	}
	return r
}

func StampaContatto(c Contatto) {
	Println(c.nome, c.cognome, c.indirizzo.via, c.indirizzo.numeroCivico, c.indirizzo.citta, c.indirizzo.cap, "- Tel.", c.telefono)
}

func StampaRubrica(r Rubrica) {
	Println("Rubrica:")
	var i int
	for _, v := range r.contatti {
		Print("[", i+1, "] - ")
		StampaContatto(v)
		i++
	}
}

func AggiornaContatto(rubrica Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica {
	cont := rubrica.contatti[cognome+nome]

	cont.cognome = cognome
	cont.nome = nome
	cont.indirizzo.via = via
	cont.indirizzo.numeroCivico = numero
	cont.indirizzo.cap = cap
	cont.indirizzo.citta = città
	cont.telefono = telefono

	rubrica.contatti[cognome+nome] = cont

	return rubrica
}
