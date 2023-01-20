// esercizio 4.1

package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(a) }
func m() {
	a := "0"
	print(a)
}

// OUTPUT:
//	G0G

// a in m() ha uno scope locale, viene creata una nuova variabile a
// diversamente da a dichiarata e inizializzata
// con un scope globale o a livello di package
