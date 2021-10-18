/*
arrivano pazienti al HO
due code a seconda della gravità
4 medici : 2 esperti, 2 no
casi gravi solo medico esperto
caso non grave : 1 sec da esperto, 2 da non esperto
caso grave: 1 secondo visita e 3 in sala
ho una sola sala!
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Paziente struct {
	nome  string
	grave bool
}

type Medico struct {
	nome    string
	esperto bool
}

func arrivoPazienti(codaBianca chan Paziente, codaGrave chan Paziente) {
	for i := 1; i < 51; i++ {
		if rand.Intn(10) < 3 {
			pat := Paziente{fmt.Sprint(i), true}
			fmt.Println("Il paziente ", i, "è grave")
			codaGrave <- pat
		} else {
			pat := Paziente{fmt.Sprint(i), false}
			codaBianca <- pat

		}
	}
}

func lavoroMedico(codaBianca chan Paziente, codaGrave chan Paziente, med Medico, salaOperatoria chan int) {
	for {
		if med.esperto {
			select {
			case pat := <-codaGrave:
				time.Sleep(time.Second)
				lock := <-salaOperatoria
				time.Sleep(3 * time.Second)
				salaOperatoria <- lock
				fmt.Println("il medico ", med.nome, " ha curato il paziente ", pat.nome)
			case pat := <-codaBianca:
				time.Sleep(time.Second)
				fmt.Println("il medico ", med.nome, " ha curato il paziente ", pat.nome)
			}
		} else {
			pat := <-codaBianca
			time.Sleep(2 * time.Second)
			fmt.Println("il medico ", med.nome, " ha curato il paziente ", pat.nome)
		}
	}

}

func stampaTempo() {
	t := 0
	time.Sleep(500 * time.Millisecond)
	for {
		fmt.Println(t, "5")
		time.Sleep(time.Second)
		t = t + 1
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	med1 := Medico{"Alice", true}
	med2 := Medico{"Bob", true}
	med3 := Medico{"Carla", false}
	med4 := Medico{"Davide", false}

	codaBianca := make(chan Paziente, 20)
	codaGrave := make(chan Paziente, 20)

	salaOperatoria := make(chan int, 1)
	salaOperatoria <- 1

	go lavoroMedico(codaBianca, codaGrave, med1, salaOperatoria)
	go lavoroMedico(codaBianca, codaGrave, med2, salaOperatoria)
	go lavoroMedico(codaBianca, codaGrave, med3, salaOperatoria)
	go lavoroMedico(codaBianca, codaGrave, med4, salaOperatoria)

	go stampaTempo()

	go arrivoPazienti(codaBianca, codaGrave)

	time.Sleep(time.Minute)
}
