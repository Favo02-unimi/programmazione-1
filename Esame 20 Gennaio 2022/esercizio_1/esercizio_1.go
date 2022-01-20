package main

import (
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	ret1 := os.Args[1]
	ret2 := os.Args[2]

	//Println(ret1)
	//Println(ret2)

	//ax+by+c=0

	a1, a2 := extractCoeff(ret1, ret2, "null", "x") //non splitta su nulla dato che null non è in nessuna stringa
	//Println("A:", a1, a2)

	b1, b2 := extractCoeff(ret1, ret2, "x", "y")
	//Println("B:", b1, b2)

	var m1, m2 float64
	m1 = float64(-a1) / float64(b1)
	m2 = float64(-a2) / float64(b2)

	//Println("m1:", m1, "m2:", m2)

	const EPSILON float64 = 0.0001

	if math.Abs(m1-m2) < EPSILON {
		Println("le due rette sono parallele")
	} else {
		Println("le due rette non sono parallele")
	}
}

func extractCoeff(ret1, ret2, splitterInizio, splitterFine string) (int, int) {
	inizio1 := strings.Split(ret1, splitterInizio)
	inizio2 := strings.Split(ret2, splitterInizio)

	//Println(inizio1, inizio2)

	var coeff1, coeff2 []string

	if len(inizio1) > 1 { //se è riuscito a splittare prende solo la seconda parte per scartare la x
		coeff1 = strings.Split(inizio1[1], splitterFine)
		coeff2 = strings.Split(inizio2[1], splitterFine)
	} else { //altrimenti si prende la prima perchè non ha splittato
		coeff1 = strings.Split(inizio1[0], splitterFine)
		coeff2 = strings.Split(inizio2[0], splitterFine)
	}

	//Println(coeff1[0], coeff2[0])

	num1, _ := strconv.Atoi(coeff1[0])
	num2, _ := strconv.Atoi(coeff2[0])

	//Println(a1n, a2n)

	return num1, num2
}
