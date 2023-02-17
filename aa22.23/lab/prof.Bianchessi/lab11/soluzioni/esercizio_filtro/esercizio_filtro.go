package main

import "fmt"
import "os"
import "strconv"
import "strings"

func main() {

    for i := 1; i < len(os.Args); i += 2 {
        ch := os.Args[i]
        n, _ := strconv.Atoi(os.Args[i+1])
        fmt.Print(strings.Repeat(ch, n))
    }

    fmt.Println()

}
