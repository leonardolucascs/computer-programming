package main

import (
	"fmt"
	"unicode"
	"bufio"
	"os"
)

func palindroma(str string) bool {
	// caso base
	if len(str)<=1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}

	// passo ricorsivo
	return palindroma(str[1:len(str)-1])
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	inStr, _ := reader.ReadString('\n')

	// Nella lettura ReadString() salva il delimitatore dentro la variabile
	// fmt.Println(inStr, len(inStr))
	// Per questo occorre effettuare un reslicing
	inStr = inStr[:len(inStr)-1]
	
	//fmt.Println(inStr, len(inStr))

	var tmp string

	for _, v := range inStr {
		if unicode.IsSpace(v) {
			fmt.Println("Presenza di spazi nella riga letta...")
			return
		} else {
			tmp += string(unicode.ToLower(v))
		}
	}

	//fmt.Println(tmp)
	fmt.Println("Sottostringhe:")

	var sottoStr string

	maxStr := ""

	// Attenzione: questo tipo di controllo tramite for indici 
	// 				vale però solo per caratteri ASCII
	
	// (*) Svolgere tentativo implementazione anche tramite for range
	// (*)	non riuscito durante esame

	for i:=0; i<len(tmp); i++{
		for j:=i+1; j<len(tmp); j++{
			// Laddove trovi una sottostringa, rappresentata da caratteri uguali
			if tmp[i] == tmp[j] { 
				sottoStr = tmp[i:j+1]	// l'ultimo carattere non deve essere escluso dallo slicing
										
										// Attenzione: stringa tmp convertita in ToLower, l'ouput
										// 				expected deve essere l'originale
				if len(sottoStr)>=2 && palindroma(sottoStr) {
					fmt.Println(i, j, sottoStr)
					if len(sottoStr) > len(maxStr) {
						maxStr = sottoStr
					}
				}
			}
		}
	}

	fmt.Println("\nSottostringa di lunghezza massima:\n", maxStr, len(maxStr))
}