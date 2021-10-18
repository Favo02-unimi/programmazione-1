package main

import . "fmt"

func main() {
	var x, y, sum int
	Print("Inserire due numeri interi: ")
	Scan(&x, &y)
	for i := x + 1; i < y; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	Println("Somma =", sum)
}
