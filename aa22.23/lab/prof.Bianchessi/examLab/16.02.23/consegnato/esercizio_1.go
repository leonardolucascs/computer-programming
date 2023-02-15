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

	fmt.Println(inStr, len(inStr))
	inStr = inStr[:len(inStr)-1]
	fmt.Println(inStr, len(inStr))

}