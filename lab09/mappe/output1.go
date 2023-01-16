package main

import "fmt"

func main() {

	var m map[string]int

	if m == nil { 
		fmt.Println("'nil' è lo zero-value di una variabile di tipo " +
		"'reference type' non inizializzata") 
	} 

	//m["mamma"] = 5 /*  panic: assignment to entry in nil map */
	
	m = make(map[string]int)

	for _, s := range []string{"questo", "é", "un", "test"} {
		m[s] = len([]rune(s))
		//m[s] = len(s)	// ? xké non così
	}

	// Print map for-range cycle: key -> value
	for k, v := range m {
		fmt.Println(k, "->", v)
	}

	// Another print map for-range cycle

}