package main

import "fmt"

func main() {
	for {
		fmt.Println("1. Exercitiul 1 (Transfer mesaj de la goroutina la main prin channel)")
		fmt.Println("2. Exercitiul 2 (Transfer mesaj de la main la o goroutina prin channel)")
		fmt.Println("3. Exercitiul 3 (Receptie a doua mesaje pe acelasi canal in main)")
		fmt.Println("4. Exercitiul 4 (Notificare terminare procesare catre main fara WaitGroup)")
		fmt.Println("5. Exercitiul 5 (Pipeline de mesaje: send -> receive -> afisare in main)")
		fmt.Println("6. Exercitiul 6 (Sistem distribuit: cursa a doi roboti cu select in observer)")
		fmt.Println("7. Exercitiul 7 (Sistem distribuit: 2 roboti cu schimb de coordonate 2D si observer)")
		fmt.Println("0. Iesire")
		fmt.Print("Alege exercitiul dorit (0-7): ")

		var optiune int
		_, err := fmt.Scanln(&optiune)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Optiune invalida! Te rog sa introduci un numar intre 0 si 7.")
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
			fmt.Println("Optiune inexistenta! Alege un numar din intervalul 0-7.")
		}
	}
}
