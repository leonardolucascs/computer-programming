/*
	(D1) Definizione. Il massimo comun divisore di due o più numeri, indicato con il simbolo MCD, è 
	il più grande divisore comune dei numeri considerati.

	(D2) Definzione. Il minimo comune multiplo di due o più numeri, indicato con il simbolo mcm, è 
	il più piccolo numero divisibil per ciascuno dei numeri considerati.

	Dati due numeri naturali "a" e "b" non entrambi nulli, si ha che: 

	(1) MCD(a,b) = a 					se b=0
				 = MCD(b, a % b)		altrimenti

	(2) mcm(a, b) = (a * b) / MCD(a, b)


	(3) mcm(a, b, c, ....) = mcm(mcm(mcm(a,b), c),...)


	Scrivere un programma che:

	i) legge da riga di comando una sequenza di numeri interi positivi(maggiori di 0);
	ii) stampi come mostrato nell'esempio di esecuzione:
	 - il valore del mcm tra i numeri che compaiono nella sequenza letta un'unica volta
	 - -1 se non esistono almeno due numeri che compaiono un'unica volta nella sequenza letta

	Si assuma che i valori letti da riga di comando siano nel formato corretto.

	Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:

	- una funzione < contaOccorrenze(numeri []uint) map[uint]uint >, che riceve in input un valore
	[]uint nel parametro numeri e restituisce un valore map[uint]uint in cui, per ogni numero 
	presente in numeri è memorizzato il numero delle sue occorrenze in numeri.

	- una funzione < genera(occorrenze map[uint]uint) []uint > che riceve in input un valore
	map[uint]uint nel parametro occorrenze e restituisce un valore []uint in cui è presente ogni 
	numero associato ad un valore di occorrenze pari a 1 in occorrenze

	- una funzione < MCD(a,b uint) uint > che riceve in input due valori uint nei parametri a e b, e 
	restituisce un valore uint pari al MCD di a e b

	- una funzione < mcm(a,b uint) uint > che riceve in input due valori uint nei parametri a e b, e
	restituisce un valore uint pari al mcm di a e b

	- una funzione mcm_sequenza(numeri []uint)int che riceve in input un valore []uint nel parametro
	numeri, se in numeri sono presenti almento 2 numeri la funzione restituisce un valore 
	int pari al mcm tra tutti i numeri presenti in numeri, -1 altrimenti 

*/

package main

import "fmt"
import "os"
import "strconv"

func main() {

	var numeri []uint

	input := os.Args[1:]

	for _, v := range input {
		tmp, _ := strconv.Atoi(string(v))
		numeri = append(numeri, uint(tmp))
	}

//	fmt.Println(numeri)
//	fmt.Println(contaOccorrenze(numeri)
//	fmt.Println(genera(contaOccorrenze(numeri)))
//	fmt.Println(MCD(14, 250))
//	fmt.Println(mcm(13, 78))

	fmt.Print("mcm = ")

	// calcola mcm
	if len(genera(contaOccorrenze(numeri))) >= 2{
		fmt.Println(mcm_sequenza(genera(contaOccorrenze(numeri))))
	} else {
		fmt.Println(-1)
	}
}


func contaOccorrenze(numeri []uint) map[uint]uint {
	var mapOcc = make(map[uint]uint)

	for _, v := range numeri {
		if _, ok := mapOcc[v]; ok {
			mapOcc[v]++
		} else {
			mapOcc[v] = 1
		}
	}

	return mapOcc
}

func genera(occorrenze map[uint]uint) []uint {
	var output []uint

	for key, value := range occorrenze {
		if value == 1 {
			output = append(output, key)
		}
	}

	return output
}


func MCD(a, b uint) uint {
	if b == 0 {
		return a
	} else {
		return MCD(b, a % b)
	}
}

func mcm(a, b uint) uint {
	return (a * b) / MCD(a, b)
}


func mcm_sequenza(numeri []uint) int {

	/*
	// non serve alcun tipo di ciclo !!!
	// Le "iterazioni cicliche" vengono "svolte" dalle chiamate ricorsive
	
	var mcm int

	for i:=0; i<len(numeri)-1; i++{
		// caso base mcm(a, b)
		if len(numeri) == 2 {
			return int(mcm(numeri[i], numeri[i+1]))	// ERROR: Cannot call non-function mcm => value type int
													// shadowing variable name & function name!!!
		} else {
			mcm = mcm_sequenza(numeri[i:])
		}
	}
	*/


//	fmt.Println("\nstato iniziale: ", numeri)

	if len(numeri) < 2 {
		return -1	// necessario se non entra nel for per len(numeri)<2
	} else {
		// caso base mcm(n1, n2)
		if len(numeri) == 2 {
//			fmt.Printf("...attivazione caso base: mcm(%d, %d)\n\n", numeri[0], numeri[1])
			
			output := mcm(numeri[0], numeri[1])

//			fmt.Println("mcm(", numeri[0], ", ", numeri[1] ,") => ", output)

			return int(output)	// int( mcm(numeri[0], numeri[1]) )
		} else {	// passo induttivo:
//			fmt.Println("...chiamata funzione ricorsiva: mcm(", numeri[:len(numeri)-1], ")")

			tmp := uint(mcm_sequenza( numeri[:len(numeri)-1] ) )
			op := mcm( numeri[len(numeri)-1], tmp )
			
//			fmt.Println("mcm(", numeri[len(numeri)-1], ", mcm(", numeri[:len(numeri)-1] ,") )")
//			fmt.Println("mcm(", numeri[len(numeri)-1],", ", tmp ,") =>", op)
			
			return int(op)	// int( mcm( numeri[len(numeri)-1], uint(mcm_sequenza( numeri[:len(numeri)-1] ) ) ) )
		}
	}
}