package main

import . "fmt"

func main() {
	var x int
	Print("Inserire un numero intero: ")
	Scan(&x)
	for i := 1; i <= 10; i++ {
		Println(i, "x", x, "=", i*x)
	}
}
