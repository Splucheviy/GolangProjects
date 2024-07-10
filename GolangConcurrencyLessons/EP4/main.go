package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// открыли канал, читаем, что было изночально: numRounds задавали в main, затем перенесли в функцию, что есть гуд, но нам теперь нужно выводить несколько раз сообщение из канала, а мы не знаем сколько в канале поступает сообщений
	// тогда мы задали переменную message, и сказали что она равна длине канала, если в канал что то поступило - делай! Но!! лучше сделать булевскую переменную open на проверку доступен ли ещё пока канал и пока он доступен, считывать поступающие
	// в него данные

	channel := make(chan string)
	go throwingNinjaStar(channel)

	// for message := range channel {
	for {
		message, open := <-channel
		if !open{
			break
		}
		fmt.Println(message)
	}

}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 4
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("Your scored ", score)
	}
	close(channel)
}
