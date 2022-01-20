package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

type Prodotto struct {
	nome, codice string
}

type Magazzino struct {
	quantita map[string]int
	nomi     map[string]string
}

func main() {

	var istruzioni string = leggiRighe()

	var istruzioniSep []string = strings.Split(istruzioni, "\n")

	var mag Magazzino
	mag = nuovoMagazzino()

	for i, istruzione := range istruzioniSep {
		codice, nome, qnt := estraiCampi(istruzione)
		if codice == "" && nome == "" && qnt == 0 { //per skippare ultima riga che è solo una \n ma non riesco a togliere dallo scanner
			continue
		}
		//Println(codice, nome, qnt)
		magNuovo, ok := modificaProdotto(mag, Prodotto{nome, codice}, qnt)
		if ok {
			mag = magNuovo
		} else {
			Print("Errore alla riga ", i+1, ": Prodotto {", codice, " ", nome, "} - disponibilità ", mag.quantita[codice], ", richiesta ", qnt)
			return
		}
	}

	stampaMagazzino(mag)

}

func leggiRighe() (res string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res += scanner.Text()
		res += "\n"
	}
	return res
}

func nuovoMagazzino() Magazzino {
	var mag Magazzino
	mag.quantita = make(map[string]int)
	mag.nomi = make(map[string]string)
	return mag
}

func modificaProdotto(m Magazzino, p Prodotto, variazione int) (Magazzino, bool) {

	if _, ok := m.quantita[p.codice]; !ok { //se non esiste lo inserisco
		return inizializzaProdotto(m, p, variazione)
	}

	if m.quantita[p.codice]+variazione > 0 {
		//Println("Valida")
		m.quantita[p.codice] += variazione
		return m, true
	} else if m.quantita[p.codice]+variazione == 0 {
		//Println("Valida, rimuovere prodotto")
		delete(m.quantita, p.codice)
		delete(m.nomi, p.codice)
		return m, true
	} else {
		//Println("Non valido")
		return m, false
	}
}

func inizializzaProdotto(m Magazzino, p Prodotto, variazione int) (Magazzino, bool) {
	//controllo che la qnt sia positiva
	if variazione <= 0 {
		return m, false
	}
	m.nomi[p.codice] = p.nome
	m.quantita[p.codice] = variazione
	return m, true
}

func estraiCampi(s string) (codice, nome string, qnt int) {
	if !strings.Contains(s, ";") {
		return "", "", 0
	}
	pezzi := strings.Split(s, ";")
	codice = pezzi[0]
	nome = pezzi[1]
	qnt, _ = strconv.Atoi(pezzi[2])
	return
}

func stampaMagazzino(m Magazzino) {
	for k, v := range m.nomi {
		Print("Codice: ", k, ", Prodotto: ", v, ", Quantità:", m.quantita[k])
		Println()
	}
}
