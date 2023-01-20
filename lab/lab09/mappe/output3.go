package main

import (
	"fmt"
)

var ( 
	nominativi = map[string]string { "010203445": "Mario Rossi", 
	"0202302343": "Carlo Bianchi", "030230423": "Matteo Giallo", 
	"0539353353": "Carlo Cracco", "23884433": "Andrea Verde", 
	"8374299244":"Jin Min", "998834422": "Joe Rae"} 
	)

func main() {
	numeriTelefonici := make(map[string]string)

	for k, v := range nominativi {
		numeriTelefonici[v] = k
	}

	for k, v := range numeriTelefonici {
		fmt.Printf("Nominativo: %v\nNumero telefonico: %v\n\n", k, v)
	}
}