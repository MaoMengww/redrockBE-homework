package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	count int
}

func (c *Counter) Increment(){
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Value() int{
	c.mu.Lock()
	now := c.count
	c.mu.Unlock()
	return now
}

func main(){
	var lock sync.Mutex
	count := 0
	var counter = Counter{
		mu: lock,
		count: count,
	}
	for i := 0; i < 100; i++{
		for i := 0; i < 10; i++{
			go counter.Increment()
		}
	}
	fmt.Println("启动一百个协程，每个协程累加10次...")
	fmt.Println("最终计数：", counter.count)
}