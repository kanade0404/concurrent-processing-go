package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	num := rand.Intn(10)
	c <- num // 結果をチャネルに送信
}

func generator(done chan struct{}) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case result <- 1:

			}
		}
	}()
	return result
}

func main() {
	done := make(chan struct{})
	result := generator(done)
	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}
	close(done)
}
