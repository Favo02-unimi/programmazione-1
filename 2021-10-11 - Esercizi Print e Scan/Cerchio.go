package main

import . "fmt";
import "math";

func main()  {
	var raggio float64;
	Print("Inserisci il raggio: ");
	Scan(&raggio);
	Printf("Perimetro = %.3f", raggio*2*math.Pi)
	Println();
	Printf("Area = %.3f", raggio*raggio*math.Pi)
}