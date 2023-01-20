package main

import "fmt"

func main() {
	//var arrAge = [5]int{18, 20, 15, 22, 16}
	//var arrAge2 = []int{18, 20, 15, 22, 16} //slice

	var arrLaze = [...]int{5, 6, 7, 8, 22} //stessa cosa sotto, no?!
	//arrLaze := [...]int{5, 6, 7, 8, 22} //inizializzazione + dichiarazione
	//var arrLaze2 = [...]int{5, 6, 7, 8, 22} //slice

	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	//var arrKeyValue2 = []string{3: "Chris", 4: "Ron"} //slice

	fmt.Printf("Primo elemento di arrLaz: %d\n", arrLaze[0])
	fmt.Println("Ultimo elemento di arrLaze: ", arrLaze[len(arrLaze)-1])

	a := [...]string{"a", "b", "c", "d"}
	for i := range a { //for i, val := range a {
		fmt.Println("Array item", i, "is ", a[i])
		//fmt.Println("Array item", i, "is ", val)
	}
}
