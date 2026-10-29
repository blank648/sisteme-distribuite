package main

import "fmt"

func goroutineSend(canal chan<- string, mesaj string) {
	fmt.Printf("[goroutine send] Trimit mesajul '%s' catre goroutina receive...\n", mesaj)
	canal <- mesaj
}

func goroutineReceive(canalIn <-chan string, canalOut chan<- string) {
	mesaj := <-canalIn
	fmt.Printf("[goroutine receive] Am receptionat mesajul '%s'. Il returnez catre main()...\n", mesaj)
	canalOut <- mesaj
}

func exercitiul5() {
	fmt.Println("\n--- Exercitiul 5 ---")
	canalTransfer := make(chan string)
	canalRaspuns := make(chan string)

	go goroutineReceive(canalTransfer, canalRaspuns)

	mesajOriginal := "Salut transmis prin reteaua de goroutine!"
	go goroutineSend(canalTransfer, mesajOriginal)

	mesajReceptionat := <-canalRaspuns
	fmt.Printf("[main] Mesaj afisat pe ecran in main(): '%s'\n", mesajReceptionat)
}
