package main

import (
	"fmt"
	// "time"
)

func main() {
	// BUFFERED CHANNELS
	// Принцип FIFO (first in first out)


	// now := time.Now()
	channel := make(chan string, 2)
	channel <- "First message"  // Когда даем CAPASITY, можно без функции, так
	channel <- "Second message"  
	
	// go func() {
	// 	channel <- "First message"
	// }()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
