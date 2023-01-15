package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	stampa(a)
	a = modifica(a)
	stampa(a)
	modificaConPuntatore(&a)
	stampa(a)
}

func stampa(a [6]int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func modifica(a [6]int) [6]int{
	for i := range a {
		a[i] *= 2
	}

	return a
}

func modificaConPuntatore(a *[6]int) {
	fmt.Println("a: ", a)
	fmt.Println("(*a): ", (*a))
	for i := range (*a) {
		(*a)[i] *= 2
	}

	/*
	// stesso effetto del for sopra
	// iterare con range su a o *a provoca lo stesso comportamento


	for i := range a {
		a[i] *= 2
	}
	*/
}