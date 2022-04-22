package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()
	taskStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}
	// ある人がある方向に向けて動いて、うまく動けたかどうか
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "%v", dirName)
		atomic.AddInt32(dir, 1) // ある方向への移動
		taskStep()              // 各人間は同じスピード、同じ歩調で動かなければならない
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprintf(out, ". Success!")
			return true
		}
		taskStep()
		atomic.AddInt32(dir, -1) // ある方向に進めないことが分かって戻って対応
		return false
	}
	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }
	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
