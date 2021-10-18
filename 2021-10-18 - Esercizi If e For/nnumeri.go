package main

import . "fmt"

func main() {
	var numeroNumeri int
	Scan(&numeroNumeri)
	Println("Inserisci", numeroNumeri, "interi")
	var num, sum, min, max, npos, nneg, nnull int
	for i := 0; i < numeroNumeri; i++ {
		Scan(&num)
		sum += num
		if i == 0 {
			min = num
			max = num
		}
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		if num == 0 {
			nnull++
		}
		if num > 0 {
			npos++
		}
		if num < 0 {
			nneg++
		}
	}
	Println("somma =", sum)
	Println("valore minimo =", min)
	Println("valore massimo =", max)
	Println("interi > 0 =", npos)
	Println("interi < 0 =", nneg)
	Println("interi = 0 =", nnull)
}
