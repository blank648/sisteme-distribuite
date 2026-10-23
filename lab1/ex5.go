package main

import "fmt"

func inverseazaSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func exercitiul5() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	fmt.Println("--- Exercitiul 5 ---")
	fmt.Printf("Slice initial: %#v\n", slice)

	inverseazaSlice(slice)

	fmt.Printf("Rezultatul (inversat in-place): %#v\n", slice)
}
