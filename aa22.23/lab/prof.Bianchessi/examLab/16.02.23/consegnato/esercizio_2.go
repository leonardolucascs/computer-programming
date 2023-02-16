package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	
	// multiIns := LeggiMultiInsiemi()
	// fmt.Println(multiIns)

	A := map[int]int{4:1, 3:1, 2:1, 1:1}
	B := map[int]int{5:1, 4:1, 3:2}

	fmt.Println(IntersecaMultiInsiemi(A, B))
	fmt.Println(UnisciMultiInsiemi(A,B))

}

func LeggiMultiInsiemi() []map[int]int{
	var multiIns []map[int]int // (?) perché solo dichiarandola non esce come nel caso delle map 
							   // 		panic: assignment to entry in nil map

	scanner := bufio.NewScanner(os.Stdin)

	// Ciclo for continuo...
	for {
		// Controlla fintanto che c'è qualcosa da scannerizzare e che sia len() > 0 
		if scanner.Scan() && len(scanner.Text())>0 {
			input := strings.Split(scanner.Text(), " ")
			// Ogni lettura(riga) viene salvata in multiIns...
			// Si ha bisogno prima però di una conversione dei dati letti, nel formato corretto
			multiIns = append(multiIns, convertiInput(input)) 
		} else {
			break
		}
		// ...Dovrei inserire un controllo errori sulla gestione lettura
	}

	// fmt.Println(multiIns)
	return multiIns
}

// Parser per stringa letta in Mappa valori-occorrenze
func convertiInput(str []string) map[int]int {
	var tmpMap map[int]int

	// Uncomment riga sotto per evitare errore-> panic: assignment to entry in nil map
	// tmpMap = map[int]int{}	// ho usato questa nel esame
	// ...Altrimenti 
	// tmpMap := make(map[int]int)
	// You have to initialize the map using the make function before you can add any elements
	// A nil map is equivalent to an empty map except that no elements may be added. 

	for _, v := range str {
		n, _ := strconv.Atoi(v)

		// pattern comma-ok per verificare presenza chiave, in questo caso è un numero la chiave
		if _, ok := tmpMap[n]; ok {
			tmpMap++
		} else {
			tmpMap[n] = 1
		}
	}

	return tmpMap
}

func IntersecaMultiInsiemi(A, B map[int]int)(AB map[int]int) {
	AB = map[int]int{}

	// In questo caso ho considerato solo le chiavi
	for k1, _ := range A {
		for k2, _ := range B {
			if k1 == k2 {
				AB[k1]=1 // assegno un valore 1 costante per ognuna di esse, non viene specifato 
						 // in che modalità vengono salvati i valori associati alle chiavi...
			}
		}
	}

	return
}

func UnisciMultiInsiemi(A, B map[int][int])(AB map[int]int) {
	AB = map[int]int{}

	// In questo caso valori associati alle chiavi servono per confronto
	for k1, v1 := range A {
		for k2, v2 := range B {
			// Per chiavi uguali salva valore maggiore
			if k1 == k2 {
				if v1 > v2 {
					AB[k1]=v1
				} else {
					AB[k1]=v2
				}
			} else { /* if k1 != k2 non necessaria specifica */
				// Altrimenti: salva entrambe coppie chiave-valore
				AB[k1]=v1
				AB[k2]=v2
			}
		}
	}

	return
}

func StampaMultiInsieme(A map[int]int) {
	if len(A) < 0 { // <- controllo non preciso:  len()<=0
		fmt.Println("VUOTO")
	} else {
		// MANCANZA: 
		// 	Implementazione meccanismo ordinamento per chiave e stampa in ord. non decrescente
		fmt.Println()
	}
}