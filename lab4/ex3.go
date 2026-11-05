package main

import (
	"fmt"
	"math"
	"sync"
)

type VoteMsg struct {
	round int
	val   float64
}

func avgVoteNode(
	id int,
	initialVote int,
	maxRounds int,
	inChans []<-chan VoteMsg,
	outChans []chan<- VoteMsg,
	decision *int,
	finalVal *float64,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	x := float64(initialVote)
	inDegree := len(inChans)

	for t := 0; t < maxRounds; t++ {
		for _, out := range outChans {
			out <- VoteMsg{round: t, val: x}
		}

		sumReceived := 0.0
		for _, in := range inChans {
			msg := <-in
			sumReceived += msg.val
		}

		x = (x + sumReceived) / float64(inDegree+1)
	}

	*finalVal = x
	sgn := 0.0
	diff := x - 0.5
	if diff > 1e-9 {
		sgn = 1.0
	} else if diff < -1e-9 {
		sgn = -1.0
	}

	*decision = int(math.Round(0.5 + 0.5*sgn))
}

func RingMajorityFlooding(votes []int, rounds int) []int {
	n := len(votes)
	if n == 0 {
		return nil
	}

	chans := make([]chan VoteMsg, n)
	for i := 0; i < n; i++ {
		chans[i] = make(chan VoteMsg, rounds+1)
	}

	decisions := make([]int, n)
	finalVals := make([]float64, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		next := (i + 1) % n

		inC := []<-chan VoteMsg{chans[i]}
		outC := []chan<- VoteMsg{chans[next]}

		go avgVoteNode(i, votes[i], rounds, inC, outC, &decisions[i], &finalVals[i], &wg)
	}

	wg.Wait()

	fmt.Printf("Valori convergente calculate de noduri: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%.4f ", finalVals[i])
	}
	fmt.Println()

	return decisions
}

func exercitiul3() {
	fmt.Println("\n--- Exercitiul 3: Algoritmul de Flooding de Medie pentru Vot Majoritar ---")

	testVotes1 := []int{1, 0, 1, 1, 0, 1, 0, 1}
	rounds := 20

	fmt.Println("Test 1:")
	fmt.Printf("Voturi initiale: %v\n", testVotes1)
	count1 := 0
	for _, v := range testVotes1 {
		count1 += v
	}
	fmt.Printf("Total voturi '1': %d din %d (majoritate: 1)\n", count1, len(testVotes1))

	decisions1 := RingMajorityFlooding(testVotes1, rounds)
	fmt.Printf("Decizii finale luate de noduri: %v\n\n", decisions1)

	testVotes2 := []int{0, 0, 1, 0, 0, 1, 0}
	fmt.Println("Test 2:")
	fmt.Printf("Voturi initiale: %v\n", testVotes2)
	count2 := 0
	for _, v := range testVotes2 {
		count2 += v
	}
	fmt.Printf("Total voturi '1': %d din %d (majoritate: 0)\n", count2, len(testVotes2))

	decisions2 := RingMajorityFlooding(testVotes2, rounds)
	fmt.Printf("Decizii finale luate de noduri: %v\n\n", decisions2)
}

