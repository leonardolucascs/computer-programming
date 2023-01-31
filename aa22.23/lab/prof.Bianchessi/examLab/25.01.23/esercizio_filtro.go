/*

	Scrivere un programma che legga da riga di comando due numeri a e b e stampi a video la media
	dei numeri compresi tra a e b (a e b inclusi)la cui ultima cifra è dispari (cioè.. la cui
	ultima cifra è 1, 3, 5, 7 oppure 9)

	Si assuma che la coppia di valori specificata a riga di comando sia nel formato corretto.

*/

package main

import "os"
import "fmt"
import "strconv"

func main() {
	var sum, media, nValori int

	input := os.Args[1:]

	inA, _ := strconv.Atoi(string(input[0]))
	inB, _ := strconv.Atoi(string(input[1]))

	for i:=inA; i<=inB; i++ {
		if i % 2 != 0 {
			sum += i
			nValori++
		}
	}

	if nValori == 0 {
		fmt.Println("NaN")
	} else {
		media = sum/nValori

		fmt.Println(media)
	}
}