package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Coordonata2D struct {
	X float64
	Y float64
}

func robotExplorator(
	nume string,
	traseu []Coordonata2D,
	chOut chan<- Coordonata2D,
	chIn <-chan Coordonata2D,
	raportObserver chan<- string,
) {
	fmt.Printf("[%s] Pornit de observer. Lungime traseu de parcurs: %d puncte.\n", nume, len(traseu))

	var pozitiiPartener []Coordonata2D
	var wgReceptie sync.WaitGroup
	wgReceptie.Add(1)

	go func() {
		defer wgReceptie.Done()
		for poz := range chIn {
			pozitiiPartener = append(pozitiiPartener, poz)
			fmt.Printf("   -> [%s] A receptionat pozitia partenerului: (%.1f, %.1f)\n", nume, poz.X, poz.Y)
		}
		fmt.Printf("[%s] Notificare primita: Partenerul s-a oprit (canalul s-a inchis). Toate cele %d pozitii au fost citite!\n",
			nume, len(pozitiiPartener))
	}()

	for i, punct := range traseu {
		time.Sleep(time.Duration(120+rand.Intn(100)) * time.Millisecond)
		fmt.Printf("[%s] Deplasare la pasul %d/%d -> Pozitie (%.1f, %.1f) - transmitere coordonate...\n",
			nume, i+1, len(traseu), punct.X, punct.Y)
		chOut <- punct
	}

	fmt.Printf("[%s] Traseu complet parcurs! Se notifica partenerul ca s-a oprit (close channel).\n", nume)
	close(chOut)

	wgReceptie.Wait()

	fmt.Printf("[%s] Raportare finala catre observer...\n", nume)
	raportObserver <- fmt.Sprintf("Robotul %s a incheiat cu succes (a parcurs %d puncte si a receptionat %d puncte de la partener)",
		nume, len(traseu), len(pozitiiPartener))
}

func exercitiul7() {
	fmt.Println("\n--- Exercitiul 7 ---")

	traseuRobot1 := []Coordonata2D{
		{X: 0.0, Y: 0.0},
		{X: 1.5, Y: 2.0},
		{X: 3.0, Y: 4.5},
		{X: 5.0, Y: 6.0},
		{X: 7.5, Y: 8.0},
	}

	traseuRobot2 := []Coordonata2D{
		{X: 10.0, Y: 10.0},
		{X: 8.5, Y: 8.0},
		{X: 6.0, Y: 5.5},
		{X: 4.0, Y: 3.0},
	}

	ch1to2 := make(chan Coordonata2D, 10)
	ch2to1 := make(chan Coordonata2D, 10)

	raportObserver := make(chan string, 2)

	fmt.Println("[Observer] Pornesc cei doi roboti in sistemul distribuit...")
	go robotExplorator("Robot-1", traseuRobot1, ch1to2, ch2to1, raportObserver)
	go robotExplorator("Robot-2", traseuRobot2, ch2to1, ch1to2, raportObserver)

	fmt.Println("[Observer] Astept rapoartele finale de la ambii roboti...")
	raport1 := <-raportObserver
	fmt.Printf("[Observer] Raport receptionat 1: %s\n", raport1)

	raport2 := <-raportObserver
	fmt.Printf("[Observer] Raport receptionat 2: %s\n", raport2)

	fmt.Println("\n[Observer] Ambii roboti au raportat finalizarea cu succes a sarcinilor!")
}
