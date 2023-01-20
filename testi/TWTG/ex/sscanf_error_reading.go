package main

import "fmt"

func main() {
	
	var name, currency string

	_, err := fmt.Sscanf("transaction benson: dollars", "transaction %s: %s", &name, &currency)

	fmt.Println(err, name, currency)
}

// Output get: "input does not match format benson:"

/*
	Sscanf legge la stringa passata come parametro, e salva i corrispondenti valori delimitati dallo spazio secondo il formato specificato.

	Sscanf scans the argument string, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed.

	Cosa succede in questo caso? Il problema è legato al fatto che %s legge fino a ':'[viene letto e contato nella prima stringa], dopo aver processato il primo %s la funzione cerca ':' nella stringa di lettura ma risulta che il "cursore di lettura" sia già andato al carattere successivo pertanto non trovando più ':', la funzione termina e ritorna un errore.

	Errore che non si verifica se, invece, sopra dichiarassimo la seguente riga di codice:
		
		_, err := fmt.Sscanf("transaction benson : dollars", "transaction %s : %s", &name, &currency)

	Separando il ':' dalla parola che lo precede, evita l'errore causato in precedenza, consentendo una corretta lettura

*/
