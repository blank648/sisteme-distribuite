package main

import (
	"fmt"
	"time"
)

func exercitiul3() {
	fmt.Println("\n--- Exercitiul 3 ---")
	canalComun := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		canalComun <- "Primul mesaj: Salut de la Goroutina A"
	}()

	go func() {
		time.Sleep(120 * time.Millisecond)
		canalComun <- "Al doilea mesaj: Salut de la Goroutina B"
	}()

	fmt.Println("[main] Se asteapta primul mesaj pe canalul comun...")
	primulMesaj := <-canalComun
	fmt.Printf("[main] Receptie 1: '%s'\n", primulMesaj)

	fmt.Println("[main] Se asteapta al doilea mesaj pe ACELASI canal comun...")
	alDoileaMesaj := <-canalComun
	fmt.Printf("[main] Receptie 2: '%s'\n", alDoileaMesaj)

	fmt.Println("[main] Ambele mesaje au fost receptionate cu succes pe acelasi canal!")
}
