package main

import "fmt"

func main() {
	for {
		fmt.Println("1. Exercitiul 1 (Algoritmul lui Cristian: sincronizare ceasuri)")
		fmt.Println("2. Exercitiul 2 (Ceasuri logice Lamport pe inel de 3 procese)")
		fmt.Println("0. Iesire")
		fmt.Print("Alege exercitiul dorit: ")

		var optiune int
		_, err := fmt.Scanln(&optiune)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Optiune invalida!")
			continue
		}

		switch optiune {
		case 1:
			exercitiul1()
		case 2:
			exercitiul2()
		case 0:
			fmt.Println("Inchidere program.")
			return
		default:
			fmt.Println("Optiune inexistenta!")
		}
	}
}
