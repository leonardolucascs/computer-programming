package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	LeggiEStampa(scanner)
}

// Lettura multiple riga da std.input
func LeggiEStampa(scanner *bufio.Scanner) {
	// Fin tanto che c'è qualcosa da leggere, stampa
	if scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		
		// La chiamata ricorsiva permette la continua lettura da std.input
		LeggiEStampa(scanner)

		// Premendo la combinazioni tasti CTRL+D, viene inserito da std.input 
		// l'indicatore di End-Of-File(EOF) definendo così la fine della 
		// lettura
	}
}

// La redirezione dell'output avviene tramite " > " sul terminale
// Mentre la redirezione dell'input avviene tramite " < "