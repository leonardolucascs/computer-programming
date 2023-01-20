package main

import "fmt"

func main() {
	var a []int
	a = []int{0, 1, 2, 3, 4, 5, 6}

	var b []int
	b = a[2:4]	// b = []int{2,3}



	b[0] = a[0]	// b[0] = 0, ma anche 
	b[len(b)-1] = a[0] // b[1] = 0


	//b[0] += 1
	//b[len(b)-1] -= 1
	// Comportamento di questa istruzione su a
	//		b[0], in cui b che fa riferimento all'array a, viene assegnato al valore a[0]
	// 		che si riflette su a, se b per la sua dichiarazione parte dal terzo elemento di a 
	//		tale assegnamento si riflette contro a stesso
	//		Modifiche successive alle variabile di b influenzeranno l'array a


	fmt.Println("a: ",a)
	fmt.Println("b: ", b)
}