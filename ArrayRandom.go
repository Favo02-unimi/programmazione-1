package main

import . "fmt"
import "math/rand"

func main()  {
	Print("Quanti array? ");
	var x int;
	Scan(&x);
	for i := 1; i < x; i++ {
		array := make([]int, i)
		for j := 0; j < len(array); j++ {
			array[j] = rand.Intn(10)
			Print(array[j])
		}
		Println();
	}
}