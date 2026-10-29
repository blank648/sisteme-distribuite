package main

import (
	"fmt"
	"math/rand"
	"time"
)

func robotCursa(nume string, start <-chan struct{}, finish chan<- string) {
	<-start
	fmt.Printf("[%s] A luat startul!\n", nume)

	timpRulare := time.Duration(250+rand.Intn(400)) * time.Millisecond
	time.Sleep(timpRulare)

	finish <- fmt.Sprintf("[%s] a trecut linia de sosire in %v", nume, timpRulare)
}

func exercitiul6() {
	fmt.Println("\n--- Exercitiul 6 ---")
	startSignal := make(chan struct{})
	finishRobot1 := make(chan string)
	finishRobot2 := make(chan string)

	go robotCursa("Robot 1", startSignal, finishRobot1)
	go robotCursa("Robot 2", startSignal, finishRobot2)

	fmt.Println("[Observer] Ambii roboti sunt la linia de start. Numaratoare inversa:")
	for i := 3; i >= 1; i-- {
		fmt.Printf("[Observer] %d...\n", i)
		time.Sleep(150 * time.Millisecond)
	}

	fmt.Println("[Observer] START! (Semnalul de pornire a fost transmis simultan)")
	close(startSignal)

	fmt.Println("[Observer] Se asteapta trecerea liniei de sosire (folosind instructiunea select)...")

	select {
	case msg := <-finishRobot1:
		fmt.Printf("\n🏆 [Observer] CÂȘTIGĂTORUL ESTE: Robot 1!\n")
		fmt.Printf("   Detalii sosire: %s\n", msg)

		msg2 := <-finishRobot2
		fmt.Printf("🥈 [Observer] Locul 2: Robot 2\n")
		fmt.Printf("   Detalii sosire: %s\n", msg2)

	case msg := <-finishRobot2:
		fmt.Printf("\n🏆 [Observer] CÂȘTIGĂTORUL ESTE: Robot 2!\n")
		fmt.Printf("   Detalii sosire: %s\n", msg)

		msg1 := <-finishRobot1
		fmt.Printf("🥈 [Observer] Locul 2: Robot 1\n")
		fmt.Printf("   Detalii sosire: %s\n", msg1)
	}

	fmt.Println("\n[Observer] Cursa s-a incheiat cu succes.")
}
