# Esercizio 3

## Parte 1

Scrivere un programma che legga da **standard input** una sequenza di righe di testo, ognuna delle quali descrive un punto sul piano cartesiano.

Ogni riga di testo è una stringa che specifica l'etichetta del punto (ad es.: `A`, `B`, ...), l'ascissa e l'ordinata del punto nel seguente formato:
```text
etichetta;x;y
```
>*Esempio*:
>Si ipotizzi che vengano inserite da **standard input** le seguenti di righe di testo:
>```text
>A;10.0;2.0
>B;11.5;3.0
>C;8.0;1.0
>```
>tali righe specificano 3 punti: A(10, 2), B(11.5, 3) e C(8, 1).

Definire la struttura `Punto` per memorizzare l'`etichetta`, l'`ascissa` e l'`ordinata` di un punto sul piano cartesiano.

Implementare una funzione `NuovoTragitto() (tragitto []Punto)` che:
1. legge da **standard input** una sequenza di righe di testo nel formato *etichetta*;*x*;*y*, terminando la lettura quando viene letto l'indicatore End-Of-File (EOF);
2. restituisce un valore `[]Punto` nella variabile `tragitto` in cui è memorizzata la sequenza di istanze del tipo `Punto` inizializzate con i valori letti da **standard input**.
L'ordine dei punti all'interno della slice `tragitto` deve rispettare l'ordine in cui i corrispondenti valori sono stati letti da **standard input**.

Si assuma che:
* le righe di testo lette da **standard input** siano nel formato corretto;
* la tripla di valori presente in ogni riga specifichi correttamente un punto sul piano cartesiano;
* vengano lette da **standard input** almeno 2 righe di testo.

## Parte 2

I punti letti da **standard input** nella parte `Parte 1` dell'esercizio definiscono un *tragitto*.
Ogni coppia di punti consecutivi definisce una *tratta* del tragitto.

La lunghezza di ciascuna tratta del tragitto è pari alla distanza euclidea tra i due punti che definiscono la tratta. 
Per esempio, la lunghezza della tratta `AB`, è pari alla distanza euclidea tra i punti `A` e `B`: ((x<sub>A</sub>-x<sub>B</sub>)<sup><small>2</small></sup> + (y<sub>A</sub>-y<sub>B</sub>)<sup><small>2</small></sup>)<sup><small>1/2</small></sup>.

La lunghezza del tragitto è data dalla somma delle lunghezze delle sue tratte.

Una volta terminata la fase di lettura (**Parte 1**), il programma deve:
* stampare a video la lunghezza totale del tragitto;
* stampare a video la rappresentazione `string` del primo punto del tragitto che si incontra dopo aver percorso più della metà della lunghezza del tragitto.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* `Distanza(p1, p2 Punto) float64` che riceve in input due punti nei parametri `p1` e `p2` e restituisce un valore `float64` pari alla distanza euclidea tra i punti rappresentati da `p1` e `p2`;
* `String(p Punto) string` che riceve in input un punto nel parametro `p` e restituisce un valore `string` che corrisponde alla rappresentazione `string` di `p` nel formato `etichetta = (x, y)`;
* `Lunghezza(tragitto []Punto) float64` che riceve in input una slice di punti nel parametro `tragitto` e restituisce un valore `float64` pari alla lunghezza del tragitto rappresentato da `tragitto`. 

##### Esempio d'esecuzione:

```text
$ cat punti1
A;10.0;2.0
B;11.5;3.0
C;8.0;1.0
D;3;4
E;1;0
F;-1;5
$ go run esercizio_3.go < punti1
Lunghezza percorso: 21.522
Punto oltre metà: D = (3.0, 4.0)

$ cat punti2
A;0;0
B;4;0
C;4;4
D;0;4
E;0;0
$ go run esercizio_3.go < punti2
Lunghezza percorso: 16.000
Punto oltre metà: D = (0.0, 4.0)

```