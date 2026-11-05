package main

import (
	"fmt"
	"sync"
)

type GenNodeMsg struct {
	round int
	maxID int
}

func genNode(
	id int,
	diam int,
	inChans []<-chan GenNodeMsg,
	outChans []chan<- GenNodeMsg,
	isLeader *bool,
	finalMaxID *int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	maxID := id
	*isLeader = false

	for r := 0; r < diam; r++ {
		for _, out := range outChans {
			out <- GenNodeMsg{round: r, maxID: maxID}
		}

		for _, in := range inChans {
			msg := <-in
			if msg.maxID > maxID {
				maxID = msg.maxID
			}
		}
	}

	*finalMaxID = maxID
	if maxID == id {
		*isLeader = true
	}
}

func GraphFlooding(ids []int, adj [][]int, diam int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	inChans := make([][]chan GenNodeMsg, n)
	outChans := make([][]chan GenNodeMsg, n)

	for u := 0; u < n; u++ {
		for _, v := range adj[u] {
			ch := make(chan GenNodeMsg, diam+1)
			outChans[u] = append(outChans[u], ch)
			inChans[v] = append(inChans[v], ch)
		}
	}

	isLeaders := make([]bool, n)
	finalMaxIDs := make([]int, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		inC := make([]<-chan GenNodeMsg, len(inChans[i]))
		for k, ch := range inChans[i] {
			inC[k] = ch
		}

		outC := make([]chan<- GenNodeMsg, len(outChans[i]))
		for k, ch := range outChans[i] {
			outC[k] = ch
		}

		go genNode(ids[i], diam, inC, outC, &isLeaders[i], &finalMaxIDs[i], &wg)
	}

	wg.Wait()

	leaderID := -1
	for i := 0; i < n; i++ {
		if isLeaders[i] {
			leaderID = ids[i]
			fmt.Printf("Nodul de la indexul %d (cu ID = %d) a fost declarat LIDER!\n", i, ids[i])
		}
	}

	return leaderID
}

func exercitiul2() {
	fmt.Println("\n--- Exercitiul 2: Algoritmul de Flooding pentru Grafuri Tare Conexe ---")

	ids := []int{14, 19, 22, 17, 8}
	adj := [][]int{
		0: {1},
		1: {2},
		2: {3, 0},
		3: {4},
		4: {1},
	}
	diam := 4

	fmt.Printf("ID-urile nodurilor din graful tare conex: %v\n", ids)
	fmt.Printf("Diametrul grafului (numar de runde): %d\n", diam)
	fmt.Println("Se ruleaza algoritmul Flooding_Gen(max)...")

	lider := GraphFlooding(ids, adj, diam)
	fmt.Printf("Rezultat final: Lider ales = ID %d (corect: 22)\n\n", lider)
}

