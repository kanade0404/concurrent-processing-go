package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Millisecond * 1000)
	defer t.Stop()
	for i := 0; i < 5; i++ {
		select {
		case <-t.C:
			fmt.Println("tick")
		}
	}
}
