package main

import (
	"fmt"
	"strings"
)

func main() {
	var numero string
	var numeroInizializzato bool

	var operazione int

	var errore bool
	for {
	
		operazione = sceltaMenù()
		
		switch operazione {
			case 1:
				var n string
				n, errore = LeggiNumero()
				if errore { // comma, ok pattern
					fmt.Println("La lettura del numero non è andata a buon fine.")
					break
				}
				numero = n
				numeroInizializzato = true
				fmt.Println("La lettura del numero è andata a buon fine.")
			case 2:
				if numeroInizializzato {
					fmt.Print("Numero corrente:",numero,"\n")
				} else {
					fmt.Println("Il numero corrente non è stato ancora inizializzato.")
				}
			case 3:
			if numeroInizializzato {
				fmt.Print("Numero corrente: ")
				fmt.Println(numero)
				fmt.Print("Introduci il valore della base in cui convertire il numero: ")  
				// Si assuma che il valore della base in cui convertire il numero sia inserito correttamente
				var b int
				fmt.Scan(&b)
				// ... si effettui la conversione ...
				fmt.Print("Numero corrente convertito: ")
				fmt.Println(numero)
			} else {
				fmt.Println("Il numero corrente non è stato ancora inizializzato.")
			}			
			case 4:
				fmt.Println("Terminazione.")
		}
		
		if operazione == 4 {
			break
			//return
		}
		
	}
	
}

func LeggiNumero() (n string, errore bool) {

	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)
	errore = false
	
	// controlli sintattici
		
	if !strings.Contains(n,"BASE:"){
		fmt.Println("Formato non corretto.\nLa stringa deve contenere il token 'BASE:'.")
		errore = true
		//return 
	}

	if !strings.Contains(n,"_"){
		if (!errore) {fmt.Println("Formato non corretto.") }
		fmt.Println("La stringa deve contenere il token '_'.")
		errore = true
		//return 
	}
	
	if !strings.Contains(n,"VALORE:"){
		if (!errore) {fmt.Println("Formato non corretto.") }
		fmt.Println("La stringa deve contenere il token 'VALORE:'.")
		errore = true
		//return 
	}
	
	// ...
	

	if len(n) < len("BASE:_VALORE:") + 2 {
		if (!errore) {fmt.Println("Formato non corretto.") }
		fmt.Println("Nella string devono essere specificati i valori di BASE e VALORE.")
		errore = true
		//return 
	}
	
	if (errore) { 
		return	// controllo sintattico non superato
	}
	
	
	// controllo semantico sui valori di BASE e VALORE
	
	s_BASE := estraiBASE(n)
		
	for i:=0; i<len(s_BASE); i++ {
		if s_BASE[i] < '0' || s_BASE[i] > '9' {
			fmt.Println("Formato non corretto.")
			fmt.Println("Il valore della BASE deve essere specificato da cifre decimali.")
			errore = true			
			break
		}
	}
	
	var i_BASE int
	
//	i_BASE = ConvertiAlfanumericoInB_b_AInteroInB_10(s_BASE)

	if i_BASE[i] < 2 || i_BASE[i] > 36 {
		fmt.Println("Il valore della BASE deve essere compreso tra 2 e 36, 2 e 36 inclusi.")
		errore = true			
	}
		
	if (errore) { 
		return	// senza un valore corretto per la BASE non è possibile effettuare il controllo sul valore di VALORE
	}
	
	s_VALORE := estraiVALORE(n)
	
	if i_BASE <= 10 {	
		for i:=0; i<len(s_VALORE); i++ {
			if (s_VALORE[i] < '0' || s_VALORE[i] > '9') &&
			   (s_VALORE[i] - '0' >= i_BASE)
				if (!errore) {fmt.Println("Formato non corretto.") }
				fmt.Println("Il valore del VALORE deve essere specificato da cifre decimali minori di BASE.")
				errore = true			
				break				
			}
		}
	}
	else {
		// ...
	}
	
	//return n, errore
	return
}

/* 

Per il numero correntemente in memoria, se inizializzato, le funzioni:

func estraiBASE(n string) string { ... }
func estraiVALORE(n string) string { ... }
func ConvertiAlfanumericoInB_b_AInteroInB_10(n string, b int) (v int) { ... }

non dovrebbero mai generare errore. 

*/

func estraiBASE(n string) string {
	return n[5:strings.Index(n,"_")]
}

func estraiVALORE(n string) string {
	return n[strings.Index(n,"VALORE:")+7:]
}

func ConvertiAlfanumericoInB_b_AInteroInB_10(n string, b int) (v int) {
	// ...
	
	return v
}

func sceltaMenù() (op int) {
	fmt.Println()
	fmt.Println("1. Inserisci un numero nel formato BASE:*_VALORE:*.")
	fmt.Println("2. Stampa a video il numero corrente.")
	fmt.Println("3. Converti il numero corrente in un'altra base.")
	fmt.Println("4. Termina.")
	fmt.Print("> ")

	fmt.Scan(&op)

	return
}
