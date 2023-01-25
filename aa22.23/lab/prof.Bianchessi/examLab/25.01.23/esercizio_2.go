package main

import "fmt"
import "os"
import "strconv"

func main() {

	var numeri []int


	input := os.Args[1:]

	for _, v := range input {
		tmp, _ := strconv.Atoi(string(v))
		numeri = append(numeri, uint(tmp))
	}

	fmt.Println(numeri)

	fmt.Println(mcm(13, 78))
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

	var mcm int

	for i:=0; i<len(numeri)-1; i++{
		// caso base mcm(a, b)
		if len(numeri) == 2 {
			return int(mcm(numeri[i], numeri[i+1]))
		} else {
			mcm = mcm_sequenza(numeri[i:])
		}
	}

}