/*

	ESERCIZIO FILTRO: Media cifre
	Letto da riga di comando un dato numero > 10:
		i. ricavare la media M delle singole cifre che compongono numero
		ii. stampare le singole cifre > M

	go run esercizio_filtro.go 35367
	Expected output: 567

*/

package main

import(
	"fmt"
	"os"
	"strconv"
)

func main() {

	var sum int
	var slIn []int

	input := os.Args[1]
	
	slIn = make([]int, len(input))

	for index, value := range input {
		//fmt.Printf("%T %v\n", value, value)
		tmpDigit, _ := strconv.Atoi(string(value))
		slIn[index] = tmpDigit
		sum += tmpDigit
	}

	M := sum/len(input)
	//fmt.Println("M: ",M)
	
	/*
	for i:=0; i<len(input); i++ {
		if slIn[i] > M {
			fmt.Print(slIn[i])
		}
		
	}
	*/

	for _, v := range slIn {
		if v > M {
			fmt.Print(v)
		}
	} 

	fmt.Println()
}