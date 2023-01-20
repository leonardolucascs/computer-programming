package main

import "fmt"

func main() {
	items := []int{10, 20, 30, 40, 50}

	// a)
	for _, item := range items {
		item *= 2

		if item==100 {
			fmt.Println("Items inside for range:\n",items)
		}
		// item rappresenta una copia dell'elemento della slice all'indice x
		// non può essere utilizzato per modificare tale elemento direttamente
	}


	// b)
	for i := 0; i < len(items); i++{
		items[i] *= 2
	}

	// Certamente funziona, ma non ha molto senso occupo "più" spazio
	// uso 1 variabile in più rispetto al for sopra
	//for ix, item := range items {
	//	items[ix] = item * 2
	//}

	fmt.Println("Items outside for range:\n",items)
}