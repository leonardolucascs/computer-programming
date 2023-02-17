# Esercizio 2

Scrivere un programma che:
* legga da **riga di comando** una stringa `s` definita da caratteri
  appartenenti all'alfabeto inglese (e quindi codificati all'interno dello
  standard US-ASCII (integrato nello standard Unicode));
* stampi a video tutte le sottostringhe di caratteri presenti in `s`
  che:
  1. iniziano e finiscono con lo stesso carattere;
  2. sono formate da almeno 3 caratteri.

Ciascuna sottostringa deve essere stampata un'unica volta, riportando il
relativo numero di occorrenze della sottostringa in `s` (cfr.
**Esecuzione d'esecuzione**).

Le sottostringhe devono essere stampate in ordine di lunghezza (dalla
più lunga alla più corta).

Se non esistono sottostringhe che soddisfano le condizioni 1 e 2, il
programma non deve stampare nulla.

Si noti che una sottostringa può essere contenuta in un'altra
sottostringa più lunga.

Si assuma che la stringa specificata a riga di comando sia nel formato
corretto e includa almeno 3 caratteri. 

##### Esempio d'esecuzione:

```text
$ go run esercizio_2.go abbabba
abbabba -> Occorrenze: 1
bbabb -> Occorrenze: 1
babb -> Occorrenze: 1
abba -> Occorrenze: 2
bbab -> Occorrenze: 1
bab -> Occorrenze: 1

$ go run esercizio_2.go abcacba
abcacba -> Occorrenze: 1
bcacb -> Occorrenze: 1
abca -> Occorrenze: 1
acba -> Occorrenze: 1
cac -> Occorrenze: 1

$ go run esercizio_2.go eabcacf
abca -> Occorrenze: 1
cac -> Occorrenze: 1

$ go run esercizio_2.go eabbcabcbf
bbcabcb -> Occorrenze: 1
bcabcb -> Occorrenze: 1
abbca -> Occorrenze: 1
bbcab -> Occorrenze: 1
cabc -> Occorrenze: 1
bcab -> Occorrenze: 1
bcb -> Occorrenze: 1

$ go run esercizio_2.go abcce

```