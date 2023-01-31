package main

import "fmt"

func main() {
	var sl []int = []int{54, 23, 88, 90, 32, 17, 55, 78}

	for i := len(sl); i > 0; i-- {
		//stampa(sl[:i])
	}

	sl1 := []int{2, 4, 5, 7}

	fmt.Println()
	fmt.Println(sum(sl1))

}

func stampa(sl []int) {
	fmt.Printf("%v\n", sl)
}

// Voglio sviluppare una chiamata a sum({2,3,4,5}) -> e poi in func sum(args ...int)
// func sumGeneric(..args int){
//	for _, v := range
//		... sum each element
//}

// Somma ricorsiva di N numeri
func sum(sl []int) int {
	fmt.Println(sl)

	// NON SERVE NESSUN TIPO DI CICLO!
	// Le "iterazioni cicliche" vengono "svolte" dalle chiamate ricorsive
	// caso base: somma tra 2 numeri
	if len(sl) == 2 {
		fmt.Println(sl[0], "+", sl[1])
		return sl[0] + sl[1]
	} else { // passo induttivo:

		// ricorsione partendo a sommare gli elementi in coda
		// i primi ad essere sommati saranno quelli in posizione len(n)-1 e len(n)-2
		// poi per retroattività verranno ricavati gli altri risultati

		//return sl[0] + sum(sl[1:])

		tmpSum := sum(sl[:len(sl)-1])

		fmt.Println(tmpSum, "+", sl[len(sl)-1])

		// ricorsione partendo a sommare gli elementi in cima

		return sl[len(sl)-1] + tmpSum
	}

}

// SOMMA RICORSIVA +1 PER NUMERO DI ELEMENTI
func countEl(numeri []uint) int {

	for i := len(numeri); i > 0; i-- {

		fmt.Println("Stato iniziale: ", numeri)

		// casi mutuamente esclusivi
		// caso base: 2 elementi
		if len(numeri) == 2 {
			return 2
		} else if len(numeri) == 1 {
			// caso base: 1 elemento
			return 1
		} else {
			fmt.Println("i:", i)
			fmt.Println("Chiamata funzione ricorsiva: countEl(", numeri[:i-1], ")")
			return 1 + countEl(numeri[:i-1])
		}
	}
	return 0
}

// SOMMA RICORSIVA ELEMENTI IN UNA SLICE
func countEl1(numeri []uint) int {

	for i := len(numeri); i > 0; i-- {

		fmt.Println("Stato iniziale: ", numeri)

		// caso base:
		if len(numeri) == 2 {
			return int(numeri[0]) + int(numeri[1])
		} else { // passo induttivo:
			fmt.Println("i:", i)
			//i--
			fmt.Println("Chiamata funzione ricorsiva: countEl1(", numeri[:i-1], ")")
			return int(numeri[i-1]) + countEl1(numeri[:i-1])
		}
	}

	return 0 // necessario se non entra nel for per len(numeri)==0
	// somma = 0
}



// Dati due numeri cacolo mcm
func mcm(a, b uint) uint {
	return (a * b) / MCD(a, b)
}

// Data una sequenza di numeri ricavare mcm ricorsivo
func mcm_sequenza(numeri []uint) int {

	// non serve alcun tipo di ciclo !!!
	// Le "iterazioni cicliche" vengono "svolte" dalle chiamate ricorsive

	/*
	var mcm int

	for i:=0; i<len(numeri)-1; i++{
		// caso base mcm(a, b)
		if len(numeri) == 2 {ccd
			return int(mcm(numeri[i], numeri[i+1]))	// ERROR: Cannot call non-function mcm => value type int
													// shadowing variable name & function name!!!
		} else {
			mcm = mcm_sequenza(numeri[i:])
		}
	}
	*/


	fmt.Println("\nstato iniziale: ", numeri)

		// caso base mcm(n1, n2)
		if len(numeri)==2{
			fmt.Printf("...attivazione caso base: mcm(%d, %d)\n\n", numeri[0], numeri[1])
			
			output := mcm(numeri[0], numeri[1])

			fmt.Println("mcm(", numeri[0], ", ", numeri[1] ,") => ", output)

			return int(output)	// int( mcm(numeri[0], numeri[1]) )
		}else{	// passo induttivo:
			fmt.Println("...chiamata funzione ricorsiva: mcm(", numeri[:len(numeri)-1], ")")

			tmp := uint(mcm_sequenza( numeri[:len(numeri)-1] ) )
			op := mcm( numeri[len(numeri)-1], tmp )
			
			fmt.Println("mcm(", numeri[len(numeri)-1], ", mcm(", numeri[:len(numeri)-1] ,") )")
			fmt.Println("mcm(", numeri[len(numeri)-1],", ", tmp ,") =>", op)
			
			return int(op)	// int( mcm( numeri[len(numeri)-1], uint(mcm_sequenza( numeri[:len(numeri)-1] ) ) ) )
		}
	
	return 0	// necessario se non entra nel for per len(numeri)==0
}
