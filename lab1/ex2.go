package main

import "fmt"

func bubbleSort(arr [8]int) [8]int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return arr
}

func exercitiul2() {
	arr := [8]int{64, 34, 25, 12, 22, 11, 90, 5}

	fmt.Println("--- Exercitiul 2 ---")
	fmt.Println("Array initial:", arr)

	sortat := bubbleSort(arr)
	fmt.Println("Array sortat crescator (Bubble Sort):", sortat)
}
