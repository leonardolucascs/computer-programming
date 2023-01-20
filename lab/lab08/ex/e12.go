// esercizio12 - Filtra e moltiplica
package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	res := a * b
	fmt.Printf("Il risultato della moltiplicazione tra i numeri interi è %d\n", res)
}