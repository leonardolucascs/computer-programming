/*
	
	Topic: STRING PADDING IN GO
	https://gosamples.dev/string-padding/

*/

package main

import (
	"fmt"
	"strings"
)

func printTable(table [][]string) {
	// get number of columns from the first table row
	columnLengths := make([]int, len(table[0]))
	for _, line := range table {
		for i, val := range line {
			if len(val) > columnLengths[i] {
				columnLengths[i] = len(val)
			}
		}
	}

	var lineLength int
	for _, c := range columnLengths {
		lineLength += c + 3 // +3 for 3 additional characters before and after each field: "| %s "
	}
	lineLength += 1 // +1 for the last "|" in the line

	for i, line := range table {
		if i == 0 { // table header
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
		for j, val := range line {
			fmt.Printf("| %-*s ", columnLengths[j], val)
			if j == len(line)-1 {
				fmt.Printf("|\n")
			}
		}
		if i == 0 || i == len(table)-1 { // table header or last line
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
	}
}

func main() {
	var table = [][]string{
		{"vegetables", "fruits", "rank"},
		{"potato", "strawberry", "1"},
		{"lettuce", "raspberry", "2"},
		{"carrot", "apple", "3"},
		{"broccoli", "pomegranate", "4"},
	}
	printTable(table)
}
