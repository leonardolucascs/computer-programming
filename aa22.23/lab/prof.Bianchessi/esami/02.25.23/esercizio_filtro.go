package main

import(
	"fmt"
	"os"
	"strconv"
)

func main() {
	var counter int

	for i:=1; i<len(os.Args); i+=2 {

		inStr := os.Args[i]	// str: contenuto da ripetere
		tmp, _ := strconv.Atoi(os.Args[i+1])	// tmp: numero ripetizioni

		if tmp > 0 {
			for j:=0; j<tmp; j++{
				fmt.Print(inStr)
			}
			counter += tmp
		}
	}

	fmt.Print(" ")

	if counter > 0 {
		fmt.Println(counter)
	}
}