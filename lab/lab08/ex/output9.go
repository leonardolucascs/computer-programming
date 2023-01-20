package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4, 5} // len(b)=5=cap(b)
	stampa(b)

	modifica(b)
	stampa(b)

	//eliminaUltimoElemento(&b)	// sol.1
	// b = eliminaUltimoElemento(b)	// sol.2
	eliminaUltimoElemento(b)

	stampa(b)
}

func stampa(sl []int) {
	for _, v := range sl {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// Le modifiche che avvengono all'interno vengono viste anche
// all'esterno sebbene la slice sia passata per copia
// il ptr al suo interno permette di agire sull'array sottostante
func modifica(sl []int) {
	for i := range sl {
		sl[i] *= 2
	}
}

// Questa funzione elimina l'ultimo elemento della slice passata come argomento
// Ma RICORDO CHE: IN GO TUTTO VIENE PASSATO PER COPIA O VALORE
// Ciò avviene anche per le slice, costituite da un header (formato da un ptr, len, cap)
// dunque passato per copia
// In questo caso le modifche non vengono prese in carico dal caller perché
// non va ad agire sull'array sottostante ma bensì agisce sui valori locali del
// header della slice, ovvero sulla sua lunghezza, facendo un resciling
func eliminaUltimoElemento(sl []int) {
	sl = sl[:len(sl)-1]
}

/*
// sol.1
// Utilizzo di un puntatore
func eliminaUltimoElemento(sl *[]int) {
	(*sl) = (*sl)[:len(*sl)-1]
}

// sol.2
// Utilizzo di un return della slice modificata al caller
func eliminaUltimoElemento(sl []int) sl []int{
	sl = sl[:len(sl)-1]
	return
}
*/
