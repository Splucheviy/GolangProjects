package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkMutexAdd(t *testing.B) {
	var wg sync.WaitGroup
	wg.Add(iterations)
	var sum int64
	var mu sync.Mutex
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				mu.Lock()
				sum++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
}

func BenchmarkAtomicAdd(t *testing.B) {
	var wg sync.WaitGroup
	wg.Add(iterations)
	var sum int64
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&sum, 1)
			}
		}()
	}
	wg.Wait()
}
