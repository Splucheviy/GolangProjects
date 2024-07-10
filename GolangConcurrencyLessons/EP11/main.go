package main

import (
	"fmt"
	"sync"
)

func main() {

	// sync.Map

	// говорим что не можем сихронно работать с мапой, если пытаемся - то перед операцией необходимо ввести мутекс и залочить и после разлочить операцию, операция не атомарна
	// однако в стандартном пакете sync есть метод для синхронной работы с мапой sync.Map{} и свой набор методов


	regularMap := make(map[int]interface{})
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		regularMap[0] = i
	// 	}()
	// }

	syncMap := sync.Map{}

	// put
	regularMap[0] = 0
	regularMap[1] = 1
	regularMap[2] = 2

	syncMap.Store(0, 0)
	syncMap.Store(1, 1)
	syncMap.Store(2, 2)

	// get
	regularValue, regularOk := regularMap[2]
	fmt.Println(regularValue, regularOk)

	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(syncValue, syncOk)

	// delete
	regularMap[1] = nil
	syncMap.Delete(1)

	// get and delete
	syncValue, loaded := syncMap.LoadAndDelete(2)
	
	mu := sync.Mutex{}
	mu.Lock()
	regularValue = regularMap[2]
	delete(regularMap, 2)
	mu.Unlock()

	fmt.Println(syncValue, loaded, regularValue)

	// get and put
	syncValue, loaded = syncMap.LoadOrStore(1, 1)
	
	mu = sync.Mutex{}
	mu.Lock()
	regularValue, regularOk = regularMap[1]
	if regularOk {
		regularMap[1] = 1
		regularValue = regularMap[1]
	}
	mu.Unlock()
	fmt.Println(syncValue, regularValue)

	// range
	for key, value := range regularMap{
		fmt.Print(key, value, " | ")
	}
	fmt.Println()

	syncMap.Range(func(key, value interface{}) bool{
		fmt.Print(key, value, " | ")
		return true
	})
}
