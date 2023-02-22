package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

// Definizione tipo Punto
type Punto struct{
	etichetta string
	ascissa float64
	ordinata float64
}

func main() {
	fmt.Println(NuovoTragitto())
	fmt.Println()

	// ATTENZIONE: FUNZIONI CREATE CORRETTE E OPERANTI MA MAIN() INCOMPLETO
	// 				NON SI RIESCE A VERIFICARE L'EFFETTIVO FUNZIONAMENTO DELLE FUNZIONI
	//				NON SEGUE L'ESECUZIONE DATA DALLA TRACCIA D'ESAME


	//punti := []Punto{{"A", 10, 11.3}, {"B", 11, -5.2}, {"C", 5, 3.4}}
	//fmt.Println("\ndx:", Distanza(punti[0], punti[1]))
}

func NuovoTragitto()(tragitto []Punto){

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() && len(scanner.Text())>0 {
			// Scansione, Salvataggio e Split valori letti separati da ';' in variabile []string
			input := strings.Split(scanner.Text(),";")

			// Salvataggio lettura in tragitto
			tragitto =  append(tragitto, convertiInput(input))
		} else {
			break
		}
	}

	//fmt.Println(tragitto)
	return tragitto
}

// Parser per stringa letta in variabile di tipo Punto
func convertiInput(str []string) Punto{
	
	var tmp Punto

	tmp.etichetta = str[0]
	tmp.ascissa, _ = strconv.ParseFloat(str[1], 64)
	tmp.ordinata, _ = strconv.ParseFloat(str[2], 64)

	return tmp
}

// Calcola la distanza tra 2 punti
func Distanza(p1, p2 Punto) float64 {
	tmp1 := p1.ascissa - p2.ascissa
	fmt.Println("Delta ascisse:", p1.ascissa, p2.ascissa, tmp1)

	tmp2 := p1.ordinata - p2.ordinata
	fmt.Println("Delta ordinate:", p1.ordinata, p2.ordinata, tmp2)

	return math.Pow(math.Pow(tmp1, 2) + math.Pow(tmp2, 2), 0.5)
}

// Calcola lunghezza tragitto
func Lunghezza(tragitto []Punto) float64{
	var lenOutput float64

	// Expected error:impostazione ciclo for errata, [limite e contatore]
	for i:=0; i<len(tragitto)-1; i=i+2 {
		lenOutput += Distanza(tragitto[i], tragitto[i+1])
	}

	return lenOutput
}

/*	Spiegazione: impostazione errata for

		somma += Distanza(tragitto[i], tragitto[i+1])

		Per n pari
		-1 10 4 5 -> linearmente 9+4+5 13+5 18
					 a coppie	 9 + 9 18

		Per n dispari, non funziona più
		-1 10 4 5 6  -> 
						linearmente 18 + 6 24
						a coppie 	9 + 9 + (6+ @index out of range)

		In più, in questo caso il mio ciclo for si ferma a len(tragitto)-1, dunque l'ultimo punto
		non viene nemmeno considerato!
*/




// MANCANZA: funzione StringPunto(p Punto) string
//			 funzione StampaTragitto(tragitto []Punto)