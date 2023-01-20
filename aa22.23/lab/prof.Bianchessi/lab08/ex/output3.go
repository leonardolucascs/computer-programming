package main

import "fmt"

func main(){

	a := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("a - %T: %v\n\n", a, a)

	sl1 := a[:]	// slicing	-> sl1 = copia a
	sl2 := sl1[1:3] // -> sl2 = []int{2,3} [i_start, i_end)

	fmt.Printf("len(sl1) = %d, cap(sl1) = %d\n", len(sl1), cap(sl1))
	fmt.Printf("sl1 - %T: %v\n", sl1, sl1)	// len=7 cap=7
	fmt.Println()

	fmt.Printf("len(sl2) = %d, cap(sl2) = %d\n", len(sl2), cap(sl2))
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)	// len=2 cap=6
	fmt.Println()

	sl2 = sl2[:len(sl2)+1]	// reslicing
	// ricordiamo che sl2 è una slice, e come tale contiene un tipo riferimento
	// a sl1, per la sua dichiarazione, che a sua volta fa riferimento ad: a [7]int 
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)	//len=3 cap=6

	sl2 = sl2[:cap(sl2)]
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)	//len=6 cap=6

	elem, sl1 := sl1[0], sl1[1:]
	// Prima viene valutato sl1[0] che è uguale a 1 ed assegnato ad elem
	// Poi fatta un reslicing e creata una nuova sl1 a partire dal secondo elemento di sl1 stesso
	fmt.Printf("\nelem - %T: %v\n", elem, elem)
	fmt.Printf("sl1 - %T: %v\n", sl1, sl1)


    /* una slice s non può essere modificata per accedere ad elementi
    	dell'array (a cui si riferisce) che precedono quello contenuto in
    	s[0]; l’istruzione s = s[-1:] genera un errore */
}