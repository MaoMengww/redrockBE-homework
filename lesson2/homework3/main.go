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
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main(){
	var wg sync.WaitGroup
	var counter = Counter{}
	wg.Add(100)
	for range 100{
		go func(){
			defer wg.Done()
			for range 10{
				counter.Increment()
		}
		}()
	}
	wg.Wait()
	fmt.Println("启动一百个协程,每个协程累加10次...")
	fmt.Println("最终计数：", counter.Value())
}
