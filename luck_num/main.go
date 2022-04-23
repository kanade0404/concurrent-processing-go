package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckNum(c chan<- int) {
	fmt.Println("...")
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	num := rand.Intn(10)
	c <- num // 結果をチャネルに送信
}

func main() {
	fmt.Println("what is today's lucky number?")
	c := make(chan int) // チャネル作成
	go getLuckNum(c)    // 引数にチャネルを渡す
	num := <-c          // チャネルから結果を受信
	fmt.Printf("Today's your lucky number is %d!\n", num)
	close(c)
}
