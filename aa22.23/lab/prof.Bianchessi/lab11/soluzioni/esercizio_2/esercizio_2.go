package main

import (
	"fmt"
	"os"
)

func main() {

	sottosequenze := Sottosequenze(os.Args[1])

	StampaSottosequenze(sottosequenze, len(os.Args[1]))
}

func Sottosequenze(sequenza string) (sottosequenze map[string]int) {
	sottosequenze = make(map[string]int)
	for i, c1 := range sequenza {
		for j, c2 := range sequenza[i+1:] {
			if c1 == c2 {
				sottosequenze[sequenza[i:i+j+2]]++
			}
		}
	}
	return
}

func StampaSottosequenze(sottosequenze map[string]int, lunghezzaMassima int) {

	for lunghezza := lunghezzaMassima; lunghezza >= 3; lunghezza-- {

		for sequenza, occorrenze := range sottosequenze {
			if len(sequenza) == lunghezza {
				fmt.Printf("%s -> Occorrenze: %d\n", sequenza, occorrenze)
			}
		}

	}
}
