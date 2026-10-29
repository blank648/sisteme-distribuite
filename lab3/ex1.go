package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// TODO:
// Calculeaza decalajul theta si round-trip-ul delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {

	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func exercitiul1() {

	serverClock := &Clock{offset: 0}
	clientClock := &Clock{offset: 500 * time.Millisecond}

	fmt.Printf("Offset initial: %v\n\n", clientClock.offset)

	thetaA, deltaA := measure(clientClock, serverClock)
	clientClock.offset += thetaA
	fmt.Printf("--- Punctul a ---\n")
	fmt.Printf("Theta: %v, Delta: %v\n", thetaA, deltaA)
	fmt.Printf("Offset dupa corectia (a): %v\n\n", clientClock.offset)

	clientClock.offset = 500 * time.Millisecond

}
