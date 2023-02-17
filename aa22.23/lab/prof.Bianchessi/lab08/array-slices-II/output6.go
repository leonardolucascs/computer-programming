package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}

	stampa(a)
	modifica(a)
	//modifica(&a)
	stampa(a)
}

func stampa(a [6]int) {
	for _, v := range a {	// indice scartato
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// all'interno di questa funzione l'array passato dal main
// risulta essere una copia di a, dunque le modifiche non 
// si manifesteranno in a dichiarata nel main
func modifica(a [6]int) {
	for i := range a {
		a[i] *= 2
	}
}


// Per far si che a venga modificato si dovrà 
// 1. cambiare la signature/firma della funzione modifica ricevendo un indirizzo(quello di a)
//		func ....(a *[6]int){}
// 2. nel main a dovrà essere passata per riferimento, con &a, dunque passando il riferimento
// 		al suo indirizzo in memoria
// NON AVVIENE PIÚ IL PASSAGGIO PER COPIA (O VALORE) DELL'ARRAY A
/*
func modifica(a *[6]int) {
	for i := range a {
		a[i] *= 2
	}
}
*/
