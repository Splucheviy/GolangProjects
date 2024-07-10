package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// gorouting
	go func() {
		fmt.Println("New gorouting")
	}()
	fmt.Println("This is the old one")

	// channels
	ch := make(chan string)
	go func() {
		ch <- "Hello World"
	}()
	fmt.Println(<-ch)
	ch = make(chan string, 2)
	ch <- "1"
	ch <- "2"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("New gorouting")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("This is the old one")

	// Mutex
	iterations := 1000
	sumFirst := 0
	sumSecond := 0
	for i := 0; i < iterations; i++ {
		go func() {
			sumFirst++
		}()
	}

	wg.Add(iterations)
	var mu sync.Mutex
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			sumSecond++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(sumFirst, sumSecond)

	// Once
	sum := 0
	wg.Add(iterations)
	var once sync.Once
	for i := 0; i < iterations; i++ {
		go func() {
			once.Do(func() {
				sum++
			})
		}()
		wg.Done()
	}
	wg.Wait()
	fmt.Println(sum)

	// pool
	memPool := &sync.Pool{
		New: func() interface{} {
			mem := make([]byte, 1024)
			return &mem
		},
	}
	mem := memPool.Get().(*[]byte)
	// ...
	memPool.Put(mem)

	// cond
	c := sync.NewCond(&sync.Mutex{})
	go func() {
		c.L.Lock()
		// changing some condition
		c.L.Unlock()
		c.Signal()
		c.Broadcast()
	}()
	// checking condition
	c.L.Lock()
	c.Wait()
	c.L.Unlock()

	// gorouting safe-Map
	// regularMap := make(map[int]interface{})
	syncMap := sync.Map{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			// regularMap[0] = 1
			syncMap.Store(0, i)
			wg.Done()
		}()
	}
	wg.Wait()

	// atomic
	var i int64
	atomic.AddInt64(&i, 1)
	mu.Lock()
	i += 1
	mu.Unlock()

	var av atomic.Value
	type ninja struct{
		name string
	}
	av.Store(ninja{"Wallace"})
	av.Load()
}
