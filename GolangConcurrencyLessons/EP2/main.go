package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool, 0)
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal)
	// attack(evilNinja)
	// smokeSignal <- false //deadlock
	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja start at", target)
	attacked <- true
}
