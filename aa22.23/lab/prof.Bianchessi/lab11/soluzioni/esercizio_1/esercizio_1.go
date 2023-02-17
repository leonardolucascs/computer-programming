package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	for posizione, parola := range os.Args[1:] {
		fmt.Print(TrasformaParola(parola, posizione), " ")
	}
	fmt.Println()
}

func TrasformaParola(parola string, posizione int) (risultato string) {
	partenza := posizione % 2
	for i, c := range parola {
		if (partenza+i)%2 == 0 {
			risultato += string(unicode.ToUpper(c))
		} else {
			risultato += string(unicode.ToLower(c))
		}
	}
	return
}
