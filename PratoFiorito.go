package main

import (
	"crypto/rand"
	. "fmt"
	"math"
	"math/big"
)

var dimX, dimY, bombs int      //dimX, dimY: dimensione X e Y del campo, bombs: bombe
var number [][]int8            //numero contenuto dalla cella, 0 se vuota e -1 se bomba
var isOpen, isFlagged [][]bool //isOpen: se la cella è coperta o visibile, isFlagged: se la cella è flaggata come bomba dal giocatore

func main() {

	Println("\u2797\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2797")
	Println("\U0001F44B Benvenuto nel prato fiorito di Favo")
	Println("\u2797\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2796\u2797")

	Println("\u2757 Selezionare difficoltà:")
	Println("1 - Facile: campo 9x9 con 10 bombe")
	Println("2 - Medio: campo 16x16 con 40 bombe")
	Println("3 - Difficile: campo 30x16 con 99 bombe")
	Println("4 - Custom: dimensioni e bombe personalizzabili")

	var difficulty int
	var loop bool = true

	for loop {
		Print("\u2757 Inserire difficoltà [1-4]: ")
		Scan(&difficulty)

		switch difficulty {
		case 1:
			dimY, dimX = 9, 9
			bombs = 10
			loop = false
		case 2:
			dimY, dimX = 16, 16
			bombs = 40
			loop = false
		case 3:
			dimY, dimX = 30, 16
			bombs = 99
			loop = false
		case 4:
			Print("Inserire dimensione del campo [larghezza altezza]")
			Scan(&dimY, &dimX)
			Print("Inserire numero di bombe [bombe]")
			Scan(&bombs)
			loop = false
		default:
			Println("\u274C Errore: seleziona la difficoltà inserendo un numero da 1 a 4")
		}
	}

	number = make([][]int8, dimX)
	isOpen = make([][]bool, dimX)
	isFlagged = make([][]bool, dimX)
	for i := 0; i < dimX; i++ {
		number[i] = make([]int8, dimY)
		isOpen[i] = make([]bool, dimY)
		isFlagged[i] = make([]bool, dimY)
	}

	placeBombs()

	placeNumbers()

	for {
		Println()
		printField(false)
		Println()
		Println("\U0001F6A9 Per flaggare/deflaggare una bomba scrivere le coordinate in negativo [-riga -colonna]")
		Println("\U0001F9E0 Pro-tip: se hai già flaggato le bombe intorno ad una cella puoi ri-aprirla")
		Print("\U0001F4C8 Inserire coordinate da aprire [riga colonna]: ")
		var x, y int
		Scan(&x, &y)

		isValid, isBomb := open(x, y)
		if !isValid {
			Println()
			Println("\u274C La mossa non è valida: la cella è già aperta, è flaggata o è fuori range")
		} else if isBomb {
			Println()
			printField(true)
			Println()
			Println("\U0001F4A3 Hai preso una bomba, hai perso.")
			break
		}

		if checkWin() {
			Println()
			Println("\u2705 Hai vinto, woooo ")
			Println()
			printField(false)
			break
		}
	}
}

func printField(debug bool) {
	for x := -1; x < dimX; x++ {
		// <Coordinate
		if x == -1 {
			Print("      ")
		} else {
			Printf("%2d    ", x+1)
		}
		// Coordinate>

		for y := 0; y < dimY; y++ {
			// <Coordinate
			if x == -1 {
				Printf("%2d ", y+1)
				if y == dimY-1 {
					Println()
				}
				continue
			}
			// Coordinate>

			if debug { //se il gioco è finito o debug
				if isFlagged[x][y] { //se l'utente ha flaggato
					if number[x][y] == -1 { //se è effettivamente una bomba
						Print("\U0001F6A9 ")
					} else { //se non è una bomba
						Print("\U0001F3F3  ")
					}
				} else {
					switch number[x][y] {
					case -1:
						Print("\U0001F4A3 ")
					case 0:
						Print("   ")
					default:
						Printf("%2d ", number[x][y])
					}
				}
			} else {
				if isFlagged[x][y] {
					Print("\U0001F6A9 ")
				} else if isOpen[x][y] {
					switch number[x][y] {
					case 0:
						Print("   ")
					default:
						Printf("%2d ", number[x][y])
					}
				} else {
					Print("-- ")
				}
			}
		}
		Println()
	}
	if debug {
		Println(" --- ")
		Println("Legenda:")
		Println("\U0001F6A9: Bomba flaggata e corretta")
		Println("\U0001F3F3: Flag messa ma non c'era la bomba")
		Println("\U0001F4A3: Bomba non flaggata")
		Println(" --- ")
	}
}

func placeBombs() {
	var x, y int
	for i := 0; i < bombs; i++ {

		randX, _ := rand.Int(rand.Reader, big.NewInt(int64(dimX)))
		randY, _ := rand.Int(rand.Reader, big.NewInt(int64(dimY)))

		x = int(randX.Int64())
		y = int(randY.Int64())

		//Println(x, y)

		if number[x][y] == -1 {
			i--
			continue
		}
		number[x][y] = -1
	}
}

func placeNumbers() {
	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			if number[x][y] == -1 {
				increaseNumber(x, y)
			}
		}
	}
}

func increaseNumber(x, y int) {
	for i := -1; i <= 1; i++ {
		if x+i >= dimX || x+i < 0 {
			continue
		}
		for j := -1; j <= 1; j++ {
			if y+j >= dimY || y+j < 0 {
				continue
			}
			if number[x+i][y+j] != -1 {
				number[x+i][y+j]++
			}
		}
	}
}

func open(x, y int) (isValid bool, isBomb bool) {
	if x < 0 && y < 0 {
		x++
		y++
		isValid = flag(x, y)
		return
	}
	x--
	y--
	if x >= dimX || y >= dimY || x < 0 || y < 0 {
		isValid = false
		return
	}
	if isFlagged[x][y] {
		isValid = false
		return
	}
	isValid = true
	if isOpen[x][y] {
		isBomb = openBombFlagged(x, y)
		return
	}
	if number[x][y] != -1 {
		isOpen[x][y] = true
		if number[x][y] == 0 {
			openRecursive(x, y)
		}
	} else {
		isBomb = true
		return
	}
	isBomb = false
	return
}

func openRecursive(x, y int) {
	for i := -1; i <= 1; i++ {
		//controllo out of bounds
		if x+i >= dimX || x+i < 0 {
			continue
		}
		for j := -1; j <= 1; j++ {
			//controllo out of bounds
			if y+j >= dimY || y+j < 0 {
				continue
			}
			//Println("x", x, "i", i, "x+i", x+i, "y", y, "j", j, "y+j", y+j)

			if isOpen[x+i][y+j] {
				continue
			} else if number[x+i][y+j] == 0 {
				isOpen[x+i][y+j] = true
				openRecursive(x+i, y+j)
			} else {
				isOpen[x+i][y+j] = true
			}
		}
	}
}

func flag(x, y int) (isValid bool) {
	x = int(math.Abs(float64(x)))
	y = int(math.Abs(float64(y)))
	if x >= dimX || y >= dimY || x < 0 || y < 0 {
		isValid = false
		return
	}
	if isOpen[x][y] {
		isValid = false
		return
	}
	isValid = true
	isFlagged[x][y] = !isFlagged[x][y]
	return
}

func openBombFlagged(x, y int) (isBomb bool) {
	var flagSurraunding int8
	for i := -1; i <= 1; i++ {
		//controllo out of bounds
		if x+i >= dimX || x+i < 0 {
			continue
		}
		for j := -1; j <= 1; j++ {
			//controllo out of bounds
			if y+j >= dimY || y+j < 0 {
				continue
			}
			if isFlagged[x+i][y+j] {
				flagSurraunding++
			}
		}
	}
	if flagSurraunding == number[x][y] {
		for i := -1; i <= 1; i++ {
			//controllo out of bounds
			if x+i >= dimX || x+i < 0 {
				continue
			}
			for j := -1; j <= 1; j++ {
				//controllo out of bounds
				if y+j >= dimY || y+j < 0 {
					continue
				}
				if !isFlagged[x+i][y+j] {
					if number[x+i][y+j] == -1 {
						isBomb = true
						return
					}
					isOpen[x+i][y+j] = true
					if number[x+i][y+j] == 0 {
						openRecursive(x+i, y+j)
					}
				}
			}
		}
		isBomb = false
		return
	} else {
		Println()
		Println("\u274C Non hai ancora flaggato le bombe intorno, non puoi aprire")
		return
	}
}

func checkWin() (win bool) {
	for i := 0; i < dimX; i++ {
		for j := 0; j < dimY; j++ {
			if !(isOpen[i][j] || number[i][j] == -1) {
				return false
			}
		}
	}
	return true
}
