package main

import . "fmt"

func main() {
	var x int
	Print("Inserire numero: ")
	Scan(&x)
	if x > 0 {
		Println("+" + Sprint(x))
	} else {
		Println(Sprint(x))
	}
}
