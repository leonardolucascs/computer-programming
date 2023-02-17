package main

import "fmt"

func main() {
	s1 := make([]int, 4, 8)
	// s1 == []int{0, 0, 0, 0}, len=4, cap=8

	for i := 0; i < len(s1); i++ {
		s1[i] = i + 1
	}
	// s1 == {1, 2, 3, 4}

	fmt.Println("1)\ns1:")
	stampa(s1)

	s2 := append(s1[2:], []int{10, 20}...)
	// se si vuole concatenare una slice y a una slice x,
	// i ... servono per espandere il secondo argomento ad un elenco di argomenti
	// x = append(x, y...)
	// s2 == {3, 4, 10, 20}, len=4 cap=6?? perché

	//  CON QUALE CRITERIO CAMBIA LA CAPACITA' NUOVA PER UNA SLICE??
	fmt.Println("\n2)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2[0] = -1
	s2 = append(s2, 100)
	// s2 == {-1, 4, 10, 20, 100}, len=5, cap=
	fmt.Println("\n3)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2 = append(s2, 1000, 2000)
	fmt.Println("\n4)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2 = append(s2[:2], s2[len(s2)-2:]...)
	fmt.Println("\n5)\ns2:")
	stampa(s2)
}

func stampa(s []int) {
	var lunghezza int

	fmt.Println("a) ", s, len(s), cap(s))
	lunghezza = len(s)

	s = s[:cap(s)] // reslicing di s fino alla sua capacità
	fmt.Println("b) ", s, len(s), cap(s))
	s = s[:lunghezza] // reslicing, restore s allo stato "originale"
}
