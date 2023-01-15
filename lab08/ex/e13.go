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

	return float64(res)/float64(len(sl))
}

func main() {
	slIn := make([]int, len(os.Args)-1)
	// slIn == []int{0, ... }  tanti 0 x n elementi di os.Args-1 
	// che sarebbero i numeri dati da riga di comando 

	// Fill slIn
	for i:=1; i <= len(os.Args)-1; i++{
		slIn[i-1], _ = strconv.Atoi(os.Args[i])
	}

	//fmt.Println(slIn, len(os.Args))

	fmt.Println("Minimo: ", Minimo(slIn))
	fmt.Println("Massimo: ", Massimo(slIn))
	fmt.Printf("Media: %.2f", Media(slIn))
	fmt.Println()
}