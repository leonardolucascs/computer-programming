// esercizio 4.3

package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
	//print(a)
}

func f1() {
	a := "0"
	print(a)
	f2()
	//print(a) //priorità e visibilità alla variabile con scope locale
}

func f2() {
	print(a)
	//a = "F"
}

// OUTPUT
// G0G

/*
	In questo esempio a in f1() viene dichiarata ed inizializzata come
	una nuova variabile visibile solo a livello locale della funzione
	mentre f2() prenderà sì come riferimento a dichiarata a livello globale,
	ma inizializzata nella funzione main(), nel momento in cui viene
	chiamata la funzione f2() in f1(), non conosce alcuna variabile a
	se non quella a livello di package

	Assegnando un altro valore alla variabile a prima di uscire da f2(),
	tale valore sarà visibile e noto di nuovo a livello globale.

	Nota bene che identificatori omonimi dichiarati nuovamente in blocchi
	di codice differenti (solo al suo interno) assumono la priorità
	(effetto SHADOWING) e addombrano le variabili esterne
*/
