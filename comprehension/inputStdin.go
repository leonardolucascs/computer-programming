package main

import (
	//"os"
	"fmt"
	"log"
)

func main() {

	var a string

	fmt.Println("value a: ", a)

	fmt.Print("Enter value a: ")
	// Lettura singola stringa
	//nInput, err := fmt.Scan(&a)
	nInput, err := fmt.Scan(a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("value a: ", a)
	fmt.Println("n input: ", nInput)
	fmt.Println("err: ", err)

	/*
		var slA []string

		fmt.Print("Enter a string: ")

		for i:=0; i<len(slA); i++ {
			//fmt.Scan(&slA[i])
			fmt.Scan(slA[i])
		}

		fmt.Println(slA)
	*/

}
