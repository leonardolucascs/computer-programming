package main

import (
	//"os"
	"fmt"
	"log"
)

func main() {

	var a string

	fmt.Println("Initial value a: ", a)

	fmt.Print("Enter value a: ")
	// Lettura singola stringa
	//nInput, err := fmt.Scan(&a)
	nInput, err := fmt.Scan(&a)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nFinal value a:", a)
	fmt.Println("n input:", nInput)
	fmt.Println("err: ", err)

	// DIGRESSION ON PRINT

	// Println relies on doPrint(args, true, true), where first argument is
	// addspace and second is addnewline. So Prinln ith multiple arguments will 
	// always print space.
	fmt.Println("a","b")

	// No space added in this case
	fmt.Print("a", "b\n")

	// With Printf output results alike as right as you expected
	fmt.Printf("%s %s\n", "a", "b")


	var slA []string

	fmt.Println("Enter multiple words[5]\nScan scans text read from standard input, storing successive space-separated values\ninto successive arguments.Newlines count as space.\n")
	fmt.Print("Enter a string: ")

	for i:=0; i<len(slA); i++ {
		//fmt.Scan(&slA[i])
		fmt.Scan(&slA[i])
	}

	fmt.Println(slA)

}
