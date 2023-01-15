package main

import (
	"fmt"
)

const Dimensione = 10

func main() {

	var a []int

	// creazione di una slice con numeri disposti in ordine decrescente
	for i := 0; i < Dimensione; i++ {
		a = append([]int{i + 1}, a...) // inserimento in cima ad a[]int di un nuovo elemento
	}

	fmt.Println(a)

	b := make([]int, len(a))
	copy(b, a)     // copy(dst, src []T) int -> n of elements copied
	fmt.Println(b) // b = a

	// Modifiche in b si riflettono in a?
	b[0] = 20
	fmt.Println("dopo modifica ad b[0]:\na:", a, "\nb:", b)

	// e viceversa?
	a[0] = 12
	fmt.Println("dopo modifica ad a[0]:\na:", a, "\nb:", b)

	c := make([]int, Dimensione)
	copy(c[:], a[Dimensione/2:]) // c[:] equivale a c

	// sarà composto solo dalla metà degli elementi in a, a partire dalla metà in poi
	fmt.Println(c)

}
