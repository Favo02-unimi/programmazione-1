package main

import (
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	eq := os.Args[1]

	a, b, c := parsing(eq)
	//Println(a, b, c)

	delta := calcolaDelta(a, b, c)
	//Println(delta)

	sol1, sol2 := calcolaRes(a, b, c)

	if delta < 0 {
		Println("Non ci sono soluzioni reali")
	} else if delta == 0 {
		soglia, _ := strconv.ParseFloat(os.Args[2], 64)
		epsilon, _ := strconv.ParseFloat(os.Args[3], 64)
		Printf("Esiste un’unica soluzione reale: %.3f\n", sol1)
		if sol1 > (soglia + epsilon) {
			Printf("La soluzione %.3f è maggiore della soglia", sol1)
		}
	} else {
		soglia, _ := strconv.ParseFloat(os.Args[2], 64)
		epsilon, _ := strconv.ParseFloat(os.Args[3], 64)
		Printf("Esistono due soluzioni reali: %.3f e %.3f\n", sol1, sol2)
		if sol1 > (soglia + epsilon) {
			Printf("La soluzione %.3f è maggiore della soglia", sol1)
		}
		if sol2 > (soglia + epsilon) {
			Printf("La soluzione %.3f è maggiore della soglia", sol2)
		}
	}

}

func parsing(eq string) (int, int, int) {
	divisoPerPiu := strings.Split(eq, "+")
	divisoPerPiu = aggiungiSegno(divisoPerPiu, "+")
	//Println(divisoPerPiu)

	var divisoPerMeno []string
	for _, v := range divisoPerPiu {
		divisoPerMeno = append(divisoPerMeno, strings.Split(v, "-")...)
	}
	divisoPerMeno = aggiungiSegno(divisoPerMeno, "-")
	//Println(divisoPerMeno)

	var divisoPerUguale []string
	for _, v := range divisoPerMeno {
		divisoPerUguale = append(divisoPerUguale, strings.Split(v, "=")...)
	}
	//Println(divisoPerUguale)

	a := divisoPerUguale[0]
	b := divisoPerUguale[1]
	c := divisoPerUguale[2]

	apulito := a[:strings.Index(a, "x")]
	bpulito := b[:strings.Index(b, "x")]

	an, _ := strconv.Atoi(apulito)
	bn, _ := strconv.Atoi(bpulito)
	cn, _ := strconv.Atoi(c)

	return an, bn, cn
}

func aggiungiSegno(strs []string, segno string) (res []string) {
	res = make([]string, len(strs))
	for i, v := range strs {
		if (string(v[0]) != "+") && (string(v[0]) != "-") {
			res[i] = segno + v
		} else {
			res[i] = v
		}
	}
	return
}

func calcolaDelta(a, b, c int) float64 {
	return float64(b*b - 4*a*c)
}

func calcolaRes(a, b, c int) (float64, float64) {
	delta := calcolaDelta(a, b, c)

	if delta < 0 {
		return 0, 0
	} else if delta == 0 {
		res1 := (-1*float64(b) + math.Sqrt(delta)) / (2 * float64(a))
		return res1, 0
	} else {
		res1 := (-1*float64(b) + math.Sqrt(delta)) / (2 * float64(a))
		res2 := (-1*float64(b) - math.Sqrt(delta)) / (2 * float64(a))
		return res1, res2
	}
}
