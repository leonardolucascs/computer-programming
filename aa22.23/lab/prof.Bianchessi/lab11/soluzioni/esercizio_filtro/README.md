# Esercizio filtro

**Definizione:** La codifica compressa di una stringa di caratteri è definita da una sequenza di coppie di valori `ch n`, dove `ch` rappresenta un carattere, mentre `n>0` è il numero di volte che il carattere `ch` viene ripetuto.

> *Esempio:* Alla sequenza di valori `e 3 a 4`, corrisponde la stringa `eeeaaaa`, ovvero il carattere `'e'`, ripetuto 3 volte, seguito dal carattere `'a'`, ripetuto 4 volte.

Scrivere un programma che legga da **riga di comando** la codifica compressa di una stringa di caratteri e stampi a video la stringa non compressa.  

Si assuma che ogni carattere dato in input appartenenga all'alfabeto inglese e sia quindi codificato all'interno dello standard US-ASCII (integrato nello standard Unicode).

Si assuma inoltre che la sequenza di valori specificata a riga di comando sia nel formato corretto e includa almeno una coppia `ch n`.

Verificare il corretto funzionamento delle funzioni utilizzando le unit di test allegate.

##### Esempi d'esecuzione:

```text
$ go run esercizio_filtro.go e 3 r 4 Y 3
eeerrrrYYY

$ go run esercizio_filtro.go R 3 A 1 B 10
RRRABBBBBBBBBB

$ go run esercizio_filtro.go W 2
WW

$ go run esercizio_filtro.go r 3 P 4
rrrPPPP
```