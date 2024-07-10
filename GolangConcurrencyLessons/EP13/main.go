package main

import (
	"fmt"
	"sync"
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
			once.Do(func ()  {
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
		c.Signal()
	}()
	c.Wait()
}