package main

import "fmt"

func main() {
	for {
		fmt.Println("1. Exercitiul 1 (Algoritmul de Flooding/LCR pe inel unidirectional)")
		fmt.Println("2. Exercitiul 2 (Algoritmul de Flooding pentru grafuri tare conexe)")
		fmt.Println("3. Exercitiul 3 (Algoritmul de Flooding de medie pentru vot majoritar)")
		fmt.Println("0. Iesire")
		fmt.Print("Alege exercitiul dorit: ")

		var optiune int
		_, err := fmt.Scanln(&optiune)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Optiune invalida! Te rog sa introduci un numar.")
			continue
		}

		switch optiune {
		case 1:
			exercitiul1()
		case 2:
			exercitiul2()
		case 3:
			exercitiul3()
		case 0:
			fmt.Println("Inchidere program.")
			return
		default:
			fmt.Println("Optiune inexistenta! Alege un numar din meniu.")
		}
	}
}
