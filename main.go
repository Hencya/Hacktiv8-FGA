package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func OrderingData(data interface{}, iteration int, mutex *sync.Mutex, c chan string) {
	mutex.Lock()
	c <- fmt.Sprintf("%v %d\n", data, iteration)
	mutex.Unlock()
}

func RandomingData(data interface{}, iteration int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%v %d\n", data, iteration)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func main() {
	fmt.Println("Hello World")

	fmt.Printf("\nData Random\n\n")
	coba := []interface{}{"coba1", "coba2", "coba3"}
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}

	wg := sync.WaitGroup{}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go RandomingData(coba, i, &wg)
		go RandomingData(bisa, i, &wg)
	}

	wg.Wait()

	fmt.Printf("\nData Terurut\n\n")
	chanCoba := make(chan string)
	chanBisa := make(chan string)

	lockCoba := sync.Mutex{}
	lockBisa := sync.Mutex{}

	for i := 1; i <= 4; i++ {
		go OrderingData(coba, i, &lockCoba, chanCoba)
		go OrderingData(bisa, i, &lockBisa, chanBisa)
	}

	for i := 1; i <= 4; i++ {
		fmt.Printf(<-chanCoba)
		fmt.Printf(<-chanBisa)
	}
}
