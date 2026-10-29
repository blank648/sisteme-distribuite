package main

import (
	"fmt"
	"time"
)

func proceseazaDate(notificare chan<- struct{}) {
	fmt.Println("[goroutina] Incep procesarea datelor...")
	for i := 1; i <= 3; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("[goroutina] Procesare etapa %d/3 finalizata.\n", i)
	}

	fmt.Println("[goroutina] Toata procesarea s-a incheiat! Trimit notificare catre main()...")
	notificare <- struct{}{}
}

func exercitiul4() {
	fmt.Println("\n--- Exercitiul 4 ---")
	notificare := make(chan struct{})

	fmt.Println("[main] Se lanseaza goroutina de procesare...")
	go proceseazaDate(notificare)

	fmt.Println("[main] Goroutina main() asteapta blocata pe canal notificarea de finalizare...")
	<-notificare

	fmt.Println("[main] Notificare receptionata! Procesarea s-a terminat cu succes.")
}
