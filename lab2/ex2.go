package main

import "fmt"

func primesteMesajEx2(canal <-chan string, confirmare chan<- bool) {
	fmt.Println("[goroutina secundara] Astept sa primesc mesajul din main()...")
	mesaj := <-canal
	fmt.Printf("[goroutina secundara] Mesaj receptionat cu succes din main(): '%s'\n", mesaj)
	confirmare <- true
}

func exercitiul2() {
	fmt.Println("\n--- Exercitiul 2 ---")
	canal := make(chan string)
	confirmare := make(chan bool)

	go primesteMesajEx2(canal, confirmare)

	mesajDeTrimis := "salut"
	fmt.Printf("[main] Se trimite mesajul '%s' catre goroutina secundara...\n", mesajDeTrimis)
	canal <- mesajDeTrimis

	<-confirmare
	fmt.Println("[main] Transferul a fost realizat cu succes!")
}
