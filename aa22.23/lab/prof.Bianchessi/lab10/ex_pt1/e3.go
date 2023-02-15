package main

import "fmt"

type Indirizzo struct {
	via, cap, città string
	numeroCivico uint
}

type Contatto struct {
	cognome, nome, telefono string
	indirizzo Indirizzo
}

type Rubrica []Contatto

func main() {

	/*

	var indirizzo Indirizzo = Indirizzo{"via MonteRosa", "0010", "Milano", 10}
	fmt.Println(indirizzo)

	var contatto Contatto = Contatto{"Sanchez", "Joe", "000331144", indirizzo}
	fmt.Println(contatto)

	var rubrica Rubrica = Rubrica{contatto, contatto}
	fmt.Println(rubrica)
	
	*/
	rubrica := NuovaRubrica()

	rubrica = InserisciContatto(rubrica, "Rossi", "Mario", "Roma", 12, "0010", "Milano", "003302966")

	fmt.Println(rubrica)
}

func NuovaRubrica() Rubrica {
	return Rubrica{}
}

func InserisciContatto(r Rubrica, cognome, nome, via string, numero uint, cap, città string, telefono string) Rubrica {

	// Controllo rubrica: cerca persone con stesso nome e cogonome
	for _, contatto := range r {
		if contatto.nome == nome && contatto.cognome == cognome{
			return r
		}
	}

	return append(r, Contatto{cognome, nome, telefono, Indirizzo{via, cap, città, numero}})
}