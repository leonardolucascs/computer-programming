package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//import "strconv"
//import "strings"
import "math"

const EPSILON = 1.0e-6

func main() {

	punti := NuovoPercorso()
	lunghezzaPercorso := Lunghezza(punti)
	fmt.Printf("Lunghezza percorso: %.3f\n", lunghezzaPercorso)

	for i := range punti {
		if ÈXMaggioreDiY(Lunghezza(punti[:i+1]), lunghezzaPercorso/2) {
			fmt.Printf("Punto oltre metà: %s\n", String(punti[i]))
			break
		}
	}
}

type Punto struct {
	Nome string
	X, Y float64
}

func NuovoPercorso() (punti []Punto) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			var p Punto
			comando := strings.Split(scanner.Text(), ";")
			p.Nome = comando[0]
			p.X, _ = strconv.ParseFloat(comando[1], 64)
			p.Y, _ = strconv.ParseFloat(comando[2], 64)

			punti = append(punti, p)
		}
	}
	return
}

func Distanza(p1, p2 Punto) float64 {
	dX := p1.X - p2.X
	dY := p1.Y - p2.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func Lunghezza(punti []Punto) (lunghezza float64) {
	precedente := punti[0]
	for _, p := range punti[1:] {
		lunghezza += Distanza(precedente, p)
		precedente = p
	}
	return
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%.1f, %.1f)", p.Nome, p.X, p.Y)
}

func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}
