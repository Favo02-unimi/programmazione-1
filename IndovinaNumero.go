package main

import . "fmt"
import "math/rand"

func main()  {
	Print("Ho generato un numero a caso da 0 a 100, indovinalo! ");
	var generata int = rand.Intn(100);
	var tentativi int = 0;
	var tentativo int;
	for true {
		Scan(&tentativo);
		tentativi++;
		if tentativo == generata {
			Println("Hai indovinato in ", tentativi, " tentativi");
			break;
		} else if tentativo < generata {
			Print("Sbagliato, il numero da indovinare è più grande. Riprova: ")
		} else if tentativo > generata {
			Print("Sbagliato, il numero da indovinare è più piccolo. Riprova: ")
		}
	}
}