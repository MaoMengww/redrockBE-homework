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
	var counter = Counter{}
	for range 100{
		go func(){
			for range 10{
			go counter.Increment()
		}
		}()
	}
	fmt.Println("启动一百个协程,每个协程累加10次...")
	fmt.Println("最终计数：", counter.Value())
}
