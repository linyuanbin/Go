package main

import (
	"fmt"
	"sync"
)
func main()  {
	var once sync.Once
	onceBody := func() {
		fmt.Println("执行了一次")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-done,"  ",i+1)
	}
}