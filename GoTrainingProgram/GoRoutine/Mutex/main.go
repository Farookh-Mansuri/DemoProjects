package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int // Common Resource
var mutex sync.Mutex

func increment(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

//go run -race main.go
func main() {
	wg.Add(2)
	go increment("Anamika")
	go increment("Monika")
	wg.Wait()
}
