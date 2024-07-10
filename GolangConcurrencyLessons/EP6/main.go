package main

import (
	"fmt"
	"sync"
)

func main() {

	// В этом примере работаем конкретно с горутинками, и нам важно, чтобы все эти горутинки отработали, прежде чем дальше пройдёт наш main и закончит работу программы преждевременно
	// Для этого применяем sync.WaitGrooup. Через Add, мы добавляем сколько горутинок у нас будет запущено, и сколько мы из них будем ждать
	// Собственно в примере мы добавили кол-во горутинок по длинне Slice'a (3 штуки) и через метод Wait будем эти горутиночки ждать
	// Что важно отметить, что когда горутиночка отработала, нам обязательно выделить что работа горутинки окончена, поэтому ставим для рутинки состояние (метод) Done
	// Done ставить обязательно, иначе вызов Wait даст ошибку

	var beeper sync.WaitGroup
	evilNinjas := []string{"Tommy", "Clark", "Kent"}
	beeper.Add(len(evilNinjas))

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, &beeper)
	}

	beeper.Wait()

	fmt.Println("Mission competed")
}

func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", evilNinja)

	beeper.Done()
}
