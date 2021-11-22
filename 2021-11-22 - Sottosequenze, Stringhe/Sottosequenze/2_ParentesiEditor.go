package main

/*
Uno dei compiti più importanti di un compilatore è quello di controllare se le parentesi sono ben bilanciate. Ad ogni parentesi aperta deve corrispondere una parentesi chiusa, e le coppie di parentesi devono essere innestate propriamente.

Un esempio di parentesi tonde ben bilanciate è il seguente: (())()

Un esempio di parentesi tonde non ben bilanciate è il seguente: ())(

Notare che in quest’ultimo esempio, anche se ad ogni parentesi aperta corrisponde una parentesi chiusa, le coppie di parentesi non sono propriamente innestate (viene chiusa una parentesi di troppo ed una parentesi rimane aperta senza mai essere chiusa).

Parte 1
Leggere da riga di comando una stringa. Se la stringa contiene caratteri diversi da parentesi tonda aperta ( e parentesi tonda chiusa ) (che ricordiamo sono caratteri ASCII) terminare l’esecuzione del programma, altrimenti stampare bilanciata se la stringa ha parentesi bilanciate o non bilanciata altrimenti. A tal fine implementare e usare una funzione isBalanced(sequence string) bool che riceve in input una stringa sequence composta di sole parentesi (aperte e chiuse) e che restituisce il valore booleano true se la stringa sequence è ben bilanciata, o false altrimenti.

Nota: poichè le parentesi sono caratteri speciali nell'esecuzione dei programmi da riga di comando, per inserirle bisogna o fare l'escape di ciascuna oppure inserire la stringa tra virgolette.

Parte 2
Stampare tutte le sottostringhe ben bilanciate.
*/

import (
	. "fmt"
	"os"
)

func main() {
	var str string = os.Args[1]
	for _, r := range str {
		if !(r == '(' || r == ')') {
			Println("L'input fornito non aveva esclusivamente parentesi tonde.")
			return
		}
	}
	if isBalanced(str) {
		Println("bilanciata")
	} else {
		Println("non bilanciata")
	}

	Println("---")
	Println("Sottosequenze bilanciate:")

	sott := sott(str)
	if len(sott) == 0 {
		Println("Nessuna.")
	}
	for i, v := range sott {
		Print(i, ") ", v)
		Println()
	}
}

func isBalanced(sequence string) bool {
	var balance int
	for _, r := range sequence {
		if r == '(' {
			balance++
		} else if r == ')' {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

func sott(s string) (sott []string) {
	var sub string
	for lungh := 2; lungh <= len(s); lungh++ {
		for start := 0; start < len(s)-lungh+1; start++ {
			sub = s[start : start+lungh]

			if isBalanced(sub) {
				sott = append(sott, sub)
			}
		}
	}
	return
}
