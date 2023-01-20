package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "asSASA ddd dsjkdsjs dk"
	s1 := "asSASA ddd dsjkdsjs∂ßµ dk"

	fmt.Printf("numero byte s: %d\n", len(s))
	fmt.Printf("numero byte s1: %d", len(s1))

	fmt.Printf("\n\n\nnumero rune s: %d\n", utf8.RuneCount([]byte(s)))
	fmt.Printf("numero rune s1: %d\n", utf8.RuneCount([]byte(s1)))
}
