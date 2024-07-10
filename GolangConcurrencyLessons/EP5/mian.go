package main

import (
	"fmt"
	"time"
)

func main() {
	// channel statement

	// В общем и целом - не совсем понимаю почему в roughlyFair сразу закрывается канал,но в этом примере разбирается конкретно Select
	// Это кейс, когда у нас есть не 1, а несколько каналов, и нам нужно работать с обоими по доступности. Т.е не 1 канал а два
	// Мне пример понравился так себе, так как на оба канала отправляются две стринги, и оба канала по сути отвечают за одно действие
	// В хорошем примере преставим, что канала у нас два, на один отправляется/ются стринги, на другой инты
	// Закинуть это дело в 1 канал так себе идея, будем делить на два канала и пусть работают по доступности 

	ninja1, ninja2 := make(chan string, 0), make(chan string, 0)

	go captainElect(ninja1, "Ninja 1")
	go captainElect(ninja2, "Ninja 2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	default:
		fmt.Println("Neither")
	}

	roughlyFair()
}

func captainElect(ninja chan string, message string) {
	time.Sleep(time.Second * 3)
	ninja <- message
}

func roughlyFair() {
	ninja1 := make(chan interface{}, 0)
	close(ninja1)
	ninja2 := make(chan interface{}, 0)
	close(ninja2)

	var ninja1Count, ninja2Count int
	for i := 0; i < 1000; i++ {
		select {
		case <-ninja1:
			ninja1Count++
		case <-ninja2:
			ninja2Count++
		}
	}
	fmt.Printf("ninja1Count: %d, ninja2Count: %d", ninja1Count, ninja2Count)
}
