package main

import "fmt"

func main() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4"); fallthrough;
	case 5:
		fmt.Println("was <= 5"); fallthrough;
	case 6:
		fmt.Println("was <= 6"); fallthrough;
	case 7:
		fmt.Println("was <= 7"); fallthrough;
	case 8:
		fmt.Println("was <= 8"); fallthrough;
	default:
		fmt.Println("default case")
	}

}

/* La keyword fallthrough permette di considerare gli altri casi dello
   switch, trasferendo il controllo alla prima riga del case successivo.
   Fallthrough funziona in ogni caso persino nei casi valutati come falsi.

   Se utilizzato nell'ultimo case viene creato un errore in fase di compilazione.

   Nello switch in Golang l'uscita da ogni ramo (i diversi rami sono mutuamente esclusivi) è
   implicita, a differenza di linguaggi come C++, Java dove si usa break
*/
