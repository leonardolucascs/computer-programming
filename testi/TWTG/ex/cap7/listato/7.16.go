package main

import "fmt"

func main() {
	s := "\u00ff\u754c"
	d := "stringa casuale"

	for i, c := range s {
		fmt.Printf("%d: %c ", i, c)
	}

	fmt.Printf("\n%d\n", len(s))
	fmt.Print(len(d))
}

// \u754c => 界 occupa 3byte
// \u00ff => ÿ occupa 1byte
