/*
Scrivete un programma che simuli un lavoro fatto da tre operai, ognuno dei quali deve usare un
martello, un cacciavite e un trapano per fare un lavoro. Devono usare il cacciavite DOPO il trapano e
il martello in un qualsiasi momento. Ovviamente, possono fare solo un lavoro alla volta. Gli attrezzi a
disposizione sono: due trapani, un martello e un cacciavite, quindi I tre operai devono aspettare di
avere a disposizione gli attrezzi per usarli. Modellate questa situazione minimizzando il più possibile le
attese.
● Creare la struttura Operaio col relativo campo “nome”.
● Creare la strutture Martello, Cacciavite e Trapano che devono essere “prese” dagli operai.
● Nelle function che creerete, inserite una stampa che dica quando l’operaio x ha preso l’oggetto y e
quando ha finito di usarlo.
● Hint sulla logica: ogni operaio può avere solo un oggetto alla volta e ogni oggetto può essere in mano
a un solo operaio.
● Per assicurarmi che ogni operaio abbia in mano un solo oggetto, posso mettere ogni operaio in un
channel, e prima di cercare di prendere un oggetto...
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Operaio struct {
	nome         string
	TrapanoUsato bool
}


type Martello struct {
	nome   string
	numero int
}

//va usato dopo il trapano
type Cacciavite struct {
	nome   string
	numero int
}

type Trapano struct {
	nome   string
	numero int
}

func usaTrapano(squadra chan Operaio, cassetta chan Trapano, wg *sync.WaitGroup) {

	defer wg.Done()
	operaio := <-squadra
	arnese := <-cassetta

	if arnese.numero > 0 {
		arnese.numero -= 1
		fmt.Println("l'operaio ", operaio.nome, "ha preso il ", arnese.nome)
		arnese.numero += 1
		cassetta <- arnese
		fmt.Println("l'operaio ", operaio.nome, "ha finito di usare il  ", arnese.nome)
		operaio.TrapanoUsato = true
		squadra <- operaio
	} else {
		fmt.Println("l'operaio ", operaio.nome, " non può usare il ", arnese.nome, " perchè non è disponibile")
		cassetta <- arnese
		squadra <- operaio
	}
}

func usaCacciavite(squadra chan Operaio, cassetta chan Cacciavite, wg *sync.WaitGroup) {

	defer wg.Done()
	operaio := <-squadra
	arnese := <-cassetta
	if operaio.TrapanoUsato == true {
		if arnese.numero > 0 {
			arnese.numero -= 1
			fmt.Println("l'operaio ", operaio.nome, "ha preso il ", arnese.nome)
			arnese.numero += 1
			cassetta <- arnese
			fmt.Println("l'operaio ", operaio.nome, "ha finito di usare il  ", arnese.nome)
			squadra <- operaio
		} else {
			fmt.Println("l'operaio ", operaio.nome, " non può usare il ", arnese.nome, " perchè non è disponibile")
			cassetta <- arnese
			squadra <- operaio
		}
	} else {
		fmt.Println("l'operaio ", operaio.nome, " non ha mai usato il trapano quindi non può usare il cacciavite")
		cassetta <- arnese
		squadra <- operaio
	}

}

func usaMartello(squadra chan Operaio, cassetta chan Martello, wg *sync.WaitGroup) {

	defer wg.Done()
	operaio := <-squadra
	arnese := <-cassetta

	if arnese.numero > 0 {
		arnese.numero -= 1
		fmt.Println("l'operaio ", operaio.nome, "ha preso il ", arnese.nome)
		arnese.numero += 1
		cassetta <- arnese
		fmt.Println("l'operaio ", operaio.nome, "ha finito di usare il  ", arnese.nome)
		squadra <- operaio
	} else {
		fmt.Println("l'operaio ", operaio.nome, " non può usare il ", arnese.nome, " perchè non è disponibile")
		cassetta <- arnese
		squadra <- operaio
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	operai := []Operaio{
		{"Mario", false},
		{"Paolo", false},
		{"Franco", false},
	}

	cassetta1 := make(chan Trapano, 2)
	cassetta2 := make(chan Cacciavite, 1)
	cassetta3 := make(chan Martello, 1)

	squadra := make(chan Operaio, len(operai))

	cassetta1 <- Trapano{"trapano", 2}
	cassetta2 <- Cacciavite{"cacciavite", 1}
	cassetta3 <- Martello{"martello", 1}

	for i := 0; i < len(operai); i++ {
		squadra <- operai[i]
	}

	wg.Add(len(operai))

	go usaTrapano(squadra, cassetta1, &wg)
	go usaMartello(squadra, cassetta3, &wg)
	go usaCacciavite(squadra, cassetta2, &wg)

	wg.Wait()
}
