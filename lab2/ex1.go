package main

import "fmt"

func trimiteMesajEx1(canal chan string) {
	canal <- "salut"
}

func exercitiul1() {
	fmt.Println("\n--- Exercitiul 1 ---")
	canal := make(chan string)

	go trimiteMesajEx1(canal)

	mesaj := <-canal
	fmt.Println("Mesaj receptionat cu succes in main():", mesaj)
}
