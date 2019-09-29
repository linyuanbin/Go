package main

import (
	"fmt"
	"sync"
)
var lock sync.WaitGroup
func sayHello()  {
	fmt.Println("hello")
	lock.Done()
}

func main()  {
	lock.Add(1)
	go sayHello()
	lock.Wait()
	fmt.Println("main")
}

