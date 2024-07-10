package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	lock   sync.Mutex
	rwlock sync.RWMutex // read | write lock
)

func main() {

	// Golang Mutex / RWMutex

	// Пояснение: короч, помимо того, что операция инкрементирования не атомарна, здесь главный смысл в том, что есть операции, к примеру глобальные, или ещё какие то, в которых нам важно что бы кокретно в данный момент доступ к ним имела только
	// одна горутинка. Так же с операцией чтения и записи. Чтобы не читали | записывали все подряд в разнобой, лочим операцию только для одной рутинки

	basics() // EXAMPLE of sync.Mutex

	readAndWrite() // EXAMPLE of sync.RWMutex
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}

func basics() {

	// EXAMPLE of sync.Mutex

	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("Resulted count is: ", count)
}

func readAndWrite() {
	go read()
	go write()

	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}

func read() {
	rwlock.RLock()
	defer rwlock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(time.Second * 1)
	fmt.Println("Reading unlocking")
}

func write() {
	rwlock.Lock()
	defer rwlock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(time.Second * 1)
	fmt.Println("Write unlocking")
}
