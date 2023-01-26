package main

import "fmt"
import "os"
import "strconv"

func main() {

	var numeri []uint


	input := os.Args[1:]

	for _, v := range input {
		tmp, _ := strconv.Atoi(string(v))
		numeri = append(numeri, uint(tmp))
	}

	//fmt.Println(numeri)

	//fmt.Println(mcm(13, 78))
	fmt.Println(mcm_sequenza(numeri))

}



func MCD(a, b uint) uint {
	if b == 0 {
		return a
	} else {
		return MCD(b, a % b)
	}
}

func mcm(a, b uint) uint {
	return (a * b) / MCD(a, b)
}


func mcm_sequenza(numeri []uint) int {

	// non serve alcun tipo di ciclo !!!
	// Le "iterazioni cicliche" vengono "svolte" dalle chiamate ricorsive

	/*
	var mcm int

	for i:=0; i<len(numeri)-1; i++{
		// caso base mcm(a, b)
		if len(numeri) == 2 {
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