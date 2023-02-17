// Minimo, massimo e valor medio v.2
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

	return float64(res)/float64(len(sl))
}

func LeggiNumeri() (numeri []int) {
	// errore:
	// numeri []int non ha bisogno di una dichiarazione 
	// poiché già dichiarato come variabile di ritorno
	// 		var numeri []int

	// A differenza di e13.go questo costrutto con il controllo dell'errore
	// sulla conversione .Atoi() mi evita l'operazione di reslicing che 
	// avveniva nel precedente esercizio per eliminare il primo elemento 
	// di os.Args relativo al nome del programma in esecuzione preso dalla 
	// linea di comando
	for _, valItem := range os.Args {
		if intValue, err := strconv.Atoi(valItem); err == nil {
			numeri = append(numeri, intValue)
		}
	}

	return
}

func main() {
	var sliceN []int = []int{}

	sliceN = LeggiNumeri()

	fmt.Println("Minimo: ", Minimo(sliceN))
	fmt.Println("Massimo: ", Massimo(sliceN))
	fmt.Printf("Media: %.2f", Media(sliceN))
	fmt.Println()
}