package main

import . "fmt";

func main()  {
	var base, altezza int;
	Print("Inserisci la base: ");
	Scan(&base);
	Print("Inserisci l'altezza: ");
	Scan(&altezza);
	Print("Perimetro = ", base*2 + altezza*2)
	Println();
	Print("Area = ", (base*altezza))
}