/*
Scrivete un programma che simuli l’ordinazione, la cottura e l’uscita dei piatti in un ristorante. 10 clienti
ordinano contemporaneamente i loro piatti. In cucina vengono preparati in un massimo di 3 alla volta,
essendoci solo 3 fornelli. Il tempo necessario per preparare ogni piatto è fra i 4 e i 6 secondi. Dopo che
un piatto viene preparato, viene portato fuori da un cameriere, che impiega 3 secondi a portarlo fuori. Ci
sono solamente 2 camerieri nel ristorante.
● Creare la strutture Piatto e Cameriere col relativo campo “nome”.
● Creare le funzioni ordina che aggiunge il piatto a un buffer di piatti da fare; creare la function cucina che
cucina ogni piatto e lo mette in lista per essere consegnato; creare la function consegna che fa uscire
un piatto dalla cucina.
● Ogni cameriere può portare solo un piatto alla volta.
● Usate buffered channels per svolgere il compito.
● Attenzione: se per cucinare un piatto lo mandate nel buffer fornello di capienza 3 e lo ritirate dopo 3
secondi, non è detto che ritiriate lo stesso piatto che avete messo sul fornello. Tenetelo in memoria.
Ovviamente la vostra soluzione potrebbe differire dalla mia e questo hint potrebbe non servirvi.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cliente struct {
	nome     string
	pietanza Piatto
	id       int
}

type Piatto struct {
	nome string
	id   int
}

type Cameriere struct {
	nome string
}

func ordina(persona Cliente, ordinazioni chan Piatto) {

	ordine1 := persona.pietanza
	ordine1.id = persona.id
	ordinazioni <- ordine1
	fmt.Println("il cliente ", persona.nome, " ha ordinato ", ordine1.nome)
}

func cucina(ordinazioni chan Piatto, fornello1 chan int, fornello2 chan int, fornello3 chan int, piattiPronti chan Piatto) {

	select {
	case lock := <-fornello1:
		pietanza := <-ordinazioni
		x := rand.Intn(2)+4
		time.Sleep(time.Duration(x) * time.Second)
		fornello1 <- lock
		fmt.Println("il piatto ", pietanza.nome, " è stato cucinato nel fornello 1")
		piattiPronti <- pietanza
	case lock := <-fornello2:
		pietanza := <-ordinazioni
		x := rand.Intn(2)+4
		time.Sleep(time.Duration(x) * time.Second)
		fornello2 <- lock
		fmt.Println("il piatto ", pietanza.nome, " è stato cucinato nel fornello 2")
		piattiPronti <- pietanza
	case lock := <-fornello3:
		pietanza := <-ordinazioni
		x := rand.Intn(2)+4
		time.Sleep(time.Duration(x) * time.Second)
		fornello3 <- lock
		fmt.Println("il piatto ", pietanza.nome, " è stato cucinato nel fornello 3")
		piattiPronti <- pietanza
	}
	
}



func consegna(piattiPronti chan Piatto, camerieri chan Cameriere, persona[] Cliente) {
		cameriere := <-camerieri
		piatto := <-piattiPronti
		for i := 0; i < len(persona); i++ {
			if piatto.id == persona[i].id{
				time.Sleep(3 * time.Second)
				fmt.Println("il cameriere ", cameriere.nome, " ha consegnato ", piatto.nome, " a ", persona[i].nome)
				camerieri <- cameriere
			}
		}
}





func main() {

	rand.Seed(time.Now().UnixNano())


	piatto1 := Piatto{"pizza", 0}
	piatto2 := Piatto{"pasticcio", 0}
	piatto3 := Piatto{"pasta", 0}
	piatto4 := Piatto{"frittura", 0}
	piatto5 := Piatto{"tiramisù", 0}

	cameriere1 := Cameriere{"Superman"}
	cameriere2 := Cameriere{"Batman"}

	ordinazioni := make(chan Piatto,1)

	piattiPronti := make(chan Piatto,3)

	camerieri :=make(chan Cameriere,2)
	camerieri<-cameriere1
	camerieri<-cameriere2

	fornello1 := make(chan int, 1)
	fornello2 := make(chan int, 1)
	fornello3 := make(chan int, 1)
	fornello1 <- 1
	fornello2 <- 1
	fornello3 <- 1

	clienti := []Cliente{
		{"Iron Man", piatto1, 1},
		{"Spider Man", piatto2, 2},
		{"Hulk", piatto3, 3},
		{"Thor", piatto4, 4},
		{"Dr Strange", piatto5, 5},
		{"Captain America", piatto1, 6},
		{"Black Widow", piatto2, 7},
		{"Black Panther", piatto3, 8},
		{"Groot", piatto4, 9},
		{"Thanos", piatto5, 10},
	}

	for i := 0; i < len(clienti); i++ {
		go ordina(clienti[i], ordinazioni)
	}

	for i := 0; i < len(clienti); i++ {
		go cucina(ordinazioni,fornello1,fornello2,fornello3,piattiPronti)
	}

	for i := 0; i < len(clienti); i++ {
		go consegna(piattiPronti,camerieri,clienti)
	}


	time.Sleep(time.Minute)
}
