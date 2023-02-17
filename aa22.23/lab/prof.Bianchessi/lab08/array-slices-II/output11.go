package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("TIPO os.Args: %T\n", os.Args)

	for i, v := range os.Args{
		fmt.Printf("os.Args[%d]:\n\tTIPO = %T - VALORE = %s\n", i, v, v)
	}
}