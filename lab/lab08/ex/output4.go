package main

import "fmt"

func main() {
	var a [6]int

	// ciclo per inizializzare array a
	for i := range a {
		a[i] = i
	}

	var b []int
	b = a[:]	// inizializzazione slice b
	// b fa riferimento all'array a
	
	// prova: "non modifica ad a"
	copyA := a 	//copyA è un array distinto con gli elementi copiati da a
	b = copyA[:]

	for i := range b {
		b[i] = i * 2
	}

	fmt.Println(a)
	// le modifiche avvenute nell'ultimo ciclo for
	// vanno ad influenzare l'array a cui fa riferimento b
	// dunque 
}