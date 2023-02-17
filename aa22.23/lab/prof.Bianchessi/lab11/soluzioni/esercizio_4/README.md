# Esercizio 4

Scrivere un programma che legga da **riga di comando** due numeri interi `n` e `d`.

Il programma deve calcolare e stampare il più piccolo numero ottenibile rimuovendo `d` cifre decimali da `n`.

>*Esempio*:
>Si ipotizzi che i valori di `n` e `d` siano rispettivamente 4567 e 2.
>
>45 è il più piccolo numero ottenibile rimuovendo 2 cifre da 4567. Gli altri numeri ottenibili rimuovendo 2 cifre da 4567 sono 46, 47, 56, 57 e 67. Si noti che non è possibile considerare permutazioni delle cifre di `n`: rimuovendo 2 cifre da 4567, non è possibile, ad esempio, ottenere 76.

Si assuma che i valori letti da **riga di comando**  siano nel formato corretto e che il numero di cifre di `n` sia maggiore di `d`. 

##### Esempio d'esecuzione:

```text
$ go run esercizio_4.go 9452 2
numero migliore: 42

$ go run esercizio_4.go 1324 3
numero migliore: 1

$ go run esercizio_4.go 4612 1
numero migliore: 412

$ go run esercizio_4.go 4213 3
numero migliore: 1
```