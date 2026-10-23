package main

import (
	"fmt"
)

func eliminaElement(slice []float32, index int) ([]float32, error) {
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("index invalid: %d (trebuie sa fie intre 0 si %d)", index, len(slice)-1)
	}

	return append(slice[:index], slice[index+1:]...), nil
}

func exercitiul4() {
	slice := []float32{10.5, 20.25, 33.75, 42.0, 55.5, 68.1}

	fmt.Println("--- Exercitiul 4 ---")
	fmt.Println("Slice initial:", slice)
	fmt.Printf("Introdu indexul elementului de eliminat (0 - %d): ", len(slice)-1)

	var index int
	_, err := fmt.Scanln(&index)
	if err != nil {
		fmt.Println("Eroare: te rog introdu un numar intreg valid.")
		return
	}

	rezultat, err := eliminaElement(slice, index)
	if err != nil {
		fmt.Println("Eroare:", err)
		return
	}

	fmt.Println("Slice dupa eliminarea elementului:", rezultat)
}
