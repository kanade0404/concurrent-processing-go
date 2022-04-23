package main

import (
	"fmt"
	"sync"
)

// 複数のチャネルから受信した値を一つの受信用チャネルの中にまとめる
func fanIn1(done chan struct{}, c1, c2 <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer fmt.Println("closed fanin")
		defer close(result)
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case num := <-c1:
				fmt.Println("send 1")
				result <- num
			case num := <-c2:
				fmt.Println("send 2")
				result <- num
			default:
				fmt.Println("continue")
				continue
			}
		}
	}()
	return result
}
func fanIn2(done chan struct{}, cs ...<-chan int) <-chan int {
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for i, c := range cs {
		go func(c <-chan int, i int) {
			defer wg.Done()
			for num := range c {
				select {
				case <-done:
					fmt.Println("wg.Done", i)
					return
				case result <- num:
					fmt.Println("send", i)
				}
			}
		}(c, i)
	}
	go func() {
		wg.Wait()
		fmt.Println("closing fanin")
		close(result)
	}()
	return result
}
func generator(done chan struct{}, num int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case result <- num:

			}
		}
	}()
	return result
}

func main() {
	done := make(chan struct{})
	gen1 := generator(done, 1)
	gen2 := generator(done, 2)
	result := fanIn1(done, gen1, gen2)
	for i := 0; i < 5; i++ {
		<-result
		fmt.Println(i)
	}
	close(done)
	fmt.Println("main close done")
	for {
		if _, ok := <-result; !ok {
			break
		}
	}
}
