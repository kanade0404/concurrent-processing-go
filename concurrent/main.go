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

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}
	c := make(chan int)

	// srcの要素毎にある何か処理をして、結果をdstにいれる
	for _, s := range src {
		go func(s int, c chan int) {
			// 何か(重い)処理をする
			time.Sleep(time.Second)
			result := s * 2
			c <- result
		}(s, c)
	}
	for _ = range src {
		num := <-c
		dst = append(dst, num)
	}

	fmt.Println(dst)
	close(c)
}
