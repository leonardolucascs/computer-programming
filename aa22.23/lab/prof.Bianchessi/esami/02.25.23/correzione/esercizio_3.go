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
	
	tragitto := []Punto{{"A", 10, 11.3},{"B", 11, -5.2},{"C", 5, 3.4}}
	
	//fmt.Println(NuovoTragitto())
	//fmt.Println("\ndx:", Distanza(punti[0], punti[1]))

	fmt.Println("Lunghezza tragitto: ", Lunghezza(tragitto))

	//fmt.Println(StringPunto(punti[0]))

	//StampaTragitto(punti)

	// Per redirezione input da file oppure dall'utente, uncomment riga sotto
	//StampaTragitto(NuovoTragitto())
}

// Leggi input e restituisci tragitto
func NuovoTragitto() (tragitto []Punto) {

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

// Parser per stringa letta in "oggetto" di tipo Punto
func convertiInput(str []string) Punto {
	
	var tmp Punto

	tmp.etichetta = str[0]
	tmp.ascissa, _ = strconv.ParseFloat(str[1], 64)
	tmp.ordinata, _ = strconv.ParseFloat(str[2], 64)

	return tmp
}

// Calcola distanza tra 2 punti
func Distanza(p1, p2 Punto) float64 {
	tmp1 := p1.ascissa - p2.ascissa
	//fmt.Println("Delta ascisse:", p1.ascissa, p2.ascissa, tmp1)

	tmp2 := p1.ordinata - p2.ordinata
	//fmt.Println("Delta ordinate:", p1.ordinata, p2.ordinata, tmp2)

	x := math.Pow(math.Pow(tmp1, 2) + math.Pow(tmp2, 2), 0.5)
	
	arrotondaN(&x)	// Utilizzo pattern indirizzo-puntatore, al posto che ricevere un output
	//fmt.Println(x)
	
	return x
}

// Calcola lunghezza tragitto
func Lunghezza(tragitto []Punto) float64 {
	var lenOutput float64

	for i:=0; i<len(tragitto)-1; i++{
		lenOutput += Distanza(tragitto[i], tragitto[i+1])
	}

	arrotondaN(&lenOutput)

	return lenOutput
}

// Stampa punto come stringa
func StringPunto(p Punto) string {
	// Use the fmt.Sprintf method to format a floating-point number as a string.
	// "%.1f" : specify default width for decimal point and with precision 2
	//return p.etichetta + " = (" + fmt.Sprintf("%.1f", p.ascissa) + " , " + fmt.Sprintf("%.1f", p.ordinata) + ")"
	return fmt.Sprintf("%s = (%.1f , %.1f)", p.etichetta, p.ascissa, p.ordinata)
}

// Stampa tragitto
func StampaTragitto(tragitto []Punto) {

	fmt.Println("Tragitto:")
	for _, punto := range tragitto {
		fmt.Println(StringPunto(punto))
	}
}