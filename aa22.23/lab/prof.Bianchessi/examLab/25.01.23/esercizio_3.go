/*
	Parte 1:
	Scrivere un programma che:
		- legga da standard input una sequenza di righe di testo
		- termini la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da standard input 
			l'indicatore End-Of-File(EOF)

	Ogni riga di testo avrà il formato:
		<inizio fine nome>

	dove:
		- inizio e fine sono stringhe nella forma 'hh:mm' (ore:minuti)e, in particolare, 'hh' rappresenta un intero 
			da 0 a 23 e 'mm' rappresenta un intero da 0 a 59;
		- nome è una stringa arbitraria in cui non è presente nessun carattere di spaziatura.

	Ogni riga di testo descrive un evento annotato in un'ipotetica agenda.

	a) Definire la struttura Evento che modella l'entità evento considerandone come proprietà
		- Inizio(ora e minuto)
		- Fine(ora e minuto)
		- Nome

	b) Implementare una funzione NuovaAgenda()(agenda []Evento) che:
		1. legge da standard input una sequenza di righe di testo nel formato <inizio fine nome>, terminando la 
			leetura quando viene letto l'indicatore End-Of-File(EOF)
		2. restituisce un valore []Evento nella variabile agenda, in cui è memorizzato la sequenza di valori di tipo
			Evento inizializzati con i valori letti da standard input


	Si assuma che:
		- le righe di testo lette da standard input siano nel formato corretto
		- vengano lette da standard input almeno 2 righe di testo

	* - * - * - * - * * - * - * - * - *

	Parte 2:
	Una volta terminata la lettura(Parte 1), il programma deve:
		1. ordinare gli eventi memorizzati in agenda in ordine cronologico rispetto al valore di Inizio(se due eventi in 
			agenda hanno il medesimoo Inizio, l'evento memorizzato per primo in agenda precede l'altro)
		2. come mostrato nell'esempio d'esecuzione stampare gli eventi ordinati cronologicamente evitando di stampare 
			quelli che si sovvrappongono ad eventi che li precedono, più precisamente

			i. il primo evento deve essere sempre stampato
			ii. un evento successivo al primo deve essere stampato se e solo se non è sovrapposto ad un evento 
				precedente che viene stampato (un evento B è sovvrapposto ad un evento A se il valore di inizio di B è 
				inferiore o uguale, in senso temporale, al valore di fine di A; per esempio un evento che inizia 
				alle 10:30 si sovvrappone ad uno che termina alle 10:30 ma non ad uno che termina alle 10:29)

		Si noti che ciascun evento deve essere stampato nel formato Inizio-Fine Nome (in particolare, si noti il
			carattere - tra Inizio e Fine)

	Oltre alla funzione main(), devono essere definite ed utilizzate almeno le seguenti funzioni:
		- una funzione <isPrecedente(e1, e2 Evento) bool> che riceve in input due istanze di tipo Evento nei parametri
			e1 ed e2 e restituisce true se e solo se e1 precede e2 (un evento A precede un evento B se il valore di
				inizio A è strettamente inferiore, in senso temporale, al valore di inizio B)
		- una funzione <isSovrapposto(e1, e2 Evento) bool che riceve in input due istanze di tipo Evento nei parametri 
			e1 ed e2, tali che e1 precede e2, e restituisce true se e solo se e2 si sovrappone ad e1 (confronta ii)
		- una funzione <StringaEvento(e Evento) string> che riceve in input un'istanza di tipo Evento nel parametro e 
			e restituisce un valore string che corrisponde alla rappresentazione string di e nel formato Inizio-Fine Nome

	Esempio d'esecuzione:
		$ cat agenda_1.txt
		10:00 14:00 lezione
		10:00 18:00 shopping
		08:00 08:30 doccia
		08:45 09:15 colazione

		$ go run esercizio_3.go < agenda_1.txt
		08:00 08:30 doccia
		08:45 09:15 colazione
		10:00 14:00 lezione

		$ cat agenda_2.txt
		15:00 18:00 boxe
		16:00 18:30 pianoforte
		13:20 13:35 caffè
		14:50 15:40 pisolino
		09:30 12:00 lezione

		$ go run esercizio_3.go < agenda_2.txt
		09:30 12:00 lezione
		13:20 13:35 caffè
		14:50 15:40 pisolino
		16:00 18:30 pianoforte
		
*/

package main

import "os"
import "fmt"
import "bufio"
import "io"

inputReader := bufio.NewReader(os.Stdin)
input, _ inputReader.ReadString('\n')

type Orario struct {
	ora uint
	minuto uint
}

type Evento struct {
	inizio Orario
	fine Orario
	nome string
}

func main() {

	inputFile, _ := os.Open("nome_file.txt")

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for{
		inputString, readError := inputReader.ReadString('\n')
		if readError == io.EOF {
			return
		}
		fmt.Printf("input %s", inputString)
	}
}

func NuovaAgenda() (agenda []Evento) {
	var tmp Evento

	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('CTRL+D')

	agenda = append(agenda, tmp)
}
