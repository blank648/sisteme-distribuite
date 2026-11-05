package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool
}

func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	out <- Msg{Value: id, Done: false}

	for msg := range in {
		if msg.Done {
			if *isLeader {
				return
			}
			out <- msg
			return
		}

		if msg.Value > id {
			out <- msg
		} else if msg.Value == id {
			*isLeader = true
			out <- Msg{Value: id, Done: true}
		}
	}
}

func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	channels := make([]chan Msg, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan Msg, 1)
	}

	isLeaders := make([]bool, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go node(ids[i], channels[i], channels[(i+1)%n], &isLeaders[i], &wg)
	}

	wg.Wait()

	leaderId := -1
	for i := 0; i < n; i++ {
		if isLeaders[i] {
			leaderId = ids[i]
			fmt.Printf("Nodul de pe pozitia %d (cu ID = %d) a fost declarat LIDER!\n", i, ids[i])
		}
	}

	return leaderId
}

func exercitiul1() {
	fmt.Println("\nExercitiul 1: Algoritmul de Flooding / LCR pe Inel Unidirectional")

	testRing1 := []int{7, 4, 9, 2, 5, 33, 10, 1}
	fmt.Println("Test 1 (Configuratia din Cursul 3):")
	fmt.Printf("Inelul de noduri: %v\n", testRing1)
	lider1 := RingFlooding(testRing1)
	fmt.Printf("Lider ales: ID = %d (corect: 33)\n\n", lider1)

	testRing2 := []int{15, 82, 3, 44, 91, 27}
	fmt.Println("Test 2 (Configuratie arbitrara):")
	fmt.Printf("Inelul de noduri: %v\n", testRing2)
	lider2 := RingFlooding(testRing2)
	fmt.Printf("Lider ales: ID = %d (corect: 91)\n\n", lider2)
}
