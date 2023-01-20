package main

import "fmt"

func main() {

	d := []byte{'r', 'o', 'a', 'd'}

	// Con questa dichiarazione eventuali modifiche
	// su e[] andavano a modificare anche d[]
	// e := d[2:]
	// e == []byte{'a', 'd'}

	// Dichiarazione slice e
	e := make([]byte, len(d))
	// Questa operazione mi permette di lavorare su e[]byte
	// senza modificare d[]byte

	// Copia contenuti di slice d a slice e
	// copy(dst, src []Type) int
	// return number elements copied
	copy(e, d)
	// e == []byte{'r','o','a', 'd'}

	e[3] = 'm'

	// e == []byte{'r','o','a', 'm'}
	// d == []byte{'r', 'o', 'a', 'd'}

	fmt.Printf("d: %c\n", d)
	fmt.Printf("e: %c", e)
}
