package main

import . "fmt"

var campo [3][3]int

func main() {
	stampaCampo(campo)
	var turno int
	for ; !controllaFine(); turno++ {
		Print("Tocca al giocatore " + Sprint(turno%2+1) + ", scrivere coordinate mossa [riga colonna]: ")
		var x, y int
		Scan(&x, &y)
		if !mossaValida(x, y) {
			Println("Mossa invalida, riprovare")
			turno--
			continue
		}
		faiMossa(x, y, turno)
		stampaCampo(campo)
	}
}

func stampaCampo(campo [3][3]int) {
	for i := 0; i < len(campo); i++ {
		for j := 0; j < len(campo[i]); j++ {
			if campo[i][j] == 0 {
				Print("-")
			} else if campo[i][j] == 1 {
				Print("X")
			} else if campo[i][j] == 2 {
				Print("O")
			}
			Print(" ")
		}
		Println()
	}
}

func mossaValida(x int, y int) bool {
	if x > 3 || y > 3 || x < 1 || y < 1 {
		return false
	}
	if campo[x-1][y-1] != 0 {
		return false
	}
	return true
}

func faiMossa(x int, y int, turno int) {
	campo[x-1][y-1] = turno%2 + 1
}

func controllaFine() bool {
	var fattoTris bool
	//orizzontale
	for i := 0; i < 3; i++ {
		if campo[i][0] != 0 && campo[i][0] == campo[i][1] && campo[i][1] == campo[i][2] {
			fattoTris = true
			Println("Giocatore", campo[i][0]+1, "ha fatto tris,", campo[i][0]+1, "HA VINTO")
			return fattoTris
		}
	}
	//verticale
	for i := 0; i < 3; i++ {
		if campo[0][i] != 0 && campo[0][i] == campo[1][i] && campo[1][i] == campo[2][i] {
			fattoTris = true
			Println("Giocatore", campo[i][0]+1, "ha fatto tris,", campo[0][i]+1, "HA VINTO")
			return fattoTris
		}
	}
	//diagonale 1
	if campo[0][0] != 0 && campo[0][0] == campo[1][1] && campo[1][1] == campo[2][2] {
		fattoTris = true
		Println("Giocatore", campo[0][0]+1, "ha fatto tris,", campo[0][0]+1, "HA VINTO")
		return fattoTris
	}
	//diagonale 2
	if campo[0][2] != 0 && campo[0][2] == campo[1][1] && campo[1][1] == campo[2][0] {
		fattoTris = true
		Println("Giocatore", campo[0][2]+1, "ha fatto tris,", campo[0][2]+1, "HA VINTO")
		return fattoTris
	}

	//campo vuoto
	var campoPieno bool = true
	for i := 0; i < len(campo); i++ {
		for j := 0; j < len(campo[i]); j++ {
			if campo[i][j] == 0 {
				campoPieno = false
			}
		}
	}
	if campoPieno {
		Println("Il campo è pieno - PAREGGIO")
	}

	return campoPieno
}
