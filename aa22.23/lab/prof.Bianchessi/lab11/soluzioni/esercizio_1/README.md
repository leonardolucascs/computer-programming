# Esercizio 1

Scrivere un programma che legga da **riga di comando** una sequenza di stringhe di caratteri (parole di senso compiuto rispetto alla lingua italiana).
 
 La **prima parola** specificata nella sequenza letta è in **posizione 0** (nella sequenza letta), la **seconda parola** specificata nella sequenza letta è in **posizione 1**, la **terza parola** specificata nella sequenza letta è in **posizione 2**, etc.
  
 Il programma deve ristampare a video la sequenza di parole lette come mostrato nell'**Esempio d'esecuzione**:
 * ogni parola viene ristampata alternando caratteri in maiuscolo a caratteri in minisculo;
 * le parole specificate in posizione pari nella sequenza letta vengono ristampate a video incominciando con un carattere maiuscolo, mentre quelle in posizione dispari incominciano con un carattere minuscolo.  

Si assuma inoltre che la sequenza di valori specificata a riga di comando sia nel formato corretto e includa almeno una stringa.

Oltre alla funzione `main()` il programma deve includere le seguenti funzioni:
* `TrasformaParola(parola string, posizione int) (parolaTrasformata string)` che riceve in input un valore `string` ed un valore `int` nei parametri `parola` e `posizione`, e restituisce nella variabile `parolaTrasformata` un valore `string` in cui i caratteri di `parola` compaiono alternativamente in maiuscolo e in minuscolo: se `posizione` è pari, il primo carattere in `parolaTrasformata` è maiuscolo, altrimenti il primo carattere in `parolaTrasformata` è minuscolo,

##### Esempio d'esecuzione:

```text
$ go run esercizio_1.go ciao mondo
CiAo mOnDo 

$ go run esercizio_1.go esame di programmazione
EsAmE dI PrOgRaMmAzIoNe 

$ go run esercizio_1.go gennaio febbraio marzo aprile maggio giugno luglio agosto
GeNnAiO fEbBrAiO MaRzO aPrIlE MaGgIo gIuGnO LuGlIo aGoStO 
``` 