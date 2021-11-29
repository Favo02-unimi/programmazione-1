package main

/*
Si consideri la rubrica dell'esercizio 3 Rubrica (1).

Si modifichi il programma in modo tale che le operazioni sulla rubrica vengano gestite tramite dei comandi letti da standard input.

Il programma deve:

creare una rubrica vuota;
leggere da standard input un testo su più righe (ogni riga è un comando);
terminare la lettura quando viene inserito da standard input l'indicatore EOF (CTRL+D);
eseguire in sequenza i comando letti eseguendo le funzioni corrispondenti.
Ogni comando è una riga di testo in cui il primo carattere determina il tipo dell'operazione:

I: inserimento di un contatto
E: cancellazione di un contatto
S: stampa della rubrica
A: aggiornamento di un contatto
I valori successivi rappresentano i valori da assegnare agli argomenti delle funzioni che effettuano l'operazione richiesta (le funzioni implementate nell'esercizio 3 Rubrica (1). I valori sono sempre separati tra di loro dal carattere ;.

Inserimento
Il comando per l'inserimento di un nuovo contatto ha il seguente formato:

I;cognome;nome;via;numeroCivico;cap;città;telefono
Cancellazione
Il comando per la cancellazione di un nuovo contatto ha il seguente formato:

E;cognome;nome
Stampa
Il comando per la stampa della rubrica è il seguente:

S
Aggiornamento di un contatto
A;cognome;nome;via;numeroCivico;cap;città;telefono
Suggerimento: per separare i valori presenti all'interno di una riga di testo utilizzate la funzione strings.Split() (usate go doc strings.Split per leggerne la documentazione).

Esempio d'esecuzione:
$ cat comandi.txt
I;Mario;Rossi;Via Celoria;18;20122;Milano;02503111
I;Mario;Rossi;Via Celoria;18;20122;Milano;
S
I;Elena;Bianchi;Via Celoria;18;20122;Milano;02503111
S
E;Mario;Rossi
S
A;Elena;Bianchi;Via Festa del perdono;7;20122;Milano;02503111
S

$ go run rubrica.go < comandi.txt
Rubrica:
[1] - Mario Rossi: Via Celoria 18, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Mario Rossi: Via Celoria 18, 20122, Milano - Tel. 02503111
[2] - Elena Bianchi: Via Celoria 18, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Elena Bianchi: Via Celoria 18, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Elena Bianchi: Via Festa del perdono 7, Milano, 20122 - Tel. 02503111
*/
import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
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
	var commands string = LeggiTesto()
	lines := strings.Split(commands, "\n")
	for _, line := range lines {
		tokens := strings.Split(line, ";")
		switch tokens[0] {
		case "I":
			//inserisci contatto
			//I;cognome;nome;via;numeroCivico;cap;città;telefono
			cognome := tokens[1]
			nome := tokens[2]
			via := tokens[3]
			numeroCivico := tokens[4]
			cap := tokens[5]
			citta := tokens[6]
			telefono := tokens[7]
			nCivico, _ := strconv.ParseUint(numeroCivico, 10, 64)
			var c = Contatto{cognome, nome, telefono, Indirizzo{via, cap, citta, uint(nCivico)}}
			r = InserisciContatto(r, c)
		case "E":
			//elimina contatto
			//E;cognome;nome
			cognome := tokens[1]
			nome := tokens[2]
			r = EliminaContatto(r, cognome, nome)
		case "S":
			//stampa rubrica
			StampaRubrica(r)
		case "A":
			//aggiorna contatto
			//A;cognome;nome;via;numeroCivico;cap;città;telefono
			cognome := tokens[1]
			nome := tokens[2]
			via := tokens[3]
			numeroCivico := tokens[4]
			cap := tokens[5]
			citta := tokens[6]
			telefono := tokens[7]
			nCivico, _ := strconv.ParseUint(numeroCivico, 10, 64)
			r = AggiornaContatto(r, cognome, nome, via, uint(nCivico), cap, citta, telefono)
		}
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
