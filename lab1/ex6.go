package main

import (
	"fmt"
	"sort"
)

func sorteazaDupaInaltime(persoane map[string]int) []string {
	nume := make([]string, 0, len(persoane))
	for n := range persoane {
		nume = append(nume, n)
	}

	sort.Slice(nume, func(i, j int) bool {
		return persoane[nume[i]] > persoane[nume[j]]
	})

	return nume
}

func exercitiul6() {
	persoane := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	fmt.Println("--- Exercitiul 6 ---")
	fmt.Println("Map-ul initial (nume: inaltime):")
	for nume, inaltime := range persoane {
		fmt.Printf("  %s: %d\n", nume, inaltime)
	}

	numeOrdonate := sorteazaDupaInaltime(persoane)

	fmt.Println("\nNumele persoanelor in ordinea descrescatoare a inaltimilor:")
	for i, nume := range numeOrdonate {
		fmt.Printf("%d. %s (%d cm)\n", i+1, nume, persoane[nume])
	}

	fmt.Println("\nSlice rezultat cu numele:", numeOrdonate)
}
