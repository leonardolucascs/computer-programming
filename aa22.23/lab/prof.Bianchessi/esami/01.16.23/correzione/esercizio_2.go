/*

	ESERCIZIO2: Trova mcm
	Dati 2 numeri da riga di comando, implementare:
		i. func Eprimo(int) bool, ritorna true-false se il numero è primo
		ii. func Scomponi(n int) map[int]int, scompone n nei suoi fattori primi
			-> stampare a video scomposizione
		iii. func mcm (n1 map, n2 map [int]int) -> restituisce mcm calcolato tra n1 e n2
		
	
	go run esercizio_2.go 350 100

	Expected output:
		350: 7^2 x 5^1 x 2^1
		100: 5^2 x 2^2
	
		mcm: 700

*/

package main

import (
	"os"
	"fmt"
	"strconv"
	"sort"
)

func main(){
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	//fmt.Printf("%d è primo: %t\n", n1, Eprimo(n1))
	//fmt.Printf("%d è primo: %t\n", n2, Eprimo(n2))

	//Scomponi(n1)

	//fmt.Println(nPow(2,3))

	fmt.Print(n1, ": "); stampaFattori(Scomponi(n1))
	fmt.Print(n2, ": "); stampaFattori(Scomponi(n2))
	
	fmt.Print("\nmcm:\t"); fmt.Println(mcm(Scomponi(n1), Scomponi(n2)))
}


func Eprimo(n int) bool{

	for i := 2; i<n; i++ {
		// Ottimizzazione del ciclo, condizione i<=n/2

		// Porre la condizione di cercare i divisori fino a n/2
		// in questo caso non ha senso, perché il ciclo si interrompe
		// già a partire da i:2, se fosse il caso 
		if n%i==0 {
			return false
		}
	} 

	return true
}


func Scomponi(n int) map[int]int {
	var fattoriN map[int]int

	fattoriN = make(map[int]int)

	// fattoriN := make(map[int]int)

	for i:=2; n>1; i++ {	// controllo i divisori fino a n > 1, siccome lavoro su n dentro
							// Ragiono sul valore di n come limite per uscire dal ciclo
							// = minor numero di iterazioni, rispetto a un controllo i < n
		for Eprimo(i) && n % i == 0 {	// itero divisione per lo stesso divisore finché % != 0
				n /= i
				fattoriN[i]++
		}
	}

	//fmt.Println("fattori di ", n, ":", fattoriN)
	
	return fattoriN
}

// By default Golang prints the map with sorted keys but while
// iterating over a map, it follows the order of the keys 
// appearing as it is.
func stampaFattori(fattoriMap map[int]int) {

	// Print elements map reverse sorted:
	// I version: sort ascending + print reverse
	// II version: sort descending + print [more accurated]
	
	// Sort a map by key
	keys := make([]int, 0, len(fattoriMap))

	// i. Save key into slide
	for k := range fattoriMap {
		keys = append(keys, k)
	}

	//fmt.Println("keys: ", keys)

	// I-VERSION
	// The Ints function sorts a slice of integer in ascending order.
	sort.Ints(keys)
	// [key] case -> strings: 		sort.Strings(keys)

	// print elements in map in reverse mode
	for i:=len(keys)-1; i >= 0; i-- {
		fmt.Print(keys[i], "^", fattoriMap[keys[i]])
		if i > 0 {
			fmt.Print(" x ")
		}
	}

	/*
	// II- VERSION
	sort.Sort(sort.Reverse(sort.Ints(keys)))
	
	// print elements in map
	for i, key := range keys {
		fmt.Print(key, "^", fattoriMap[key])
		if i != len(keys); fmt.Print(" x ")
	}
	*/

	fmt.Println()
}

func nPow(b, e int) int{
	
	tmp := 1

	for i:=0; i<e; i++{
		tmp *= b
	}

	return tmp
}

func mcm(n1Map map[int]int, n2Map map[int]int) int {

	//fmt.Println(n1Map)
	//fmt.Println(n2Map)
	
	// calcolo mcm
	// i. salvare per ogni mappa le proprie chiavi
	//    vincolo: ricerca e salva solo le chiavi non presenti
	// ii. utilizzo una terza mappa ausiliaria
	// iii. effettuare una ricerca per chiave
	//			se presente non fare nulla, 
	//			altrimenti salva
	//			per chiavi uguali, salva quella con value maggiore
	//			
	//			->posso usare la mappa ausiliaria per salvarmi direttamente
	// 			i fattori utili per l'mcm

	mcmMap := make(map[int]int)
	
	// copy key-value n1Map
	for k , v := range n1Map {
		mcmMap[k] = v
	}

	// iterate over n2Map, check equal key or save new value
	for k, v := range n2Map {
		if _, isPresent := mcmMap[k]; isPresent {	// se presente la stessa chiave
			// Svolgimento tramite switch...
			switch(true) {
			case n1Map[k] < v :
				mcmMap[k] = v
			// other cases unnecessary for this problem,
			// only stored v in case n2Map[k] > n1Map[k]
			// n1Map[k] already saved
			
			//case n1Map[k] > v:
				// do-something...


			// II version with if-else statement
				// if n1Map[k] < v{
				// 	...aggiorna valore relativo alla chiave k in mcmMap, valore n2Map[k] > n1Map[k]
				// } else if n1Map[k] > v {
				//	...
				// } else {
				//	...
				// }
			}
		} else { // altrimenti salva valore per chiave diversa
			mcmMap[k] = v
		}
	}

	//fmt.Println(mcmMap)

	mcm := 1

	for k, v := range mcmMap {
		mcm *= nPow(k,v)
	}

	return mcm
}











