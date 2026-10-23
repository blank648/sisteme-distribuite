package main

import "fmt"

func trimiteMesaj(canal chan string) {
	canal <- "salut"

}

func main() {
	canal := make(chan string)

	go trimiteMesaj(canal)

	mesaj := <-canal

	fmt.Println("Mesaj receptionat cu succes in main():", mesaj)
}
