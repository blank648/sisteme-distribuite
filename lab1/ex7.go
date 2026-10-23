package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func esteDoarLitereMici(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func citesteCuvinteDinFisier(numeFisier string) ([]string, error) {
	continut, err := os.ReadFile(numeFisier)
	if err != nil {
		return nil, err
	}

	cuvinteBrute := strings.Fields(string(continut))
	var cuvinte []string

	for _, cuv := range cuvinteBrute {
		if esteDoarLitereMici(cuv) {
			cuvinte = append(cuvinte, cuv)
		} else {
			cuvinte = append(cuvinte, strings.ToLower(cuv))
		}
	}

	return cuvinte, nil
}

func numaraAparitii(cuvinte []string) map[string]int {
	frecvente := make(map[string]int)
	for _, cuv := range cuvinte {
		frecvente[cuv]++
	}
	return frecvente
}

func exercitiul7() {
	numeFisier := "cuvinte.txt"

	fmt.Println("--- Exercitiul 7 ---")

	if _, err := os.Stat(numeFisier); os.IsNotExist(err) {
		exemplu := "mar para mar strugure banana para mar strugure portocala"
		err := os.WriteFile(numeFisier, []byte(exemplu), 0644)
		if err != nil {
			fmt.Printf("Eroare la crearea fisierului demonstrativ '%s': %v\n", numeFisier, err)
			return
		}
		fmt.Printf("Fisierul '%s' a fost creat automat cu datele de test.\n", numeFisier)
	}

	fmt.Printf("Citire din fisierul: %s\n", numeFisier)

	cuvinte, err := citesteCuvinteDinFisier(numeFisier)
	if err != nil {
		fmt.Printf("Eroare la citirea fisierului: %v\n", err)
		return
	}

	fmt.Println("Slice de string-uri citit din fisier:")
	fmt.Printf("%#v\n", cuvinte)

	frecvente := numaraAparitii(cuvinte)

	fmt.Println("\nNumarul de aparitii pentru fiecare cuvant:")
	for cuvant, numarAparitii := range frecvente {
		fmt.Printf("  %s: %d\n", cuvant, numarAparitii)
	}
}
