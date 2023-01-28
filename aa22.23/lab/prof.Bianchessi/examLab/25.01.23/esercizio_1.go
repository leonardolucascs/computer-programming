/*

	
	Scrivere un programma che legga da standard input una stringa 
	di caratteri arbitrari in cui  possono essere presenti caratteri di spaziatura.

	Il programma deve stampare a video il messaggio "la frase è palindroma" nel caso in cui
	la frase sia palindroma e "la frase non è palindroma" altrimenti.

	
*/



package main

import "fmt"
import "os"
import "bufio"
import "unicode"

func main(){
	var ok bool = true

	// Lettura str da std input
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	fmt.Println(input)

	var str string

	// Trasforma la stringa in caratteri lowercase
	for _, v := range input {
		//fmt.Println(string(v))
		str += string(unicode.ToLower(v))
	}

	// str = str[1:]
	fmt.Println("str: \"",str,"\"")

	// indici ricerca
	inizio := 0
	fine := len(str)-1

	for i:=0; i<len(str)/2 && ok; i++ {
		fmt.Println(string(str[i]), unicode.IsLetter(rune(str[i])))

		if unicode.IsLetter(rune(str[inizio])) {
			if unicode.IsLetter(rune(str[fine])) {
				// fmt.Println(string(str[inizio])), " vs ", string(str[fine])
				if str[inizio] != str[fine] {
					ok = false
				} else {
					inizio++
					fine--
				}
			} else {
				fine--
			}
		} else {
				inizio++
			}
	}


	if ok {
		fmt.Println("la frase è palindroma")
	} else {
		fmt.Println("la frase non è palindroma")
	}










}















