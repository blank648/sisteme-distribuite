package main

import "fmt"

func filtreazaPare(numere []int) []int {
	var pare []int

	for _, numar := range numere {
		if numar%2 == 0 {
			pare = append(pare, numar)
		}
	}

	return pare
}

func exercitiul3() {
	numere := []int{12, 7, 9, 14, 22, 5, 8, 3, 16, 21, 40}

	fmt.Println("--- Exercitiul 3 ---")
	fmt.Println("Sirul initial de numere:", numere)

	pare := filtreazaPare(numere)
	fmt.Println("Slice cu numerele pare:", pare)
}
