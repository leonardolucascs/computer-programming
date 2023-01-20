package main

import "fmt"

func main() {
	a, b := 10, 5

	sum, mul, sub := getResult_(a, b)

	fmt.Printf("The sum beetween %d and %d is equal to %d\n", a, b, sum)
	fmt.Printf("The multiplication beetween %d and %d is equal to %d\n", a, b, mul)
	fmt.Printf("The substraction beetween %d and %d is equal to %d\n", a, b, sub)
}

// funzione con variabili di rientro senza nome
func getResult(x1, x2 int) (int, int, int) {
	return x1 + x2, x1 * x2, x1 - x2
}

// funzione con variabili di rientro senza nome
func getResult_(x1, x2 int) (sum int, mul int, sub int) {
	sum = x1 + x2
	mul = x1 * x2
	sub = x1 - x2

	return
}
