package main

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	if x%10 == 0 {
		Println(x, "è multiplo di 10")
	} else {
		Println(x, "non è multiplo di 10")
	}
}
