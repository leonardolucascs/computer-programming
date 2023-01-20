// Minimo, massimo e valor medio
package main

import (
	"fmt"
	"strconv"
	"os"
)

func Minimo(sl []int) int{
	var min int = sl[0]

	for i:=1; i<len(sl); i++{
		if (sl[i] < min){
			min = sl[i]
		}
	}

	return min
}

func Massimo(sl []int) int{
	max := sl[0]


	for i:=1; i<len(sl); i++{
		if (sl[i] > max){
			max = sl[i]
		}
	}
	
	return max
}

func Media(sl []int) float64{
	var res int

	for i:=0; i<len(sl); i++ {
		res += sl[i]
	}
	
	// questo da errore di risultato:
	// nel caso 10/4 il compilatore interpreta questa come una divisione tra interi
	// il che è giusto ma volendo un risultato di tipo float fare la conversione di tipo
	// sul risultato non ci darà il risultato atteso, ovvero 2.5, ma bensi 10/4=2<-float()==2.0
	// quindi bisogna assicurarsi che entrambi dividendo e divisore siano float affinché
	// si voglia avere un risultato corretto
	//		return float64(res/len(sl))

	// Questo avviene solo per le variabili, con costanti date in pasto al compilatore si ha un
	// comportamento diverso e non viene segnalato un errore
	// fmt.Println(10/4.00) == 2.5
	// fmt.Println(10.00/4) == 2.5
	// fmt.Println(10/4)	== 2

	return float64(res)/float64(len(sl))
}

func main() {
	var slIn []int = []int{}
	
	// Se uso make() non devo usare append() nel for che segue
	// perchè dopo make() slIn == [0 0 0]
	//slIn := make([]int, len(os.Args)-1)

	// Fill slIn
	for _, valItem := range os.Args {
		value, _ := strconv.Atoi(valItem)
		slIn = append(slIn, value)
	}
	
	/*

	// Another for loop to fill slIn
	for i:=1; i <= len(os.Args)-1; i++{
		slIn[i-1], _ = strconv.Atoi(os.Args[i])
	}
	
	*/

	// Deleting first element os.Args
	// Necessario solo se uso for-range loop
	slIn = slIn[1:]

	fmt.Println(slIn, len(os.Args))

	fmt.Println("Minimo: ", Minimo(slIn))
	fmt.Println("Massimo: ", Massimo(slIn))
	fmt.Printf("Media: %.2f", Media(slIn))
	fmt.Println()
}