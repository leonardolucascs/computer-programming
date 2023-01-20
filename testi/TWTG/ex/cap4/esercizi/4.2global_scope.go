// esercizio 4.2

package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(a) }
func m() {
	a = "0"
	print(a)
}

// OUTPUT
// G00

// diversamente da local_scope.go, in m() ora viene assegnata alla
// variabile a che ha uno scope globale un nuovo valore, e non viene di
// fatto creata un'altra varibaile omnonima in m() come succedeva in
// precedenza
