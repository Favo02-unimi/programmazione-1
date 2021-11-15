package main

/*
Scrivere un programma che legga da riga di comando una sequenza di valori intervallati da caratteri di spaziatura.

Il primo valore che definisce la sequenza (da sinistra verso destra) è in posizione 0, il secondo in posizione 1, etc.

La sequenza è valida se:

Tutti i valori letti rappresentano dei numeri interi.
Ciascun numero che appare in una posizione dispari all'interno della sequenza è minore del numero che lo precede.
Fatta eccezione per il numero che appare in posizione 0, ciascun numero che appare in una posizione pari all'interno della sequenza è maggiore del numero che lo precede.
Nel caso in cui la sequenza letta sia valida, il programma deve stampare:

Sequenza valida.

In caso contrario, il programma deve stampare:

Valore in posizione POSIZIONE non valido.

dove POSIZIONE è la posizione del primo valore che invalida la sequenza.

Ad esempio, se la sequenza di valori letta da riga di comando fosse:

5 4 9 abc 6

il programma deve stampare:

Valore in posizione 3 non valido.

Si assuma che la sequenza di valori letta da riga di comando sia definita da almeno un valore.
*/

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args
	var nums []int

	for i, s := range args {
		if i == 0 {
			continue
		}
		i--

		if n, ok := getNum(s); !ok {
			Println("Valore in posizione", i, "non valido.")
			return
		} else {
			nums = append(nums, n)
		}
	}

	for i, n := range nums {
		if i%2 == 0 {
			if i != 0 && n <= nums[i-1] {
				Println("Valore in posizione", i, "non valido.")
				return
			}
		} else {
			if n >= nums[i-1] {
				Println("Valore in posizione", i, "non valido.")
				return
			}
		}
	}

	Println("Sequenza valida.")

}

func getNum(s string) (int, bool) {
	if n, ok := strconv.Atoi(s); ok != nil {
		return 0, false
	} else {
		return n, true
	}
}
