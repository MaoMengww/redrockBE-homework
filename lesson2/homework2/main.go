package main

import (
	"fmt"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	results <- fmt.Sprintf("%v 下载完成", filename)
}

func main(){
	var files = []string{"file1.zip", "file2.pdf", "file3.mp4"}
	results := make(chan string, 3)
	var wg sync.WaitGroup
	fmt.Println("开始下载 3 个文件...")
	for _, filename := range files {
		wg.Add(1)
		go download(filename, &wg, results)
	}
	wg.Wait()
	close(results)
	for result := range results {
		fmt.Println(result)
	}
	fmt.Println("所有文件下载完成!")
	
	
}
