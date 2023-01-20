package main

import "fmt"
import "unicode/utf8"

func main() {
	
	var s string="s∆ƒa🦍"
	
	fmt.Println(s)
	fmt.Println("len s: ", len(s))
	fmt.Println("Rune in s: ", utf8.RuneCountInString(s))

	for i, c := range s{
		fmt.Printf("i: %d c: %c; len(): %d", i, c, len(string(c)))
		fmt.Println()
	}

}