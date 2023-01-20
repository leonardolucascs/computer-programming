package main

import ( 
      "fmt"
      "strconv"
)

func main() {
var numeroIntero int = 7
      fmt.Printf("Tipo: %T\n", numeroIntero)
      fmt.Printf("- formato default: %v\n", numeroIntero)
      fmt.Printf("- formato base 10: %d\n", numeroIntero)
      fmt.Printf("- formato base 10 (larghezza = 6): %6d\n", numeroIntero)
      fmt.Printf("- formato base 2: %b\n", numeroIntero)
      fmt.Println()

var numeroReale float64 = 14.5674
      fmt.Printf("Tipo: %T\n", numeroReale)
      fmt.Printf("- formato default: %v\n", numeroReale)
      fmt.Printf("- formato con decimali: %f\n", numeroReale)
      fmt.Printf("- formato con numero fissato di decimali: %.2f\n", numeroReale)
      fmt.Printf("- formato esponenziale: %e\n", numeroReale)
      fmt.Println()

var valoreLogico bool = false
      fmt.Printf("Tipo: %T\n", valoreLogico)
      fmt.Printf("- formato default: %v\n", valoreLogico)
      fmt.Printf("- formato true/false: %t\n", valoreLogico)
      fmt.Println()

var carattere rune = 'A'
      fmt.Printf("Tipo: %T\n", carattere)
      fmt.Printf("- formato default: %v\n", carattere)
      fmt.Printf("- formato carattere: %c\n", carattere)
      fmt.Printf("- formato intero: %d\n", carattere)
      fmt.Printf("- formato unicode: %U\n", carattere)
      fmt.Println()

var stringaTesto string = "Hello\tworld!" 
      fmt.Printf("Tipo: %T\n", stringaTesto) 
      fmt.Printf("- formato default: %v\n", stringaTesto)
      fmt.Printf("- formato stringa: %s\n", stringaTesto)
      fmt.Println()

var posizioniDecimali = 2
      fmt.Printf("Printf con formato creato dinamicamente %." + strconv.Itoa(posizioniDecimali) + "f\n", 10.458)
}
