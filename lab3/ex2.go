package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct{ counter int }

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() int {
	l.counter++
	return l.counter
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func runProcess(
	id int,
	in <-chan Message,
	out chan<- Message,
	rounds int,
	eventsChan chan<- Event,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	clock := &LamportClock{}

	for r := 0; r < rounds; r++ {
		numTicks := rand.Intn(3) + 1
		for t := 0; t < numTicks; t++ {
			clock.Tick()
			eventsChan <- Event{
				process:          id,
				internal_counter: clock.counter,
				kind:             "internal",
			}
			time.Sleep(time.Duration(rand.Intn(10)+1) * time.Millisecond)
		}

		sendTs := clock.Send()
		eventsChan <- Event{
			process:          id,
			internal_counter: sendTs,
			kind:             "send",
		}
		out <- Message{sender: id, ts: sendTs}

		msg := <-in
		clock.Receive(msg.ts)
		eventsChan <- Event{
			process:          id,
			internal_counter: clock.counter,
			kind:             "recv",
		}
	}
}

func exercitiul2() {
	fmt.Println("\n--- Exercitiul 2: Ceasuri Logice Lamport pe Inel P1 -> P2 -> P3 -> P1 ---")

	rounds := 2
	numProcesses := 3

	channels := make([]chan Message, numProcesses)
	for i := 0; i < numProcesses; i++ {
		channels[i] = make(chan Message, 10)
	}

	eventsChan := make(chan Event, 100)
	var wg sync.WaitGroup

	for i := 0; i < numProcesses; i++ {
		wg.Add(1)
		next := (i + 1) % numProcesses
		go runProcess(i+1, channels[i], channels[next], rounds, eventsChan, &wg)
	}

	wg.Wait()
	close(eventsChan)

	var events []Event
	for ev := range eventsChan {
		events = append(events, ev)
	}

	fmt.Printf("Total evenimente generate: %d\n", len(events))
	fmt.Println("\nEvenimentele inainte de sortare (ordinea colectarii):")
	for _, ev := range events {
		fmt.Printf("P%d | ceas: %2d | tip: %-8s\n", ev.process, ev.internal_counter, ev.kind)
	}

	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	fmt.Println("\nEvenimentele ordonate crescator (Ordine totala Lamport):")
	for _, ev := range events {
		fmt.Printf("P%d | ceas: %2d | tip: %-8s\n", ev.process, ev.internal_counter, ev.kind)
	}
	fmt.Println()
}

