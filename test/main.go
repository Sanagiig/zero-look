package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func work(i *int32) {
	atomic.AddInt32(i, 1)
}

func main() {
	i := int32(0)
	for j := 0; j < 1000; j++ {
		work(&i)
	}
	time.Sleep(time.Millisecond * 500)
	fmt.Println("done", i)
}
