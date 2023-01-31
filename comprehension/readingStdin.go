
/*
	(1) Go read input with NewReader
	7
	The bufio package implements buffered I/O. Buffered I/O has much better performance than non-buffered. The package wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

*/
package main

import (
     "os"
     "bufio"
     "fmt"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')
	// The ReadString reads until the first occurrence of the specified delimiter 
	// (new-line in our case) in the input, returning a string containing the data up to and
	// including the delimiter.
	
	fmt.Printf("Hello %s\n", name)
}


/*
	
	(2) Go read input with Scanf

	The Scanf function scans text read from standard input, storing successive space-separated 
	values into successive arguments as determined by the format. It returns the number of items 
	successfully scanned.

*/

func main() {

    var name string

    fmt.Print("Enter your name: ")

    fmt.Scanf("%s", &name)
    
    // The entered value is stored into the name variable.

    fmt.Println("Hello", name)
}


/*
	(3) Go read input with NewReader

	The bufio package implements buffered I/O. Buffered I/O has much better performance than non-buffered. The package wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

*/
	

func main() {

    names := make([]string, 0)

    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Print("Enter name: ")
        
        scanner.Scan()
        
        text := scanner.Text()

        if len(text) != 0 {

            fmt.Println(text)
            names = append(names, text)
        } else {
            break
        }
    }

    fmt.Println(names)
}