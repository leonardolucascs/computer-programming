/*
	
	ESERCIZIO1: Codifica messaggio segreto
	Data da riga di comango una stringa

	Codificare una fuzione che restituisca il messaggio nascosto tra le parentesi tonde.

	Vincoli:
		i.	Si considerino sottostringhe del messaggio, tutte le stringhe contenute dalle parentesi '(' e ')'.
		ii.	Sono validi solo i messaggi che iniziano racchiusi tra '(' e terminano con ')', laddove ci siano parentesi
			all'interno del messaggio questo è da considerare non valido.
		iii.	Qualora ci fossero "()", ossia stringhe vuote, si considera lo spazio " " come sottostringa
		iv.	[ ...non implementato] Non sono e non devono essere presenti spazi nella stringa di input, 
			la singola runa deve essere passata alla funzione del package unicode isSpace(runa) bool 


		input 	=>	"fd(((sw(we))()(o)(s)(ciao)"
		output	=>	"we o s ciao"

		go run esercizio1.go "()ssas(sd(:-D)sdwqw(t)(r)(o)(v)((a)))rer((to)" 
		Expected output: " :-D t r o v a to"

		go run es.go "(ciao)èfgv(()ffg((tsa(Sara)df-()fsf(:-D)"
		Expected output: "ciao  Sara  :-D"

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var strOut, tmp string
	var isOpen bool = false
	input := os.Args[1]	// ?! string(os.Args[1]) è già una stringa os.Args
	

	for inizio:=0; inizio < len(input); inizio++ {
		if input[inizio] == '(' {
			tmp = ""
			isOpen = true
			
			for fine := range input[inizio+1:] {
				if isOpen {
					if input[inizio+1+fine] != '(' && input[inizio+1+fine] != ')' {
						tmp += string(input[inizio+1+fine])
					} else if input[inizio+1+fine]==')' {
						tmp += " "
						isOpen = false
					} else if input[inizio+1+fine] == '(' {
						tmp = ""
						isOpen = false
					}
				}
			}
			inizio += len(tmp)
			strOut += tmp
		}
	}

	fmt.Println(strOut)
}







