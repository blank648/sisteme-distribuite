package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

type Point struct {
	X float64
	Y float64
}

type Robot = Point

func distance(a Point, b Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Hypot(dx, dy)
}

func simulateRobot(
	id int,
	initialPos Point,
	radius float64,
	rounds int,
	inChans []<-chan Point,
	outChans []chan<- Point,
	finalPositions []Point,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	currentPos := initialPos

	for r := 0; r < rounds; r++ {
		for _, out := range outChans {
			out <- currentPos
		}

		sumX := currentPos.X
		sumY := currentPos.Y
		count := 1.0

		for _, in := range inChans {
			neighborPos := <-in
			if distance(currentPos, neighborPos) <= radius {
				sumX += neighborPos.X
				sumY += neighborPos.Y
				count++
			}
		}

		currentPos = Point{
			X: sumX / count,
			Y: sumY / count,
		}
	}

	finalPositions[id] = currentPos
}

func runAveraging(pos []Point, radius float64, rounds int) []Point {
	n := len(pos)
	if n == 0 {
		return nil
	}

	channels := make([][]chan Point, n)
	for i := 0; i < n; i++ {
		channels[i] = make([]chan Point, n)
		for j := 0; j < n; j++ {
			if i != j {
				channels[i][j] = make(chan Point, 2)
			}
		}
	}

	finalPositions := make([]Point, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)

		inC := make([]<-chan Point, 0, n-1)
		outC := make([]chan<- Point, 0, n-1)

		for j := 0; j < n; j++ {
			if i != j {
				inC = append(inC, channels[j][i])
				outC = append(outC, channels[i][j])
			}
		}

		go simulateRobot(i, pos[i], radius, rounds, inC, outC, finalPositions, &wg)
	}

	wg.Wait()
	return finalPositions
}

func main() {
	if file, err := os.Open("input.txt"); err == nil {
		defer file.Close()
		os.Stdin = file
	}

	var n int
	fmt.Scan(&n)

	pos := make([]Point, n)
	for i := range pos {
		fmt.Scan(&pos[i].X, &pos[i].Y)
	}

	fmt.Printf("Numar de roboti incarcati: %d\n", len(pos))
	fmt.Println("Pozitii initiale:")
	sumInitX := 0.0
	sumInitY := 0.0
	for i, p := range pos {
		fmt.Printf("Robot %d: (%.4f, %.4f)\n", i, p.X, p.Y)
		sumInitX += p.X
		sumInitY += p.Y
	}
	avgInitX := sumInitX / float64(len(pos))
	avgInitY := sumInitY / float64(len(pos))
	fmt.Printf("Centru de masa teoretic (media globala): (%.4f, %.4f)\n\n", avgInitX, avgInitY)

	radius := 16.0
	rounds := 30

	fmt.Printf("Rulare algoritm de averaging (raza vizibilitate = %.1f, runde = %d)...\n\n", radius, rounds)
	finalPositions := runAveraging(pos, radius, rounds)

	fmt.Println("Pozitii finale dupa averaging:")
	for i, p := range finalPositions {
		fmt.Printf("Robot %d: (%.4f, %.4f)\n", i, p.X, p.Y)
	}
}
