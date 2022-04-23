package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getLuckNum() {
	fmt.Println("...")
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	num := rand.Intn(10)
	fmt.Printf("Today's your lucky number is %d!\n", num)
}

func main() {
	fmt.Println("what is today's lucky number?")
	var wg sync.WaitGroup // 初期化時点でのカウンタは0
	wg.Add(1)             // 内部カウントの値を+1する
	go func() {
		defer wg.Done() // ゴールーチンが終了した時にwgの内部カウントの値を-1するよう設定
		getLuckNum()
	}()
	wg.Wait() // wgの内部カウンタが0になるまでメインゴールーチンをブロックして待つ
}
