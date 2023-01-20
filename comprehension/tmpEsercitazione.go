package main

import "fmt"

func main() {

	input := "Prova stringa"

	// Snippet code preso dal video l11: simulazione prova esame 20.21
	for inizio := range input {
		fmt.Println("inizio: ", inizio,"\n input: ", input)
		for fine := range input[inizio+1:] {
			fmt.Println("fine: ", fine, " \nsottostr: ", input[inizio+1:])
		}
	}
}