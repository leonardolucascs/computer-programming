package main

import "fmt"

const (
	WIDTH = 4// 1920	
	HEIGHT = 3 // 1080
)

type pixel int //alias pixel per int

// Sapendo che: matrice[righe][colonne]

var screen [WIDTH][HEIGHT]pixel
// impostazione schermo verticale


// impostazione schermo orizzontale: screen[HEIGHT][WIDTH]

func main() {
	var i pixel

	for y := 0; y < HEIGHT; y++ {
		i=0

		for x := 0; x < WIDTH; x++ {
			screen[x][y] = i
			i++
			// riempimento verticale
			// [y] colonna fissa, [x] riga variabile
		}
	}

	/*for range case use
	for row := range screen {
		for column := range screen[0] {
			screen[row][column] = 6
			fmt.Println("\nI:",column,"\n", screen)
		}
	}
	*/

	/* (?) Cosa succede se al posto di screen[0], ci fosse screen[row] */
	for row := range screen {
		fmt.Println("row: ", screen[row])
		for column := range screen[row] {
			screen[row][column] = 9
			fmt.Println("I:",column ,"\n", screen)
		}
		fmt.Println()
	}

	// comportamento identico del for-range sopra 
	// che si metta screen[0] o screen[row] equivalgono
	// gli array interni(righe della matrice) hanno sempre
	// la stessa dimensione
	
}