package main

import "fmt"

func este_prim(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func exercitiul1() {
	var n int
	fmt.Println("--- Exercitiul 1 ---")
	fmt.Print("Introdu un numar natural: ")
	fmt.Scanln(&n)
	if n <= 0 {
		fmt.Println("Numar incorect")
		return
	}
	if este_prim(n) {
		fmt.Println("Numarul este prim")
	} else {
		fmt.Println("Numarul nu este prim")
	}
}
