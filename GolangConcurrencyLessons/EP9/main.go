package main

import (
	"fmt"
	"sync"
)

func main() {

	// Resource Pool

	// sync.Pool for Get() and Put() methods
	
	// Короче, что то дико сложное. Смысл в выделении памяти для какой либо структуры. Сборщик мусора в GO не постоянно собирает мусор, а через определённые промежутки времени.
	// В случае если код выделяет память под некоторые структуры данных, а потом освобождает их - и так по кругу - короч, мы просим выделить runtime больше памяти на операцию, хотя предыдущий пул уже отработал, и ждёт
	// пока его утилизируют, и пока он ждёт и лежит в памяти, мы такие - чел плз, выдели ещё памяти, он такой - конечно ок, но ждать будешь долго. Такого нам не надо, для этого используем пул

	var numMemPieces int
	memPool := &sync.Pool{
		New: func() interface{} {
			numMemPieces++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	const numWorkers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			mem := memPool.Get().(*[]byte)
			memPool.Put(mem)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("%d numMemPieces were created", numMemPieces)
}
