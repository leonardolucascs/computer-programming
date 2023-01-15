# S0_Shell - Strumenti Go
## 1 Hello World

Usare gli strumenti `go run` e `go build` per eseguire il codice `helloworld.go`.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

## 2 Formattazione automatica

Utilizzate lo strumento `go fmt` per formattare in modo automatico il codice `formattazione.go`.

```go
package main
import "fmt"
func main(){
var x = 10
y := 5
fmt.Println(x + y)
}
```

## 3 Formattazione con errori

Lo strumento `go fmt` non è in grado di formattare il codice quando sono presenti degli errori.
È quindi importante formattare il proprio codice durante la scrittura.

Correggere gli errori del seguente codice, formattarlo ed eseguirlo.

```go
package main

import fmt

func main {

var x = 10
var y = 15

sum = x + y

fmt.Println("La somma è ", sum

}
}
```


## 4 Documentazione

Utilizzare lo strumento `go doc` per visualizzare la documentazione riguardante il package `fmt` e in particolare le funzioni `Print` e `Println`.



