package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func queue(id int, tipe *int, wg *sync.WaitGroup) {

	defer wg.Done() 

	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)

	fmt.Printf("Queue %d %d selesai dalam %v\n", id, &tipe, delay)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		func(i int) {
			queue(i,&i, &wg)
			wg.Wait()
		}(i)
	}
}
