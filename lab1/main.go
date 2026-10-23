package main

import "fmt"

func main() {
	for {
		fmt.Println("1. Exercitiul 1 (Verificare numar prim)")
		fmt.Println("2. Exercitiul 2 (Sortare array crescator - Bubble Sort)")
		fmt.Println("3. Exercitiul 3 (Filtrare numere pare din sir)")
		fmt.Println("4. Exercitiul 4 (Eliminare element dupa index din slice float32)")
		fmt.Println("5. Exercitiul 5 (Inversare in-place slice string-uri)")
		fmt.Println("6. Exercitiul 6 (Sortare descrescatoare map dupa inaltime)")
		fmt.Println("7. Exercitiul 7 (Frecventa cuvinte citite din fisier)")
		fmt.Println("0. Iesire")
		fmt.Print("Alege exercitiul dorit (0-7): ")

		var optiune int
		_, err := fmt.Scanln(&optiune)
		if err != nil {
			fmt.Println("Optiune invalida!")
			continue
		}

		switch optiune {
		case 1:
			exercitiul1()
		case 2:
			exercitiul2()
		case 3:
			exercitiul3()
		case 4:
			exercitiul4()
		case 5:
			exercitiul5()
		case 6:
			exercitiul6()
		case 7:
			exercitiul7()
		case 0:
			fmt.Println("Inchidere program.")
			return
		default:
			fmt.Println("Optiune inexistenta!")
		}
	}
}
