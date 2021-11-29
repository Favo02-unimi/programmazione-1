package main

/*
Scrivere un programma che:

legga da riga di comando una sequenza s di valori che rappresentano caratteri appartenenti all'alfabeto inglese (e quindi codificati all'interno dello standard US-ASCII (integrato nello standard Unicode));
stampi a video tutte le sottosequenze di caratteri presenti in s che:
iniziano e finiscono con lo stesso carattere;
sono formate da almeno 3 caratteri.
Ciascuna sottosequenza deve essere stampata un'unica volta, riportando il relativo numero di occorrenze della sottosequenza in s (cfr. Esecuzione d'esecuzione).

Se non esistono sottosequenze che soddisfano le condizioni 1 e 2, il programma non deve stampare nulla.

Si noti che una sottosequenza può essere contenuta in un'altra sottosequenza più grande.

Si assuma che la sequenza di valori specificata a riga di comando sia nel formato corretto e includa almeno 3 caratteri.
*/

import (
	. "fmt"
	"os"
)

func main() {
	var val []string = os.Args[1:]
	var str string
	for _, v := range val {
		str += string(v)
	}
	sott := Sottosequenze(str)
	var mappa map[string]int = make(map[string]int)
	for _, v := range sott {
		if v[0] == v[len(v)-1] {
			mappa[v]++
		}
	}
	for k, v := range mappa {
		for _, v := range k {
			Print(string(v), " ")
		}
		Println("-> Occorrenze:", v)
	}
}

func Sottosequenze(s string) (res []string) {
	for lungh := 3; lungh <= len(s); lungh++ {
		for start := 0; start < len(s)-lungh+1; start++ {
			res = append(res, s[start:start+lungh])
		}
	}
	return res
}
