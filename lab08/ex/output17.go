package main

import "fmt"

func main() {

	//strFrom := "ADEF"
	//strFrom[2] = 'C' /* error: "cannot assign to strFrom[2]" */

	//		Go strings are immutable and behave like read-only byte slices
	//  	una slice di rune ha proprio come dimensione il numeri di caratteri
	//		Unicode presenti nella stringa

	// Vediamo due modi possibili per moificare una stringa:

	/* (1) */
	strFrom1 := "ABÈDEF"
	// len(sliceDiRune) è pari al numero di caratteri Unicode presenti in strFrom1

	sliceDiRune := []rune(strFrom1)
	fmt.Println("ABÈDEF []rune:\t", sliceDiRune)
	sliceDiRune[2] = 'C'
	fmt.Println("\n...dopo modifica\nABCÈDEF []rune:\t", sliceDiRune)
	strTo1 := string(sliceDiRune)
	fmt.Println("\nas string:\t", strTo1)
	/* (1) - end */

	// Questa alternativa funziona solo per stringhe contenenti caratteri ASCII
	// Se il carattere da sostituire occupa una dimensione maggiore a 1 byte questo non funziona
	/* (2) */
	strFrom2 := "ABÈDEF"
	fmt.Println([]rune(strFrom2)) // [65 66 200 68 69 70] come atteso
	//strTo2 := strFrom2[:2] + string('C') + strFrom2[3:]
	strTo2 := strFrom2[:2] + "C" + strFrom2[3:]
	fmt.Println(strTo2)         // output non come atteso: ABC�DEF
	fmt.Println([]rune(strTo2)) // [65 66 67 65533 68 69 70] ??xkè cambia 200(?)
	/* (2) - end */
}
