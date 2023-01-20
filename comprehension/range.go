package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// cicla su pow ma stampa solo l'indice
	for v := range pow { 
		fmt.Printf("%d ", v)
	}
	
	fmt.Println()
	// stampa valori presenti in pow, acesso tramite indice pow[v]
	for v := range pow {
		fmt.Printf("%d ", pow[v])
	}

	fmt.Println()
	// stampa valori presenti in pow, acesso diretto a valore tramite x, e indice tramite i
	for i, x := range pow{
		fmt.Printf("\nidx %d: %d",i, x)
	}
}
