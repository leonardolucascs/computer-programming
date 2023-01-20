package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	slIn := make([]int, len(os.Args)-1)


	// Fill slIn
	for i:=1; i < len(os.Args); i++ {
		slIn[i-1] , _ = strconv.Atoi(os.Args[i])
	}

	min := MinimoDispari(slIn)
	//fmt.Println(min)
	//fmt.Println(slIn)

	// Print output
	for i:=0; i < len(slIn); i++ {
		if (slIn[i] % 2 == 0 && slIn[i] > min) {
			fmt.Print(slIn[i], " ")
		}
	}
	fmt.Println()

}


func MinimoDispari(sl []int) int {
	var min int
	var firstOdd bool = true

	for i:=0; i < len(sl); i++ {
		if (sl[i] % 2 != 0) {
			if firstOdd {
				min = sl[i]
				firstOdd = false
			}

			temp := sl[i]
			
			if temp < min {
				min = temp
			}
		}




	}

	return min
}