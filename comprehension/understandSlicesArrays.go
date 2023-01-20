package main

import (
	"fmt"
)

func main() {
	x := [3]string{"Hola", "Hello", "Halo"} // "x" is an array

	y := x[:] // slice "y" points to the underlying array "x"

	w := x // "w" is an array copy of "x", it won't change "x"

	z := make([]string, len(x)) // "z" is a slice

	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "Adios" // the value at index 1 is now "Adios" for both "y" and "x"

	w[2] = "Bye" // this won't change nothing to "x"

	x[0] = "Aloa" // the value at index 0 is now "Aloa" even for "y" cause points to "x"

	fmt.Printf("X:%T %v\n", x, x)
	fmt.Printf("Y:%T %v\n", y, y)
	fmt.Printf("W:%T %v\n", w, w)
	fmt.Printf("Z:%T %v\n", z, z)
}
