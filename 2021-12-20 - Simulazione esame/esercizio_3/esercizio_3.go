package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

type Utenza struct {
	numero string
	sim    string
}

type RegistroTelefonico = map[string][]string

func main() {
	utenze := LeggiUtenze()
	registro := InizializzaRegistro()
	for _, utenza := range utenze {
		AggiungiUtenza(registro, utenza)
	}
	for u, _ := range registro {
		if u[:3] == "338" {
			Println("Il numero", u, "è associato alla sim", Identifica(registro, u))
		}
	}
}

func LeggiUtenze() (utenze []Utenza) {
	utenze = make([]Utenza, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ";")
		telefono := split[0]
		sim := split[1]

		utenze = append(utenze, Utenza{telefono, sim})
	}

	return utenze
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(map[string][]string)
	return registro
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) {
	registro[utenza.numero] = append(registro[utenza.numero], utenza.sim)
	return registro
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	if _, ok := registro[telefono]; ok {
		return registro[telefono][len(registro[telefono])-1]
	} else {
		return ""
	}
}
