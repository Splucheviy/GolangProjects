package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Clark", "Bobby", "Kent"}

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja) // в данном случае "go" запускает рутинки, параллелит. Можно было и без го, тогда бы метод выполнялся по очереди, но мы хотим параллелить
	}

	time.Sleep(time.Second * 2) // даем рутинкам время закончится, иначе main ветка уже закончит работать, а рутинок нифига не дождётся
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at ", target)
	time.Sleep(time.Second)
}
