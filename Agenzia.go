/*
Scrivete un programma che simuli una agenzia di viaggi che deve gestire le prenotazioni per due
diversi viaggi da parte di 7 clienti. Ogni cliente fa una prenotazione per un viaggio in una delle due
mete disponibili (Spagna e Francia), ognuna delle quali ha un numero minimo di partecipanti per
essere confermata (rispettivamente 4 e 2).
● Creare la struttura Cliente col relativo campo “nome”.
● Creare la struttura Viaggio col rispettivo campo “meta”.
● Creare la function prenota, che prende come input una persona e che prenota uno a caso dei due
viaggi.
● Creare una function stampaPartecipanti che alla fine del processo stampa quali viaggi sono
confermati e quali persone vanno dove.
● Ogni persona può prenotarsi al viaggio contemporaneamente.
● Create tutte le classi e function che vi servono, ma mantenete la struttura data dalle due strutture e
le due function che ho elencato sopra.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct {
	nome string
}

type Viaggio struct {
	meta         string
	numeroMinimo int
}
type Gruppo struct {
	persone      []Cliente
	destinazione Viaggio
}

func prenota(persona Cliente, meta1 chan Gruppo, meta2 chan Gruppo, wg *sync.WaitGroup) {
	defer wg.Done()
	if rand.Intn(10) <= 5 {
		gruppo1 := <-meta1
		gruppo1.persone = append(gruppo1.persone, persona)
		meta1 <- gruppo1
	} else {
		gruppo2 := <-meta2
		gruppo2.persone = append(gruppo2.persone, persona)
		meta2 <- gruppo2
	}
}

func stampaPartecipanti(meta1 chan Gruppo, meta2 chan Gruppo) {
	gruppo1 := <-meta1
	if len(gruppo1.persone) < gruppo1.destinazione.numeroMinimo {
		fmt.Println("Il viaggio in ", gruppo1.destinazione.meta, " non è stato organizzato a causa del numero di partecipanti")
	} else {
		fmt.Println("Il viaggio in ", gruppo1.destinazione.meta, " è stato confermato con i seguenti partecipanti :")
		for i := 0; i < len(gruppo1.persone); i++ {
			fmt.Println(gruppo1.persone[i].nome)
		}
		fmt.Println("BUEN VIAJE!")
	}

	fmt.Println()

	gruppo2 := <-meta2
	if len(gruppo2.persone) < gruppo2.destinazione.numeroMinimo {
		fmt.Println("Il viaggio in ", gruppo2.destinazione.meta, " non è stato organizzato a causa del numero di partecipanti")
	} else {
		fmt.Println("Il viaggio in ", gruppo2.destinazione.meta, " è stato confermato con i seguenti partecipanti :")
		for i := 0; i < len(gruppo2.persone); i++ {
			fmt.Println(gruppo2.persone[i].nome)
		}
		fmt.Println("BON VOYAGE!")
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	clienti := []Cliente{
		{"Pisolo"},
		{"Eolo"},
		{"Mammolo"},
		{"Dotto"},
		{"Brontolo"},
		{"Cucciolo"},
		{"Gongolo"},
	}

	meta1 := Viaggio{"Spagna", 4}
	meta2 := Viaggio{"Francia", 2}

	gruppo1 := make(chan Gruppo, 1)
	gruppo2 := make(chan Gruppo, 1)

	gruppo1 <- Gruppo{[]Cliente{}, meta1}
	gruppo2 <- Gruppo{[]Cliente{}, meta2}

	wg.Add(7)

	for i := 0; i < len(clienti); i++ {
		go prenota(clienti[i], gruppo1, gruppo2, &wg)
	}

	wg.Wait()

	stampaPartecipanti(gruppo1, gruppo2)
}
