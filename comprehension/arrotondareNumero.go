// La logica dietro: 
	// 	dato un numero decimale N, le funzioni ceil-floor-round arrotondano al primo intero
	//  occerrente a seconda dei casi. 
	//  Per poter arrotondare dunque il numero decimale N a x cifre decimali, la funzione chiamata 
	//  agirà sul numero decimale N moltiplicato per 10^x_cifreDecimaliNecessarie, ottenendo così 
	// 	un "arrotondamento" provvisorio.
	
	// 	Per ritrovare il nostro numero decimale N arrotondato correttamente concluderemo 
	//	dividendo per 10^x_cifreDecimaliNecessarie

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 12.3456
	// In questo caso si vuole effettuare un arrotondamento a 2 cifre decimali

	// Floor returns the greatest integer value less than or equal to n.
	// arrotondamento per difetto (floor - pavimento): si limita ad effettuare il troncamento del 
	// numero alla cifra fissata, eliminando tutte le cifre successive
	fmt.Println(math.Floor(x*100)/100) // 12.34 (round down)


	// Ceil returns the least integer value greater than or equal to n.
	// arrotondamento per eccesso (ceil - soffitto): aggiunge 1 alla cifra fissata ed elimina le
	// successive
	fmt.Println(math.Ceil(x*100)/100)  // 12.35 (round up)	
	
	
	// Third... alternative
	// arrotondamento al più vicino [Corretta interpretazione]
	// fissata la cifra decimale, se la cifra successiva è <=4 -> per difetto 
	//													   >4  -> per eccesso
	fmt.Println(math.Round(x*100)/100) // 12.35 (round to nearest)	
	
}
