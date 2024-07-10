package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

func main() {

	// ONCE

	var wg sync.WaitGroup
	wg.Add(100)

	var once sync.Once

	for i := 0; i < 100; i++ {
		go func() {
			if foundTressure() {
				once.Do(markMissionedCompleted)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	checkMissiomCompletition() // default missionCompleted is False
}

func checkMissiomCompletition() {
	if missionCompleted {
		fmt.Println("Mission is now completed")
	} else {
		fmt.Println("Mission was a failure")
	}
}

func markMissionedCompleted() {
	missionCompleted = true
	fmt.Println("Marking mission completed")
}

func foundTressure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
