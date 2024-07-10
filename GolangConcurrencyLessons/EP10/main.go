package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Окей, ладно, судя по всему работает для синхронизации между собой горутин, которые используют одни данные, чтобы не прокидывать в один канал. Сложно - отдельно читать

func main() {
	gettingReadyForMissionWithCond()
	broadcastStartOfMission()
}

func broadcastStartOfMission() {
	beeper := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)
	standByForMission(func() {
		fmt.Println("Ninja 1 starting the mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 2 starting the mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 3 starting the mission.")
		wg.Done()
	}, beeper)
	beeper.Broadcast()
	wg.Wait()
	fmt.Println("All ninjas have started their missions")
}

func standByForMission(fn func(), beeper *sync.Cond)  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		wg.Done()
		beeper.L.Lock()
		defer beeper.L.Unlock()
		beeper.Wait()
		fn()
	}()
	wg.Wait()

}

var ready bool

func gettingReadyForMissionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workIntervals := 0

	cond.L.Lock()
	for !ready {
		// time.Sleep(time.Second * 5)
		workIntervals++
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Printf("We are now ready! After %d work intervals. \n", workIntervals)
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
