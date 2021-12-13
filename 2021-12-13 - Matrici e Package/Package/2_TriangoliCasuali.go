package main

/*
Si estenda il package triangolo definito relativamente all'esercizio 1 Perimetro e Area di un triangolo implementando la funzione:

String(t Triangolo) string che riceve in input un'instanza del tipo Triangolo nel parametro t e restituisce un valore string che corrisponde alla rappresentazione string di t nel formato Triangolo con lati L1, L2 e L3., dove L1, L2 ed L3 sono i valori ai campi lato1, lato2 e lato3 di t. La conversione in string dei valori dei tre lati deve essere effettuata utilizzando due cifre decimali.
Utilizzando le funzionalità messe a disposizione dal package triangolo, scrivere un programma che:

legga da riga di comando un numero intero n;
generi in maniera casuale n triple di valori reali compresi tra 10 e 1000; i valori l1, l2, l3 di ciascuna tripla corrispondono alle misure dei lati di un ipotetico triangolo;
stampi a video la rappresentazione string del triangolo con area più grande tra quelli corrispondenti alle triple di valori reali generate;
stampi a video la rappresentazione string del triangolo con perimetro più piccolo tra quelli corrispondenti alle triple di valori reali generate.
Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

una funzione GeneraTriangoli(n int) (tN []triangolo.Triangolo) che riceve in input un valore int nel parametro n e restituisce un valore []triangolo.Triangolo nella variabile tN in cui sono memorizzate solamente le istanze di tipo triangolo.Triangolo valide.
*/

import (
	. "fmt"
	"math/rand"
	"time"

	"triangolo.com/main/triangolo"
)

func main() {
	var n int
	Print("Inserire numero: ")
	Scan(&n)
	triangoli := GeneraTriangoli(n)
	tAreaMax := triangoli[0]
	tPeriMin := triangoli[0]
	var area, perimetro float64
	area = triangolo.Area(tAreaMax)
	perimetro = triangolo.Perimetro(tPeriMin)
	for _, tr := range triangoli {
		/*Println("MAX:", area)
		Println("MIN:", perimetro)
		Println("new max:", triangolo.Area(tr))
		Println("new min:", triangolo.Perimetro(tr))*/
		if triangolo.Area(tr) > area {
			tAreaMax = tr
		}
		if triangolo.Perimetro(tr) < perimetro {
			tPeriMin = tr
		}
	}
	Println("Triangolo con area maggiore =", triangolo.String(tAreaMax))
	Println("Triangolo con perimetro minore =", triangolo.String(tPeriMin))
}

func GeneraTriangoli(n int) (tN []triangolo.Triangolo) {
	var l1, l2, l3 float64
	rand.Seed(int64(time.Now().Unix()))
	for i := 0; i < n; i++ {
		min := float64(10)
		max := float64(1000)
		l1 = min + rand.Float64()*(max-min)
		l2 = min + rand.Float64()*(max-min)
		l3 = min + rand.Float64()*(max-min)
		t, ok := triangolo.NuovoTriangolo(l1, l2, l3)
		if ok {
			tN = append(tN, t)
		}
	}
	return tN
}
