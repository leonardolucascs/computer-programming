// Esercizio 5.6

package main

import "fmt"

func main() {
	var str, final_str string

	// VERSIONE 2: USATE UN SOLO CICLO FOR E CONCATENAZIONE STRINGHE
	for i := 0; i < 5; i++ {
		str += "G"
		fmt.Print(str)
		fmt.Println()
		final_str += str + "\n"

	}

	fmt.Println("\nFINAL STRING:\n" + final_str)
}

//	VERSIONE 1: USATE DUE CICLI ANNIDATI
/*
for i := 0; i < 10; i++ {
	for j := 0; j <= i; j++ {
		fmt.Print("G")
		if j+1 > i {
			fmt.Println()
		}
	}
}
*/
